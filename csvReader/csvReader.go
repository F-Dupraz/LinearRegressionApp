package csvReader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// ReadCsv struct
type ReadCsv struct{}

// NewReadCsv constructor
func NewReadCsv() *ReadCsv {
	return &ReadCsv{}
}

// GetTargetValues gets target values from a CSV file
func (rc *ReadCsv) GetTargetValues(filepath string, hasHeader bool, columnIndex uint) ([]float64, error) {
	// Open the CSV file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Unable to open the file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the header if it exists
	if hasHeader {
		_, err := reader.Read()
		if err != nil {
			return nil, fmt.Errorf("Error reading header: %v", err)
		}
	}

	// Read data from the CSV file
	var columnData []float64

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Ensure the record has enough fields
		if len(record) <= int(columnIndex) {
			return nil, fmt.Errorf("Error: The line does not have enough fields")
		}

		// Convert the field to a double and add it to the vector
		value, err := strconv.ParseFloat(record[columnIndex], 64)
		if err != nil {
			return nil, fmt.Errorf("Error converting to double: %v", err)
		}

		columnData = append(columnData, value)
	}

	return columnData, nil
}

// GetIndependentVariables gets independent variables from a CSV file
func (rc *ReadCsv) GetIndependentVariables(filepath string, hasHeader bool, columnIndices []uint) ([][]float64, error) {
	// Open the CSV file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Unable to open the file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the header if it exists
	if hasHeader {
		_, err := reader.Read()
		if err != nil {
			return nil, fmt.Errorf("Error reading header: %v", err)
		}
	}

	// Initialize the 2D slice to store data from multiple columns
	columnsData := make([][]float64, len(columnIndices))

	// Read data from the CSV file
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Initialize currentIndex to keep track of the column index
		currentIndex := uint(0)

		// Parse each field in the line
		for _, field := range record {
			// Check if the current column index is within the specified column indices
			if currentIndex < uint(len(columnIndices)) {
				// Check if the current index is in the specified column indices
				if contains(columnIndices, currentIndex) {
					// Convert the field to a double and add it to the corresponding vector
					value, err := strconv.ParseFloat(field, 64)
					if err != nil {
						return nil, fmt.Errorf("Error converting to double: %v - Field: %v", err, field)
					}

					columnsData[currentIndex] = append(columnsData[currentIndex], value)
				}

				// Increment currentIndex after processing each field
				currentIndex++
			} else {
				// Ignore additional fields if there are more than specified in columnIndices
				break
			}
		}
	}

	return columnsData, nil
}

// contains checks if a value is present in a slice of uint
func contains(slice []uint, value uint) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
