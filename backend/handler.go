package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type TODOHandler interface {
	CreateTODO(http.ResponseWriter, *http.Request)
	ListTODOs(http.ResponseWriter, *http.Request)
	GetTODO(http.ResponseWriter, *http.Request)
	UpdateTODO(http.ResponseWriter, *http.Request)
	DeleteTODO(http.ResponseWriter, *http.Request)
}

type todoHandler struct {
	todoRepo TODORepository
}

func NewTODOHandler(todoRepo TODORepository) TODOHandler {
	return &todoHandler{
		todoRepo: todoRepo,
	}
}

//
func (th todoHandler) CreateTODO(w http.ResponseWriter, r *http.Request) {
	// Start context
	ctx := r.Context()

	// Receive data format and decode
	var tm TODOModel
	err := json.NewDecoder(r.Body).Decode(&tm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ask domain layer and accept response data format
	todo, err := th.todoRepo.CreateTODO(ctx, &tm)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

//
func (th todoHandler) ListTODOs(w http.ResponseWriter, r *http.Request) {
	// Start context
	ctx := r.Context()

	// Receive data format and decode

	// Ask domain layer and accept response data format
	todos, err := th.todoRepo.ListTODOs(ctx)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

//
func (th todoHandler) GetTODO(w http.ResponseWriter, r *http.Request) {
	// Start context
	ctx := r.Context()

	// Receive data format and decode

	// Decode path param
	id := chi.URLParam(r, "id")

	// Ask domain layer and accept response data format
	todo, err := th.todoRepo.GetTODO(ctx, id)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

//
func (th todoHandler) UpdateTODO(w http.ResponseWriter, r *http.Request) {
	// Start context
	ctx := r.Context()

	// Receive data format and decode
	var tm TODOModel
	err := json.NewDecoder(r.Body).Decode(&tm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Decode path param
	id := chi.URLParam(r, "id")

	// Ask domain layer and accept response data format
	todo, statusCode, err := th.todoRepo.UpdateTODO(ctx, id, &tm)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Response
	switch statusCode {
	case 200:
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(todo); err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
	case 400:
		http.Error(w, "Bad Request", 400)
	case 500:
		http.Error(w, "Internal Server Error", 500)
	default:
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(todo); err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}
}

//
func (th todoHandler) DeleteTODO(w http.ResponseWriter, r *http.Request) {
	// Start context
	ctx := r.Context()

	// Receive data format and decode

	// Decode path param
	id := chi.URLParam(r, "id")

	// Ask domain layer and accept response data format
	statusCode, err := th.todoRepo.DeleteTODO(ctx, id)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Response
	switch statusCode {
	case 204:
		w.WriteHeader(http.StatusNoContent)
	case 404:
		http.Error(w, "Bad Request", 404)
	case 500:
		http.Error(w, "Internal Server Error", 500)
	default:
		w.WriteHeader(http.StatusNoContent)
	}

}
