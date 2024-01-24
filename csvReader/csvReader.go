package csvReader

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ReadCSV struct
type ReadCSV struct{}

// NewReadCSV constructor
func NewReadCSV() *ReadCSV {
	return &ReadCSV{}
}

// GetCSVs gets the list of CSV files in the "CSVs" folder.
func (rc *ReadCSV) GetCSVs() []string {
	var myCSVs []string

	myDir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("%v", err)
		return nil
	}

	CSVDir := "CSVs"
	myCSVsDir := filepath.Join(myDir, CSVDir)

	myDirDocuments, err := os.ReadDir(myCSVsDir)
	if err != nil {
		fmt.Errorf("%v", err)
		return nil
	}

	for _, doc := range myDirDocuments {
		docName := doc.Name()
		if len(docName) > 3 {
			extension := len(docName) - 3
			if docName[extension:] == "csv" {
				myCSVs = append(myCSVs, docName)
			}
		}
	}

	return myCSVs
}

// GetTargetValues gets target values from a CSV file.
func (rc *ReadCSV) GetTargetValues(filePath string, hasHeader bool) ([]float64, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
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

		// Ensure the record has at least one field
		if len(record) > 0 {
			// Convert the last field to a double and add it to the vector
			value, err := strconv.ParseFloat(record[len(record)-1], 64)
			if err != nil {
				return nil, fmt.Errorf("Error converting to double: %v", err)
			}

			columnData = append(columnData, value)
		} else {
			return nil, fmt.Errorf("Error: The line does not have any fields")
		}
	}

	return columnData, nil
}

// GetIndependentVariables gets independent variables from a CSV file.
func (rc *ReadCSV) GetIndependentVariables(filePath string, hasHeader bool) ([][]float64, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
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

	// Initialize the 2D matrix to store data from multiple columns
	var columnsData [][]float64

	// Read data from the CSV file
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Ensure the row has at least one column
		if len(record) > 0 {
			// Convert the fields to floating-point numbers and add them to the matrix
			var rowData []float64
			for i := 0; i < (len(record) - 1); i++ {
				value, err := strconv.ParseFloat(record[i], 64)
				if err != nil {
					return nil, fmt.Errorf("Error converting to floating-point number: %v - Field: %v", err, record[i])
				}
				rowData = append(rowData, value)
			}

			columnsData = append(columnsData, rowData)
		} else {
			return nil, fmt.Errorf("Error: The row does not have any fields")
		}
	}

	// Transpose the data directly into columnsData
	transposedData := make([][]float64, len(columnsData[0]))
	for i := range transposedData {
		transposedData[i] = make([]float64, len(columnsData))
		for j := range columnsData {
			transposedData[i][j] = columnsData[j][i]
		}
	}

	return transposedData, nil
}

// Get the worked data
func (rc *ReadCSV) GetData(filePath string) ([]float64, error) {
	// Abre el archivo CSV
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Lee el contenido del archivo CSV
	reader := csv.NewReader(file)
	record, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Convierte los valores de la primera línea a float64 y almacénalos en un array
	var floatArray []float64
	for _, value := range record {
		// Elimina espacios en blanco y convierte a float64
		floatValue, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
		if err != nil {
			return nil, err
		}
		floatArray = append(floatArray, floatValue)
	}

	return floatArray, nil
}


// WriteDB writes data to a CSV file.
func (rc *ReadCSV) WriteDB(csvFile string, values [][]float64) {
	file, err := os.OpenFile(csvFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening the CSV file: %v\n", err)
		return
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	for _, row := range values {
		var stringValues []string

		for _, value := range row {
			stringValues = append(stringValues, fmt.Sprintf("%f", value))
		}

		if err := csvWriter.Write(stringValues); err != nil {
			fmt.Printf("Error writing to the CSV file: %v\n", err)
			return
		}

		csvWriter.Flush()

		if err := csvWriter.Error(); err != nil {
			fmt.Printf("Error flushing to the CSV file: %v\n", err)
			return
		}
	}
}


// WriteCSV writes a single row of data to the end of a CSV file.
func (rc *ReadCSV) WriteCSV(csvFile string, values []string) {
	file, err := os.OpenFile(csvFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening the CSV file: %v\n", err)
		return
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	if err := csvWriter.Write(values); err != nil {
		fmt.Printf("Error writing to the CSV file: %v\n", err)
		return
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		fmt.Printf("Error flushing to the CSV file: %v\n", err)
		return
	}
}
