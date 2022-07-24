package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DendiAnugerah/crud/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}

func GetBook(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, book := range mocks.Books{
		if book.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(book)
			break
		}
	}
}