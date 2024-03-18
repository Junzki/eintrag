package eintrag

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	req := CreateUserRequest{
		UserName:    "username@domain.com",
		Password:    "password",
		DisplayName: "display name11111111111111111111111111111111",
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(req)
	fmt.Println(err)

	assert.NoError(t, err, "expected no error")

}
