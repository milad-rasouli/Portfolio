package sqlitedb

import (
	"context"
	"fmt"
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
		fetchedBlog, err := blogDB.GetBlogByID(context.TODO(), blog.ID)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, blog, fetchedBlog, "blog and fetchUser should be equal ")
		fmt.Printf("GetBlogByID: %+v\n", fetchedBlog)
	}

	{
		UpdateTest := func(base model.Blog, expected model.Blog) (result bool) {
			result = (base.Body == expected.Body) && (base.Title == expected.Title) &&
				(base.Caption == expected.Caption) && (base.ImagePath == expected.ImagePath)
			return result
		}
		updatedBlog := model.Blog{
			ID:        1,
			Title:     "bar barrrr",
			Body:      "bazzz bazzz",
			Caption:   "fooo fooo",
			ImagePath: "foo/bbbbarrr",
		}
		err = blogDB.UpdateBlogByID(context.TODO(), updatedBlog)
		assert.NoError(t, err, "unable to update blog")
		expectedBlog, err := blogDB.GetBlogByID(context.TODO(), updatedBlog.ID)
		assert.NoError(t, err, "unable to read user")
		fmt.Printf("UpdateBlogByID: %+v GetBlogByID: %+v\n", updatedBlog, expectedBlog)
		assert.True(t, UpdateTest(updatedBlog, expectedBlog), "parameters do not match")

	}

	{
		_, err = blogDB.GetBlogByID(context.TODO(), 5)
		assert.Error(t, err)

	}

	{
		err := blogDB.DeleteBlogByID(context.TODO(), blog.ID)
		assert.NoError(t, err, err)
		_, err = blogDB.GetBlogByID(context.TODO(), blog.ID)
		assert.Error(t, err)
	}

}
