package main

import (
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v3"
)

type base58Encoder struct{}

func (enc base58Encoder) Encode(u uuid.UUID) string {
	return base58.Encode(u[:])
}

func (enc base58Encoder) Decode(s string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(s))
}

func main() {
	enc := base58Encoder{}
	u := shortuuid.NewWithEncoder(enc)
	fmt.Println(len(u), u) // 6R7VqaQHbzC1xwA5UueGe6

	var total = 10000000
	var j int
	dm := make(map[string]interface{}, total)
	for i := 0; i < total; i++ {
		//u := shortuuid.New() // Cekw67uyMpBGZLRP2HFVbe
		u := shortuuid.NewWithEncoder(enc)
		if _, ok := dm[u]; ok {
			fmt.Println("-->", u)
			j++
		} else {
			dm[u] = struct{}{}
		}
	}

	fmt.Println("->", total, j)
}
