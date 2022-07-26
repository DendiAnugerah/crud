package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DendiAnugerah/crud/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil{
		log.Fatal(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the Book table
	if result := h.DB.Create(&book); result.Error != nil{
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}