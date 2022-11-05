package handlers

import (
	"encoding/json"
	"github.com/LittleMikle/rest2.git/pkg/mocks"
	"github.com/LittleMikle/rest2.git/pkg/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	for index, book := range mocks.Books {
		if book.ID == id {
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author

			mocks.Books[index] = book

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
