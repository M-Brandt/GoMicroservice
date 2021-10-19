package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Michael",
		Price: 1.0,
		SKU:   "b-b-b",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
