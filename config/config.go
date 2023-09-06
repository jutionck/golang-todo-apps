package config

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jutionck/golang-todo-apps/utils/commons"
	"strconv"
	"time"
)

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type FileConfig struct {
	FilePath string
	Env      string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     []byte
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type Config struct {
	ApiConfig
	DbConfig
	FileConfig
	TokenConfig
}

func (c *Config) ReadConfig() error {
	config := commons.New()
	c.DbConfig = DbConfig{
		Host:     config.Get("DB_HOST"),
		Port:     config.Get("DB_PORT"),
		Name:     config.Get("DB_NAME"),
		User:     config.Get("DB_USER"),
		Password: config.Get("DB_PASSWORD"),
	}

	c.ApiConfig = ApiConfig{
		ApiHost: config.Get("API_HOST"),
		ApiPort: config.Get("API_PORT"),
	}

	c.FileConfig = FileConfig{
		FilePath: config.Get("FILE_PATH"),
		Env:      config.Get("ENV"),
	}

	appTokenExpire, err := strconv.Atoi(config.Get("APP_TOKEN_EXPIRE"))
	if err != nil {
		return err
	}
	accessTokenLifeTime := time.Duration(appTokenExpire) * time.Hour

	c.TokenConfig = TokenConfig{
		ApplicationName:     config.Get("APP_TOKEN_NAME"),
		JwtSignatureKey:     []byte(config.Get("APP_TOKEN_KEY")),
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: accessTokenLifeTime,
	}

	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" ||
		c.DbConfig.User == "" || c.DbConfig.Password == "" || c.ApiConfig.ApiHost == "" ||
		c.ApiConfig.ApiPort == "" || c.FileConfig.FilePath == "" ||
		c.FileConfig.Env == "" {
		return fmt.Errorf("missing required environment variables")
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
