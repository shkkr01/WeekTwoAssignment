package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"
)

// Database represents the main structure for managing arrays.
type Database struct {
	Data map[string][]int
	Lock sync.Mutex
}

func NewDatabase() *Database {
	return &Database{
		Data: make(map[string][]int),
	}
}

func (db *Database) newArray(name string, elements ...string) {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	if _, exists := db.Data[name]; exists {
		fmt.Println("Error: Array already exists")
		return
	}

	var newArray []int
	for _, elem := range elements {
		// Convert elements to integers and append to the array
		// Handle error if conversion fails
		if num, err := strconv.Atoi(elem); err == nil {
			newArray = append(newArray, num)
		}
	}

	db.Data[name] = newArray
	fmt.Printf("CREATED (%d)\n", len(newArray))
}

func (db *Database) showArray(name string) {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	if array, exists := db.Data[name]; exists {
		fmt.Println(array)
	} else {
		fmt.Printf("Error: \"%s\" does not exist\n", name)
	}
}

func (db *Database) deleteArray(name string) {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	if _, exists := db.Data[name]; exists {
		delete(db.Data, name)
		fmt.Println("DELETED")
	} else {
		fmt.Printf("Error: \"%s\" does not exist\n", name)
	}
}

func (db *Database) mergeArrays(dest, src string) {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	destArray, destExists := db.Data[dest]
	srcArray, srcExists := db.Data[src]

	if destExists {
		if srcExists {
			db.Data[dest] = append(destArray, srcArray...)
			fmt.Println("MERGED")
		} else {
			fmt.Printf("Error: \"%s\" does not exist\n", src)
		}
	} else {
		fmt.Printf("Error: \"%s\" does not exist\n", dest)
	}
}

func (db *Database) loadFromFile(filePath string) error {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &db.Data)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) persistToFile(filePath string) error {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	data, err := json.Marshal(db.Data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
