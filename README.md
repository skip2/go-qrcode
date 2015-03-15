# go-qrcode #

<img src='https://skip.org/img/nyancat-youtube-qr.png' align='right'>

Package qrcode implements a QR Code encoder. [![Build Status](https://travis-ci.org/skip2/go-qrcode.svg?branch=master)](https://travis-ci.org/skip2/go-qrcode) <br>

<br>
A QR Code is a matrix (two-dimensional) barcode. Arbitrary content may be encoded, with URLs being a popular choice :)<br>
<br>
Each QR Code contains error recovery information to aid reading damaged or obscured codes. There are four levels of error recovery: Low, medium, high and highest. QR Codes with a higher recovery level are more robust to damage, at the cost of being physically larger.<br>
<br>
<h1>Usage</h1>
<pre>import qrcode "github.com/skip2/go-qrcode"
</pre>

<ul><li><b>Create a PNG image:</b><pre>var png []byte
png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
</pre></li></ul>

<ul><li><b>Create a PNG image and write to a file:</b>
<pre>err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "qr.png")
</pre></li></ul>

Both examples use the <code>qrcode.Medium</code> error Recovery Level and create a 256x256 pixel, black on white QR Code.<br>
<br>
The maximum capacity of a QR Code varies according to the content encoded and<br>
the error recovery level. The maximum capacity is 2,953 bytes, 4,296<br>
alphanumeric characters, 7,089 numeric digits, or a combination of these.<br>
<br>
<h1>Documentation</h1>

<a href='https://godoc.org/github.com/skip2/go-qrcode'><img src='https://godoc.org/github.com/skip2/go-qrcode?status.png' /></a>

<h1>Demoapp</h1>
<a href='http://go-qrcode.appspot.com'>http://go-qrcode.appspot.com</a>

<h1>Links</h1>

<ul><li><a href='http://en.wikipedia.org/wiki/QR_code'>http://en.wikipedia.org/wiki/QR_code</a>
</li><li><a href='http://www.iso.org/iso/catalogue_detail.htm?csnumber=43655'>ISO/IEC 18004:2006</a> - Main QR Code specification (approx CHF 198,00)<br>
</li><li><a href='https://github.com/qpliu/qrencode-go/'>https://github.com/qpliu/qrencode-go/</a> - alternative Go QR encoding library based on <a href='https://github.com/zxing/zxing'>ZXing</a>
