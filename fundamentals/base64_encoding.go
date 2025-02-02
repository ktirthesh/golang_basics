package main

import (
	b64 "encoding/base64"
	"fmt"
)

func base64_encode_golang() {
	data := "abc123!?$*&()'-=@~"
	sEc := b64.StdEncoding.EncodeToString(([]byte(data)))
	fmt.Println(sEc)
	sDec, _ := b64.StdEncoding.DecodeString(sEc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
