package LinearRegression

import (
	"fmt"
	"math"
)

// LinearRegression struct
type LinearRegression struct {
	coefficients []float64
	intercept float64
}

// New LinearRegression constructor
func NewLinearRegression() *LinearRegression {
	return &LinearRegression{}
}

func (lr *LinearRegression) Fit(xValues [][]float64, yValues []float64) float64 {
	if len(xValues) == 0 || len(yValues) == 0 {
		fmt.Println("Error: One or both vectors are empty.")
		return math.NaN()
	}

	numCoefficients := len(xValues)
	lr.coefficients = make([]float64, numCoefficients)

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
				xValue = xValues[i][j] // Intercepto o Xj
			}

			accumulativeNumeratorSum += (xValue - xMeans[j]) * (yValues[i] - yMean)
			accumulativeDenominatorSum += (xValue - xMeans[j]) * (xValue - xMeans[j])
		}

		// Avoid division by zero
		if accumulativeDenominatorSum == 0 {
			return math.NaN()
		}

		lr.coefficients[j] = accumulativeNumeratorSum / accumulativeDenominatorSum
	}

	lr.calculateIntercept(xMeans, yMean)

	return 1
}

func (lr *LinearRegression) Predict(xValues []float64) float64 {
	if len(xValues) != len(lr.coefficients) {
		return math.NaN()
	}

	// Calculate the predicted dependent variable value
	result := lr.intercept

	for j := 0; j < len(lr.coefficients); j++ {
		result += lr.coefficients[j] * xValues[j]
	}

	return result
}

func (lr *LinearRegression) GetSlope() []float64 {
	return lr.coefficients
}

func (lr *LinearRegression) GetIntercept() float64 {
	return lr.intercept
}

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
		fmt.Println(means[i])
	}
}

func (lr *LinearRegression) calculateIntercept(xMeans []float64, yMean float64) {
	myIntercept := yMean
	for j := 0; j < len(lr.coefficients); j++ {
		myIntercept -= lr.coefficients[j] * xMeans[j]
	}

	lr.intercept = myIntercept

	fmt.Println(lr.intercept)
}


