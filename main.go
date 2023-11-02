package main

//https://github.com/gophercises/quiz

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type QuestionFormat struct {
	NumOne int
	NumTwo int
}

func (qF QuestionFormat) EvaluateAddition() int {
	return qF.NumOne + qF.NumTwo
}

func MakeRangeOfRandomQuestionFormat(min, max int) []QuestionFormat {
	numRange := make([]QuestionFormat, max-min+1)
	for i := range numRange {
		numRange[i] = QuestionFormat{NumOne: rand.Intn(100), NumTwo: rand.Intn(100)}
	}
	return numRange
}

func MakeQuiz(questionsMin, questionsMax int) {
	RandomQuestionsRange := MakeRangeOfRandomQuestionFormat(1, 100)

	csvFile, err := os.Create("problems.csv")
	defer csvFile.Close()
	if err != nil {
		log.Fatalf("failed creating file: %s\n", err)
	}

	csvWriter := csv.NewWriter(csvFile)

	var data [][]string

	data = append(data, []string{"equation", "answer"})
	for _, questions := range RandomQuestionsRange {
		equation := strconv.Itoa(questions.NumOne) + "+" + strconv.Itoa(questions.NumTwo)
		row := []string{equation, strconv.Itoa(questions.EvaluateAddition())}
		data = append(data, row)
	}

	csvWriter.WriteAll(data)

}

func main() {
	MakeQuiz(1, 100)
}
