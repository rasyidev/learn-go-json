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

## JSON Tag
- Atribut dalam Struct akan di-_mapping_ sama persis (_case sensitif_) menjadi nama atribut JSON
- Jika ingin membuat nama atribut yang berbeda, dapat memanfaatkan tag reflection di dalam Struct. Contoh:
```go
type Product struct {
  Id        string  `json:"id"`
  Name      string  `json:"name"`
  Price     int64   `json:"price"`
  ImageUrl  string  `json:"image_url"`
}
```

## Map
- Alternatif dari Struct untuk representasi JSON yang dinamis
- Menggunakan `map[string]interface{}`
- Tidak mendukung JSON Tag

## Streaming Decoder
- Data JSON dapat berasal dari input `io.Reader` seperti File, Network, Request Body, dll.
- Package json memiliki fitur untuk membaca data tersebut dari Stream tanpa harus menyimpannya terlebih dahulu ke dalam variable
- Menggunakan `json.NewDecoder(reader)` -> `Decode(interface{})`

## Streaming Encoder
- Menulis JSON langsung ke io.Writer
- Tidak perlu menyimpan JSON ke variable untuk menyimpan string / bytes JSON-nya
- Menggunakan `json.NewEncoder(interface{})` -> `Encode(interface{})`