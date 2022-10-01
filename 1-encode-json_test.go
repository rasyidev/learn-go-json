package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Rasyidev")
	logJson(88)
	logJson(true)
	logJson([]string{"Rasyidev", "Pro", "Projects"})
}

/*
=== RUN   TestEncode
"Rasyidev"
88
true
["Rasyidev","Pro","Projects"]
--- PASS: TestEncode (0.00s)
PASS
ok      learn-go-json   0.487s
*/
