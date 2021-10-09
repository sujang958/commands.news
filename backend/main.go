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
	// res, err := http.Get("https://news.google.com/topics/CAAqJggKIiBDQkFTRWdvSUwyMHZNRGRqTVhZU0FtVnVHZ0pWVXlnQVAB?hl=en-US&gl=US&ceid=US%3Aen")
	// if err != nil {
	// 	log.Fatalln("Failure getting news")
	// }
	// io.ReadAll() res.Body.Read()

	mux := http.NewServeMux()
	log.Fatalln(http.ListenAndServe(":3000", mux))
}