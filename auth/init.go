package auth

import (
	"auth-server/config"
	"auth-server/database"
)

func Init(c config.Config) error {
	err := createTables(c)
	if err != nil {
		return err
	}

	return nil
}

func createTables(c config.Config) error {
	db, err := database.ConnectToSqlLite(c.DatabasePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(database.UserTable)
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}
