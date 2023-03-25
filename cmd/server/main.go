package main

import (
	"log"

	"github.com/yuya-morimoto/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
