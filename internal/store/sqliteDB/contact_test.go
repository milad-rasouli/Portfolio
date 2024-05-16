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

func TestContactCRUD(t *testing.T) {
	var contactDB *StoreSqlite
	ti, err := time.Parse(timeLayout, time.Now().Format(timeLayout))
	contact := model.Contact{
		ID:        1,
		Email:     "bar@baz.com",
		Message:   "foo bar baz boo faa bar baz",
		CreatedAt: ti,
	}

	d := SqliteInit{Folder: "data"}
	contactDB, cancel, err := d.Init(true, db.Config{}, nil)
	assert.NoError(t, err)
	defer cancel()

	{
		var id int64
		id, err = contactDB.CreateContact(context.TODO(), contact)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("CreateContact:contact created id is %d", id)
	}

	{
		fetchedContact, err := contactDB.GetContactByID(context.TODO(), contact.ID)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, contact, fetchedContact, "contact and fetchContact should be equal ")
		fmt.Printf("GetContactByID:gotten user is %+v", fetchedContact)

	}

	{
		_, err = contactDB.GetUserByID(context.TODO(), 5)
		assert.Error(t, err)
	}

	{
		err := contactDB.DeleteBlogByID(context.TODO(), contact.ID)
		assert.NoError(t, err, err)
		_, err = contactDB.GetUserByID(context.TODO(), contact.ID)
		assert.Error(t, err)
	}

}
