package eintrag

import (
	"eintrag/internal/eintrag"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginRequest struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := eintrag.Authenticate(req.UserName, req.Password)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exp := time.Now().Add(24 * time.Hour)
	ss, err := eintrag.GenerateTokenWithUser(user, exp)

	c.JSON(http.StatusOK, gin.H{
		"access_token": ss,
		"error":        err,
	})
	return
}
