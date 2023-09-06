package commons

import (
	"github.com/joho/godotenv"
	"github.com/jutionck/golang-todo-apps/utils/exception"
	"os"
)

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.CheckError(err)
	return &configImpl{}
}
