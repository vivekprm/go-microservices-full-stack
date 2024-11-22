package config

import (
	"fmt"
	"os"
)

const (
	postgres = "postgres"
)

type Config struct {
	DriverName string
	Port       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPass     string
	DbName     string
	DbSslMode  string
}

func GetConfig() *Config {
	return &Config{
		DriverName: os.Getenv("DRIVER_NAME"),
		Port:       os.Getenv("PORT"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPass:     os.Getenv("DB_PASS"),
		DbName:     os.Getenv("DB_NAME"),
		DbSslMode:  os.Getenv("DB_SSL_MODE"),
	}
}

func (c *Config) GetDbConnectionString() (string, error) {
	if c.DriverName == postgres {
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			c.DbHost, c.DbPort, c.DbUser, c.DbPass, c.DbName, c.DbSslMode), nil
	}
	return "", fmt.Errorf("driver not supported")
}
