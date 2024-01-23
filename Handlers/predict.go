package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/F-Dupraz/LinearRegressionApp/server"
	LR "github.com/F-Dupraz/LinearRegressionApp/LinearRegression"
	CSVReader "github.com/F-Dupraz/LinearRegressionApp/csvReader"
)

type RequestPredictInsurance struct {
	Age uint `json:"age"`
	Bmi float64 `json:"bmi"`
	Children uint `json:"children"`
}

type ResponsePredictInsurance struct {
	Charges float64 `json:"charges"`
}

func PredictInsurance(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var myPredict = RequestPredictInsurance{}
		err := json.NewDecoder(r.Body).Decode(&myPredict)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Initialization of the linear regression and CSV reader classes.
		linearRegression := LR.NewLinearRegression()
		myCSVReader := CSVReader.NewReadCSV()

		predictionData, err := myCSVReader.GetData("./DB/insurance.csv")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i := 0; i < len(predictionData); i++ {
			if i == 0 {
				linearRegression.Intercept = predictionData[i]
			} else {
				linearRegression.Coefficients = append(linearRegression.Coefficients, predictionData[i])
			}
		}

		inputData := []float64{float64(myPredict.Age), myPredict.Bmi, float64(myPredict.Children)}

		prediction := linearRegression.Predict(inputData)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponsePredictInsurance{
			Charges: float64(prediction),
		})
	}
}