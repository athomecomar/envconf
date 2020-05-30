package envconf

import "os"

const (
	Production  env = "production"
	Staging     env = "staging"
	Development env = "development"
)

type env string

func InProduction() bool {
	return GetENV() == Production
}

func InDevelopment() bool {
	return GetENV() == Development
}

func InStaging() bool {
	return GetENV() == Staging
}

func NotInDevelopment() bool {
	return InProduction() || InStaging()
}

func GetENV() env {
	return env(Get("ENV", string(Development)))
}

func Get(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		v = def
	}
	return v
}
