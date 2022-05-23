package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	width, height = 1920, 1080          // Размер конвы в пикселях
	cells         = 100                 // Количество ячеек сетки
	XYrange       = 30.0                // Диапазон осей
	XYscale       = width / 2 / XYrange // Пикселей в единице X и Y
	Zscale        = height * 0.4        // Пикселей в единице Z
	angle         = math.Pi / 6         // Углы осей X, Y (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {

	widthStr, heightStr := strconv.Itoa(width), strconv.Itoa(height)
	file, err := os.Create("Surface.svg")
	if err != nil {
		fmt.Println("Error create file", err)
		os.Exit(1)
	}
	defer file.Close()
	startSVG := `<?xml version="1.0"?>` + "\n" + `<svg xmlns="http://www.w3.org/2000/svg" ` +
		`xmlns:xlink="http://www.w3.org/1999/xlink" ` +
		`style="stroke: grey; fill: white; stroke-width: 0.7" ` +
		`width="` + widthStr + `" height="` + heightStr + `">` + "\n"

	file.WriteString(startSVG)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			AXstr, AYstr := fmt.Sprintf("%g", ax), fmt.Sprintf("%g", ay)
			BXstr, BYstr := fmt.Sprintf("%g", bx), fmt.Sprintf("%g", by)
			CXstr, CYstr := fmt.Sprintf("%g", cx), fmt.Sprintf("%g", cy)
			DXstr, DYstr := fmt.Sprintf("%g", dx), fmt.Sprintf("%g", dy)

			middleString := `<polygon points="` + AXstr + `,` + AYstr + " " + BXstr + `,` + BYstr + ` ` +
				CXstr + `,` + CYstr + ` ` + DXstr + `,` + DYstr + `"/>` + "\n"
			file.WriteString(middleString)
		}
	}
	file.WriteString(`</svg>`)
}

func corner(i, j int) (float64, float64) {
	// Ищем угловую точку (X, Y) ячейки (j, i).
	x := XYrange * (float64(i)/cells - 0.5)
	y := XYrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности Z
	z := f(x, y)
	// Изометрически проецируем (x, y, z) на двумерную канву SVG (sx, sy)
	sx := width/2 + (x-y)*cos30*XYscale
	sy := height/2 + (x+y)*sin30*XYscale - z*Zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от (0, 0)
	return math.Sin(r)
}
