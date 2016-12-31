// go-qrcode
// Copyright 2014 Tom Harwood

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	outFile := flag.String("o", "", "out PNG file prefix, empty for stdout")
	size := flag.Int("s", 256, "image size (pixel)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `qrcode -- QR Code encoder in Go
https://github.com/skip2/go-qrcode

Flags:
`)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Usage:
  1. Arguments except for flags are joined by " " and used to generate QR code.
     Default output is STDOUT, pipe to imagemagick command "display" to display
     on any X server.

       qrcode hello word | display

  2. Save to file if "display" not available:

       qrcode "homepage: https://github.com/skip2/go-qrcode" > out.png

`)
	}
	flag.Parse()

	if *size <= 0 {
		checkError(fmt.Errorf("Error: value of -s should > 0"))
	}
	if len(flag.Args()) == 0 {
		flag.Usage()
		checkError(fmt.Errorf("Error: no content given"))
	}

	content := strings.Join(flag.Args(), " ")

	var err error
	var q *qrcode.QRCode
	q, err = qrcode.New(content, qrcode.Highest)
	checkError(err)

	var png []byte
	png, err = q.PNG(*size)
	checkError(err)

	if *outFile == "" {
		os.Stdout.Write(png)
	} else {
		var fh *os.File
		fh, err = os.Create(*outFile + ".png")
		checkError(err)
		defer fh.Close()
		fh.Write(png)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(0)
	}
}
