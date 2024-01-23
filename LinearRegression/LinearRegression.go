package LinearRegression

import (
	"fmt"
	"math"
)

// LinearRegression structure representing a linear regression model.
type LinearRegression struct {
	Coefficients []float64
	Intercept    float64
}

// NewLinearRegression is a constructor to create a new instance of LinearRegression.
func NewLinearRegression() *LinearRegression {
	return &LinearRegression{}
}

// Fit trains the linear regression model with input and output data.
// xValues: Matrix of input values.
// yValues: Vector of corresponding output values.
// Returns 1 if the fit is successful, NaN in case of an error.
func (lr *LinearRegression) Fit(xValues [][]float64, yValues []float64) float64 {
	if len(xValues) == 0 || len(yValues) == 0 {
		fmt.Println("Error: One or both vectors are empty.")
		return math.NaN()
	}

	numCoefficients := len(xValues)
	lr.Coefficients = make([]float64, numCoefficients)

	xMeans := make([]float64, numCoefficients)
	lr.calculateMeans(xValues, xMeans)

	yMean := lr.meanFinder(yValues)

	for j := 0; j < numCoefficients; j++ {
		var accumulativeNumeratorSum float64 = 0.0
		var accumulativeDenominatorSum float64 = 0.0

		for i := 0; i < len(xValues); i++ {
			xValue := 0.0
			if j == 0 {
				xValue = 1
			} else {
				xValue = xValues[i][j] // Intercept or Xj
			}

			accumulativeNumeratorSum += (xValue - xMeans[j]) * (yValues[i] - yMean)
			accumulativeDenominatorSum += (xValue - xMeans[j]) * (xValue - xMeans[j])
		}

		// Avoid division by zero
		if accumulativeDenominatorSum == 0 {
			return math.NaN()
		}

		lr.Coefficients[j] = accumulativeNumeratorSum / accumulativeDenominatorSum
	}

	lr.calculateIntercept(xMeans, yMean)

	return 1
}

// Predict makes a prediction using the trained model.
// xValues: Vector of input values for making the prediction.
// Returns the predicted value.
func (lr *LinearRegression) Predict(xValues []float64) float64 {
	if len(xValues) != len(lr.Coefficients) {
		return math.NaN()
	}

	// Calculate the predicted dependent variable value
	result := lr.Intercept

	for j := 0; j < len(lr.Coefficients); j++ {
		result += lr.Coefficients[j] * xValues[j]
	}

	return result
}

// GetSlope returns the coefficients (slopes) of the linear regression model.
func (lr *LinearRegression) GetSlope() []float64 {
	return lr.Coefficients
}

// GetIntercept returns the Intercept term of the linear regression model.
func (lr *LinearRegression) GetIntercept() float64 {
	return lr.Intercept
}

// meanFinder calculates the mean of a vector of numbers.
func (lr *LinearRegression) meanFinder(findMean []float64) float64 {
	if len(findMean) == 0 {
		fmt.Println("Error: The vector was empty.")
		return math.NaN()
	}

	var accumulativeSum float64
	for _, value := range findMean {
		accumulativeSum += value
	}

	meanSum := accumulativeSum / float64(len(findMean))
	return meanSum
}

// calculateMeans calculates the means of the columns in an input matrix and stores them in the provided slice.
func (lr *LinearRegression) calculateMeans(xValues [][]float64, means []float64) {
	for i := 0; i < len(xValues); i++ {
		var accumulatedMeans float64
		for j := 0; j < len(xValues[i]); j++ {
			if j < len(xValues[i]) {
				accumulatedMeans += xValues[i][j]
			} else {
				fmt.Printf("Warning: Column index out of range for row %d\n", i)
			}
		}
		means[i] = accumulatedMeans / float64(len(xValues[i]))
	}
}

// calculateIntercept calculates the Intercept term of the linear regression model.
func (lr *LinearRegression) calculateIntercept(xMeans []float64, yMean float64) {
	myIntercept := yMean
	for j := 0; j < len(lr.Coefficients); j++ {
		myIntercept -= lr.Coefficients[j] * xMeans[j]
	}

	lr.Intercept = myIntercept
}
