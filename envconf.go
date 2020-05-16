package envconf

import "os"

const (
	Production  env = "production"
	Staging     env = "staging"
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
