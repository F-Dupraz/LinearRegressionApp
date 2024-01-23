package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/F-Dupraz/LinearRegressionApp/server"
	LR "github.com/F-Dupraz/LinearRegressionApp/LinearRegression"
	CSVReader "github.com/F-Dupraz/LinearRegressionApp/csvReader"
)

// Insurance REQUEST and RESPONSE
type RequestPredictInsurance struct {
	Age      uint    `json:"age"`
	Bmi      float64 `json:"bmi"`
	Children uint    `json:"children"`
}

type ResponsePredictInsurance struct {
	Charges float64 `json:"charges"`
}

// Heart REQUEST and RESPONSE
type RequestPredictHeart struct {
	Age      uint    `json:"age"`
	Sex      uint    `json:"sex"`      // 1 -> Male, 0 -> Female
	Cp       uint    `json:"cp"`
	Trestbps uint    `json:"trestbps"`
	Chol     uint    `json:"chol"`
	Fbs      uint    `json:"fbs"`
	Restecg  uint    `json:"restecg"`
	Thalach  uint    `json:"thalach"`
	Exang    uint    `json:"exang"`
	Oldpeak  float64 `json:"oldpeak"`
	Slope    uint    `json:"slope"`
	Ca       uint    `json:"ca"`
	Thal     uint    `json:"thal"`
}

type ResponsePredictHeart struct {
	Target float64 `json:"target"`
}

// Candy REQUEST and RESPONSE
type RequestPredictCandy struct {
	Chocolate        uint    `json:"chocolate"`
	Fruity           uint    `json:"fruity"`
	Caramel          uint    `json:"caramel"`
	Peanutyalmondy   uint    `json:"peanutyalmondy"`
	Nougat           uint    `json:"nougat"`
	Crispedricewafer uint    `json:"crispedricewafer"`
	Hard             uint    `json:"hard"`
	Bar              uint    `json:"bar"`
	Pluribus         uint    `json:"pluribus"`
	SugarPercent     float64 `json:"sugarpercent"`
	PricePercent     float64 `json:"pricepercent"`
}

type ResponsePredictCandy struct {
  WinPercent float64 `json:"winpercent"`
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

		inputData := []float64{
			float64(myPredict.Age),
			myPredict.Bmi,
			float64(myPredict.Children),
		}

		prediction := linearRegression.Predict(inputData)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponsePredictInsurance{
			Charges: float64(prediction),
		})
	}
}

func PredictHeart(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var myPredict = RequestPredictHeart{}
		err := json.NewDecoder(r.Body).Decode(&myPredict)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Initialization of the linear regression and CSV reader classes.
		linearRegression := LR.NewLinearRegression()
		myCSVReader := CSVReader.NewReadCSV()

		predictionData, err := myCSVReader.GetData("./DB/heart.csv")
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

		inputData := []float64{
			float64(myPredict.Age),
			float64(myPredict.Sex),
			float64(myPredict.Cp),
			float64(myPredict.Trestbps),
			float64(myPredict.Chol),
			float64(myPredict.Fbs),
			float64(myPredict.Restecg),
			float64(myPredict.Thalach),
			float64(myPredict.Exang),
			myPredict.Oldpeak,
			float64(myPredict.Slope),
			float64(myPredict.Ca),
			float64(myPredict.Thal),
		}

		prediction := linearRegression.Predict(inputData)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponsePredictHeart{
			Target: float64(prediction),
		})
	}
}

func PredictCandy(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var myPredict = RequestPredictCandy{}
		err := json.NewDecoder(r.Body).Decode(&myPredict)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Initialization of the linear regression and CSV reader classes.
		linearRegression := LR.NewLinearRegression()
		myCSVReader := CSVReader.NewReadCSV()

		predictionData, err := myCSVReader.GetData("./DB/candy.csv")
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

		inputData := []float64{
			float64(myPredict.Chocolate),
			float64(myPredict.Fruity),
			float64(myPredict.Caramel),
			float64(myPredict.Peanutyalmondy),
			float64(myPredict.Nougat),
			float64(myPredict.Crispedricewafer),
			float64(myPredict.Hard),
			float64(myPredict.Bar),
			float64(myPredict.Pluribus),
			myPredict.SugarPercent,
			myPredict.PricePercent,
		}

		prediction := linearRegression.Predict(inputData)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponsePredictCandy{
			WinPercent: float64(prediction),
		})
	}
}
