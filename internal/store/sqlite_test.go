package store

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/db"
	"github.com/Milad75Rasouli/portfolio/internal/model"
)

func TestUserDB(t *testing.T) {
	cfg := db.Config{
		IsSqlite:          true,
		ConnectionTimeout: time.Millisecond * 200,
	}
	dbPool, err := db.New(cfg)
	if err != nil {
		t.Error(err)
	}

	err = CreateSqliteTable(dbPool, cfg)
	if err != nil {
		t.Error(err)
	}

	user := model.User{
		FullName:   "foo bar",
		Email:      "bar@baz.com",
		Password:   "foobarbaz",
		IsGithub:   0,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
		OnlineAt:   time.Now(),
	}
	db := NewUserSqlite(dbPool, nil)
	err = db.Create(context.TODO(), user)
	if err != nil {
		t.Error(err)
	}
	fetchedUser, err := db.GetByID(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}

	reflect.DeepEqual(fetchedUser, user)
}
