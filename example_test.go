// go-qrcode
// Copyright 2014 Tom Harwood
/*
	Amendments Thu, 2017-December-14:
	- test integration (go test -v)
	- idiomatic go code
*/
package qrcode

import (
	"fmt"
	"image/color"
	"os"
	"testing"
)

func TestExampleEncode(t *testing.T) {
	if png, err := Encode("https://example.org", Medium, 256); err != nil {
		t.Errorf("Error: %s", err.Error())
	} else {
		fmt.Printf("PNG is %d bytes long", len(png))
	}
}

func TestExampleWriteFile(t *testing.T) {
	filename := "example.png"
	if err := WriteFile("https://example.org", Medium, 256, filename); err != nil {
		if err = os.Remove(filename); err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	}
}

func TestExampleEncodeWithColourAndWithoutBorder(t *testing.T) {
	q, err := New("https://example.org", Medium)
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	// Optionally, disable the QR Code border.
	q.DisableBorder = true

	// Optionally, set the colours.
	q.ForegroundColor = color.RGBA{R: 0x33, G: 0x33, B: 0x66, A: 0xff}
	q.BackgroundColor = color.RGBA{R: 0xef, G: 0xef, B: 0xef, A: 0xff}

	err = q.WriteFile(256, "example2.png")
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}
}
