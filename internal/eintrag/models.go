package eintrag

import "github.com/google/uuid"

type User struct {
	Uid         *uuid.UUID
	Username    *string
	Password    *string
	DisplayName *string
}
