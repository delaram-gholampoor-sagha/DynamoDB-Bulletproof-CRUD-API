package routes

import (
	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/internal/repository/adapter"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(serviceConfig.GetConfig().timeout()),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRoutes(respository adapter.Interface) *chi.Mux {
	r.setConfigRoutes()
	r.RouterHealth(respository)
	r.RouterProduct(respository)

	return r.router
}

func (r *Router) setConfigRoutes() {
	r.EnableCORS()
	r.EnableLogger()
	r.EnableTimeout()
	r.EnableRecover()
	r.EnableRequestID()
	r.EnableRealIP()

}

func (r *Router) RouterHealth(repository adapter.Interface) {

	handler := HealthHandler.newHandler(repository)

	r.router.Route("/health", func(route chi.Router) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/", handler.Put)
		route.Delete("/", handler.Delete)
		route.Options("/", handler.Options)

	})

}

func (r *Router) RouterProduct(repository adapter.Interface) {

	handler := ProductHandler.newHandler(repository)

	r.router.Route("/product", func(route chi.Router) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/{ID}", handler.Put)
		route.Delete("/{ID}", handler.Delete)
		route.Options("/", handler.Options)

	})

}

func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router) EnableCORS() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableRecover() *Router {
     r.router.Use(middleware.Recoverer)
	 return r
}

func (r *Router) EnableRequestID() *Router {
       r.router.Use(middleware.RequestID)
	   return r
}

func (r *Router) EnableRealIP() *Router {
     r.router.Use(middleware.RealIP)
	 return r
}
