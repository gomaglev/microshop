package util

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads all env variables from .env path provided, relative to the
// source of execution. If no .env path is provided, .env will be loaded from
// directory of source of execution.
func LoadEnv(file ...string) error {
	return godotenv.Load(file...)
}

// Debug returns boolean indicating if service is in debug mode by specifying
// `env` or by log level set to trace or debug.
func Debug() bool {
	var debug = false
	env := os.Getenv("ENV")
	level := os.Getenv("LOG_LEVEL")

	if env == "debug" || level == "trace" || level == "debug" {
		debug = true
	}

	return debug
}
