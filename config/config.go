package config

import (
	"errors"
	"os"
)

const PROTOCOL_TCP = "tcp"
const PORT_8080 = "8080"

var DB_HOST string
var DB_PORT string
var DB_USER string
var DB_PASSWORD string
var DB_DATABASE string

func Get() error {

	DB_HOST = os.Getenv("HOMEWORK_DB_HOST")
	DB_PORT = os.Getenv("HOMEWORK_DB_PORT")
	DB_USER = os.Getenv("HOMEWORK_DB_USER")
	DB_PASSWORD = os.Getenv("HOMEWORK_DB_PASSWORD")
	DB_DATABASE = os.Getenv("HOMEWORK_DB_DATABASE")

	if DB_HOST == "" {
		return errors.New("env varialble HOMEWORK_DB_HOST has not filled")
	}

	if DB_PORT == "" {
		return errors.New("env varialble HOMEWORK_DB_PORT has not filled")
	}

	if DB_USER == "" {
		return errors.New("env varialble HOMEWORK_DB_USER has not filled")
	}

	if DB_PASSWORD == "" {
		return errors.New("env varialble HOMEWORK_DB_PASSWORD has not filled")
	}

	if DB_DATABASE == "" {
		return errors.New("env varialble HOMEWORK_DB_NAME has not filled")
	}

	return nil

}
