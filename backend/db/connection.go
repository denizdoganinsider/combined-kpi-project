package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error occured. While loading .env file")
	}

	config := DatabaseConfig{
		Username:     os.Getenv("DB_username"),
		Password:     os.Getenv("DB_password"),
		Protocol:     os.Getenv("DB_protocol"),
		Host:         os.Getenv("DB_host"),
		Port:         os.Getenv("DB_port"),
		DatabaseName: os.Getenv("DB_database_name"),
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s", config.Username, config.Password, config.Protocol, config.Host, config.Port, config.DatabaseName))

	if err != nil {
		panic(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
	fmt.Println("Database connection is successfully completed")

	return db
}
