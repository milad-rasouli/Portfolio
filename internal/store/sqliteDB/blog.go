package sqlitedb

import (
	"context"
	"time"

	"crawshaw.io/sqlite"
	"crawshaw.io/sqlite/sqlitex"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/Milad75Rasouli/portfolio/internal/store"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type BlogSqlite struct {
	dbPool *sqlitex.Pool
	logger *zap.Logger
}

func NewBlogSqlite(dbPool *sqlitex.Pool, logger *zap.Logger) *BlogSqlite {
	return &BlogSqlite{
		dbPool: dbPool,
		logger: logger,
	}
}

func (b BlogSqlite) parseToUser(stmt *sqlite.Stmt) (model.Blog, error) {
	var (
		blog model.Blog
		err  error
	)
	blog.ID = stmt.GetInt64("id")
	blog.Title = stmt.GetText("title")
	blog.Body = stmt.GetText("body")
	blog.Caption = stmt.GetText("caption")
	blog.ImagePath = stmt.GetText("image_path")
	blog.CreatedAt, err = time.Parse(timeLayout, stmt.GetText("created_at"))
	if err != nil {
		return blog, err
	}
	blog.ModifiedAt, err = time.Parse(timeLayout, stmt.GetText("modified_at"))

	if err != nil {
		return blog, err
	}
	blog.CreatedAt, err = time.Parse(timeLayout, stmt.GetText("created_at"))
	if err != nil {
		return blog, err
	}
	blog.ModifiedAt, err = time.Parse(timeLayout, stmt.GetText("modified_at"))
	return blog, err
}

func (b *BlogSqlite) CreateBlog(ctx context.Context, blog model.Blog) (int64, error) {
	var rowID int64
	conn := b.dbPool.Get(ctx)
	defer b.dbPool.Put(conn)

	stmt, err := conn.Prepare(`INSERT INTO post (title, body, caption, image_path,created_at,modified_at)
	VALUES ($1, $2, $3, $4, $5, $6);`)
	if err != nil {
		return rowID, errors.Errorf("unable to create the new blog %s", err.Error())
	}
	defer stmt.Finalize()
	stmtSelect, err := conn.Prepare(`SELECT last_insert_rowid();`)
	if err != nil {
		return 0, errors.Errorf("unable to prepare the select statement: %s", err.Error())
	}
	defer stmtSelect.Finalize()
	stmt.SetText("$1", blog.Title)
	stmt.SetText("$2", blog.Body)
	stmt.SetText("$3", blog.Caption)
	stmt.SetText("$4", blog.ImagePath)
	stmt.SetText("$5", blog.CreatedAt.Format(timeLayout))
	stmt.SetText("$6", blog.ModifiedAt.Format(timeLayout))

	_, err = stmt.Step()
	if err != nil {
		e := err.Error()[18:42]
		if e == "SQLITE_CONSTRAINT_UNIQUE" {
			return rowID, store.DuplicateUserError
		}
		return rowID, err
	}
	hasRow, err := stmtSelect.Step()
	if err != nil {
		return rowID, err
	}

	if hasRow {
		rowID = conn.LastInsertRowID()
	}
	return rowID, err
}

func (b *BlogSqlite) GetByID(ctx context.Context, id int64) (model.Blog, error) {
	var blog model.Blog
	conn := b.dbPool.Get(ctx)
	defer b.dbPool.Put(conn)
	stmt, err := conn.Prepare(`SELECT * FROM post WHERE id=$1 LIMIT 1;`)
	if err != nil {
		return blog, errors.Errorf("unable to get the post %s from id", err.Error())
	}
	defer stmt.Finalize()
	stmt.SetInt64("$1", id)

	var hasRow bool
	hasRow, err = stmt.Step()
	if hasRow == false {
		return blog, store.UserNotFountError
	}
	if err != nil {
		return blog, err
	}
	blog, err = b.parseToUser(stmt)
	return blog, err
}

func (b *BlogSqlite) GetAll(ctx context.Context) ([]model.Blog, error) {
	var blog []model.Blog
	conn := b.dbPool.Get(ctx)
	defer b.dbPool.Put(conn)
	stmt, err := conn.Prepare(`SELECT * FROM post;`)
	if err != nil {
		return blog, errors.Errorf("unable to get all blog %s", err.Error())
	}
	defer stmt.Finalize()

	for {
		var (
			swapUser model.Blog
			hasRow   bool
		)
		hasRow, err = stmt.Step()
		if hasRow == false {
			break
		}
		swapUser, err = b.parseToUser(stmt)
		if err != nil {
			return blog, errors.Errorf("getting the blog from database error %s", err.Error())
		}
		blog = append(blog, swapUser)
	}
	return blog, err
}

func (u *BlogSqlite) DeleteByID(ctx context.Context, id int64) error {
	conn := u.dbPool.Get(ctx)
	defer u.dbPool.Put(conn)
	stmt, err := conn.Prepare(`DELETE FROM post WHERE id=$1;`)
	if err != nil {
		return err
	}
	defer stmt.Finalize()
	stmt.SetInt64("$1", id)
	_, err = stmt.Step()
	return err
}

// func (u *BlogSqlite) Update(ctx context.Context, id int64, password string, fullname string) error {
// 	conn := u.dbPool.Get(ctx)
// 	defer u.dbPool.Put(conn)
// 	var s string
// 	if len(password) != 0 && len(fullname) != 0 {
// 		s = fmt.Sprintf(`UPDATE user
// 		SET password='%s', full_name='%s'
// 		WHERE id=%d;`, password, fullname, id)
// 	} else if len(password) != 0 {
// 		s = fmt.Sprintf(`UPDATE user
// 		SET password='%s'
// 		WHERE id=%d;`, password, id)
// 	} else if len(fullname) != 0 {
// 		s = fmt.Sprintf(`UPDATE user
// 		SET full_name='%s'
// 		WHERE id=%d;`, fullname, id)
// 	}
// 	stmt, err := conn.Prepare(s)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Finalize()
// 	var hasRow bool
// 	hasRow, err = stmt.Step()
// 	if hasRow {
// 		return store.UserNotFountError
// 	}
// 	return err
// }
