package main

import "encoding/json"

type Product struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Vendors []string `json:"vendors"`
}

func main() {
	str := `{ "name": "my product", "id": 1 }`
	product := Product{}
	product2 := Product{
		Id:      5,
		Name:    "Test product 2",
		Vendors: []string{"Test vendor 1", "Test vendor 2", "Test vendor 3"},
	}
	_ = json.Unmarshal([]byte(str), &product)
	res, _ := json.Marshal(product2)
	println(product.Id, "-", product.Name)
	println(string(res))
}
