package controllers

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lihulab/xiaolongbaoproxy-mgt/models"
)

func Login(c *gin.Context) {
	var u models.UserInfo
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json request")
		return
	}

	if u.Username != "lyy" || u.Password != "1234" {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}

	token, err := genToken(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "unknown error")
		return
	}

	c.JSON(http.StatusOK, token)
}

func Logout(c *gin.Context) {

}

type CustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func genToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer:    "xiaolongbao",
		},
	})

	return token.SignedString("sss")
}
