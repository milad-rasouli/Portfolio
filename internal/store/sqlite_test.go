package store

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/db"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/stretchr/testify/assert"
)

type DatabaseInit struct {
	folder string
}
type cancelDB func()

func (d *DatabaseInit) Init() (*UserSqlite, cancelDB, error) {
	var userDB *UserSqlite

	os.Mkdir(d.folder, 0777)
	cfg := db.Config{
		IsSqlite:          true,
		ConnectionTimeout: time.Millisecond * 200,
	}
	dbPool, err := db.New(cfg)
	if err != nil {
		return nil, nil, err
	}
	err = CreateSqliteTable(dbPool, cfg)
	if err != nil {
		return nil, nil, err
	}
	userDB = NewUserSqlite(dbPool, nil)
	return userDB, func() {
		err := dbPool.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
		err = os.RemoveAll(d.folder)
		if err != nil {
			log.Printf("Error removing database folder: %v", err)
		}
	}, nil
}

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

	d := DatabaseInit{folder: "data"}
	userDB, cancel, err := d.Init()
	assert.NoError(t, err)
	defer cancel()

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

}

func BenchmarkCreateUser(b *testing.B) {
	d := DatabaseInit{folder: "data"}
	userDB, cancel, err := d.Init()
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
		err = userDB.Create(context.TODO(), user)
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
