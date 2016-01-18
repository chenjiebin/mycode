package main

import (
	"bufio"
	//	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

var (
	dpi        float64 = 150
	fontfile   string  = "./msyh2.ttf"
	hinting    string  = "none"
	size       float64 = 12
	spacing    float64 = 1.5
	bgfilepath string  = "./meiqu_watermark.png"
)

func createWatermark(w http.ResponseWriter, r *http.Request) {
	var keyword string
	r.ParseForm()
	if len(r.Form["k"]) > 0 {
		keyword = r.Form["k"][0]
	}

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	bgfile, _ := os.Open(bgfilepath)
	bg, _ := png.Decode(bgfile)

	rgba := image.NewRGBA(image.Rect(0, 0, 331, 61))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	c2 := freetype.NewContext()
	c2.SetDPI(dpi)
	c2.SetFont(f)
	c2.SetFontSize(size)
	c2.SetClip(rgba.Bounds())
	c2.SetDst(rgba)
	c2.SetSrc(image.Black)
	c2.SetHinting(font.HintingNone)

	// Draw the text.
	pt2 := freetype.Pt(11, 2+int(c2.PointToFixed(size)>>6))
	_, err = c2.DrawString(keyword, pt2)
	if err != nil {
		log.Println(err)
		return
	}
	pt2.Y += c2.PointToFixed(size * spacing)

	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(f)
	c.SetFontSize(size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.White)
	c.SetHinting(font.HintingNone)
	// Draw the text.
	pt := freetype.Pt(10, 1+int(c.PointToFixed(size)>>6))
	_, err = c.DrawString(keyword, pt)
	if err != nil {
		log.Println(err)
		return
	}
	pt.Y += c.PointToFixed(size * spacing)

	b := bufio.NewWriter(w)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/cw", createWatermark)
	err := http.ListenAndServe(":19000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
