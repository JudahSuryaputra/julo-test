package main

import (
	"julo-case-study/cfg"
	"julo-case-study/http"
	"julo-case-study/repositories"
	"log"

	_ "github.com/lib/pq"
)

func init() {
	cfg.Init()
}

func main() {
	app := http.App{}

	conn, err := repositories.Conn()
	if err != nil {
		log.Fatalf("Cannot initialize connection to database: %v", err)
	}

	app.Initialize(conn)
	app.RunServer()
}
