package api

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"shw/model"
	"shw/utils"
	"time"
)

func MethodOverride(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		if req.Method == "POST" {
			method := req.PostFormValue("_method")
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				req.Method = method
			}
		}
		return next(c)
	}
}

func SignUp(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	u := model.User{Email: email, Password: password}
	if err := u.Create(nil); err != nil {
		return echo.ErrUnauthorized
	}
	t, cErr := GenerateJWT(u.ID)
	if cErr != nil {
		return echo.ErrUnauthorized
	}
	return c.JSON(200, map[string]string{
		"token": t,
	})
}

func SignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	u := model.User{Email: email, Password: password}
	_, err := u.Get()
	if err != nil {
		return echo.ErrUnauthorized
	}
	t, cErr := GenerateJWT(u.ID)
	if cErr != nil {
		return echo.ErrUnauthorized
	}
	domain, _ := os.LookupEnv("DOMAIN_NAME")
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   t,
		Domain:  domain,
		Expires: time.Now().Add(24 * time.Hour),
	})
	return c.JSON(200, map[string]string{
		"message": "success",
	})
}

func SignOut(c echo.Context) error {
	return nil
}
func Restricted(c echo.Context) error {
	return nil
}

func CreateUser(username, email, password string) *utils.CustomError {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 15)
	password = string(hashed)

	user := model.User{Email: email, Name: username, Password: password}
	if err := user.Create(nil); err != nil {
		return utils.Error(err, 2)
	}
	return nil
}

func GenerateJWT(userId uint) (string, *utils.CustomError) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userId
	claims["iat"] = time.Now().Unix()
	claims["exp"] = utils.GetExpireTime().Unix()
	t, err := token.SignedString([]byte(utils.GetJWTSecret()))
	if err != nil {
		return "", utils.Error(err, 2)
	}
	return t, nil
}
