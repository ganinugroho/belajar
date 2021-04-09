package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/ganinugroho/belajar/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal data product", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}

func (p *Products) EditProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}
	p.l.Printf("success to convert id %#v", _id)
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	ep, errSave := data.EditProduct(_id, prod)
	if errSave == data.ProductNotFound {
		http.Error(rw, data.ProductNotFound.Error(), http.StatusNotFound)
		return
	}
	if errSave != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	ep.ToJSONSingle(rw)
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Error deserialize product", http.StatusBadRequest)
			return
		}
		p.l.Println("success serialize product")
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
