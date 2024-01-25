package main

import (
	"math"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/F-Dupraz/LinearRegressionApp/handlers"
)

func assertEquals(t *testing.T, expected, actual float64, tolerance float64, message string) {
	if math.Abs(expected-actual) > tolerance {
		t.Errorf("%s. Esperado: %f, Obtenido: %f", message, expected, actual)
	}
}

func TestPredictInsurance(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handlers.PredictInsurance(nil)))
	defer ts.Close()

	requestBody := handlers.RequestPredictInsurance{
		Age: 30,
		Bmi: 26.4,
		Children: 2,
	}

	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Errorf("Error al serializar la solicitud: %v", err)
		return
	}

	req, err := http.NewRequest("POST", ts.URL+"/predict/insurance", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		t.Errorf("Error al crear la solicitud HTTP: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Error al enviar la solicitud HTTP: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado código de estado 200, pero obtuvo %d", resp.StatusCode)
		return
	}

	var responseBody handlers.ResponsePredictInsurance

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("Error al decodificar la respuesta JSON: %v", err)
		return
	}

	// Comparación de valores flotantes con tolerancia
	assertEquals(t, 11044.56551, responseBody.Charges, 0.00001, "Esperado prediccion de 11044.56551")
}

func TestPredictHeart(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handlers.PredictHeart(nil)))
	defer ts.Close()

	requestBody := handlers.RequestPredictHeart{
		Age: 30,
		Sex: 1,
		Cp: 0,
		Trestbps: 110,
		Chol: 300,
		Fbs: 1,
		Restecg: 1,
		Thalach: 125,
		Exang: 0,
		Oldpeak: 2,
		Slope: 2,
		Ca: 0,
		Thal: 3,
	}

	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Errorf("Error al serializar la solicitud: %v", err)
		return
	}

	req, err := http.NewRequest("POST", ts.URL+"/predict/heart", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		t.Errorf("Error al crear la solicitud HTTP: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Error al enviar la solicitud HTTP: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado código de estado 200, pero obtuvo %d", resp.StatusCode)
		return
	}

	var responseBody handlers.ResponsePredictHeart

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("Error al decodificar la respuesta JSON: %v", err)
		return
	}

	// Comparación de valores flotantes con tolerancia
	assertEquals(t, 0.383427, responseBody.Target, 0.00001, "Esperado prediccion de 11044.56551")
}

func TestPredictCandy(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handlers.PredictCandy(nil)))
	defer ts.Close()

	requestBody := handlers.RequestPredictCandy{
		Chocolate: 1,
		Fruity: 0,
		Caramel: 0,
		Peanutyalmondy: 1,
		Nougat: 0,
		Crispedricewafer: 0,
		Hard: 0,
		Bar: 1,
		Pluribus: 0,
		SugarPercent: 0.72,
		PricePercent: 0.55,
	}

	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Errorf("Error al serializar la solicitud: %v", err)
		return
	}

	req, err := http.NewRequest("POST", ts.URL+"/predict/candy", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		t.Errorf("Error al crear la solicitud HTTP: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Error al enviar la solicitud HTTP: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado código de estado 200, pero obtuvo %d", resp.StatusCode)
		return
	}

	var responseBody handlers.ResponsePredictCandy

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("Error al decodificar la respuesta JSON: %v", err)
		return
	}

	// Comparación de valores flotantes con tolerancia
	assertEquals(t, 43.73594982, responseBody.WinPercent, 0.00001, "Esperado prediccion de 11044.56551")
}
