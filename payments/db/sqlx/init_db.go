package sqlx

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB;

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
    connection_string := fmt.Sprintf("postgres://%s:%s@localhost:5432/bank?sslmode=disable", postgres_username, postgres_password)
        db, err := sqlx.Open("postgres", connection_string)
        if err != nil {
            log.Fatalln(err)
        }        
		Db = db;
}