package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: csvfilter <inputfile.csv>")
		return
	}

	inputFile := os.Args[1]

	// Open input CSV
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1 // allow variable columns

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
		return
	}

	if len(records) == 0 {
		fmt.Println("CSV file is empty")
		return
	}

	// Header row mapping (so we don't rely on fixed column order)
	headers := records[0]
	colIndex := make(map[string]int)
	for i, h := range headers {
		colIndex[h] = i
	}

	// Prepare output CSV
	outFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Printf("Error creating output.csv: %v\n", err)
		return
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	// Write new header
	writer.Write([]string{"Date", "Payee", "Memo", "Amount"})

	// Actions to skip
	skipActions := map[string]bool{
		"Interest on cash":  true,
		"Spending cashback": true,
		"Market buy":        true,
		"Market sell":       true,
	}

	// Process rows
	for _, row := range records[1:] {
		action := row[colIndex["Action"]]
		if skipActions[action] {
			continue
		}

		// Trim time down to just date
		timeFull := row[colIndex["Time"]]
		dateOnly := strings.Split(timeFull, " ")[0]

		merchant := row[colIndex["Merchant"]]
		category := row[colIndex["Category"]]
		total := row[colIndex["Total"]]

		newRow := []string{dateOnly, merchant, category, total}
		writer.Write(newRow)
	}

	fmt.Println("Processing complete. Output written to output.csv")
}
