package sqlitedb

import (
	"context"
	"testing"

	"github.com/Milad75Rasouli/portfolio/internal/db"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestAboutMeCRUD(t *testing.T) {
	var (
		aboutMeDB *StoreSqlite
		err       error
	)

	aboutMe := model.AboutMe{
		Content: "this is a very long text.",
	}

	d := SqliteInit{Folder: "data"}
	aboutMeDB, cancel, err := d.Init(true, db.Config{}, nil)
	assert.NoError(t, err)
	defer cancel()

	{
		err = aboutMeDB.UpdateAboutMe(context.TODO(), aboutMe)
		assert.NoError(t, err)
	}

	{
		fetchedAboutMe, err := aboutMeDB.GetAboutMe(context.TODO())
		assert.NoError(t, err)
		assert.Equal(t, aboutMe, fetchedAboutMe, "aboutMe and fetchAboutMe should be equal")
	}

	{
		aboutMe2 := model.AboutMe{
			Content: "this is a very long text 1. this is a very long text 2. this is a very long text 3.",
		}
		err = aboutMeDB.UpdateAboutMe(context.TODO(), aboutMe2)
		assert.NoError(t, err)

		fetchedAboutMe, err := aboutMeDB.GetAboutMe(context.TODO())
		assert.NoError(t, err)
		assert.NotEqual(t, fetchedAboutMe, aboutMe, " aboutMe and fetchAboutMe2 shouldn't br equal")

	}
}
