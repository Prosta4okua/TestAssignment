package swg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const csvFileName string = "assets/emails.csv"

func main() {
	//CsvAdd("danyavdele@gmail.com")
	//print(CsvRead())
	fmt.Print("[DEBUG] Manipulating with .csv file\n\n")
	for {
		fmt.Print(
			"Available categories:\n" +
				"	1. Add email\n" +
				"	2. Delete email\n" +
				"	3. Show emails\n" +
				"	4. Check email\n" +
				"	0. Exit\n" +
				"Enter the category:\t")
		var input int32
		fmt.Scanln(&input)
		switch input {
		case 0:
			os.Exit(0)
		case 1:
			var str string
			fmt.Println("Enter the email that you want to add")
			fmt.Scanln(&str)
			CsvAdd(str)
		case 2:
			var delThisEmail string
			fmt.Println("Enter the email that you want to delete")
			fmt.Scanln(&delThisEmail)
			CsvDelValue(delThisEmail)
		case 3:
			CsvShow()
		case 4:
			var checkEmail string
			fmt.Println("Enter the email that you want to check if it exists")
			fmt.Scanln(&checkEmail)
			if CheckCsvValue(checkEmail) == true {
				fmt.Println("This email already exists in the file!")
			} else {
				fmt.Println("This email doesn't exist in the file!")
			}
		}

	}
}

func CsvShow() {
	array := CsvRead()
	for index, item := range array {
		fmt.Println("[" + strconv.Itoa(index) + "] " + item)
	}
}

func CsvRead() []string {
	//csvLines, err := os.OpenFile(csvFileName, os.O_RDONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	panic(err)
	//}

	csvLines, err := os.ReadFile(csvFileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)

	}

	// fmt.Print("Hehe: ")
	// fmt.Println(string(csvLines))
	fmt.Println(strings.Split(string(csvLines), ","))
	return strings.Split(strings.Trim(string(csvLines), ","), ",")

}

func CsvAdd(newEmail string) {
	//f, err1 := os.OpenFile(csvFileName, os.O_APPEND, 0666)

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
	array := CsvRead()
	for _, item := range array {
		if item == email {
			return true
		}
	}
	return false
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
