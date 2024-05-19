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

const (
	HomeID = 1
)

type HomeSqlite struct {
	dbPool *sqlitex.Pool
	logger *zap.Logger
}

func NewHomeSqlite(dbPool *sqlitex.Pool, logger *zap.Logger) *HomeSqlite {
	return &HomeSqlite{
		dbPool: dbPool,
		logger: logger,
	}
}
func (h HomeSqlite) parseToHome(stmt *sqlite.Stmt) model.Home {
	var (
		home model.Home
	)

	home.Name = stmt.GetText("name")
	home.Slogan = stmt.GetText("slogan")
	home.ShortIntro = stmt.GetText("short_intro")
	home.GithubUrl = stmt.GetText("github_url")
	return home
}
func (h *HomeSqlite) createHome(ctx context.Context, home model.Home) error {
	conn := h.dbPool.Get(ctx)
	defer h.dbPool.Put(conn)
	stmt, err := conn.Prepare(`INSERT INTO home (id, name, slogan, short_intro, github_url)
	VALUES($1,$2,$3,$4,$5);`)
	if err != nil {
		return errors.Errorf("unable to create home %s", err.Error())
	}
	defer stmt.Finalize()

	stmt.SetInt64("$1", HomeID)
	stmt.SetText("$2", home.Name)
	stmt.SetText("$3", home.Slogan)
	stmt.SetText("$4", home.ShortIntro)
	stmt.SetText("$5", home.GithubUrl)

	_, err = stmt.Step()
	return err
}

func (h *HomeSqlite) GetHome(ctx context.Context) (model.Home, error) {
	var home model.Home
	conn := h.dbPool.Get(ctx)
	defer h.dbPool.Put(conn)
	stmt, err := conn.Prepare(`SELECT * FROM home WHERE id=$1;`)
	if err != nil {
		return home, errors.Errorf("unable to get home %s", err.Error())
	}
	defer stmt.Finalize()
	stmt.SetInt64("$1", HomeID)

	var hasRow bool
	hasRow, err = stmt.Step()
	if err != nil {
		return home, err
	}
	if hasRow == false {
		return home, store.HomeNotFountError
	}
	home = h.parseToHome(stmt)
	return home, err
}

func (h *HomeSqlite) UpdateHome(ctx context.Context, home model.Home) error {
	conn := h.dbPool.Get(ctx)
	defer h.dbPool.Put(conn)

	stmt, err := conn.Prepare(`UPDATE home
		SET name=$1, slogan=$2, short_intro=$3, github_url=$4
		WHERE id=$5;`)
	if err != nil {
		return err
	}
	defer stmt.Finalize()
	stmt.SetInt64("$5", HomeID)
	stmt.SetText("$1", home.Name)
	stmt.SetText("$2", home.Slogan)
	stmt.SetText("$3", home.ShortIntro)
	stmt.SetText("$4", home.GithubUrl)

	_, err = stmt.Step()
	if conn.Changes() == 0 {
		return h.createHome(ctx, home)
	}
	return err
}
