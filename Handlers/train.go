package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/F-Dupraz/LinearRegressionApp/server"
	CSVReader "github.com/F-Dupraz/LinearRegressionApp/csvReader"
)

// Insurance REQUEST
type RequestTrainInsurance struct {
	Age      string `json:"age"`
	Bmi      string `json:"bmi"`
	Children string `json:"children"`
	Charges  string `json:"charges"`
}

// Heart REQUEST
type RequestTrainHeart struct {
	Age      string `json:"age"`
	Sex      string `json:"sex"`
	Cp       string `json:"cp"`
	Trestbps string `json:"trestbps"`
	Chol     string `json:"chol"`
	Fbs      string `json:"fbs"`
	Restecg  string `json:"restecg"`
	Thalach  string `json:"thalach"`
	Exang    string `json:"exang"`
	Oldpeak  string `json:"oldpeak"`
	Slope    string `json:"slope"`
	Ca       string `json:"ca"`
	Thal     string `json:"thal"`
	Target   string `json:"target"`
}

// Candy REQUEST
type RequestTrainCandy struct {
	Chocolate        string `json:"chocolate"`
	Fruity           string `json:"fruity"`
	Caramel          string `json:"caramel"`
	Peanutyalmondy   string `json:"peanutyalmondy"`
	Nougat           string `json:"nougat"`
	Crispedricewafer string `json:"crispedricewafer"`
	Hard             string `json:"hard"`
	Bar              string `json:"bar"`
	Pluribus         string `json:"pluribus"`
	SugarPercent     string `json:"sugarpercent"`
	PricePercent     string `json:"pricepercent"`
  WinPercent       string `json:"winpercent"`
}

// General RESPONSE
type ResponseTrain struct {
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
			myTrain.Age,
			myTrain.Bmi,
			myTrain.Children,
			myTrain.Charges,
		}

		myCSVReader.WriteCSV("./CSVs/insurance.csv", inputData)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ResponseTrain{
			Message: "Data added succesfully!",
		})
	}
}

func TrainHeart(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var myTrain = RequestTrainHeart{}
		err := json.NewDecoder(r.Body).Decode(&myTrain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		myCSVReader := CSVReader.NewReadCSV()

		inputData := []string{
			myTrain.Age,
			myTrain.Sex,
			myTrain.Cp,
			myTrain.Trestbps,
			myTrain.Chol,
			myTrain.Fbs,
			myTrain.Restecg,
			myTrain.Thalach,
			myTrain.Exang,
			myTrain.Oldpeak,
			myTrain.Slope,
			myTrain.Ca,
			myTrain.Thal,
			myTrain.Target,
		}

		myCSVReader.WriteCSV("./CSVs/heart.csv", inputData)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ResponseTrain{
			Message: "Data added succesfully!",
		})
	}
}

func TrainCandy(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var myTrain = RequestTrainCandy{}
		err := json.NewDecoder(r.Body).Decode(&myTrain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		myCSVReader := CSVReader.NewReadCSV()

		inputData := []string{
			myTrain.Chocolate,
			myTrain.Fruity,
			myTrain.Caramel,
			myTrain.Peanutyalmondy,
			myTrain.Nougat,
			myTrain.Crispedricewafer,
			myTrain.Hard,
			myTrain.Bar,
			myTrain.Pluribus,
			myTrain.SugarPercent,
			myTrain.PricePercent,
			myTrain.WinPercent,
		}

		myCSVReader.WriteCSV("./CSVs/candy.csv", inputData)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ResponseTrain{
			Message: "Data added succesfully!",
		})
	}
}