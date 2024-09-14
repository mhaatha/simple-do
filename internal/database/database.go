package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetConnection() *sql.DB {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	dbPassword := os.Getenv("DB_PASSWORD")

	// Open connection to database
	db, err := sql.Open("mysql", "root:"+dbPassword+"@tcp(localhost:3306)/simple_do?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// Setting db pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
