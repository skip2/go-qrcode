// go-qrcode

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	qrcode "github.com/skip2/go-qrcode"
)

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getQRCode(c *gin.Context) {

	size, err := strconv.Atoi(c.Param("size"))
	checkError(err)

	// outFile := randSeq(10)

	data := c.Param("data")

	var q *qrcode.QRCode
	q, err = qrcode.New(data, qrcode.Highest)
	checkError(err)

	var png []byte
	png, err = q.PNG(size)
	checkError(err)

	/*
		var fh *os.File
		fh, err = os.Create(outFile + ".png")
		checkError(err)
		defer fh.Close()
		fh.Write(png)
	*/

	c.Header("Content-Disposition", "attachment; filename=qrcode.png")
	c.Data(http.StatusOK, "application/octet-stream", png)
}

func main() {
	router := gin.New()

	envFile, _ := godotenv.Read(".env")
	var authorized *gin.RouterGroup
	authorized = router.Group("/api", gin.BasicAuth(gin.Accounts{
		envFile["basicAuthUsername"]: envFile["basicAuthPassword"],
	}))

	// authorized.GET("/qrcode/:data/:size/:border/:color/:bgcolor", getQRCode)
	authorized.GET("/qrcode/:data/:size", getQRCode)

	router.Run(":6868")
	// router.RunTLS(":6868", envFile["SSL_CERTFILE_FULLCHAIN"], envFile["SSL_PRIVATE_KEYFILE"])
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
