package handlers

import (
	"bank-poc/src/internal/domain"
	"bank-poc/src/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

type TransactionHandler struct {
	Service domain.TransactionService
}

func (h *TransactionHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", utils.ContentTypeJSON)

	var t domain.TransactionDto
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		apiResponse := new(domain.ApiResponse)
		apiResponse.Failure(err.Error())
		jsonResp, _ := json.Marshal(apiResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	res := h.Service.Deposit(t)
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

func (h *TransactionHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", utils.ContentTypeJSON)

	var t domain.TransactionDto
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		apiResponse := new(domain.ApiResponse)
		apiResponse.Failure(err.Error())
		jsonResp, _ := json.Marshal(apiResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	res := h.Service.Withdraw(t)
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

func (h *TransactionHandler) Balance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", utils.ContentTypeJSON)

	res := h.Service.Balance()
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

func (h *TransactionHandler) TransactionHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", utils.ContentTypeJSON)

	res := h.Service.TransactionHistory()
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
