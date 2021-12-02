package main

import (
	"fmt"

	//Package yang mempunyai fungsi untuk HTTP Protocol
	"net/http"

	"github.com/gorilla/mux"

	"database/sql"

	"log"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func routing() {

	//Kode dibawah untuk request router
	//ROUTER

	r := mux.NewRouter()

	//Registering a Request Handler

	/*
		Kode dibawah untuk
		Pertama membuat Handler untuk menerima koneksi HTTP masuk dari browser
		klien HTTP , atau API

		http.ResponseWrite adalah parameter tempat dimana tulisan anada di respon dimana
		http.Request adalah paramater informasi HTTP Request termasuk URL dan HEADER FIELDS
	*/
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello , you have requested: %s\n", r.URL.Path)

	})

	//Request handler URL
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {

		//Get Data
		//Fungsi dibawah untuk mendapat data dari segmene

		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, " %s and %s ", title, page)
	})

	ShowTitle := func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		title := vars["title"]

		fmt.Fprintf(w, " Title : %s", title)

	}
	r.HandleFunc("/books/{title}", ShowTitle).Methods("POST")

	//Serving Static Asset
	//Kode dibawah untuk pointing URL PATH

	//fs := http.FileServer(http.Dir("static/"))

	//Kode dibawah untuk jalur nama direktori tempat file berada

	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	//Untuk Server HTTP mendegarkan port untuk meneruskan koneksi
	//Port 80 adalah port lalu lintas HTTP
	//Mendengarkan koneksi port 80 dan meneruskan ke browser ke localhsot dan pemintaan
	http.ListenAndServe(":80", r)

}
func assetFile() {

	fs := http.FileServer(http.Dir("assets/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
func connectSQL() {
	//Connection to MYSQL DATABASE

	db, err := sql.Open("mysql", "root:@(localhost:3306)/dbname?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	//Create FIRST DATABASE Table
	{
		query := `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

		//Ekesekusi SQL QUERY , and check error unutuk mengetahui tidak ada error
		if _, err = db.Exec(query); err != nil {
			log.Fatal(err)
		}

	}

	{ //Insert a new user
		username := "rakha"
		password := "elang"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		//Untuk mendapatkan id baru yang dibuat
		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{

		//Deklarasi variables untuk menyimpan data dari database
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id , username , password , created_at FROM users WHERE id = ?"

		//Query Dtabase untuk mendapatkan nilai dan assign to variable
		//QueryRow untuk specific row
		//Query untuk semua row
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)

	}
	{

		//Deklasi variabel untuk menyimpan data dan menyimpan data ke database
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id , username , password , created_at FROM users`)

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user

		for rows.Next() {
			var u user

			//Membaca data
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)

			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)

		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v", users)
	}
	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}

	}
}
