// go-qrcode
// Copyright 2014 Tom Harwood

package qrcode

import (
	"fmt"
	"testing"

	bitset "code.google.com/p/go-qrcode/bitset"
)

func TestBuildRegularSymbol(t *testing.T) {
	for k := 0; k <= 7; k++ {
		var v *qrCodeVersion = getQRCodeVersion(Low, 1)
		var s *symbol

		var data *bitset.Bitset = bitset.New()
		for i := 0; i < 26; i++ {
			data.AppendNumBools(8, false)
		}

		s, err := buildRegularSymbol(*v, k, data)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			_ = s
			//fmt.Print(m.string())
		}
	}
}
