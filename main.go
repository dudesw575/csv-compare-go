package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func removeQuotesAndBrackets(s string) string {
	s = strings.ReplaceAll(s, `"`, "") // Remove quotes
	s = strings.ReplaceAll(s, "[", "") // Remove opening square bracket
	s = strings.ReplaceAll(s, "]", "") // Remove closing square bracket
	return s
}

func removeQuotes(s []string) []string {
	for i := range s {
		s[i] = removeQuotesAndBrackets(s[i])
	}
	return s
}

func main() {
	// Import the first CSV file.
	file1, err := os.Open("")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader1 := csv.NewReader(file1)
	records1, _ := reader1.ReadAll()

	// Import the second CSV file.
	file2, err := os.Open("")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader2 := csv.NewReader(file2)
	records2, _ := reader2.ReadAll()

	// Compare the two CSV files on the first column.
	matches := []string{}
	for i := range records1 {
		for j := range records2 {
			if records1[i][0] == records2[j][0] {
				match := fmt.Sprintf("%v, %v", records1[i], records2[j])
				match = removeQuotesAndBrackets(match)
				matches = append(matches, match)
			}
		}
	}

	// Export the matching records to a CSV file.
	file, err := os.Create("matching_records.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := csv.NewWriter(file)
	writer.Write([]string{file1.Name(), file2.Name()})
	for _, match := range matches {
		writer.Write([]string{match})
	}
	writer.Flush()
}
