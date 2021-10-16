package main

import (
	"log"

	"github.com/tooSadman/de-tasks/replicated-log/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":9001", "slave")
	log.Fatal(srv.ListenAndServe())
}
