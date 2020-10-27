package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:"hrithik",
		Price: 12,
		SKU: "asd-asd-asd",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
