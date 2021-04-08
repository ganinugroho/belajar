package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ganinugroho/belajar/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		p.l.Println(r.Method)
		p.l.Println(r.URL.Path)
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		p.l.Println(len(g))
		if len(g) != 1 {
			http.Error(rw, "Invalid URI 1", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI 2", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI 3", http.StatusBadRequest)
		}
		p.l.Printf("Got id %#v", id)
		p.editProducts(id, rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal data product", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusInternalServerError)
	}
	p.l.Printf("Product: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) editProducts(_id int, rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusInternalServerError)
	}
	p.l.Printf("Product: %#v", prod)
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
