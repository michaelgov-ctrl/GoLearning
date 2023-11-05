package main

//https://github.com/gophercises/quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Uneccessary but still used struct for formatting questions
type questionFormat struct {
	NumOne int
	NumTwo int
}

// 'method' to add NumOne & NumTwo
func (qF questionFormat) EvaluateAddition() int {
	return qF.NumOne + qF.NumTwo
}

func makeRangeOfRandomQuestionFormat(min, max int) []questionFormat {
	numRange := make([]questionFormat, max-min+1)
	for i := range numRange {
		numRange[i] = questionFormat{NumOne: rand.Intn(100), NumTwo: rand.Intn(100)}
	}
	return numRange
}

func makeQuiz(questionsMin, questionsMax int) {
	randomQuestionsRange := makeRangeOfRandomQuestionFormat(questionsMin, questionsMax)
	filePath := "problems.csv"

	csvFile, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("failed creating file: %s\n", err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	var data [][]string

	data = append(data, []string{"equation", "answer"})
	for _, questions := range randomQuestionsRange {
		equation := strconv.Itoa(questions.NumOne) + "+" + strconv.Itoa(questions.NumTwo)
		row := []string{equation, strconv.Itoa(questions.EvaluateAddition())}
		data = append(data, row)
	}

	err = csvWriter.WriteAll(data)
	if err != nil {
		log.Fatalf("unable to write to file: %s %s\n", filePath, err)
	}
}

func getQuiz(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to read input file: %s %s\n", filePath, err)
	}
	defer f.Close()

	csvReader := *csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("unable to parse csv as file: %s %s\n", filePath, err)
	}

	return records

}

func takeQuiz(qs [][]string) {
	reader := bufio.NewReader(os.Stdin)
	correctCount := 0
	for index, val := range qs {
		if index == 0 {
			continue
		}
		fmt.Print("what is: ", val[0], "= ")
		input, _ := reader.ReadString('\n')
		SplitInput := strings.Trim(input, "\r\n")
		atoiInput, err := strconv.Atoi(SplitInput)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to convert input to integer: %s", err))
			continue
		}
		answer, err := strconv.Atoi(val[1])
		if atoiInput == answer {
			fmt.Println("correct")
			correctCount++
		} else {
			fmt.Println("incorrect")
		}
	}
	fmt.Printf("%d/%d correct\n", correctCount, len(qs)-1)
}

func main() {
	makeQuiz(1, 25)
	takeQuiz(getQuiz("problems.csv"))
}
