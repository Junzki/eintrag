package eintrag

import (
	"eintrag/internal/eintrag"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CreateUserRequest struct {
	UserName    string `form:"username" json:"username" binding:"required" validate:"email"`
	Password    string `form:"password" json:"password" binding:"required" validate:"required"`
	DisplayName string `form:"display_name" json:"display_name" validate:"required,lte=32"`
}

var validate *validator.Validate

func (r CreateUserRequest) Clean() *eintrag.User {

	if "" == r.DisplayName {
		r.DisplayName = r.UserName
	}

	return &eintrag.User{
		Username:    &r.UserName,
		Password:    &r.Password,
		DisplayName: &r.DisplayName,
	}
}

func CreateUser(c *gin.Context) {
	var req eintrag.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := eintrag.CreateUser(&req)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": nil})
	return
}
