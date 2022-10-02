package learngojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingEncoder(t *testing.T) {
	writer, err := os.Create("9-smartphone-out.json")
	if err != nil {
		panic(err.Error())
	}
	encoder := json.NewEncoder(writer)
	smartphone := Smartphone{
		Name:              "Iphone 13 Pro Max",
		Brand:             "Apple",
		BatterySize:       5000,
		CPUClock:          3.5,
		ReleasedCountries: []string{"Indonesia", "India", "Korea", "Malaysia"},
		Credentials: []Credential{
			{IMEI: "239KLLKWER90238048", BuildCode: "XJ230283X"},
			{IMEI: "722FHDSUASE2793723", BuildCode: "DJ328083S"},
		},
	}

	err = encoder.Encode(&smartphone)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Berhasil membuat file 9-smartphone-out.json")
}

/*
=== RUN   TestStreamingEncoder
Berhasil membuat file 9-smartphone-out.json
--- PASS: TestStreamingEncoder (0.00s)
PASS
ok      learn-go-json   0.374s
*/
