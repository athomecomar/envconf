package coreconf

import "os"

const (
	Production  env = "production"
	Development env = "development"
)

type env string

func GetENV() env {
	osEnv := os.Getenv("ENV")
	if osEnv != "" {
		return env(osEnv)
	}
	return "development"
}
