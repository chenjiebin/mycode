package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	dpi        float64 = 138
	fontfile   string  = "./msyh.ttf"
	hinting    string  = "none"
	size       float64 = 12
	spacing    float64 = 1.5
	bgfilepath string  = "./meiqu_watermark.png"
)

func createWatermark(w http.ResponseWriter, r *http.Request) {
	var keyword string
	var scr uint8 = 70
	var scg uint8 = 67
	var scb uint8 = 68
	var t int
	var err error
	r.ParseForm()
	if len(r.Form["k"]) > 0 {
		keyword = r.Form["k"][0]
	}
	if len(r.Form["scr"]) > 0 {
		t, err = strconv.Atoi(r.Form["scr"][0])
		scr = uint8(t)
	}
	if len(r.Form["scg"]) > 0 {
		t, err = strconv.Atoi(r.Form["scg"][0])
		scg = uint8(t)
	}
	if len(r.Form["scb"]) > 0 {
		t, err = strconv.Atoi(r.Form["scb"][0])
		scb = uint8(t)
	}
	if err != nil {
		fmt.Fprintln(w, "阴影色参数错误，数值0~255之间")
		return
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

	//阴影
	d2 := &font.Drawer{
		Dst: rgba,
		Src: &image.Uniform{color.RGBA{scr, scg, scb, 255}},
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingNone,
		}),
	}
	y2 := 2 + int(math.Ceil(size*dpi/72))
	d2.Dot = fixed.Point26_6{
		X: fixed.I(331) - d2.MeasureString(keyword) - fixed.I(2),
		Y: fixed.I(y2),
	}
	d2.DrawString(keyword)

	//画白字
	d := &font.Drawer{
		Dst: rgba,
		Src: image.White,
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingNone,
		}),
	}
	y := int(math.Ceil(size * dpi / 72))
	d.Dot = fixed.Point26_6{
		X: fixed.I(331) - d.MeasureString(keyword) - fixed.I(3),
		Y: fixed.I(y),
	}
	d.DrawString(keyword)

	//输出内容
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
