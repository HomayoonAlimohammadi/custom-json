package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Name  string
	Price float64
}

type Cargo struct {
	Products []Product
}

func (c *Cargo) UnmarshalJSON(d []byte) error {
	type temp struct {
		Products []string  `json:"products"`
		Prices   []float64 `json:"prices"`
	}

	tmp := &temp{}
	err := json.Unmarshal(d, tmp)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON in temp: %w", err)
	}

	fmt.Printf("first temp: %+v\n", tmp)

	tmp2 := make(map[string]any)
	_ = json.Unmarshal(d, &tmp2)
	fmt.Printf("second temp: %+v\n", tmp2)

	if len(tmp.Prices) != len(tmp.Products) {
		return fmt.Errorf("length of products (%d) and prices (%d) does not match", len(tmp.Products), len(tmp.Prices))
	}

	for i := 0; i < len(tmp.Products); i++ {
		c.Products = append(c.Products, Product{Name: tmp.Products[i], Price: tmp.Prices[i]})
	}

	return nil
}

func main() {
	d := GetProductPrices()

	cargo := &Cargo{}
	_ = json.Unmarshal(d, cargo)

	b, _ := json.MarshalIndent(cargo, "", "\t")
	fmt.Println(string(b))
}

func GetProductPrices() []byte {
	d, _ := os.ReadFile("products.json")
	return d
}
