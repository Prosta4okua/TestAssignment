package swg

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const csvFileName string = "assets/emails.csv"

func CsvShow() {
	array := CsvRead()
	fmt.Print("[")
	for _, item := range array {
		fmt.Print(item + ", ")
	}
	fmt.Print("]")
}

func CsvRead() []string {

	csvLines, err := os.ReadFile(csvFileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)

	}

	return strings.Split(strings.Trim(string(csvLines), ","), ",")

}

func CsvAdd(newEmail string) {

	f, err := os.OpenFile(csvFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	if _, err = f.WriteString("," + newEmail); err != nil {
		panic(err)
	}
}

func CheckCsvValue(email string) bool {
	if checkIfFileExists() == false {
		createCsvFile()
	}
	array := CsvRead()
	for _, item := range array {
		if item == email {
			return true
		}
	}
	return false
}

func createCsvFile() {
	file, e := os.Create(csvFileName)
	if e != nil {
		log.Fatal(e)
	}
	err := file.Close()
	if err != nil {
		return
	}
}

func checkIfFileExists() bool {
	_, err := os.ReadFile(csvFileName)
	if err != nil {
		return false
	}
	return true
}
