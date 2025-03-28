package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsv() {
	file, err := os.Open("sample.csv")
	if err != nil {
		log.Fatal("Error opening CSV file")
	}

	// create a new csv reader
	reader := csv.NewReader(file)

	// Read all the data from the file
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error while reading from the file", err)
	}

	for idx, row := range data {
		// this is to avoid header assignment
		if idx == 0{
			continue
		} 
		name := row[0]
		age := row[1]
		fmt.Printf("Name: %s, Age: %s\n", name, age)
	}

}
