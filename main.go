package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("couldnt import .env file")
	}

	mux := http.NewServeMux()
	log.Fatalln(http.ListenAndServe(":3000", mux))
}