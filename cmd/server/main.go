package main

import (
	"log"

	"github.com/sc7639/prolog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
