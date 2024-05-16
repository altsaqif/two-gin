package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host       string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	Driver     string
}

type ApiConfig struct {
	ApiPort string
}

type Config struct {
	DbConfig
	ApiConfig
}

func (c *Config) Configuration() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("missing env file %v", err.Error())
	}
	c.DbConfig = DbConfig{
		Host:       os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		Driver:     os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	if c.Host == "" || c.DbPort == "" || c.DbUser == "" || c.DbPassword == "" || c.DbName == "" || c.Driver == "" {
		return fmt.Errorf("missing environment")
	}
	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}

	if err := config.Configuration(); err != nil {
		return nil, err
	}
	return config, nil
}
