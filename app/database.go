package app

import (
	"database/sql"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/"+os.Getenv("DB_NAME")+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	//db.SetMaxIdleConns(5)
	//db.SetMaxOpenConns(20)
	//db.SetConnMaxLifetime(60 * time.Minute)
	//db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
