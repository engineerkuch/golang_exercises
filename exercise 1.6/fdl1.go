package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{255, 0, 0, 0},
	color.RGBA{0, 255, 0, 0xff},
	color.RGBA{0, 0, 255, 0},
}



const (
	firstIndex = iota
	secondIndex
	thirdIndex
	fourthIndex
)

func main() {
	f, err := os.Create("lissajous.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		// cycles  = 5     // number of complex x oscillator revolutions
		res     = 0.01 // angular resolution
		size    = 100  // image canvas covers [-size..+size]
		nframes = 100  // number pf animation frames
		delay   = 8    // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; /* t < cycles*2*math.Pi; is same as --->*/ t < 6.29; t += res {
			x := math.Log10(t)
			y := math.Log10(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), thirdIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
