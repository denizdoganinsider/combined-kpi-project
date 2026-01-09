package app

import (
	"log"
	"os"
	"strconv"

	"myapp-backend/common/mysql"

	"github.com/joho/godotenv"
)

type ConfigurationManager struct {
	MySqlConfig mysql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	MySqlConfig := getMySqlConfig()
	return &ConfigurationManager{
		MySqlConfig: MySqlConfig,
	}
}

func getMySqlConfig() mysql.Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf(".env file wasn't loaded: %v", err)
	}

	dbUser := os.Getenv("DB_username")
	dbPassword := os.Getenv("DB_password")
	dbHost := os.Getenv("DB_host")
	dbPort := os.Getenv("DB_port")
	dbName := os.Getenv("DB_database_name")

	dbMaxConnections, err := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	if err != nil {
		dbMaxConnections = 10
	}

	dbMaxConnectionIdleTime, err := strconv.Atoi(os.Getenv("MAX_CONNECTION_IDLE_TIME"))
	if err != nil {
		dbMaxConnectionIdleTime = 300
	}

	return mysql.Config{
		Host:                  dbHost,
		Port:                  dbPort,
		UserName:              dbUser,
		Password:              dbPassword,
		DbName:                dbName,
		MaxConnections:        dbMaxConnections,
		MaxConnectionIdleTime: dbMaxConnectionIdleTime,
	}
}
