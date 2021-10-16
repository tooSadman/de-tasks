package main

import (
	"log"

	"github.com/tooSadman/de-tasks/replicated-log/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080", "master")
	log.Fatal(srv.ListenAndServe())
}
