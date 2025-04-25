package controllers

import (
	"api/internal/models"
	"api/internal/responses"
	"encoding/json"
	"io"
	"net/http"
)

// Create transaction in database
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var transaction models.Transactions
	if err := json.Unmarshal(body, &transaction); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if transaction.Validate() != nil {
		responses.Error(w, http.StatusBadRequest, transaction.Validate())
		return
	}

	//todo: continue...

	responses.JSON(w, http.StatusCreated, transaction)
}
