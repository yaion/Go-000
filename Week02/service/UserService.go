package service

import (
	"Go-000/Week02/Dao"
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
)

func GetUser() string {
	name := "xxxx"
	user, err := Dao.GetUserByName(name)
	if err != nil {
		if errors.Is(errors.Cause(err), sql.ErrNoRows) {
			log.Fatal(err)
		} else {
			log.Fatal(err)
		}
		return ""
	}
	str, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(str)
}
