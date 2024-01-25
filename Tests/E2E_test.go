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

// --------- ARREGLAR EL ERROR DE LA REQUEST ---------- //
// E2E_test.go:39: Esperado código de estado 200, pero obtuvo 400
// E2E_test.go:45: invalid character 'o' looking for beginning of value

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
