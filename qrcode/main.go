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
	outFile := flag.String("o", "", "out file prefix, empty for stdout")
	size := flag.Int("s", 256, "image size (pixel)")
	vector := flag.Bool("v", false, "render QR code as SVG instead of PNG")
	textArt := flag.Bool("t", false, "print as text-art on stdout")
	negative := flag.Bool("i", false, "invert black and white")
	disableBorder := flag.Bool("d", false, "disable QR Code border")
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

	if len(flag.Args()) == 0 {
		flag.Usage()
		checkError(fmt.Errorf("Error: no content given"))
	}

	content := strings.Join(flag.Args(), " ")

	var err error
	var q *qrcode.QRCode
	q, err = qrcode.New(content, qrcode.Highest)
	checkError(err)

	if *disableBorder {
		q.DisableBorder = true
	}

	if *textArt {
		art := q.ToString(*negative)
		fmt.Println(art)
		return
	}

	if *negative {
		q.ForegroundColor, q.BackgroundColor = q.BackgroundColor, q.ForegroundColor
	}

	render := q.PNG
	suffix := "png"
	if *vector {
		render = q.SVG
		suffix = "svg"
	}

	data, err := render(*size)
	checkError(err)

	if *outFile == "" {
		os.Stdout.Write(data)
	} else {
		var fh *os.File
		fh, err = os.Create(fmt.Sprintf("%v.%v", *outFile, suffix))
		checkError(err)
		defer fh.Close()
		fh.Write(data)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
