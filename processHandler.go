package main

import (
	"fmt"
	"os"

	LR "github.com/F-Dupraz/LinearRegressionApp/LinearRegression"
	CSVReader "github.com/F-Dupraz/LinearRegressionApp/csvReader"
)

// main is the main function that runs when the program starts.
func main() {
	// Initialization of the linear regression and CSV reader classes.
	linearRegression := LR.NewLinearRegression()
	myCSVReader := CSVReader.NewReadCSV()

	// Get the list of CSV documents in the "CSVs" folder.
	myDocuments := myCSVReader.GetCSVs()

	var hasHeader bool = true
	var myData [][]float64

	// Iterate over each CSV document.
	for _, name := range myDocuments {
		docPath := "./CSVs/" + name

		// Read CSV data for independent (X) and dependent (Y) variables.
		XValues, err := myCSVReader.GetIndependentVariables(docPath, hasHeader)
		if err != nil || len(XValues) == 0 {
			fmt.Println("Error: Unable to read independent variables.")
			os.Exit(1)
		}

		YValues, err := myCSVReader.GetTargetValues(docPath, hasHeader)
		if err != nil || len(YValues) == 0 {
			fmt.Println("Error: Unable to read target values.")
			os.Exit(1)
		}

		// Error handling for different data lengths.
		if len(XValues[0]) != len(YValues) {
			fmt.Println("Error: Different number of data points for the variables.")
			os.Exit(1)
		}

		// Fit the linear regression model and display results.
		if linearRegression.Fit(XValues, YValues) == 1 {
			var myRowData []float64
			myRowData = append(myRowData, linearRegression.GetIntercept())
			for _, slope := range linearRegression.GetSlope() {
				myRowData = append(myRowData, slope)
			}
			myData = append(myData, myRowData)
		} else {
			fmt.Println("Error in linear regression fit.")
		}
	}

	// Write the data to a database CSV file.
	myCSVReader.WriteCSV("./DB/myDataDB.csv", myData)

	// Example prediction using the linear regression model.
	// prediction := linearRegression.Predict([]float64{63.0, 23.085, 0})
	// fmt.Println("The prediction is:", prediction, "\n")
}
