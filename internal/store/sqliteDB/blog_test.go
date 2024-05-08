package sqlitedb

import (
	"context"
	"testing"
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/db"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestBlogCRUD(t *testing.T) {
	var blogDB *BlogSqlite
	ti, err := time.Parse(timeLayout, time.Now().Format(timeLayout))
	blog := model.Blog{
		ID:         1,
		Title:      "foo bar",
		Body:       "bar baz",
		Caption:    "foo",
		ImagePath:  "/foo/bar",
		CreatedAt:  ti,
		ModifiedAt: ti,
	}

	d := SqliteInit{Folder: "data"}
	_, blogDB, cancel, err := d.Init(true, db.Config{}, nil)
	assert.NoError(t, err)
	defer cancel()

	{
		_, err = blogDB.CreateBlog(context.TODO(), blog)
		if err != nil {
			t.Error(err)
		}
	}

	{
		fetchedBlog, err := blogDB.GetByID(context.TODO(), blog.ID)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, blog, fetchedBlog, "blog and fetchUser should be equal ")
	}

	// {
	// 	type caseType struct {
	// 		Password   string
	// 		FullName   string
	// 		ID         int64
	// 		testTarget func(model.User, caseType) bool
	// 	}
	// 	cases := []caseType{{

	// 		Password: "password1234567890",
	// 		FullName: "fullname1234567890",
	// 		ID:       1,
	// 		testTarget: func(u model.User, c caseType) bool {
	// 			return u.Password == c.Password && u.FullName == c.FullName
	// 		},
	// 	}, {
	// 		Password: "0987654321password",
	// 		ID:       1,
	// 		testTarget: func(u model.User, c caseType) bool {
	// 			return u.Password == c.Password
	// 		},
	// 	}, {
	// 		FullName: "0987654321Fullname",
	// 		ID:       1,
	// 		testTarget: func(u model.User, c caseType) bool {
	// 			return u.FullName == c.FullName
	// 		},
	// 	},
	// 	}

	// 	for _, item := range cases {
	// 		err = userDB.UpdatePasswordFullName(context.TODO(), item.ID, item.Password, item.FullName)
	// 		assert.NoError(t, err, "unable to update user")
	// 		expectedUser, err := userDB.GetByID(context.TODO(), item.ID)
	// 		assert.NoError(t, err, "unable to read user")
	// 		assert.True(t, item.testTarget(expectedUser, item), "parameters do not match")
	// 	}
	// }

	{
		_, err = blogDB.GetByID(context.TODO(), 5)
		assert.Error(t, err)
	}

	{
		err := blogDB.DeleteByID(context.TODO(), blog.ID)
		assert.NoError(t, err, err)
		_, err = blogDB.GetByID(context.TODO(), blog.ID)
		assert.Error(t, err)
	}

}
