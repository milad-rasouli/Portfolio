package sqlitedb

import (
	"context"

	"crawshaw.io/sqlite"
	"crawshaw.io/sqlite/sqlitex"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/Milad75Rasouli/portfolio/internal/store"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// type AboutMe interface {
// 	UpdateAboutMe(context.Context, model.AboutMe) error
// 	GetAboutMe(context.Context) (model.AboutMe, error)
// }

const (
	AboutMeID = 1
)

type AboutMeSqlite struct {
	dbPool *sqlitex.Pool
	logger *zap.Logger
}

func NewAboutMeSqlite(dbPool *sqlitex.Pool, logger *zap.Logger) *AboutMeSqlite {
	return &AboutMeSqlite{
		dbPool: dbPool,
		logger: logger,
	}
}
func (am AboutMeSqlite) parseToAboutMe(stmt *sqlite.Stmt) model.AboutMe {
	var (
		aboutMe model.AboutMe
	)
	aboutMe.ID = stmt.GetInt64("id")
	aboutMe.Content = stmt.GetText("content")
	return aboutMe
}
func (am *AboutMeSqlite) createAboutMe(ctx context.Context, aboutMe model.AboutMe) error {
	conn := am.dbPool.Get(ctx)
	defer am.dbPool.Put(conn)
	//TODO: Get the id 1 then create it
	stmt, err := conn.Prepare(`INSERT INTO about_me (id, content)
	VALUES($1,$2);`)
	if err != nil {
		return errors.Errorf("unable to create about me %s", err.Error())
	}
	defer stmt.Finalize()

	stmt.SetInt64("$1", 1)
	stmt.SetText("$2", aboutMe.Content)

	_, err = stmt.Step()
	if err != nil {
		return err
	}

	return err
}

func (am *AboutMeSqlite) GetAboutMe(ctx context.Context) (model.AboutMe, error) {
	var aboutMe model.AboutMe
	conn := am.dbPool.Get(ctx)
	defer am.dbPool.Put(conn)
	stmt, err := conn.Prepare(`SELECT * FROM about_me WHERE id=$1;`)
	if err != nil {
		return aboutMe, errors.Errorf("unable to get about me %s", err.Error())
	}
	defer stmt.Finalize()
	stmt.SetInt64("$1", AboutMeID)

	var hasRow bool
	hasRow, err = stmt.Step()
	if hasRow == false {
		return aboutMe, store.UserNotFountError
	}
	if err != nil {
		return aboutMe, err
	}
	aboutMe = am.parseToAboutMe(stmt)
	return aboutMe, err
}

func (am *AboutMeSqlite) UpdateAboutMe(ctx context.Context, aboutMe model.AboutMe) error {
	conn := am.dbPool.Get(ctx)
	defer am.dbPool.Put(conn)
	var s = `UPDATE about_me
		SET content=$1
		WHERE id=$2;`
	stmt, err := conn.Prepare(s)
	if err != nil {
		return err
	}
	defer stmt.Finalize()
	stmt.SetInt64("$2", AboutMeID)
	stmt.SetText("$1", aboutMe.Content)

	_, err = stmt.Step()
	if conn.Changes() == 0 {
		return am.createAboutMe(ctx, aboutMe)
	}
	return err
}
