package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapMarshal(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P3823",
		"name":  "Apple Macbook Pro Max",
		"price": 40000000,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(bytes))
}

/*
=== RUN   TestJsonMapMarshal
{"id":"P3823","name":"Apple Macbook Pro Max","price":40000000}
--- PASS: TestJsonMapMarshal (0.00s)
PASS
ok      learn-go-json   0.382s
*/

func TestJsonMapUnmarshal(t *testing.T) {
	stringProduct := `{"id":"P3823","name":"Apple Macbook Pro Max","price":40000000}`

	product := map[string]interface{}{}
	err := json.Unmarshal([]byte(stringProduct), &product)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v\n", product)
	fmt.Println(product["id"])
	fmt.Println(product["name"])
	fmt.Println(product["price"])
}

/*
=== RUN   TestJsonMapUnmarshal
map[string]interface {}{"id":"P3823", "name":"Apple Macbook Pro Max", "price":4e+07}
P3823
Apple Macbook Pro Max
4e+07
--- PASS: TestJsonMapUnmarshal (0.00s)
PASS
ok      learn-go-json   0.368s
*/
