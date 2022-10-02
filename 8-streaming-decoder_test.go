package learngojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDicoder(t *testing.T) {
	reader, err := os.Open("8-smartphone.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(reader)

	smartphone := Smartphone{}
	decoder.Decode(&smartphone)
	fmt.Printf("%#v\n", smartphone)
}

/*
=== RUN   TestStreamDicoder
learngojson.Smartphone{Name:"Iphone 13 Pro Max", Brand:"Apple", BatterySize:5000, CPUClock:3.5, ReleasedCountries:[]string{"Indonesia", "India", "Korea", "Malaysia"}, Credentials:[]learngojson.Credential{learngojson.Credential{IMEI:"239KLLKWER90238048", BuildCode:"XJ230283X"}, learngojson.Credential{IMEI:"722FHDSUASE2793723", BuildCode:"DJ328083S"}}}
--- PASS: TestStreamDicoder (0.00s)
PASS
ok      learn-go-json   0.466s
*/
