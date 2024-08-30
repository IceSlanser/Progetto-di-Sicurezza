package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
)

func (db *appdbimpl) RegisterUser(name string, password string) (uint64, error) {
	query := fmt.Sprintf("INSERT INTO profiles (Username, Password) VALUES ('%s', '%s')", name, password)
	res, err := db.c.Exec(query)
	if err != nil {
		if errors.Is(err, sqlite3.ErrConstraintUnique) {
			return 0, err
		}
		return 0, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ID), nil
}

func (db *appdbimpl) LoginUser(name string, password string) (User, error) {
	var user User
	query := fmt.Sprintf("SELECT * FROM profiles WHERE Username = '%s' AND Password = '%s'", name, password)
	err := db.c.QueryRow(query).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, err
		}
		return user, err
	}
	return user, nil
}
