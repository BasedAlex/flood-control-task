package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnect string
	Threshold int
	TimeLimit int
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	threshold, err := strconv.Atoi(os.Getenv("THRESHOLD"))
	if err != nil {
		return nil, err
	}
	timeLimit, err := strconv.Atoi(os.Getenv("TIME_LIMIT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		DBConnect: os.Getenv("DB_CONNECT"), 
		Threshold: threshold,
		TimeLimit: timeLimit,
		}, nil
}