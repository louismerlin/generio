package main

import "github.com/fogleman/gg"
import "math/rand"
import "time"

//import "fmt"

const (
	imW    = 2000
	imH    = 2000
	w      = 100
	h      = 10
	lambda = 1500
	fill   = 0.6
	bgR    = 126
	bgG    = 163
	bgB    = 204
	r      = 0
	g      = 7
	b      = 45
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

func simpleRand(mo [h][w]float64) [h][w]float64 {
	// Play with mo
	for i := range mo {
		for j := range mo[i] {
			top := ((lambda*r1.Float64() - lambda/2) + float64(i)*float64(imH))
			down := (1 / fill * float64(h))
			mo[i][j] = top / down
		}
	}
	return mo
}

func fromLast(mo [h][w]float64) [h][w]float64 {
	for i := range mo {
		mo[i][0] = (float64(i) * float64(imH)) / (1 / fill * float64(h))
		for j := range mo[i][1:] {
			mo[i][j+1] = mo[i][j] + 0.05*(lambda*r1.Float64()-lambda/2)
		}
	}
	return mo
}

func main() {
	// Init
	var mo [h][w]float64
	dc := gg.NewContext(imW, imH)
	dc.SetRGB255(bgR, bgG, bgB)
	dc.Clear()

	//mo = simpleRand(mo)
	mo = fromLast(mo)

	// Fill mountain
	dc.RotateAbout(gg.Radians(180.), float64(imW)/2., float64(imH)/2.)
	var c float64 = 255 * 2 / float64(h)
	for _, l := range mo {
		dc.SetRGBA255(r, g, b, int(c))
		for j, s := range l {
			dc.LineTo(float64(j)*float64(imW)/float64(w-1), float64(s))
		}
		dc.LineTo(imW, 0)
		dc.LineTo(0, 0)
		dc.FillPreserve()
		dc.SetRGB255(r, g, b)
		dc.SetLineWidth(0.)
		dc.Stroke()
	}

	dc.SavePNG("out.png")
}
