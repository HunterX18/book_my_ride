package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)


var Db *sql.DB;


func loadEnv() {
    err := godotenv.Load(".env");
    if err != nil {
        log.Fatalln("Error loading env file")
    }
}

func InitDB() {
	loadEnv();
    postgres_username := os.Getenv("POSTGRES_USERNAME")
    postgres_password := os.Getenv("POSTGRES_PASSWORD")
    connection_string := fmt.Sprintf("postgres://%s:%s@localhost:5432/garage?sslmode=disable", postgres_username, postgres_password)
        db, err := sql.Open("postgres", connection_string)
        if err != nil {
            log.Fatalln(err.Error())
        }        
		Db = db;
}
func main() {
	InitDB();
    for {
        _, err := Db.Exec("UPDATE rides SET status = 'free' WHERE status = 'reserved' AND status_changed_at < NOW() - INTERVAL '5 seconds'")
        if err != nil {
            fmt.Println("Error updating records:", err)
        }

        time.Sleep(5 * time.Second)
    }
}

