package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Credential struct {
	IMEI      string
	BuildCode string
}

func TestJsonArrayComplexMarshal(t *testing.T) {
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

	bytes, _ := json.Marshal(&smartphone)
	fmt.Println(string(bytes))
}

/*
=== RUN   TestJsonArrayComplexMarshal
{"Name":"Iphone 13 Pro Max","Brand":"Apple","BatterySize":5000,"CPUClock":3.5,"ReleasedCountries":["Indonesia","India","Korea","Malaysia"],"Credentials":[{"IMEI":"239KLLKWER90238048","BuildCode":"XJ230283X"},{"IMEI":"722FHDSUASE2793723","BuildCode":"DJ328083S"}]}
--- PASS: TestJsonArrayComplexMarshal (0.00s)
PASS
ok      learn-go-json   0.340s
*/

func TestJsonArrayComplexUnmarshal(t *testing.T) {
	smartphoneString := `{"Name":"Iphone 13 Pro Max","Brand":"Apple","BatterySize":5000,"CPUClock":3.5,"ReleasedCountries":["Indonesia","India","Korea","Malaysia"],"Credentials":[{"IMEI":"239KLLKWER90238048","BuildCode":"XJ230283X"},{"IMEI":"722FHDSUASE2793723","BuildCode":"DJ328083S"}]}`
	smartphone := Smartphone{}
	err := json.Unmarshal([]byte(smartphoneString), &smartphone)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(smartphone.Name)
	fmt.Println(smartphone.ReleasedCountries)

}

/*
=== RUN   TestJsonArrayComplexUnmarshal
Iphone 13 Pro Max
[Indonesia India Korea Malaysia]
--- PASS: TestJsonArrayComplexUnmarshal (0.00s)
PASS
ok      learn-go-json   0.384s
*/
