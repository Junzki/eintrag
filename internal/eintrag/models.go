package eintrag

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	Uid         *uuid.UUID
	Username    *string
	Password    *string
	DisplayName *string
}

func CreateUser(u *User) error {

	results, err := DbConn.Exec(
		"SELECT uid FROM create_user($1::varchar, $2::varchar, $3::varchar);", u.Username, u.Password, u.DisplayName,
	)

	if nil != err {
		LOG.Errorf("Error creating user with name %s, %s", *u.Username, err)
		return fmt.Errorf("error creating user")
	}

	fmt.Println(results)

	return nil
}
