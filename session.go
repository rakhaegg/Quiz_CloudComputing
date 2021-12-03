// json.go
package main

import (
	"context"
	"database/sql"
	"example/hello/model"
	"fmt"
	"log"

	"github.com/gorilla/mux"

	"encoding/json"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func MySql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@(localhost:3306)/dbname?parseTime=true")

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, table_check := db.Query("SELECT*FROM toko;")

	if table_check == nil {

	} else {
		query := `
            CREATE TABLE toko (
                id INT ,
                restaurant_id TEXT NOT NULL,
                name TEXT NOT NULL,
                description TEXT NOT NULL,
				pictureID TEXT NOT NULL,
                city TEXT NOT NULL,
                rating TEXT NOT NULL,
                PRIMARY KEY (id)
            );`
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
	return db, nil

}
func getContent(ctx context.Context) ([]model.Toko, error) {

	var tokos []model.Toko

	db, err := MySql()

	if err != nil {
		log.Fatal("Error Coneect Database", err)
	}
	queryText := fmt.Sprintf("SELECT * FROM %v ", "toko")

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var toko model.Toko

		if err = rowQuery.Scan(
			&toko.ID,
			&toko.Restaurant_id,
			&toko.Name,
			&toko.Description,
			&toko.City,
			&toko.Rating,
			&toko.PictureID); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		tokos = append(tokos, toko)
	}
	return tokos, nil

}

func Inset(ctx context.Context, toko model.Toko) error {
	db, err := MySql()

	if err != nil {
		log.Fatal("Cant connect to MySql", err)
	}
	queryText := fmt.Sprintf("INSERT INTO toko (id , restaurant_id, name, description , pictureID ,  city , rating) values( %v ,'%v','%v','%v','%v','%v','%v')",
		toko.ID,
		toko.Restaurant_id,
		toko.Name,
		toko.Description,
		toko.PictureID,
		toko.City,
		toko.Rating,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	ubahkeByte, err := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "error om", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(ubahkeByte))
}
func getToko(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		tokos, err := getContent(ctx)

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		ResponseJSON(w, tokos, http.StatusOK)
		return
	}

}

func getTokobyID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// paramas := mux.Vars(r)

	// for _, item := range tokos {
	// 	if item.Restaurant_id == paramas["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(&Toko{})
}

func createToko(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	// var toko Toko

	// body, err := ioutil.ReadAll(r.Body)

	// // toko.Restaurant_id = strconv.Itoa(rand.Intn(1000000))

	// toko
	// tokos = append(tokos, toko)

	// json.NewEncoder(w).Encode(toko)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// id, err := result.LastInsertId()
	// fmt.Println(id)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var toko model.Toko

	if err := json.NewDecoder(r.Body).Decode(&toko); err != nil {
		fmt.Println(err)
		ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := Inset(ctx, toko); err != nil {
		fmt.Println(err)
		ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	ResponseJSON(w, res, http.StatusCreated)
	return

}

func updateToko(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "apllication/json")
	// params := mux.Vars(r)

	// for index, item := range tokos {
	// 	if item.Restaurant_id == params["id"] {
	// 		tokos = append(tokos[:index], tokos[index+1:]...)

	// 		var toko Toko
	// 		_ = json.NewDecoder(r.Body).Decode(&toko)
	// 		toko.Restaurant_id = params["id"]
	// 		tokos = append(tokos, toko)

	// 		json.NewEncoder(w).Encode(toko)
	// 		return
	// 	}
	// }
}
func deleteToko(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	// params := mux.Vars(r)

	// for index, item := range tokos {
	// 	if item.Restaurant_id == params["id"] {
	// 		//append(tokos[:index]) artinya sebelum index yang akan dihapus
	// 		//tokos[index+1]... artinya sebelum index ditambahkan
	// 		tokos = append(tokos[:index], tokos[index+1:]...)
	// 		break
	// 	}

	// }
	// json.NewEncoder(w).Encode(tokos)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/toko", getToko).Methods("GET")
	r.HandleFunc("/api/toko/{id}", getTokobyID).Methods("GET")
	r.HandleFunc("/api/toko/create", createToko).Methods("POST")
	r.HandleFunc("/api/toko/{id}", updateToko).Methods("PUT")
	r.HandleFunc("/api/toko/{id}", deleteToko).Methods("DELETE")

	http.ListenAndServe(":8081", r)
}
