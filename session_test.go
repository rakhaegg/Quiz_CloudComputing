package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TextConnection(t *testing.T) {

	//
	req, err := http.NewRequest("GET", "/api/toko", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getToko)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// // Check the response body is what we expect.
	// expected := `[{"restaurant_id":"1","name":"Toko Baru","description":"This is Description","pictureID":"This is Picture ID","city":"Malang","rating":"4"},{"restaurant_id":"2","name":"Toko Baru","description":"This is Description","pictureID":"This is Picture ID","city":"Malang","rating":"3"}]`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }

}

func TestCreateToko(t *testing.T) {

	var jsonStr = []byte(`{
		"restaurant_id" : "3",
		"name" : "res name" ,
		"description" : "res desc" ,
		"pictureID" : "res pic" ,
		"city" : "res pic",
		"rating" : "res rating"
	}`)

	req, err := http.NewRequest("POST", "/api/toko/create", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createToko)
	handler.ServeHTTP(rr, req)

	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wront status code : %v want %v ", status, http.StatusOK)
	// }

	// var m map[string]interface{}
	// json.Unmarshal(response.Body.Bytes(), &m)

	expected := `{"status":"Succesfully"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body :  %v want %v", rr.Body.String(), expected)
	}

}
