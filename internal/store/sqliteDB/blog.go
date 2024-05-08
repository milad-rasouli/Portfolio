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

func (b BlogSqlite) parseToBlog(stmt *sqlite.Stmt) (model.Blog, error) {
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

func (b BlogSqlite) parseToCategory(stmt *sqlite.Stmt) (model.Category, error) {
	var (
		blog model.Category
		err  error
	)
	blog.ID = stmt.GetInt64("id")
	blog.Name = stmt.GetText("name")
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

func (b *BlogSqlite) GetBlogByID(ctx context.Context, id int64) (model.Blog, error) {
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
		return blog, store.BlogNotFoundError
	}
	if err != nil {
		return blog, err
	}
	blog, err = b.parseToBlog(stmt)
	return blog, err
}

func (b *BlogSqlite) GetAllBlog(ctx context.Context) ([]model.Blog, error) {
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
			swapBlog model.Blog
			hasRow   bool
		)
		hasRow, err = stmt.Step()
		if hasRow == false {
			break
		}
		swapBlog, err = b.parseToBlog(stmt)
		if err != nil {
			return blog, errors.Errorf("getting the blog from database error %s", err.Error())
		}
		blog = append(blog, swapBlog)
	}
	return blog, err
}

func (u *BlogSqlite) DeleteBlogByID(ctx context.Context, id int64) error {
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

func (u *BlogSqlite) UpdateBlogByID(ctx context.Context, blog model.Blog) error {
	conn := u.dbPool.Get(ctx)
	defer u.dbPool.Put(conn)
	var s string
	s = `UPDATE post
	SET title=$1, body=$2, caption=$3, image_path=$4
	WHERE id=$5;`
	stmt, err := conn.Prepare(s)
	if err != nil {
		return err
	}
	stmt.SetText("$1", blog.Title)
	stmt.SetText("$2", blog.Body)
	stmt.SetText("$3", blog.Caption)
	stmt.SetText("$4", blog.ImagePath)
	stmt.SetInt64("$5", blog.ID)
	defer stmt.Finalize()
	var hasRow bool
	hasRow, err = stmt.Step()
	if hasRow {
		return store.BlogNotFoundError
	}
	return err
}

/******************* Category *******************/

func (b *BlogSqlite) CreateCategory(ctx context.Context, category model.Category) (int64, error) {
	var rowID int64
	conn := b.dbPool.Get(ctx)
	defer b.dbPool.Put(conn)

	stmt, err := conn.Prepare(`INSERT INTO category (name) VALUES ($1);`)
	if err != nil {
		return rowID, errors.Errorf("unable to create the new category %s", err.Error())
	}
	defer stmt.Finalize()
	stmtSelect, err := conn.Prepare(`SELECT last_insert_rowid();`)
	if err != nil {
		return 0, errors.Errorf("unable to prepare the select statement: %s", err.Error())
	}
	defer stmtSelect.Finalize()
	stmt.SetText("$1", category.Name)

	_, err = stmt.Step()
	if err != nil {
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

func (b *BlogSqlite) GetCategoryByID(ctx context.Context, id int64) (model.Category, error) {
	var blog model.Category
	conn := b.dbPool.Get(ctx)
	defer b.dbPool.Put(conn)
	stmt, err := conn.Prepare(`SELECT * FROM category WHERE id=$1 LIMIT 1;`)
	if err != nil {
		return blog, errors.Errorf("unable to get the category %s from id", err.Error())
	}
	defer stmt.Finalize()
	stmt.SetInt64("$1", id)

	var hasRow bool
	hasRow, err = stmt.Step()
	if hasRow == false {
		return blog, store.CategoryNotFoundError
	}
	if err != nil {
		return blog, err
	}
	blog, err = b.parseToCategory(stmt)
	return blog, err
}

func (b *BlogSqlite) GetAllCategory(ctx context.Context) ([]model.Category, error) {
	var blog []model.Category
	conn := b.dbPool.Get(ctx)
	defer b.dbPool.Put(conn)
	stmt, err := conn.Prepare(`SELECT * FROM category;`)
	if err != nil {
		return blog, errors.Errorf("unable to get all category %s", err.Error())
	}
	defer stmt.Finalize()

	for {
		var (
			swapCategory model.Category
			hasRow       bool
		)
		hasRow, err = stmt.Step()
		if hasRow == false {
			break
		}
		swapCategory, err = b.parseToCategory(stmt)
		if err != nil {
			return blog, errors.Errorf("getting the category from database error %s", err.Error())
		}
		blog = append(blog, swapCategory)
	}
	return blog, err
}

func (u *BlogSqlite) DeleteCategoryByID(ctx context.Context, id int64) error {
	conn := u.dbPool.Get(ctx)
	defer u.dbPool.Put(conn)
	stmt, err := conn.Prepare(`DELETE FROM category WHERE id=$1;`)
	if err != nil {
		return err
	}
	defer stmt.Finalize()
	stmt.SetInt64("$1", id)
	_, err = stmt.Step()
	return err
}

func (u *BlogSqlite) UpdateCategoryByID(ctx context.Context, blog model.Category) error {
	conn := u.dbPool.Get(ctx)
	defer u.dbPool.Put(conn)
	var s string
	s = `UPDATE category SET name=$1 WHERE id=$2;`
	stmt, err := conn.Prepare(s)
	if err != nil {
		return err
	}
	stmt.SetText("$1", blog.Name)
	stmt.SetInt64("$2", blog.ID)
	var hasRow bool
	hasRow, err = stmt.Step()
	if hasRow {
		return store.BlogNotFoundError
	}
	return err
}
