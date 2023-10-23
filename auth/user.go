package auth

import (
	"auth-server/config"
	"auth-server/database"
)

func CreateUser(c config.Config, username string, password string) error {
	db, err := database.ConnectToSqlLite(c.DatabasePath)

	_, err = db.Exec(database.CreateUser, username, hashPassword(password))
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetUser(c config.Config, username string) (database.User, error) {
	db, err := database.ConnectToSqlLite(c.DatabasePath)

	var user database.User
	err = db.QueryRow(database.GetUser, username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return database.User{}, err
	}

	err = db.Close()
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}

func DeleteUser(c config.Config, id int) error {
	db, err := database.ConnectToSqlLite(c.DatabasePath)

	_, err = db.Exec(database.DeleteUser, id)
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}
