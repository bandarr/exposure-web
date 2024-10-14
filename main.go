package main

import (
	"exposure-web/rfexposure"

	//"exposure-web/rfexposure"
	"fmt"
	"html/template"
	"log"
	"net/http"
	//_ "github.com/mattn/go-sqlite3"
)

// var db *sql.DB

// func initDB() {
// 	var err error
// 	db, err = sql.Open("sqlite3", "./form.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
//         "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
//         "name" TEXT,
//         "email" TEXT
//     );`

// 	_, err = db.Exec(createTableSQL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func formHandler(w http.ResponseWriter, r *http.Request) {
	data := rfexposure.TestStub()
	tmpl := template.Must(template.ParseFiles("form.html"))
	tmpl.Execute(w, data)
}

// func submitHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	name := r.FormValue("name")
// 	email := r.FormValue("email")

// 	insertSQL := `INSERT INTO users (name, email) VALUES (?, ?)`
// 	_, err := db.Exec(insertSQL, name, email)
// 	if err != nil {
// 		http.Error(w, "Unable to save data", http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "Form submitted successfully!")
// }

func main() {
	//	initDB()
	//	defer db.Close()

	http.HandleFunc("/", formHandler)
	//	http.HandleFunc("/submit", submitHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
