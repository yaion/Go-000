package Dao

import (
	"github.com/pkg/errors"
)

type User struct {
	Id     int    `db:id`
	Name   string `db:name`
	Adress string `db:name`
}

func GetUserByName(name string) (*User, error) {
	var user *User
	stmt, err := db.Prepare("select * from User where name=?")
	if err != nil {

	}
	err = stmt.QueryRow(name).Scan(&user.Id, &user.Name, &user.Adress)
	if err != nil {
		return nil, errors.Wrap(err, "no row")
	}
	return user, nil
}
