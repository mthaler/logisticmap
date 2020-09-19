package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"net/http"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/lm", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	plot(w)
}


func logisticMap(r, y float64) float64 {
	return r * y * (1 - y)
}


func plot(out io.Writer) {
	const (
		size = 1000
		rmin = 2.4
		rmax = 4
	)

	rect := image.Rect(0, 0, 1.5 * size, size)
	img := image.NewPaletted(rect, palette)
	for x := 0; x < 1.5 * size * 10; x++ {
		r := rmin + (rmax - rmin) * float64(x) / float64(1.5 *size * 10)
		y := 0.5
		for i := 0; i < 1000; i++ {
			y = logisticMap(r, y)
		}
		for i := 1000; i < 1000 + 100; i++ {
			y = logisticMap(r, y)
			img.SetColorIndex(int(x / 10), size - int(y * size), blackIndex)
		}
	}

	gif.Encode(out, img, nil)
}