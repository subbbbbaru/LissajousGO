package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"os"
)

const (
	Xmin, Ymin, Xmax, Ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(Ymax-Ymin) + Ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(Xmax-Xmin) + Xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	Save(os.Stdout, img)
	//png.Encode(out, img)

}

func Save(out io.Writer, img image.Image) {
	png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const iters = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iters; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Gray{255}
}
