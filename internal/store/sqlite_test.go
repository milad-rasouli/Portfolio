package store

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/db"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestUserDB(t *testing.T) {
	os.Mkdir("data", 0777)

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

	ti, err := time.Parse(timeLayout, time.Now().Format(timeLayout))
	user := model.User{
		ID:         1,
		FullName:   "foo bar",
		Email:      "bar@baz.com",
		Password:   "foobarbaz",
		IsGithub:   0,
		CreatedAt:  ti,
		ModifiedAt: ti,
		OnlineAt:   ti,
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

	a := fmt.Sprintf("%+v, %+v", user, fetchedUser)
	assert.Equal(t, user, fetchedUser, "user and fetchUser should be equal "+a)
	err = os.RemoveAll("data")
	if err != nil {
		t.Error("Error removing directory:", err)
		return
	}
}
