package config

import (
	"github.com/joho/godotenv"
	"strconv"
)

func Load(path string) (*Config, error) {
	env, err := godotenv.Read(path)
	if err != nil {
		return nil, err
	}

	return ParseConfigMap(env), nil
}

func ParseConfigMap(env map[string]string) *Config {
	rate, _ := strconv.Atoi(env["RATE_LIMIT"])

	return &Config{
		env["ROUTER_PORT"],
		rate,
		env["AUTH_KEY"],
	}
}

type Config struct {

	RouterPort string
	RateLimit int
	AuthKey string

}