package main

import "github.com/fogleman/gg"

//import "fmt"

const (
	imW = 2000
	imH = 2000
	w   = 300
	h   = 200
)

func main() {
	// Init
	var mo [w][h]float64
	dc := gg.NewContext(imW, imH)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Play with mo
	for i := range mo {
		for j := range mo[i] {
			mo[i][j] = (float64(i) * float64(imH)) / (2. * float64(h))
		}
	}

	// Fill mountain
	dc.RotateAbout(gg.Radians(180.), float64(imW)/2., float64(imH)/2.)
	for _, l := range mo {
		var c float64 = 1. / float64(h)
		dc.SetRGBA(0, 0, 0, c)
		for j, s := range l {
			dc.LineTo(float64(j)*float64(imW)/float64(w-1), float64(s))
		}
		dc.LineTo(imW, 0)
		dc.LineTo(0, 0)
		dc.Fill()
	}

	dc.SavePNG("out.png")
}
