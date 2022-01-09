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

func Post() {

}

func Put() {

}

func Delete() {

}

func Options() {

}
