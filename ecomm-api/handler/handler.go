package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sir-radar/go-ecom/ecomm-api/server"
	"github.com/sir-radar/go-ecom/ecomm-api/storer"
)

type Handler struct {
	ctx    context.Context
	server *server.Server
}

func NewHandler(srv *server.Server) *Handler {
	return &Handler{
		ctx:    context.Background(),
		server: srv,
	}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p ProductReq

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	product, err := h.server.CreateProduct(h.ctx, toStorerProduct(p))

	if err != nil {
		http.Error(w, "error creating product", http.StatusInternalServerError)
		return
	}

	res := toProductRes(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}

func toStorerProduct(p ProductReq) *storer.Product {
	return &storer.Product{
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
	}
}

func toProductRes(product *storer.Product) ProductRes {
	return ProductRes{
		ID:           product.ID,
		Name:         product.Name,
		Image:        product.Image,
		Category:     product.Category,
		Description:  product.Description,
		Rating:       product.Rating,
		NumReviews:   product.NumReviews,
		Price:        product.Price,
		CountInStock: product.CountInStock,
		UpdatedAt:    product.UpdatedAt,
		CreatedAt:    product.CreatedAt,
	}
}
