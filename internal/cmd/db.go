package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/logger"
)

const (
	//DBHOST will be stored from corresponding environment
	DBHOST = "DB_HOST"
	//DBPORT will be stored from corresponding environment
	DBPORT = "DB_PORT"
	//DBUSER will be stored from corresponding environment
	DBUSER = "DB_USER"
	//DBPASSWORD will be stored from corresponding environment
	DBPASSWORD = "DB_PASSWORD"
	//DBNAME will be stored from corresponding environment
	DBNAME = "DB_NAME"
)

func prepareDB() *sql.DB {
	env := getEnv()
	psql := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		env[DBHOST],
		env[DBPORT],
		env[DBUSER],
		env[DBPASSWORD],
		env[DBNAME],
	)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	return db
}

func getEnv() map[string]string {
	env := make(map[string]string)
	host, ok := os.LookupEnv(DBHOST)
	if !ok {
		logger.Error("HOST environment variable required but not set")
		panic("Program exit with status 1")
	}
	port, ok := os.LookupEnv(DBPORT)
	if !ok {
		logger.Error("PORT environment variable required but not set")
		panic("Program exit with status 1")
	}
	user, ok := os.LookupEnv(DBUSER)
	if !ok {
		logger.Error("USER environment variable required but not set")
		panic("Program exit with status 1")
	}
	password, ok := os.LookupEnv(DBPASSWORD)
	if !ok {
		logger.Error("PASSWORD environment variable required but not set")
		panic("Program exit with status 1")
	}
	database, ok := os.LookupEnv(DBNAME)
	if !ok {
		logger.Error("DATABASE environment variable required but not set")
		panic("Program exit with status 1")
	}
	env[DBHOST] = host
	env[DBPORT] = port
	env[DBUSER] = user
	env[DBPASSWORD] = password
	env[DBNAME] = database
	return env
}
