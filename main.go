package main

import (
	"casetask/hehe/src/route"
	_ "encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print(err)
		return
	}

	createAssetsFolderIfNotExists()

	router := route.NewRouter()
	fmt.Println("Server has started")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func createAssetsFolderIfNotExists() {
	path := "assets/"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
