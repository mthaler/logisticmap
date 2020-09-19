package main

import (
	"github.com/fogleman/gg"
	"image/png"
	"io"
	"log"
	"net/http"
	"strconv"
)

const R = 4
const Size = 1000

func main() {
	http.HandleFunc("/lm", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	iterations := getQueryParameter(r, "iterations", 100)
	plot(w, iterations)
}


func logisticMap(r, y float64) float64 {
	return r * y * (1 - y)
}


func plot(out io.Writer, iterations int) {
	dc := gg.NewContext(Size, Size)
	dc.DrawRectangle(0, 0, Size, Size)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	for x := 0; x < Size; x++ {
		r := 4 * float64(x) / float64(Size)
		y := 0.1
		for i := 0; i < iterations; i++ {
			y = logisticMap(r, y)
		}
		dc.DrawCircle(float64((x)), y * Size, float64(1))
		dc.SetRGB(0, 0, 0)
		dc.Fill()
	}
	img := dc.Image()
	png.Encode(out, img)
}


func getQueryParameter( r *http.Request, name string, defaultValue int) int {
	q, ok := r.URL.Query()[name]
	if ok && len(q) > 0 {
		s := q[0]
		i, err := strconv.Atoi(s)
		if err != nil {
			return defaultValue
		} else {
			return i
		}
	} else {
		return defaultValue
	}
}