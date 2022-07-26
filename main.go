package main

import (
	"casetask/hehe/src/route"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	router := route.NewRouter()
	fmt.Println("Server has started")

	log.Fatal(http.ListenAndServe(":8080", router))

}
