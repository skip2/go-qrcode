package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func getQRCode(c *gin.Context) {

	var data string = c.Query("data")
	size, _ := strconv.Atoi(c.Query("size"))

	var q *qrcode.QRCode
	q, err := qrcode.New(data, qrcode.Highest)
	checkError(err)

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
		os.Exit(1)
	}
}
