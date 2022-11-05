package handlers

import (
	"encoding/json"
	"github.com/LittleMikle/rest2.git/pkg/mocks"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, book := range mocks.Books {
		if book.ID == id {
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
