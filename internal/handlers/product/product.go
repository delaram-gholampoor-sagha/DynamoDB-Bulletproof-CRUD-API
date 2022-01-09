package product

import (
	"net/http"

	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/internal/handlers"
	"github.com/go-playground/locales/et"
)



type Handler struct {
	handlers.Interface
	Controller product.Interface
	Rules rules.Interface
}

func NewHnadler(repository adapter.Interface) handler.Interface {
	return &Handler{
		Controller: product.NewController(repository) ,
		Rules : Rulesproduct.NewRules(),
	}
}

func Get() {

}

func GetOne() {
	
}

func GetAll() {
	
}

func (h *Handler) Post(w http.ResponseWriter , r *http.Request) {

	productBody , err := h.getBodyAndValidate(r , uuid.Nil)
	if err != nil {
		HttpStatus.StatusBadRequest(w ,r , err)
		return
	}
	ID , err := h.Controller.Create(productBody)
	if err != nil {
		HttpStatus.StatusInternalServerError( w , r , err) 
		return
	}

	HttpStatus.StatusOK(w ,r , map[string]interface{}{"id" : ID.strin() })
	
}

func Put() {
	
}

func Delete() {
	
}

func Options() {
	
}


func (h *Handler) getBodyAndValidate() () {

}