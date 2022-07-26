package swg

import (
	"errors"
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

	defer f.Close()

	if _, err = f.WriteString("," + newEmail); err != nil {
		panic(err)
	}
}

func CsvDelValue(badEmail string) {
	array := CsvRead()
	if CheckCsvValue(badEmail) == false {
		err := errors.New("this email doesn't exist")
		if err != nil {
			return
		}
	}
	for index, item := range array {
		if item == badEmail {
			array = append(array[:index], array[index+1:]...)
			break
		}
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkIfFileExists() bool {
	_, err := os.ReadFile(csvFileName)
	if err != nil {
		return false
	}
	return true
}
