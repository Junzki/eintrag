package eintrag

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DbConn *sql.DB

func InitDbConn(c string) {
	if conn, err := sql.Open("postgres", c); nil == err {
		DbConn = conn
	} else {
		panic(err)
	}
}

func Authenticate(username string, password string) (*User, error) {
	stmt, _ := DbConn.Prepare(
		"SELECT uid, found_username, display_name FROM authorize_user_with_password($1::varchar, $2::varchar);",
	)

	rows, err := stmt.Query(username, password)
	if nil != err {
		LOG.Errorf("Error authenticating user with name %s, %s", username, err)
		return nil, fmt.Errorf("user not found or password incorrect")
	}

	defer rows.Close()

	user := User{}
	for rows.Next() {
		err := rows.Scan(&user.Uid, &user.Username, &user.DisplayName)
		if nil != err {
			LOG.Errorf("Error scanning user model: %s", err)
			return nil, fmt.Errorf("user not found or password incorrect")
		}
	}

	if nil == user.Uid {
		return nil, fmt.Errorf("user not found or password incorrect")
	}

	LOG.Infof("User %s authenticated", *user.Username)

	return &user, nil
}
