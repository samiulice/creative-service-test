package dbrepo

import (
	"database/sql"
	"creative-service-test/internal/repository"
)

type postgresDBRepo struct {
	DB  *sql.DB
}

func NewDBRepo(conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB:  conn,
	}
}