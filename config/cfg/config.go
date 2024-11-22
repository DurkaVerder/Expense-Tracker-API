package cfg

import (
	"Expense-Tracker-API/internal/model"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*model.Config, error) {
	cfg := model.Config{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error open config")
	}
	defer file.Close()

	decode := yaml.NewDecoder(file)

	if err := decode.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
