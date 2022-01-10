package product

import (
	"errors"
	"net/http"
	"time"

	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/internal/handlers"
	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/internal/repository/adapter"
	HttpStatus "github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/utils/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Handler struct {
	handlers.Interface
	Controller product.Interface
	Rules      rules.Interface
}

func NewHnadler(repository adapter.Interface) handler.Interface {
	return &Handler{
		Controller: product.NewController(repository),
		Rules:      Rulesproduct.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "ID") != "" {
		h.GetOne(w, r)
	} else {
		h.GetAll(w, r)
	}
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}
	response, err := h.Controller.ListOne(ID)
	if err != nil {
		HttpStatus.StatusIntervalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)

}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := h.Controller.ListAll()
	if err != nil {
		HttpStatus.StatusIntervalServerError(w, r, err)
		return
	}
	HttpStatus.StatusOK(w, r, response)

}

// the basic job of a handler is to call a controller funciton

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {

	productBody, err := h.getBodyAndValidate(r, uuid.Nil)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}
	ID, err := h.Controller.Create(productBody)
	if err != nil {
		HttpStatus.StatusIntervalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, map[string]interface{}{"id": ID.strin()})

}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}
	productBody, err := h.getBodyAndValidate(r, ID)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	if err := h.Controller.Update(ID, productBody); err != nil {
		HttpStatus.StatusIntervalServerError(w, r, err)
		return
	}

	HttpStatus.StatusNoContent(w, r)

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	if err := h.Controller.Remove(ID); err != nil {
		HttpStatus.StatusIntervalServerError(w, r, err)
		return
	}
	HttpStatus.StatusNoContent(w, r)

}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) getBodyAndValidate(r *http.Request, ID uuid.UUID) (*EntityProduct.product, error) {

	// take the body
	productBody := &EntityProduct.Product{}

	// convert the body into struct
	body, err := h.Rules.ConvertIOReaderTOStruct(r.Body, productBody)
	if err != nil {
		return &EntityProduct.Product{}, errors.New("body is required")
	}

	productParsed, err := EntityProduct.InterfaceToModel(body)
	if err != nil {
		return &EntityProduct.Product{}, errors.New("error on converting body to model")
	}

	setDefaultValues(productParsed, ID)

	return productParsed, h.Rule.Validate(productParsed)

}

func setDefaultValues(product *EntityProduct.Product, ID uuid.UUID) {

	product.UpdatedAt = time.Now()
	if ID == uuid.Nil {
		product.ID = uuid.New()
		product.CreatedAt = time.Now()
	} else {
		product.ID = ID
	}
}
