package config

import (
	"github.com/joho/godotenv"
	"shw/utils"
)

func LoadEnv() *utils.CustomError {
	err := godotenv.Load()

	if err != nil {
		return utils.Error(err, 2)
	}
	return nil
}
