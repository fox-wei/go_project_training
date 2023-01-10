package token

import (
	"time"

	mylog "github.com/fox-wei/go_project_training/account_management/MyLog"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type Context struct {
	ID       string
	Username string
}

func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	//*没有指定加载秘钥则读取配置文件
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	//*jwt-token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})

	//sign token with the secret
	mylog.Log.Info("secret:" + secret)
	tokenString, err = token.SignedString([]byte(secret))

	return
}
