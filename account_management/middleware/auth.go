package middleware

import (
	"errors"

	mylog "github.com/fox-wei/go_project_training/account_management/MyLog"
	"github.com/fox-wei/go_project_training/account_management/myerr"
	"github.com/fox-wei/go_project_training/account_management/res"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//!jwt token verify
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")

		mylog.Log.Info("Authorization:" + auth)
		if len(auth) == 0 {
			res.SendResponse(ctx, myerr.ErrValidation, "must login")
			ctx.Abort()
			return
		}

		//*verify token
		_, err := ParseToken(auth)
		if err != nil {
			res.SendResponse(ctx, myerr.ErrToken, "token parse failed")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

//*parse jwt token
func parseToken2(token string) (*MyCustomclaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &MyCustomclaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(*MyCustomclaims); ok && jwtToken.Valid {
		return claims, nil
	}
	// if err != nil && jwtToken != nil {
	// 	if claim, ok := jwtToken.Claims.(*MyCustomclaims); ok && jwtToken.Valid {
	// 		return claim, nil
	// 	}
	// }

	return nil, errors.New("invalid token")
}

type MyCustomclaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func ParseToken(token string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
