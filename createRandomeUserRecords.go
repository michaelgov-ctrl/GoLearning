package main

import (
	"encoding/csv"
	"errors"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

type Record struct {
	Name string
	ID   int
}

func makeRecords(maxRecords int) ([]Record, error) {
	if maxRecords <= 0 {
		return nil, errors.New("cannot create a slice of Records with length <= 0")
	}

	firstNames := []string{
		"Amelia", "Ava", "Barbara", "Benjamin", "Charles", "Charlotte", "David", "Dorothy", "Elijah", "Elizabeth", "Emma", "Evelyn",
		"Henry", "Isabella", "James", "Jennifer", "John", "Joseph", "Liam", "Linda", "Lucas", "Luna", "Margaret", "Maria", "Mary",
		"Mia", "Michael", "Noah", "Oliver", "Olivia", "Patricia", "Richard", "Robert", "Sophia", "Susan", "Theodore", "Thomas", "William",
	}
	lastNames := []string{
		"Adams", "Allen", "Anderson", "Bailey", "Baker", "Barnes", "Bell", "Bennett", "Brooks", "Brown", "Butler",
		"Campbell", "Carter", "Clark", "Collins", "Cook", "Cooper", "Cox", "Cruz", "Davis", "Diaz", "Edwards",
		"Evans", "Fisher", "Flores", "Foster", "Garcia", "Gomez", "Gonzalez", "Gray", "Green", "Gutierrez", "Hall",
		"Harris", "Hernandez", "Hill", "Howard", "Hughes", "Jackson", "James", "Jenkins", "Johnson", "Jones", "Kelly",
		"King", "Lee", "Lewis", "Long", "Lopez", "Martin", "Martinez", "Miller", "Mitchell", "Moore", "Morales",
		"Morgan", "Morris", "Murphy", "Myers", "Nelson", "Nguyen", "Ortiz", "Parker", "Perez", "Perry", "Peterson", "Phillips",
		"Powell", "Price", "Ramirez", "Reed", "Reyes", "Richardson", "Rivera", "Roberts", "Robinson", "Rodriguez", "Rogers",
		"Ross", "Russell", "Sanchez", "Sanders", "Scott", "Smith", "Stewart", "Sullivan", "Taylor", "Thomas", "Thompson",
		"Torres", "Turner", "Walker", "Ward", "Watson", "White", "Williams", "Wilson", "Wood", "Wright", "Young",
	}
	recordsSlice := make([]Record, maxRecords+1)
	for i := range recordsSlice {
		randomName := lastNames[rand.Intn(len(lastNames)-1)] + ", " + firstNames[rand.Intn(len(firstNames)-1)]
		recordsSlice[i] = Record{Name: randomName, ID: i}
	}

	sort.Slice(
		recordsSlice,
		func(i, j int) bool {
			return recordsSlice[i].Name < recordsSlice[j].Name
		},
	)

	return recordsSlice, nil
}

func writeRecords(records []Record, filePath string) error {
	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	csvFile, err := os.OpenFile(filePath, flags, 0644)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	var dataToWrite [][]string

	dataToWrite = append(dataToWrite, []string{"user", "id"})
	for _, rec := range records {
		row := []string{rec.Name, strconv.Itoa(rec.ID)}
		dataToWrite = append(dataToWrite, row)
	}

	err = csvWriter.WriteAll(dataToWrite)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 3 || os.Args[1] == "-h" {
		log.Fatalf("argument position 1 is recordCount position 2 is outputfilepath, ex: <executable>.exe 100 './output.csv'")
	} else {
		recordCount, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("failed to convert command line argument to int")
		}

		records, err := makeRecords(recordCount)
		if err != nil {
			log.Fatalf("failed to create records")
		}

		if err := writeRecords(records, os.Args[2]); err != nil {
			log.Fatalf("failed to sort records")
		}
	}
}
