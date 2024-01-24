package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/F-Dupraz/LinearRegressionApp/server"
	CSVReader "github.com/F-Dupraz/LinearRegressionApp/csvReader"
)

// Insurance REQUEST and RESPONSE
type RequestTrainInsurance struct {
	Age      uint    `json:"age"`
	Bmi      float64 `json:"bmi"`
	Children uint    `json:"children"`
	Charges  float64 `json:"charges"`
}

type ResponseTrainInsurance struct {
	Message string `json:"message"`
}

func TrainInsurance(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var myTrain = RequestTrainInsurance{}
		err := json.NewDecoder(r.Body).Decode(&myTrain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		myCSVReader := CSVReader.NewReadCSV()

		inputData := []string{
			strconv.Itoa(int(myTrain.Age)),
			strconv.FormatFloat(myTrain.Bmi, 'f', -1, 64),
			strconv.Itoa(int(myTrain.Children)),
			strconv.FormatFloat(myTrain.Charges, 'f', -1, 64),
		}

		myCSVReader.WriteCSV("./CSVs/insurance.csv", inputData)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ResponseTrainInsurance{
			Message: "Data added succesfully!",
		})
	}
}