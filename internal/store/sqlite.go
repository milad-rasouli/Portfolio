package store

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"crawshaw.io/sqlite"
	"crawshaw.io/sqlite/sqlitex"
	"github.com/Milad75Rasouli/portfolio/internal/db"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const timeLayout = "2006-01-02 15:04:05"

type UserSqlite struct {
	dbPool *sqlitex.Pool
	logger *zap.Logger
}

func NewUserSqlite(dbPool *sqlitex.Pool, logger *zap.Logger) *UserSqlite {
	return &UserSqlite{
		dbPool: dbPool,
		logger: logger,
	}
}
func CreateSqliteTable(dbPool *sqlitex.Pool, cfg db.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectionTimeout)
	defer cancel()
	conn := dbPool.Get(ctx)
	defer dbPool.Put(conn)
	stmt, err := conn.Prepare(`CREATE TABLE IF NOT EXISTS user (
		id INTEGER,
		full_name TEXT NOT NULL,
		email INTEGER NOT NULL UNIQUE,
		password TEXT NOT NULL,
		is_github INTEGER DEFAULT 0,
		online_at DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		modified_at DATETIME,
		PRIMARY KEY(id)
		);`)
	if err != nil {
		return errors.Wrap(err, CannotCreateTableError.Error())
	}
	defer stmt.Finalize()
	_, err = stmt.Step()

	return err
}
func (u UserSqlite) parseToUser(stmt *sqlite.Stmt) (model.User, error) {
	var (
		usr model.User
		err error
	)
	usr.ID = stmt.GetInt64("id")
	usr.FullName = stmt.GetText("full_name")
	usr.Email = stmt.GetText("email")
	usr.Password = stmt.GetText("password")
	usr.IsGithub = stmt.GetInt64("is_github")
	usr.OnlineAt, err = time.Parse(timeLayout, stmt.GetText("online_at"))
	if err != nil {
		return usr, err
	}
	usr.CreatedAt, err = time.Parse(timeLayout, stmt.GetText("created_at"))
	if err != nil {
		return usr, err
	}
	usr.ModifiedAt, err = time.Parse(timeLayout, stmt.GetText("modified_at"))
	return usr, err
}
func (u *UserSqlite) Create(ctx context.Context, usr model.User) error {
	conn := u.dbPool.Get(ctx)
	defer u.dbPool.Put(conn)
	stmt, err := conn.Prepare(`INSERT INTO user (full_name, email, password, is_github,online_at, modified_at, created_at) 
		VALUES($1,$2,$3,$4,$5,$6,$7);`)
	if err != nil {
		return errors.Errorf("unable to create the new user %s", err.Error())
	}
	defer stmt.Finalize()
	stmt.SetText("$1", usr.FullName)
	stmt.SetText("$2", usr.Email)
	stmt.SetText("$3", usr.Password)
	stmt.SetInt64("$4", usr.IsGithub)
	stmt.SetText("$5", usr.OnlineAt.Format(timeLayout))
	stmt.SetText("$6", usr.ModifiedAt.Format(timeLayout))
	stmt.SetText("$7", usr.CreatedAt.Format(timeLayout))

	_, err = stmt.Step()
	if err != nil {
		e := err.Error()[18:42]
		log.Println(e)
		if e == "SQLITE_CONSTRAINT_UNIQUE" {
			return DuplicateUserError
		}
	}
	return err
}
func (u *UserSqlite) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var usr model.User
	conn := u.dbPool.Get(ctx)
	defer u.dbPool.Put(conn)
	stmt, err := conn.Prepare(`SELECT * FROM user WHERE email=$1 LIMIT 1;`)
	if err != nil {
		return usr, errors.Errorf("unable to get the user %s from email", err.Error())
	}
	defer stmt.Finalize()
	stmt.SetText("$1", email)

	var hasRow bool
	hasRow, err = stmt.Step()
	if hasRow == false {
		return usr, UserNotFountError
	}
	if err != nil {
		return usr, err
	}
	usr, err = u.parseToUser(stmt)
	return usr, err
}
func (u *UserSqlite) GetByID(ctx context.Context, id int64) (model.User, error) {
	var usr model.User
	conn := u.dbPool.Get(ctx)
	defer u.dbPool.Put(conn)
	stmt, err := conn.Prepare(`SELECT * FROM user WHERE id=$1 LIMIT 1;`)
	if err != nil {
		return usr, errors.Errorf("unable to get the user %s from id", err.Error())
	}
	defer stmt.Finalize()
	stmt.SetInt64("$1", id)

	var hasRow bool
	hasRow, err = stmt.Step()
	if hasRow == false {
		return usr, UserNotFountError
	}
	if err != nil {
		return usr, err
	}
	usr, err = u.parseToUser(stmt)
	return usr, err
}
func (u *UserSqlite) GetAll(ctx context.Context) ([]model.User, error) {
	var usr []model.User
	conn := u.dbPool.Get(ctx)
	defer u.dbPool.Put(conn)
	stmt, err := conn.Prepare(`SELECT * FROM user;`)
	if err != nil {
		return usr, errors.Errorf("unable to get all users %s", err.Error())
	}
	defer stmt.Finalize()

	for {
		var (
			swapUser model.User
			hasRow   bool
		)
		hasRow, err = stmt.Step()
		if hasRow == false {
			break
		}
		swapUser, err = u.parseToUser(stmt)
		if err != nil {
			return usr, errors.Errorf("getting the users from database error %s", err.Error())
		}
		usr = append(usr, swapUser)
	}
	return usr, err
}
func (u *UserSqlite) DeleteByID(ctx context.Context, id int64) error {
	conn := u.dbPool.Get(ctx)
	defer u.dbPool.Put(conn)
	stmt, err := conn.Prepare(`DELETE FROM user WHERE id=$1;`)
	if err != nil {
		return err
	}
	defer stmt.Finalize()
	stmt.SetInt64("$1", id)
	_, err = stmt.Step()
	return err
}
func (u *UserSqlite) UpdatePasswordFullName(ctx context.Context, id int64, password string, fullname string) error {
	conn := u.dbPool.Get(ctx)
	defer u.dbPool.Put(conn)
	var s string
	if len(password) != 0 && len(fullname) != 0 {
		s = fmt.Sprintf(`UPDATE user
		SET password='%s', full_name='%s'
		WHERE id=%d;`, password, fullname, id)
	} else if len(password) != 0 {
		s = fmt.Sprintf(`UPDATE user
		SET password='%s'
		WHERE id=%d;`, password, id)
	} else if len(fullname) != 0 {
		s = fmt.Sprintf(`UPDATE user
		SET full_name='%s'
		WHERE id=%d;`, fullname, id)
	}
	stmt, err := conn.Prepare(s)
	if err != nil {
		return err
	}
	defer stmt.Finalize()
	var hasRow bool
	hasRow, err = stmt.Step()
	if hasRow {
		return UserNotFountError
	}
	return err
}

type SqliteInit struct {
	Folder string
}

func (d *SqliteInit) Init(isTestMode bool, config db.Config, logger *zap.Logger) (*UserSqlite, func(), error) {
	var userDB *UserSqlite

	os.Mkdir(d.Folder, 0777)
	cfg := config
	if isTestMode == true {
		cfg = db.Config{
			IsSqlite:          true,
			ConnectionTimeout: time.Millisecond * 200,
		}
	}
	dbPool, err := db.New(cfg)
	if err != nil {
		return nil, nil, err
	}
	err = CreateSqliteTable(dbPool, cfg)
	if err != nil {
		return nil, nil, err
	}
	userDB = NewUserSqlite(dbPool, logger)
	return userDB, func() {
		err := dbPool.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
		if isTestMode == true {
			err = os.RemoveAll(d.Folder)
			if err != nil {
				log.Printf("Error removing database folder: %v", err)
			}
		}
	}, nil
}
