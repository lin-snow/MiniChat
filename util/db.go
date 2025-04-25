package util

import (
	"database/sql"
	"fmt"
	"log"
	"minichat/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.GlobalConfig.DBUser,
		config.GlobalConfig.DBPassword,
		config.GlobalConfig.DBHost,
		config.GlobalConfig.DBPort,
		config.GlobalConfig.DBName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	DB = db
}
