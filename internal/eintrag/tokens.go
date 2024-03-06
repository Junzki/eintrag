package eintrag

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateTokenWithUser(user *User, expiresAt time.Time) (string, error) {

	claims := jwt.RegisteredClaims{
		ID:        user.Uid.String(),
		Issuer:    "Die Eintrag",
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("signature-key"))
	if nil != err {
		return "failure", err
	}

	return ss, nil
}
