package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInit, downInit)
}

var initDB = `
CREATE TABLE messages (
	id uuid PRIMARY KEY,
	message text,
	from uuid NOT NULL,
	to uuid NOT NULL
);
`

func upInit(tx *sql.Tx) error {
	_, err := tx.Exec(initDB)
	if err != nil {
		return err
	}
	return nil
}

func downInit(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
