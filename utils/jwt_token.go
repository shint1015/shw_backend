package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"time"
)

type JwtCustomClaims struct {
	id uint
	jwt.RegisteredClaims
}

var signingKey = []byte(GetJWTSecret())

var JwtConfig = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(JwtCustomClaims)
	},
	SigningKey: signingKey,
}

func GenerateToken(id uint) (string, error) {
	claims := JwtCustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(getJWTExpire()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return t, nil
}

func getJWTExpire() time.Time {
	return time.Now().Add(time.Hour * 24)
}

func GetClaims(c *echo.Context) *JwtCustomClaims {
	user := (*c).Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}

func CheckToken(tokenStr string) error {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			return fmt.Errorf("token is expired")
		}
		return nil
	}
	return fmt.Errorf("token is invalid")
}
