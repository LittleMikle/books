package handlers

import (
	"encoding/json"
	"github.com/LittleMikle/rest2.git/pkg/mocks"
	"github.com/LittleMikle/rest2.git/pkg/models"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	book.ID = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")

}
