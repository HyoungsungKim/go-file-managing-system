package DBHandler

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ConnectDB() *sql.DB {
	loadEnv()
	var (
		host     = os.Getenv("DB_ADDRESS")
		port, _  = strconv.Atoi(os.Getenv("DB_PORT"))
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Print(psqlconn)

	db, err := sql.Open("postgres", psqlconn)
	checkError(err, "Open db err")
	//defer db.Close()

	err = db.Ping()
	checkError(err, "db ping err")

	fmt.Println("Connected!")
	fmt.Println("Checking tables...")

	resetDB(db)
	generateTables(db)

	return db
}

func loadEnv() {
	err := godotenv.Load("/app/.postgres_env")
	checkError(err, "load env err")
}
