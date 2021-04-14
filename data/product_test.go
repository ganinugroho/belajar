package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "gani",
		Price: 1,
		SKU:   "abc-abc-abc",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
