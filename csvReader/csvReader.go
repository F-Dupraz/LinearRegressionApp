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

func (rc *ReadCsv) GetCSVs() []string {
	var myCSVs []string

	myDir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("%v", err)
		return nil
	}

	myDirDocuments, err := os.ReadDir(myDir)
	if err != nil {
		fmt.Errorf("%v", err)
		return nil
	}

	for _, doc := range myDirDocuments {
		docName := doc.Name()
		extension := len(docName) - 3
		if docName[extension:] == "csv" {
			myCSVs = append(myCSVs, docName)
		}
	}

	return myCSVs
}

// GetTargetValues gets target values from a CSV file
func (rc *ReadCsv) GetTargetValues(filepath string, hasHeader bool) ([]float64, error) {
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

// GetIndependentVariables obtiene variables independientes desde un archivo CSV
func (rc *ReadCsv) GetIndependentVariables(filepath string, hasHeader bool) ([][]float64, error) {
	// Abrir el archivo CSV
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("No se puede abrir el archivo: %v", err)
	}
	defer file.Close()

	// Crear un lector de CSV
	reader := csv.NewReader(file)

	// Leer el encabezado si existe
	if hasHeader {
		_, err := reader.Read()
		if err != nil {
			return nil, fmt.Errorf("Error al leer el encabezado: %v", err)
		}
	}

	// Inicializar la matriz 2D para almacenar datos de múltiples columnas
	var columnsData [][]float64

	// Leer datos del archivo CSV
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Verificar que la fila tiene al menos una columna
		if len(record) > 0 {
			// Convertir los campos a números flotantes y agregarlos a la matriz
			var rowData []float64
			for _, field := range record {
				value, err := strconv.ParseFloat(field, 64)
				if err != nil {
					return nil, fmt.Errorf("Error al convertir a número flotante: %v - Campo: %v", err, field)
				}
				rowData = append(rowData, value)
			}

			columnsData = append(columnsData, rowData)
		} else {
			return nil, fmt.Errorf("Error: La fila no tiene campos")
		}
	}

	// Transponer los datos directamente en columnsData
	transposedData := make([][]float64, len(columnsData[0]))
	for i := range transposedData {
		transposedData[i] = make([]float64, len(columnsData))
		for j := range columnsData {
			transposedData[i][j] = columnsData[j][i]
		}
	}

	return transposedData, nil
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
