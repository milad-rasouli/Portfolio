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
	// 		err = userDB.UpdateUserByPasswordFullName(context.TODO(), item.ID, item.Password, item.FullName)
	// 		assert.NoError(t, err, "unable to update user")
	// 		expectedUser, err := userDB.GetUserByID(context.TODO(), item.ID)
	// 		assert.NoError(t, err, "unable to read user")
	// 		assert.True(t, item.testTarget(expectedUser, item), "parameters do not match")
	// 	}
	// }

	// {
	// 	_, err = userDB.GetUserByID(context.TODO(), 5)
	// 	assert.Error(t, err)
	// }

	// {
	// 	err := userDB.DeleteUserByID(context.TODO(), user.ID)
	// 	assert.NoError(t, err, err)
	// 	_, err = userDB.GetUserByID(context.TODO(), user.ID)
	// 	assert.Error(t, err)
	// }

}
