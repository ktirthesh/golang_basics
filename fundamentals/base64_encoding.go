package main

import (
	b64 "encoding/base64"
	"fmt"
)

func base64_encode_golang() {
	data := "abc123!?$*&()'-=@~"
	sEc := b64.StdEncoding.EncodeToString(([]byte(data)))
	fmt.Println(sEc)
}
