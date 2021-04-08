package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) ToJSONSingle(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Tahu",
		Description: "Tahu bulat 500 an",
		Price:       500,
		SKU:         "0001",
		CreatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Tempe",
		Description: "Tempe mambu",
		Price:       200,
		SKU:         "0002",
		CreatedOn:   time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func EditProduct(_id int, p *Product) (*Product, error) {
	gp, pos, err := findProduct(_id)
	if err != nil {
		return nil, err
	}
	p.ID = gp.ID
	p.CreatedOn = gp.CreatedOn
	productList[pos] = p
	return p, nil
}

var ProductNotFound = fmt.Errorf("product not found")

func findProduct(_id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == _id {
			return p, i, nil
		}
	}
	return nil, -1, ProductNotFound
}

func getNextID() int {
	p := productList[len(productList)-1]
	return p.ID + 1
}
