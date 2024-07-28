package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

var r *chi.Mux

func RegisterRoutes(handler *handler) *chi.Mux {
	r = chi.NewRouter()
	r.Route("/products", func(r chi.Router) {
		r.Post("/", handler.createProduct)
		r.Get("/", handler.listProducts)
		r.Get("/{id}", handler.getProduct)
		r.Patch("/{id}", handler.updateProduct)
		r.Delete("/{id}", handler.deleteProduct)
	})

	r.Route("/orders", func(r chi.Router) {
		r.Post("/", handler.createOrder)
		r.Get("/", handler.listOrders)
		r.Get("/{id}", handler.getOrder)
		r.Delete("/{id}", handler.deleteOrder)
	})
	return r
}

func Start(addr string) error {
	return http.ListenAndServe(addr, r)
}
