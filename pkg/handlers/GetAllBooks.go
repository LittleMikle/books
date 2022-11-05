package handlers

import (
	"encoding/json"
	"github.com/LittleMikle/rest2.git/pkg/mocks"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}
