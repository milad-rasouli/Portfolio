package sqlitedb

import (
	"context"
	"testing"

	"github.com/Milad75Rasouli/portfolio/internal/db"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestHomeCRUD(t *testing.T) {
	var (
		homeDB *StoreSqlite
		err    error
	)

	home := model.Home{
		Name:       "foo bar",
		Slogan:     "this is a very long text.",
		ShortIntro: "this is a short intro to me",
		GithubUrl:  "http://github.com/",
	}

	d := SqliteInit{Folder: "data"}
	homeDB, cancel, err := d.Init(true, db.Config{}, nil)
	assert.NoError(t, err)
	defer cancel()

	{
		err = homeDB.UpdateHome(context.TODO(), home)
		assert.NoError(t, err)
	}

	{
		fetchedHome, err := homeDB.GetHome(context.TODO())
		assert.NoError(t, err)
		assert.Equal(t, home, fetchedHome, "home and fetchHome should be equal")
	}

	{
		home2 := model.Home{
			Name:       "bar baz",
			Slogan:     "this is a very long text 1. this is a very long text 2. this is a very long text 3.",
			ShortIntro: "short",
		}
		err = homeDB.UpdateHome(context.TODO(), home2)
		assert.NoError(t, err)

		fetchedHome, err := homeDB.GetHome(context.TODO())
		assert.NoError(t, err)
		assert.NotEqual(t, fetchedHome, home, " home and fetchHome2 shouldn't br equal")

	}
}
