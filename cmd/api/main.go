package main

import (
	"log"

	"github.com/guruorgoru/learning-distributed-system/internal/server"
)

func main() {
	log.Println("Server started ta port: 8414")
	srv := server.NewHttpServer("", "8414")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
