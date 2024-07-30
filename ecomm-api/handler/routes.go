package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

var r *chi.Mux

func RegisterRoutes(handler *handler) *chi.Mux {
	r = chi.NewRouter()
	tokenMaker := handler.TokenMaker
	r.Route("/products", func(r chi.Router) {
		r.With(GetAdminMiddlewareFunc(tokenMaker)).Post("/", handler.createProduct)
		r.Get("/", handler.listProducts)
		r.Get("/{id}", handler.getProduct)
		r.With(GetAdminMiddlewareFunc(tokenMaker)).Patch("/{id}", handler.updateProduct)
		r.With(GetAdminMiddlewareFunc(tokenMaker)).Delete("/{id}", handler.deleteProduct)
	})

	r.Route("/myorder", func(r chi.Router) {
		r.Use(GetAuthMiddlewareFunc(tokenMaker))
		r.Get("/", handler.getOrder)
	})

	r.Route("/orders", func(r chi.Router) {
		r.With(GetAuthMiddlewareFunc(tokenMaker)).Post("/", handler.createOrder)
		r.With(GetAdminMiddlewareFunc(tokenMaker)).Get("/", handler.listOrders)
		r.With(GetAuthMiddlewareFunc(tokenMaker)).Delete("/{id}", handler.deleteOrder)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.createUser)
		r.Post("/login", handler.loginUser)

		r.With(GetAdminMiddlewareFunc(tokenMaker)).Get("/", handler.listUsers)
		r.With(GetAdminMiddlewareFunc(tokenMaker)).Delete("/{id}", handler.deleteUser)

		r.With(GetAuthMiddlewareFunc(tokenMaker)).Patch("/", handler.updateUser)
		r.With(GetAuthMiddlewareFunc(tokenMaker)).Post("/logout", handler.logoutUser)
	})

	r.Group(func(r chi.Router) {
		r.Use(GetAuthMiddlewareFunc(tokenMaker))
		r.Route("/tokens", func(r chi.Router) {
			r.Post("/renew", handler.renewAccessToken)
			r.Post("/revoke", handler.revokeSession)
		})
	})

	return r
}

func Start(addr string) error {
	return http.ListenAndServe(addr, r)
}
