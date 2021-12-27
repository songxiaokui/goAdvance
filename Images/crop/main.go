package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/draw"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"

	"os"
)

/*
@Time    : 2021/12/27 09:42
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

func readJpeg() {
	f, err := os.Open("/Users/austsxk/github_sxk_repo/goAdvance/goAdvance/Images/crop/water.jpeg")
	if err != nil {
		panic(err)
	}
	// decode图片
	m, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T", m)
	rgba := m.(*image.YCbCr)
	subImage := rgba.SubImage(image.Rect(829, 0, 1705, 1375)).(*image.YCbCr)
	// 保存图片
	create, _ := os.Create("new2.jpg")
	err = png.Encode(create, subImage)
}

func readImage() {
	f, err := os.Open("water.png")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	// decode图片
	m, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	rgba := m.(*image.NRGBA)
	subImage := rgba.SubImage(image.Rect(829, 0, 1705, 1375)).(*image.NRGBA)
	// 保存图片
	create, _ := os.Create("water.png")
	err = png.Encode(create, subImage)
}

func OpacityAdjust(m *image.RGBA, percentage float64) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	var (
		R, G, B, A uint32
	)
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			if i == 1 && j == 1 {
				R, G, B, A = r, g, b, a
				fmt.Println(R, G, B, A)

			}
			opacity := uint8(float64(a) * percentage)
			if r == R && g == G && b == B {
				//颜色模型转换，至关重要！
				v := newRgba.ColorModel().Convert(color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: opacity})
				//Alpha = 0: Full transparent
				rr, gg, bb, aa := v.RGBA()
				newRgba.SetRGBA(i, j, color.RGBA{R: uint8(rr), G: uint8(gg), B: uint8(bb), A: uint8(aa)})
			} else {
				v := newRgba.ColorModel().Convert(color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
				//Alpha = 0: Full transparent
				rr, gg, bb, aa := v.RGBA()
				newRgba.SetRGBA(i, j, color.RGBA{R: uint8(rr), G: uint8(gg), B: uint8(bb), A: uint8(aa)})
			}
		}
	}
	return newRgba
}

func JPGOpacityAdjust2(m *image.RGBA64, percentage float64) *image.RGBA64 {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	var (
		R, G, B, A uint32
	)
	newRgba := image.NewRGBA64(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			if i == 1 && j == 1 {
				R, G, B, A = r, g, b, a
				fmt.Println(R, G, B, A)

			}
			opacity := uint16(float64(a) * percentage)
			if r == R && g == G && b == B {
				//颜色模型转换，至关重要！
				v := newRgba.ColorModel().Convert(color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: opacity})
				//Alpha = 0: Full transparent
				rr, gg, bb, aa := v.RGBA()
				newRgba.SetRGBA64(i, j, color.RGBA64{R: uint16(rr), G: uint16(gg), B: uint16(bb), A: uint16(aa)})
			} else {
				v := newRgba.ColorModel().Convert(color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
				//Alpha = 0: Full transparent
				rr, gg, bb, aa := v.RGBA()
				newRgba.SetRGBA64(i, j, color.RGBA64{R: uint16(rr), G: uint16(gg), B: uint16(bb), A: uint16(aa)})
			}
		}
	}
	return newRgba
}

func AlphaSet() {
	f1, err := os.Open("water.png")
	defer f1.Close()
	if err != nil {
		panic(err)
	}
	// decode图片
	m1, err := png.Decode(f1)
	if err != nil {
		panic(err)
	}
	im2 := OpacityAdjust(m1.(*image.RGBA), 0)
	create2, _ := os.Create("water.png")
	err = png.Encode(create2, im2)
}

func AlphaSet2() {
	f1, err := os.Open("/Users/austsxk/github_sxk_repo/goAdvance/goAdvance/Images/crop/new2.jpg")
	if err != nil {
		panic(err)
	}
	// decode图片
	m1, err := png.Decode(f1)
	if err != nil {
		panic(err)
	}
	fmt.Println("------")
	im2 := JPGOpacityAdjust2(m1.(*image.RGBA64), 0)
	create2, _ := os.Create("waterjpg.png")
	err = png.Encode(create2, im2)
}

func main() {
	readImage()
	AlphaSet()
	// readJpeg()
	// readGif()
	// readJpeg()
	// readJpeg()
	// AlphaSet2()
}
