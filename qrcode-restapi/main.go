package main

import (
	"fmt"
	"image/color"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

// https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 6:
		_, err = fmt.Sscanf(s, "%02x%02x%02x", &c.R, &c.G, &c.B)
	case 3:
		_, err = fmt.Sscanf(s, "%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func getQRCode(c *gin.Context) {

	var data string = c.Query("data")
	size, _ := strconv.Atoi(c.Query("size"))
	var ecc string = strings.ToUpper(c.Query("ecc"))
	var recovery_level qrcode.RecoveryLevel = qrcode.Low
	// margin, _ := strconv.Atoi(c.Query("margin")), no margin is supported yet in go-qrcode
	var color string = c.Query("color")
	var bgcolor string = c.Query("bgcolor")
	// qzone, no qzone is supported yet in go-qrcode
	// format, only png format is supported in go-qrcode

	if (strings.Compare(ecc, "L") == 0) ||
		(strings.Compare(ecc, "M") == 0) ||
		(strings.Compare(ecc, "Q") == 0) ||
		(strings.Compare(ecc, "H") == 0) {

		if strings.Compare(ecc, "L") == 0 {
			recovery_level = qrcode.Low
		} else if strings.Compare(ecc, "M") == 0 {
			recovery_level = qrcode.Medium
		} else if strings.Compare(ecc, "Q") == 0 {
			recovery_level = qrcode.High
		} else if strings.Compare(ecc, "H") == 0 {
			recovery_level = qrcode.Highest
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", "ECC aka RecoveryLevel is not supported")
		}
	}

	mycolor, _ := ParseHexColor(color)
	mybgcolor, _ := ParseHexColor(bgcolor)

	var q *qrcode.QRCode
	q, err := qrcode.New(data, recovery_level)
	checkError(err)
	q.ForegroundColor = mycolor
	q.BackgroundColor = mybgcolor

	var png []byte
	png, err = q.PNG(size)
	checkError(err)

	c.Header("Content-Disposition", "inline; filename=qrcode.png")
	c.Data(http.StatusOK, "application/octet-stream", png)
}

func main() {
	router := gin.New()

	router.GET("/api/qrcode", getQRCode)

	router.Run(":6868")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		// os.Exit(1)
	}
}
