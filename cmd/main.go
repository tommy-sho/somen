package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"math/rand"
	"os"
)

// 画像を単色に染める
func fillRect(img *image.RGBA, col color.Color) {
	rect := img.Rect
	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			img.Set(v, h, col)
		}
	}
}

type Circle struct {
	p image.Point
	r int
}

func (c *Circle) drawBounds(img *image.RGBA, col color.Color) {
	for rad := 0.0; rad < 2*math.Pi; rad += 0.01 {
		x := int(float64(c.p.X) + float64(c.r)*math.Cos(rad))
		y := int(float64(c.p.Y) + float64(c.r)*math.Sin(rad))
		img.Set(x, y, col)
	}
}

func (c *Circle) drawRadius(img *image.RGBA, rad float64, col color.Color) {
	for r := 0.0; r < float64(c.r); r += 0.1 {
		x := int(float64(c.p.X) + r*math.Cos(rad-math.Pi/2))
		y := int(float64(c.p.X) + r*math.Sin(rad-math.Pi/2))
		img.Set(x, y, col)
	}
}

func main() {
	x := 0
	y := 0
	width := 500
	height := 500

	img := image.NewRGBA(image.Rect(x, y, width, height))
	fillRect(img, color.RGBA{255, 255, 255, 0})

	center := image.Point{250, 250}
	circle := Circle{center, 100}
	circle.drawBounds(img, color.RGBA{255, 0, 0, 0})
	colors := GetFullColor()
	colors.shuffle()
	for c, col := range colors {
		rad := float64(c) / float64(len(colors)) * 2 * math.Pi
		circle.drawRadius(img, rad, color.RGBA{col.Red, col.Green, col.Blue, 0})
	}
	//for c := 0.0; c < math.Pi*2; c += 0.01 {
	//	circle.drawRadius(img, float64(circle.r)*(c/math.Pi*2), color.RGBA{colors[0].Red, colors[0].Green, colors[0].Blue, 0})
	//}

	// 出力用ファイル作成(エラー処理は略)
	file, _ := os.Create("sample.jpg")
	defer file.Close()

	// JPEGで出力(100%品質)
	if err := jpeg.Encode(file, img, &jpeg.Options{100}); err != nil {
		panic(err)
	}

}

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Colors []Color

func GetFullColor() Colors {
	colors := Colors{}
	for c := 0.0; c < math.Pi*2; c += 0.01 {
		green := calculateGreen(c)
		red := calculateRed(c)
		blue := calculateBlue(c)
		c := Color{
			Red:   red,
			Green: green,
			Blue:  blue,
		}
		colors = append(colors, c)
	}
	return colors
}

func calculateRed(angle float64) uint8 {
	if (0 <= angle) && (angle <= math.Pi/3) {
		return 255
	}
	if (math.Pi/3 < angle) && (angle < math.Pi*2/3) {
		res := 255 + (angle-math.Pi/3)*-(255*3/math.Pi)
		return uint8(res)
	}
	if (math.Pi*2/3 < angle) && (angle < math.Pi*4/3) {
		return 0
	}
	if (math.Pi*4/3 < angle) && (angle < math.Pi*5/3) {
		res := (angle - math.Pi*4/3) * (255 * 3 / math.Pi)
		return uint8(res)
	}
	return 255
}

func calculateGreen(angle float64) uint8 {
	if (0 <= angle) && (angle <= math.Pi/3) {
		s := angle - (math.Pi)/3
		res := 255 * (3 / math.Pi) * s
		return uint8(res)
	}
	if (math.Pi/3 < angle) && (angle < math.Pi*2/3) {
		return 255
	}
	if (math.Pi*2/3 < angle) && (angle < math.Pi) {
		return 255
	}
	if (math.Pi < angle) && (angle < math.Pi*4/3) {
		s := angle - math.Pi
		res := 255 - 255*(3/math.Pi)*s
		return uint8(res)
	}
	if (math.Pi*4/3 < angle) && (angle < math.Pi*5/3) {
		return 0
	}
	return 0
}

func calculateBlue(angle float64) uint8 {
	if (0 <= angle) && (angle <= math.Pi/3) {
		return 0
	}
	if (math.Pi/3 < angle) && (angle < math.Pi*2/3) {
		return 0
	}
	if (math.Pi*2/3 < angle) && (angle <= math.Pi) {
		s := angle - (2*math.Pi)/3
		res := 255 * (3 / math.Pi) * s
		return uint8(res)
	}
	if (math.Pi < angle) && (angle < math.Pi*4/3) {
		return 255
	}
	if (math.Pi*4/3 < angle) && (angle < math.Pi*5/3) {
		return 255
	}
	s := angle - (5*math.Pi)/3
	res := 255 - 255*(3/math.Pi)*s
	return uint8(res)
}

func (c Colors) shuffle() {
	for i := range c {
		j := rand.Intn(i + 1)
		c[i], c[j] = c[j], c[i]
	}
}
