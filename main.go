package main

import (
	"exposure-web/rfexposure"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	data := rfexposure.TestStub()
	tmpl := template.Must(template.ParseFiles("form.html"))
	tmpl.Execute(w, data)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Uncomment and implement the following lines if database functionality is required.
	// name := r.FormValue("name")
	// email := r.FormValue("email")
	// insertSQL := `INSERT INTO users (name, email) VALUES (?, ?)`
	// _, err := db.Exec(insertSQL, name, email)
	// if err != nil {
	// 	http.Error(w, "Unable to save data", http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Fprintf(w, "Form submitted successfully!")
}

func main() {
	// Uncomment the following lines if database initialization is required.
	// initDB()
	// defer db.Close()

	http.HandleFunc("/", formHandler)
	http.HandleFunc("/submit", submitHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
