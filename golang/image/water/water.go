package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

var (
	dpi      = flag.Float64("dpi", 200, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "./water/Source_Code_Pro-YaHei0.90.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 12, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
)

var text = "不瘦立马卸载"

func main() {
	flag.Parse()

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize the context.
	fg := image.White

	bgfile, _ := os.Open("./water/meiqu_watermark.png")
	bg, _ := png.Decode(bgfile)

	rgba := image.NewRGBA(image.Rect(0, 0, 331, 61))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	c2 := freetype.NewContext()
	c2.SetDPI(*dpi)
	c2.SetFont(f)
	c2.SetFontSize(*size)
	c2.SetClip(rgba.Bounds())
	c2.SetDst(rgba)
	c2.SetSrc(image.Black)
	c2.SetHinting(font.HintingNone)

	// Draw the text.
	pt2 := freetype.Pt(11, 11+int(c2.PointToFixed(*size)>>6))
	_, err = c2.DrawString(text, pt2)
	if err != nil {
		log.Println(err)
		return
	}
	pt2.Y += c2.PointToFixed(*size * *spacing)

	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)
	// Draw the text.
	pt := freetype.Pt(10, 10+int(c.PointToFixed(*size)>>6))
	_, err = c.DrawString(text, pt)
	if err != nil {
		log.Println(err)
		return
	}
	pt.Y += c.PointToFixed(*size * *spacing)

	fmt.Println(pt.X, pt.Y, pt2.X, pt2.Y)

	// Save that RGBA image to disk.
	outFile, err := os.Create("out.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
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
	fmt.Println("Wrote out.png OK.")
}
