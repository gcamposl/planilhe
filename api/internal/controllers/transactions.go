package controllers

import (
	"api/internal/database"
	"api/internal/models"
	"api/internal/repositories"
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

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewTransactionRepository(db)
	transaction.ID, err = repo.Create(transaction)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, transaction)
}
