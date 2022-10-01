package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonRequest := `{"Name":"Iphone 13 Pro Max","Brand":"Apple","BatterySize":5000,"CPUClock":3.5}`
	jsonByte := []byte(jsonRequest)

	smartphone := Smartphone{}

	err := json.Unmarshal(jsonByte, &smartphone)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(smartphone)
}

/*
=== RUN   TestDecodeJson
{Iphone 13 Pro Max Apple 5000 3.5}
--- PASS: TestDecodeJson (0.00s)
PASS
ok      learn-go-json   0.544s
*/
