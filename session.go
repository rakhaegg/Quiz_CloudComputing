// json.go
package main

import (
	"github.com/gorilla/mux"

	"encoding/json"
	"fmt"
	"net/http"
)

var tokos []Toko

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

type Toko struct {
	Restaurant_id string `json:"restaurant_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	PictureID     string `json:"pictureID"`
	City          string `json:"city"`
	Rating        string `json:"rating"`
}
type Makanan struct {
	Restaurant_id string `json:"restaurant_id"`
	Picture       string `json:"picture"`
	Description   string `json:"description"`
	Name          string `json:"name"`
	Harga         int    `json:"harga"`
}
type Minuman struct {
	Restaurant_id string `json:"restaurant_id"`
	Picture       string `json:"picture"`
	Description   string `json:"description"`
	Name          string `json:"name"`
	Harga         int    `json:"harga"`
}

func getToko(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokos)

}

func getTokobyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	paramas := mux.Vars(r)

	for _, item := range tokos {
		if item.Restaurant_id == paramas["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Toko{})
}

func createToko(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var toko Toko

	_ = json.NewDecoder(r.Body).Decode(&toko)

	// toko.Restaurant_id = strconv.Itoa(rand.Intn(1000000))
	tokos = append(tokos, toko)

	json.NewEncoder(w).Encode(toko)

}

func updateToko(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	params := mux.Vars(r)

	for index, item := range tokos {
		if item.Restaurant_id == params["id"] {
			tokos = append(tokos[:index], tokos[index+1:]...)

			var toko Toko
			_ = json.NewDecoder(r.Body).Decode(&toko)
			toko.Restaurant_id = params["id"]
			tokos = append(tokos, toko)

			json.NewEncoder(w).Encode(toko)
			return
		}
	}
}
func deleteToko(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range tokos {
		if item.Restaurant_id == params["id"] {
			//append(tokos[:index]) artinya sebelum index yang akan dihapus
			//tokos[index+1]... artinya sebelum index ditambahkan
			tokos = append(tokos[:index], tokos[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(tokos)
}

func main() {

	r := mux.NewRouter()

	//
	tokos = append(tokos, Toko{Restaurant_id: "1", Name: "Toko Baru", Description: "This is Description",
		PictureID: "This is Picture ID", City: "Malang", Rating: "4",
	})
	tokos = append(tokos, Toko{Restaurant_id: "2", Name: "Toko Baru", Description: "This is Description",
		PictureID: "This is Picture ID", City: "Malang", Rating: "3",
	})

	//

	r.HandleFunc("/api/toko", getToko).Methods("GET")
	r.HandleFunc("/api/toko/{id}", getTokobyID).Methods("GET")
	r.HandleFunc("/api/toko", createToko).Methods("POST")
	r.HandleFunc("/api/toko/{id}", updateToko).Methods("PUT")
	r.HandleFunc("/api/toko/{id}", deleteToko).Methods("DELETE")

	///
	r.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	r.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Println(peter)
		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":8080", r)
}
