package randstr

import (
	"encoding/base64"
	"crypto/rand"
	"fmt"
)

// SOURCE: https://www.socketloop.com/tutorials/golang-how-to-generate-random-string

func RandomString() string {
	size := 32 // change the length of the generated random string here

	rb := make([]byte, size)
	_, err := rand.Read(rb)


	if err != nil {
		fmt.Println(err)
	}

	rs := base64.URLEncoding.EncodeToString(rb)

	return rs
}
