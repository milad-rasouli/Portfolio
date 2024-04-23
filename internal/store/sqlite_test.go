package store

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/db"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestUserCRUD(t *testing.T) {
	var userDB *UserSqlite
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

	os.Mkdir("data", 0777)

	{
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
		userDB = NewUserSqlite(dbPool, nil)
	}

	{
		err = userDB.Create(context.TODO(), user)
		if err != nil {
			t.Error(err)
		}
	}

	{
		fetchedUser, err := userDB.GetByID(context.TODO(), user.ID)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, user, fetchedUser, "user and fetchUser should be equal ")
	}

	{
		fetchedUser, err := userDB.GetByEmail(context.TODO(), user.Email)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, user, fetchedUser, "user and fetchUser should be equal ")
	}

	{
		type caseType struct {
			Password   string
			FullName   string
			ID         int64
			testTarget func(model.User, caseType) bool
		}
		cases := []caseType{{

			Password: "password1234567890",
			FullName: "fullname1234567890",
			ID:       1,
			testTarget: func(u model.User, c caseType) bool {
				return u.Password == c.Password && u.FullName == c.FullName
			},
		}, {
			Password: "0987654321password",
			ID:       1,
			testTarget: func(u model.User, c caseType) bool {
				return u.Password == c.Password
			},
		}, {
			FullName: "0987654321Fullname",
			ID:       1,
			testTarget: func(u model.User, c caseType) bool {
				return u.FullName == c.FullName
			},
		},
		}

		for _, item := range cases {
			err = userDB.UpdatePasswordFullName(context.TODO(), item.ID, item.Password, item.FullName)
			assert.NoError(t, err, "unable to update user")
			expectedUser, err := userDB.GetByID(context.TODO(), item.ID)
			assert.NoError(t, err, "unable to read user")
			assert.True(t, item.testTarget(expectedUser, item), "parameters do not match")
		}
	}

	{
		_, err = userDB.GetByID(context.TODO(), 5)
		assert.Error(t, err)
	}

	{
		err := userDB.DeleteByID(context.TODO(), user.ID)
		assert.NoError(t, err, err)
		_, err = userDB.GetByID(context.TODO(), user.ID)
		assert.Error(t, err)
	}

	err = os.RemoveAll("data")
	if err != nil {
		t.Error("Error removing directory:", err)
		return
	}
}
