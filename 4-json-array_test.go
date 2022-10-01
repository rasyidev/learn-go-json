package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArrayMarshal(t *testing.T) {
	smartphone := Smartphone{
		Name:              "Iphone 13 Pro Max",
		Brand:             "Apple",
		BatterySize:       5000,
		CPUClock:          3.5,
		ReleasedCountries: []string{"Indonesia", "India", "Korea", "Malaysia"},
	}

	byte, err := json.Marshal(smartphone)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(byte))
}

/*
=== RUN   TestJsonArray
{"Name":"Iphone 13 Pro Max","Brand":"Apple","BatterySize":5000,"CPUClock":3.5,"ReleasedCountries":["Indonesia","India","Korea","Malaysia"]}
--- PASS: TestJsonArray (0.00s)
PASS
ok      learn-go-json   0.411s
*/

func TestJsonArrayUnmarshal(t *testing.T) {
	smartphoneString := `{"Name":"Iphone 13 Pro Max","Brand":"Apple","BatterySize":5000,"CPUClock":3.5,"ReleasedCountries":["Indonesia","India","Korea","Malaysia"]}`

	smartphone := Smartphone{}

	err := json.Unmarshal([]byte(smartphoneString), &smartphone)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(smartphone.Name)
	fmt.Println(smartphone.ReleasedCountries)
}

/*
=== RUN   TestJsonArrayUnmarshal
Iphone 13 Pro Max
[Indonesia India Korea Malaysia]
--- PASS: TestJsonArrayUnmarshal (0.00s)
PASS
ok      learn-go-json   0.457s
*/
