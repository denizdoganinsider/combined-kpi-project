package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionPool(ctx context.Context, config Config) *sql.DB {

	fmt.Println("config: ", config)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.UserName, config.Password, config.Host, config.Port, config.DbName)

	log.Println("dsn: ", dsn)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.SetMaxOpenConns(config.MaxConnections)
	db.SetConnMaxIdleTime(time.Duration(config.MaxConnectionIdleTime) * time.Second)

	if err := db.Ping(); err != nil {
		log.Println("log test")
		log.Fatalf("Database connection failed: %v", err)
	}

	fmt.Println("MySQL connection established successfully!")

	return db
}
