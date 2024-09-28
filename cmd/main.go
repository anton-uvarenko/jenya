package main

import (
	"jenya/internal/domain_name/handlers"
	"jenya/internal/pkg/server"
	"log"
)

func main() {
	handler := handlers.NewHandler("")

	httpServer := server.NewServer(handler.RegisterRouts())
	log.Fatal(httpServer.ListenAndServe())
}
