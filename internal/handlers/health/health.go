package health

import (
	"errors"
	"net/http"

	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/internal/handlers"
	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/internal/repository/adapter"
)

type Handler struct {
	handlers.Interface
	Repository adapter.Interface
}

func NewHnadler(repository adapter.Interface) handler.Interface {
	return &Handler{
		Repository: repository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if !h.Repository.Health() {
		httpStatus.StatusInternalServerError(w, r, errors.New("Relational database not alive"))
		return
	}

	HttpStatus.StatusOK(w, r, "Service Ok")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
       HttpStatus.StatusMethodNotAllowed(w ,r)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w ,r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w ,r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w ,r)
}
