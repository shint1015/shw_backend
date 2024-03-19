package config

import (
	"github.com/joho/godotenv"
	"shwgrpc/utils"
)

func LoadEnv() *utils.CustomError {
	err := godotenv.Load("../../config/.env")

	if err != nil {
		return utils.Error(err, 2)
	}
	return nil
}
