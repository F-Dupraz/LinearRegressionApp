package main

import (
	"fmt"
	"os"

	LN "github.com/F-Dupraz/LinearRegressionApp/LinearRegression"
	Reader "github.com/F-Dupraz/LinearRegressionApp/csvReader"
)

func main() {
	// Initialization of the linear regression and CSV reader classes
	linearRegression := LN.NewLinearRegression()
	myCsvReader := Reader.NewReadCsv()

	myDocuments := myCsvReader.GetCSVs()

	var hasHeader bool = true

	for _, name := range myDocuments {
		// Read CSV data for independent (X) and dependent (Y) variables
		X_values, err := myCsvReader.GetIndependentVariables(name, hasHeader)
		if err != nil || len(X_values) == 0 {
			fmt.Println("Error: Unable to read independent variables.")
			os.Exit(1)
		}

		Y_values, err := myCsvReader.GetTargetValues(name, hasHeader)
		if err != nil || len(Y_values) == 0 {
			fmt.Println("Error: Unable to read target values.")
			os.Exit(1)
		}

		fmt.Println("This is the length of the X values: ", len(X_values))
		fmt.Println("This is the length of the X[1] values: ", len(X_values[1]))
		fmt.Println("This is the length of the Y values: ", len(Y_values))

		// // Error handling for different data lengths
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
	}

	// prediction := linearRegression.Predict([]float64{63.0, 23.085, 0})
	// fmt.Println("The prediction is:", prediction, "\n")
}