package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var db *Database

	if len(os.Args) == 1 {
		// No command line arguments, check if .wkn file exists in the current directory
		if _, err := os.Stat(".wkn"); os.IsNotExist(err) {
			fmt.Println("Error: .wkn file not found. Use 'wkn new' to create a new database.")
			return
		}
		db = NewDatabase()
	} else if len(os.Args) == 2 && os.Args[1] == "new" {
		// Create a new .wkn file and start the REPL
		if _, err := os.Stat(".wkn"); err == nil {
			fmt.Println("Error: Db file already present")
			return
		}
		db = NewDatabase()
		err := db.persistToFile(".wkn")
		if err != nil {
			fmt.Println("Error creating .wkn file:", err)
			return
		}
	} else if len(os.Args) == 3 && os.Args[1] == "--db-path" {
		// Use a custom .wkn file path and start the REPL
		filePath := os.Args[2]
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Println("Error: File is corrupted.")
			return
		}

		db = NewDatabase()
		err := db.loadFromFile(filePath)
		if err != nil {
			fmt.Println("Error loading data from file:", err)
			return
		}
	} else {
		fmt.Println("Usage:")
		fmt.Println("1. To start REPL with an existing .wkn file: wkn")
		fmt.Println("2. To create a new .wkn file and start REPL: wkn new")
		fmt.Println("3. To start REPL with a custom .wkn file: wkn --db-path ./path_to_file")
		return
	}

	fmt.Println("REPL Database - Type 'exit' to quit")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("wkn> ")
		scanner.Scan()
		command := scanner.Text()

		if command == "exit" {
			fmt.Println("Bye!")
			break
		}

		handleCommand(db, command)
		err := db.persistToFile(".wkn")
		if err != nil {
			fmt.Println("Error persisting data:", err)
		}
	}
}

func handleCommand(db *Database, command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "new":
		if len(parts) > 1 {
			db.newArray(parts[1], parts[2:]...)
		} else {
			fmt.Println("Error: Creating new array_name [element1, element2, ...]")
		}
	case "show":
		if len(parts) > 1 {
			db.showArray(parts[1])
		} else {
			fmt.Println("Error:- show array_name")
		}
	case "del":
		if len(parts) > 1 {
			db.deleteArray(parts[1])
		} else {
			fmt.Println("Error:- delete array_name")
		}
	case "merge":
		if len(parts) == 3 {
			db.mergeArrays(parts[1], parts[2])
		} else {
			fmt.Println("Error:- merge array_name1 array_name2")
		}
	default:
		fmt.Println("Error:" + parts[0] + "" + "is not a supported operation")
	}
}
