package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	PanicIfError(err)

	return os.Getenv(key)
}
