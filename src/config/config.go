package config

import (
	"os"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	SmtpHost   string
	SmtpUser   string
	SmtpPass   string
	EmailFrom  string
}

func LoadConfig() (*Config, error) {
	return &Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		SmtpHost:   os.Getenv("SMTP_HOST"),
		SmtpUser:   os.Getenv("SMTP_USER"),
		SmtpPass:   os.Getenv("SMTP_PASS"),
		EmailFrom:  os.Getenv("SMTP_USER"),
	}, nil
}
