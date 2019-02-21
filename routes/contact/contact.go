package contact

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dsummers91/go-bmdc/database"
)

type postContactDetails struct {
	Username string `json:"username,omitempty"`
	Message  string `json:"message,omitempty"`
}

func PostContactHandler(w http.ResponseWriter, r *http.Request) {
	var data postContactDetails
	json.NewDecoder(r.Body).Decode(&data)

	collection, context, cancel := database.Collection("contact")
	defer cancel()

	response, err := collection.InsertOne(context,
		data,
	)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
