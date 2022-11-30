package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

func readCsvFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to read file"+path, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+path, err)
	}

	return records
}

func main() {

	fileNamePtr := flag.String("filename", "problems.csv", "Provide the file name where the quiz data is located.")
	flag.Parse()

	records := readCsvFile(*fileNamePtr)
	var correctAnswers float64 = 0
	var wrongAnswers float64 = 0

	for i := 0; i < len(records); i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(records[i][0])
		input, err := reader.ReadString('\n')
		input = input[:len(input)-1]
		if err != nil {
			log.Fatal("Unable to read your input!")
		}

		if input == records[i][1] {
			correctAnswers++
		} else {
			wrongAnswers++
		}
	}

	score := correctAnswers / float64((len(records))) * 100

	fmt.Println("CORRECT ANSWERS: ", correctAnswers)
	fmt.Println("WRONG ASNWERS: ", wrongAnswers)
	fmt.Println("YOUR OVERALL SCORE (%): ", math.Round(score))
}
