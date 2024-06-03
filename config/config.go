package config

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}

	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) GetStartTime() (time.Time, error) {
	return time.Parse("15:04", c.StartTime)
}

func (c *Config) GetEndTime() (time.Time, error) {
	return time.Parse("15:04", c.EndTime)
}
