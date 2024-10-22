package main

import (
	"encoding/json"
	"log"
	"net/http"
	"persogo/controllers"
	"persogo/database"
)

func main() {
	http.HandleFunc("/employee", employeesHandler)
	log.Fatal(http.ListenAndServe(":8787", nil))
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("employees")

	if r.Method == http.MethodGet {
		employees, err := controllers.GetEmployees(collection)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
