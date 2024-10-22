package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/employee", employeesHandler)
	log.Fatal(http.ListenAndServe(":8787", nil))
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {

}
