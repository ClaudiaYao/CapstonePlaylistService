package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var counts int64

func main() {
	log.Println("Starting playlist service")

	// connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	//set up
	app := PlaylistService{
		DBConnection: conn,
	}

	host := goDotEnvVariable("SERVICE_HOST")

	srv := &http.Server{
		Addr:    host,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

// C: this function will connect to database and then return *DataQuery
func openDB(dsn string) (*DataQuery, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	dataquery := DataQuery{db: db}

	return &dataquery, nil
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {
	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// load .env file which is located at the root path
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// C: wrap the openDB function and provide retry mechanism
func connectToDB() *DataQuery {

	// host := goDotEnvVariable("DB_HOST")
	// port := goDotEnvVariable("DB_PORT")
	// user := goDotEnvVariable("DB_USER")
	// password := goDotEnvVariable("DB_PASS")
	// dbname := goDotEnvVariable("DB_NAME")

	// dsn := fmt.Sprintf("host=%s port=%s user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	dsn := os.Getenv("DSN")

	println("debugging line:", dsn)

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}
