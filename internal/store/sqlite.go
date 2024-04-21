package store

import (
	"context"

	"crawshaw.io/sqlite/sqlitex"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"go.uber.org/zap"
)

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

func (u *UserSqlite) Create(context.Context, model.User) error {
	return nil
}
func (u *UserSqlite) GetByEmail(context.Context, string) (model.User, error) {
	return model.User{}, nil
}
func (u *UserSqlite) GetByID(context.Context, int64) (model.User, error) {
	return model.User{}, nil
}
func (u *UserSqlite) GetAll(context.Context) (model.User, error) {
	return model.User{}, nil
}
func (u *UserSqlite) DeleteByID(context.Context, int64) error {
	return nil
}
func (u *UserSqlite) UpdateByID(context.Context, int64) error {
	return nil
}
