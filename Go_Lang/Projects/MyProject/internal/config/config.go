package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser      string	`env:"DB_USER,required"`
	DBPassword  string	`env:"DB_PASSWORD,required"`
	DBHost      string	`env:"DB_HOST,required"`
	DBPort      int		`env:"DB_PORT" envDefault:"3306"`
	DBName      string	`env:"DB_NAME,required"`
	ServerPort  int		`env:"SERVER_PORT" envDefault:"8081"`
	KafkaBroker string	`env:"KAFKA_BROKER" envDefault:"localhost:9092"`
}

func Load() (*Config, error) {
  err := godotenv.Load()
  if err != nil {
    log.Println(".env file not found in Go module")
  }

  var appConfig Config
  err = env.Parse(&appConfig)
  if err != nil {
    return nil, fmt.Errorf(".env parsing failed: %w", err)
  }
  log.Printf("Configuration loaded successfully...")
  return &appConfig, nil
}
