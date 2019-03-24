package data

import "database/sql"

//Repo contains db object and it implements every handlers
type Repo struct {
	DB *sql.DB
}

//NewRepo will return repo object
func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		DB: db,
	}
}
