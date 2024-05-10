package sqlitedb

import (
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/db"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestUserCRUD(t *testing.T) {
	var userDB *StoreSqlite
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

	d := SqliteInit{Folder: "data"}
	userDB, cancel, err := d.Init(true, db.Config{}, nil)
	assert.NoError(t, err)
	defer cancel()

	{
		_, err = userDB.CreateUser(context.TODO(), user)
		if err != nil {
			t.Error(err)
		}
	}

	{
		fetchedUser, err := userDB.GetUserByID(context.TODO(), user.ID)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, user, fetchedUser, "user and fetchUser should be equal ")
	}

	{
		fetchedUser, err := userDB.GetUserByEmail(context.TODO(), user.Email)
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
			err = userDB.UpdateUserByPasswordFullName(context.TODO(), item.ID, item.Password, item.FullName)
			assert.NoError(t, err, "unable to update user")
			expectedUser, err := userDB.GetUserByID(context.TODO(), item.ID)
			assert.NoError(t, err, "unable to read user")
			assert.True(t, item.testTarget(expectedUser, item), "parameters do not match")
		}
	}

	{
		_, err = userDB.GetUserByID(context.TODO(), 5)
		assert.Error(t, err)
	}

	{
		err := userDB.DeleteUserByID(context.TODO(), user.ID)
		assert.NoError(t, err, err)
		_, err = userDB.GetUserByID(context.TODO(), user.ID)
		assert.Error(t, err)
	}

}

func BenchmarkCreateUser(b *testing.B) {
	d := SqliteInit{Folder: "data"}
	userDB, cancel, err := d.Init(true, db.Config{}, nil)
	assert.NoError(b, err)
	defer cancel()

	b.ResetTimer()

	var totalDuration time.Duration
	var totalUsers int

	for i := 0; i < b.N; i++ {
		now := time.Now()
		user := model.User{
			FullName:   "foo",
			Password:   "barbaz123",
			Email:      strconv.FormatUint(uint64(rand.Int63()), 10),
			OnlineAt:   now,
			ModifiedAt: now,
		}

		start := time.Now()
		_, err = userDB.CreateUser(context.TODO(), user)
		assert.NoError(b, err)
		elapsed := time.Since(start)

		totalDuration += elapsed
		totalUsers++
	}

	b.StopTimer()

	b.Logf("Total Users Created: %d", totalUsers)
	b.Logf("Total Time Taken: %v", totalDuration)
	b.Logf("Avg. Time per User: %v", totalDuration/time.Duration(totalUsers))
	b.Logf("Ops/second: %f", float64(totalUsers)/totalDuration.Seconds())
}
