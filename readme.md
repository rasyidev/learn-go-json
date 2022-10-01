## Encode JSON
- Konversi data ke JSON
- `json.Marshal(interface{})(bytes, error)`

## JSON Object
- Representasi JSON dalam Go Lang menggunakan **Struct**
- Attribute di JSON Object adalah attribute di Struct

## Decode JSON
- Konversi JSON ke tipe data di Go Lang
- `Unmarshal(byte[], interface{})`

## JSON Array
- Direpresentasikan dalam bentuk slice
- Proses konversi JSON ke Object maupun sebaliknya dilakukan oleh package json secara otomatis. Hasilnya berupa slice