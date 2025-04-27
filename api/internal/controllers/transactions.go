package controllers

import (
	"api/internal/database"
	"api/internal/models"
	"api/internal/repositories"
	"api/internal/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewTransactionRepository(db)
	transactions, err := repo.Find(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, transactions)
}
