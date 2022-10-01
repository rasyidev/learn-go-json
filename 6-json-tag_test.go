package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTagMarshal(t *testing.T) {
	product := Product{
		Id:       "P9230",
		Name:     "Macbook Pro Max",
		Price:    42000000,
		ImageUrl: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

/*
=== RUN   TestJsonTagMarshal
{"id":"P9230","name":"Macbook Pro Max","price":42000000,"image_url":"http://example.com/image.png"}
--- PASS: TestJsonTagMarshal (0.00s)
PASS
ok      learn-go-json   0.427s
*/

func TestJsonTagUnmarshal(t *testing.T) {
	productString := `{"id":"P9230","name":"Macbook Pro Max","price":42000000,"image_url":"http://example.com/image.png"}`

	product := Product{}
	err := json.Unmarshal([]byte(productString), &product)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.Price)
	fmt.Println(product.ImageUrl)
}

/*
=== RUN   TestJsonTagUnmarshal
P9230
Macbook Pro Max
42000000
http://example.com/image.png
--- PASS: TestJsonTagUnmarshal (0.00s)
PASS
ok      learn-go-json   0.408s
*/
