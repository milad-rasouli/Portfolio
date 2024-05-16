package sqlitedb

import (
	"context"
	"fmt"
	"time"

	"crawshaw.io/sqlite"
	"crawshaw.io/sqlite/sqlitex"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/Milad75Rasouli/portfolio/internal/store"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ContactSqlite struct {
	dbPool *sqlitex.Pool
	logger *zap.Logger
}

func NewContactSqlite(dbPool *sqlitex.Pool, logger *zap.Logger) *ContactSqlite {
	return &ContactSqlite{
		dbPool: dbPool,
		logger: logger,
	}
}
func (c ContactSqlite) parseToContact(stmt *sqlite.Stmt) (model.Contact, error) {
	var (
		contact model.Contact
		err     error
	)
	contact.ID = stmt.GetInt64("id")
	contact.Subject = stmt.GetText("subject")
	contact.Email = stmt.GetText("email")
	contact.Message = stmt.GetText("message")
	contact.CreatedAt, err = time.Parse(timeLayout, stmt.GetText("created_at"))
	return contact, err
}
func (c *ContactSqlite) CreateContact(ctx context.Context, contact model.Contact) (int64, error) {
	var rowID int64
	conn := c.dbPool.Get(ctx)
	defer c.dbPool.Put(conn)

	stmt, err := conn.Prepare(`INSERT INTO contact (subject, email, message, created_at)
	VALUES($1,$2,$3,$4);`)
	if err != nil {
		return rowID, errors.Errorf("unable to create the new contact message %s", err.Error())
	}
	defer stmt.Finalize()

	stmtSelect, err := conn.Prepare(`SELECT last_insert_rowid();`)
	if err != nil {
		return 0, errors.Errorf("unable to prepare the select statement: %s", err.Error())
	}
	defer stmtSelect.Finalize()

	stmt.SetText("$1", contact.Subject)
	stmt.SetText("$2", contact.Email)
	stmt.SetText("$3", contact.Message)
	stmt.SetText("$4", contact.CreatedAt.Format(timeLayout))

	fmt.Println("createContact: here")
	_, err = stmt.Step()
	if err != nil {
		return rowID, err
	}
	hasRow, err := stmtSelect.Step()
	if err != nil {
		return rowID, err
	}

	if hasRow {
		rowID = conn.LastInsertRowID()
	}
	return rowID, err
}
func (c *ContactSqlite) GetContactByID(ctx context.Context, id int64) (model.Contact, error) {
	var contact model.Contact
	conn := c.dbPool.Get(ctx)
	defer c.dbPool.Put(conn)
	stmt, err := conn.Prepare(`SELECT * FROM contact WHERE id=$1 LIMIT 1;`)
	if err != nil {
		return contact, errors.Errorf("unable to get the user %s from id", err.Error())
	}
	defer stmt.Finalize()
	stmt.SetInt64("$1", id)

	var hasRow bool
	hasRow, err = stmt.Step()
	if hasRow == false {
		return contact, store.UserNotFountError
	}
	if err != nil {
		return contact, err
	}
	contact, err = c.parseToContact(stmt)
	return contact, err
}
func (u *ContactSqlite) GetAllContact(ctx context.Context) ([]model.Contact, error) {
	var contact []model.Contact
	conn := u.dbPool.Get(ctx)
	defer u.dbPool.Put(conn)
	stmt, err := conn.Prepare(`SELECT * FROM contact;`)
	if err != nil {
		return contact, errors.Errorf("unable to get all users %s", err.Error())
	}
	defer stmt.Finalize()

	for {
		var (
			swapContact model.Contact
			hasRow      bool
		)
		hasRow, err = stmt.Step()
		if hasRow == false {
			break
		}
		swapContact, err = u.parseToContact(stmt)
		if err != nil {
			return contact, errors.Errorf("getting the users from database error %s", err.Error())
		}
		contact = append(contact, swapContact)
	}
	return contact, err
}
func (c *ContactSqlite) DeleteContactByID(ctx context.Context, id int64) error {
	conn := c.dbPool.Get(ctx)
	defer c.dbPool.Put(conn)
	stmt, err := conn.Prepare(`DELETE FROM contact WHERE id=$1;`)
	if err != nil {
		return err
	}
	defer stmt.Finalize()
	stmt.SetInt64("$1", id)
	_, err = stmt.Step()
	return err
}
