package routes

import (
	"github.com/go-chi/chi"
)


type Router struct {
	config *Config
	router *chi.Mux
}



func NewRouter() *Router{
	return &Router{
		config : NewConfig().SetTimeout(serviceConfig.GetConfig().timeout())
		router:  chi.NewRouter(),
	}
}


func (r *Router) setRoutes() *chi.Mux{

}


func (r *Router) setConfigRoutes() {

}


func RouterHealth() {

}

func RouterProduct() {

}

func EnableTimeout() {

}

func EnableCORS() {

}

func EnableRecover() {

}

func EnableRequestID() {

}

func EnableRealIP() {

}



