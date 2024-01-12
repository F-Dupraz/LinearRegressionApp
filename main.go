package main

import (
	"fmt"
	"os"
	"strings"

	LN "github.com/F-Dupraz/LinearRegressionApp/LinearRegression"
	Reader "github.com/F-Dupraz/LinearRegressionApp/csvReader"
)

func main() {
	// Initialization of the linear regression and CSV reader classes
	linearRegression := LN.NewLinearRegression()
	myCsvReader := Reader.NewReadCsv()

	// User input for file path, target column index, number of independent variables, and their indices
	var filepath string
	fmt.Print("Enter the path to the file: ")
	fmt.Scanln(&filepath)

	var targetIndexY uint
	fmt.Print("Enter the column index of the dependent variable (Y): ")
	fmt.Scan(&targetIndexY)

	var numberOfValues int
	fmt.Print("How many independent variables do you have? ")
	fmt.Scan(&numberOfValues)

	var targetIndexX []uint
	for i := 0; i < numberOfValues; i++ {
		var index uint
		fmt.Print("Enter one of the indices for the independent variable (X): ")
		fmt.Scan(&index)
		targetIndexX = append(targetIndexX, index)
	}

	var hasHeader string
	fmt.Print("Does your file have a header? (y/n) ")
	fmt.Scan(&hasHeader)
	hasHeader = strings.ToLower(hasHeader)

	// Read CSV data for independent (X) and dependent (Y) variables
	X_values, err := myCsvReader.GetIndependentVariables(filepath, hasHeader == "y", targetIndexX)
	if err != nil || len(X_values) == 0 {
		fmt.Println("Error: Unable to read independent variables.")
		os.Exit(1)
	}

	Y_values, err := myCsvReader.GetTargetValues(filepath, hasHeader == "y", targetIndexY)
	if err != nil || len(Y_values) == 0 {
		fmt.Println("Error: Unable to read target values.")
		os.Exit(1)
	}

	// Error handling for different data lengths
	if len(X_values[0]) != len(Y_values) {
		fmt.Println("Error: Different number of data points for the variables.")
		os.Exit(1)
	}

	// Fit the linear regression model and display results
	if linearRegression.Fit(X_values, Y_values) == 1 {
		fmt.Println("\nIntercept:", linearRegression.GetIntercept())
		fmt.Print("Coefficients: ")
		for _, coef := range linearRegression.GetSlope() {
			fmt.Print(coef, " ")
		}
		fmt.Println("\n")
	} else {
		fmt.Println("Error in linear regression fit.")
	}

	prediction := linearRegression.Predict([]float64{63.0, 23.085, 0})
	fmt.Println("The prediction is:", prediction, "\n")
}