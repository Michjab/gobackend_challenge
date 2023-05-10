// Package handlers provides
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var response string

// Public API EchoHandler function provides handler for REST endpoint /echo for echo service given csv file
func (app *application) EchoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		app.logger.Printf("error reading file \"%s\": %s", file, err)
		return
	}
	defer file.Close()
	var response string

	app.logger.Printf("processing csv file on /echo endpoint")
	records, err := csv.NewReader(file).ReadAll() //This could be considered as a weakpoint in case sent file was very big (mem alocation twice big as could be)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		app.logger.Printf("error reading file \"%s\": %s", file, err)
		return
	}
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
	app.logger.Printf("processing succeed")
}

// Public API SumHandler function provides handler for REST endpoint /sum for csv integer matrix fields sum
func (app *application) SumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		app.logger.Printf("error reading file \"%s\": %s", file, err)
		return
	}

	defer file.Close()

	var response string
	CSVreader := csv.NewReader(file)
	CSVSum := 0
	app.logger.Printf("processing csv file on /sum endpoint")
sumLoop:
	for {
		record, err := CSVreader.Read()
		if err == io.EOF {
			break sumLoop
		}
		if err != nil {
			response := fmt.Errorf("error during reading CSV file: %w", err)
			fmt.Fprint(w, response, "\n")
			return
		}
		rowResult, err := sumRow(record)
		if err != nil {
			app.logger.Printf("error during sum operation: %s", err)
			response := fmt.Errorf("error sum operation: %w", err)
			fmt.Fprint(w, response, "\n")
			return
		}
		CSVSum += rowResult
	}
	response = fmt.Sprintf("%v\n", CSVSum)
	fmt.Fprint(w, response)
	app.logger.Printf("processing succeed")
}

// Public API MultiplyHandler function provides handler for REST endpoint /multiply for csv integer matrix fields multiplication
func (app *application) MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		app.logger.Printf("error reading file \"%s\": %s", file, err)
		return
	}
	defer file.Close()

	var response string
	CSVreader := csv.NewReader(file)
	CSVResult := 1
	app.logger.Printf("processing csv file on /multiply endpoint")
multiplyLoop:
	for {
		record, err := CSVreader.Read()
		if err == io.EOF {
			break multiplyLoop
		}
		if err != nil {
			response := fmt.Errorf("error during reading CSV file: %w", err)
			fmt.Fprint(w, response, "\n")
			return
		}
		recordInt, err := strToInt(record)
		if err != nil {
			response := fmt.Errorf("error during CSV conversion to integer data \"%v\": %w", recordInt, err)
			fmt.Fprint(w, response, "\n")
			return
		}
		rowResult, err := multiply(recordInt)
		if err != nil {
			app.logger.Printf("error during multiplication operation: %s", err)
			response := fmt.Errorf("error multiplication: %w", err)
			fmt.Fprint(w, response, "\n")
			return
		}
		CSVResult *= rowResult
		if CSVResult == 0 {
			fmt.Fprint(w, CSVResult, "\n")
			return
		}
	}
	response = fmt.Sprintf("%v\n", CSVResult)
	fmt.Fprint(w, response)
	app.logger.Printf("processing succeed")
}

// Public API FlattenHandler function provides handler for REST endpoint /flatten for csv matrix conversion 2D into 1D
func (app *application) FlattenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		app.logger.Printf("error reading file \"%s\": %s", file, err)
		return
	}
	defer file.Close()

	var response string
	CSVreader := csv.NewReader(file)
	var isFirstRow = true // flag to avoid comma prefix before first slice join
	app.logger.Printf("processing csv file on /flatten endpoint")
CSVLoop:
	for {
		record, err := CSVreader.Read()
		if err == io.EOF {
			response = fmt.Sprintf("%s\n", response)
			break CSVLoop
		}
		if err != nil {
			response := fmt.Errorf("error during reading CSV file: %w", err)
			fmt.Fprint(w, response, "\n")
			return
		}
		recordInt, err := strToInt(record) // check if csv contains integers only
		if err != nil {
			response := fmt.Errorf("error during CSV conversion to integer data \"%v\": %w", recordInt, err)
			fmt.Fprint(w, response, "\n")
			return
		}
		if isFirstRow {
			isFirstRow = false
			response = strings.Join(record, ",")
		} else {
			response = fmt.Sprintf("%s,%s", response, strings.Join(record, ","))
		}
	}
	fmt.Fprint(w, response)
	app.logger.Printf("processing succeed")
}

// Public API TranspositionHandler function provides handler for REST endpoint /invert for csv matrix
func (app *application) TranspositionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		app.logger.Printf("error reading file \"%s\": %s", file, err)
		return
	}
	defer file.Close()
	var response string
	app.logger.Printf("processing csv file on /invert endpoint")
	records, err := csv.NewReader(file).ReadAll() //This could be considered as a weakpoint in case sent file was very big (mem alocation twice big as could be)
	if err != nil {
		response := fmt.Errorf("error during reading CSV file: %w", err)
		fmt.Fprint(w, response, "\n")
		return
	}
	invertedCSV, err := transpose(records)
	if err != nil {
		response := fmt.Errorf("error during transposing: %w", err)
		fmt.Fprint(w, response, "\n")
		return
	}
	for _, row := range invertedCSV {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
	app.logger.Printf("processing succeed")
}
