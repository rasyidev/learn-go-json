package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Smartphone struct {
	Name        string
	Brand       string
	BatterySize int
	CPUClock    float32
}

func TestEncodeStructToJSON(t *testing.T) {
	smartphone := Smartphone{
		Name:        "Iphone 13 Pro Max",
		Brand:       "Apple",
		BatterySize: 5000,
		CPUClock:    3.5,
	}

	bytes, err := json.Marshal(smartphone)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(bytes))
}

/*
=== RUN   TestEncodeStructToJSON
{"Name":"Iphone 13 Pro Max","Brand":"Apple","BatterySize":5000,"CPUClock":3.5}
--- PASS: TestEncodeStructToJSON (0.00s)
PASS
ok      learn-go-json   0.431s
*/
