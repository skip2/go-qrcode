// go-qrcode
// Copyright 2014 Tom Harwood

package main

import (
	"encoding/base64"
	"fmt"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		//fmt.Printf("PNG is %d bytes long", len(png))
	}

	codeBase64 := base64.StdEncoding.EncodeToString(png)

	val := "data:image/png;base64," + codeBase64
	print(val)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
