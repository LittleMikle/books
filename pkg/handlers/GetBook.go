package handlers

import (
	"encoding/json"
	"github.com/LittleMikle/rest2.git/pkg/mocks"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, book := range mocks.Books {
		if book.ID == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
