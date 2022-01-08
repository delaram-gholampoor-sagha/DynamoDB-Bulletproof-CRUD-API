package env

import "os"

func GetEnv(env string, defaultValue string) string {

	environment := os.Getenv(env)
	if environment == "" {
		return defaultValue
	}

	return environment

}


