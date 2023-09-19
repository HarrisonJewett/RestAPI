package main

import (
	"os"

	"github.com/REST-API-Test/usecase"

	"github.com/REST-API-Test/db/psql"
	http "github.com/REST-API-Test/server"
)

func main() {
	os.Setenv("PGSSLMODE", "disable")
	os.Setenv("PGUSER", "postgres")
	db := psql.NewPsql()
	u := usecase.NewUsecase(db)
	s := http.NewHttpServer(u)
	s.Start()
}
