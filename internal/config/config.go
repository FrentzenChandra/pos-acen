package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort      string
	LogLevel     string
	LogAddSource bool
	DBHost       string
	DBPort       int
	DBUser       string
	DBPassword   string
	DBName       string
	DBDebug      bool
	BaseURLPath  string
	DBSSLMode    string
}

func LoadConfig() (*Config, error) {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error Loading .env File")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Println("Error Cannot Convert String DB_Port To Integer ")
		return nil, err
	}

	dbDebug, err := strconv.ParseBool(os.Getenv("DB_DEBUG"))

	if err != nil {
		log.Println("Error Cannot Convert String Db_Debug To Boolean ")
		return nil, err
	}

	config := &Config{
		AppPort:     os.Getenv("APP_PORT"),
		DBPort:      dbPort,
		DBHost:      os.Getenv("DB_HOST"),
		DBName:      os.Getenv("DB_NAME"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBSSLMode:   os.Getenv("DB_SSL_MODE"),
		DBDebug:     dbDebug,
		DBUser:      os.Getenv("DB_USERNAME"),
		BaseURLPath: os.Getenv("BASE_URL_PATH"),
	}

	return config, nil
}

func WriteTimeout() time.Duration {
	return 10 * time.Second
}

func ReadTimeout() time.Duration {
	return 10 * time.Second
}
