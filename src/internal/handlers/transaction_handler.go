package handlers

import (
	"bank-poc/src/internal/domain"
	"bank-poc/src/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

type TransactionHandler struct {
	TransactionService domain.TransactionService
}

func (handler *TransactionHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", utils.ContentTypeJSON)

	var trnx domain.TransactionDto
	err := json.NewDecoder(r.Body).Decode(&trnx)
	if err != nil {
		apiResponse := new(domain.ApiResponse)
		apiResponse.Failure(err.Error())
		jsonResp, _ := json.Marshal(apiResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	res := handler.TransactionService.Deposit(&trnx)
	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
	}

	if !res.Successful || res.Data == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func (handler *TransactionHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", utils.ContentTypeJSON)

	var trnx domain.TransactionDto
	err := json.NewDecoder(r.Body).Decode(&trnx)
	if err != nil {
		apiResponse := new(domain.ApiResponse)
		apiResponse.Failure(err.Error())
		jsonResp, _ := json.Marshal(apiResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	res := handler.TransactionService.Withdraw(&trnx)
	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
	}

	if !res.Successful || res.Data == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func (handler *TransactionHandler) Balance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", utils.ContentTypeJSON)

	res := handler.TransactionService.Balance()
	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
	}

	if !res.Successful || res.Data == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func (handler *TransactionHandler) TransactionHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", utils.ContentTypeJSON)

	res := handler.TransactionService.TransactionHistory()
	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
	}

	if !res.Successful || res.Data == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func (handler *TransactionHandler) Rollback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", utils.ContentTypeJSON)

	res := handler.TransactionService.Rollback()
	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
	}

	if !res.Successful || res.Data == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
