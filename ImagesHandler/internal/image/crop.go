package image

import (
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	"image"
	"image/color"
	"image/png"
	"os"
	"path"
)

/*
@Time    : 2021/12/27 14:57
@Author  : austsxk
@Email   : austsxk@163.com
@File    : crop.go
@Software: GoLand
*/

const (
	PNGType = ".png"
)

type CropInter interface {
	Crop(inputFilePath, outputFilePath string, x1, y1, x2, y2 int) error
	Transparent(inputFile, outputFile string) error
}

func NewCropInter(t string) CropInter {
	switch t {
	case PNGType:
		return NewPNG()
	default:
		return nil
	}
}

// png 类型文件处理
type PNG struct {
}

func NewPNG() PNG {
	return PNG{}
}

func (pp PNG) Crop(inputFilePath, outputFilePath string, x1, y1, x2, y2 int) error {
	f, err := os.Open(inputFilePath)
	defer f.Close()
	if err != nil {
		return err
	}
	// decode图片
	m, err := png.Decode(f)
	if err != nil {
		return err
	}
	rgba := m.(*image.NRGBA)
	subImage, ok := rgba.SubImage(image.Rect(x1, y1, x2, y2)).(*image.NRGBA)
	if !ok {
		return errors.New("image type error.\n")
	}
	// 保存图片
	create, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	err = png.Encode(create, subImage)
	if err != nil {
		return err
	}
	return nil
}

// OpacityAdjust 图片
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
				// 颜色模型转换，至关重要！
				v := newRgba.ColorModel().Convert(color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: opacity})
				// Alpha = 0: Full transparent
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

// png设置透明度
func (pp PNG) Transparent(inputFile, outputFile string) error {
	f, err := os.Open(inputFile)
	defer f.Close()
	if err != nil {
		return err
	}
	// decode图片
	m, err := png.Decode(f)
	if err != nil {
		return err
	}
	im2 := OpacityAdjust(m.(*image.RGBA), 0)
	create, _ := os.Create(outputFile)
	err = png.Encode(create, im2)
	if err != nil {
		return err
	}
	return nil
}

// 图片裁剪
func CropHandler(inputFilePath, outputFilePath string, x1, y1, x2, y2 int) {
	fileType := path.Ext(path.Base(inputFilePath))
	crop := NewCropInter(fileType)
	if crop == nil {
		panic("type is error.\n")
	}
	err := crop.Crop(inputFilePath, outputFilePath, x1, y1, x2, y2)
	if err != nil {
		log.Errorf("picture crop error : %v", err)
		panic(err)
	}
	err = crop.Transparent(outputFilePath, inputFilePath)
	if err != nil {
		log.Errorf("picture transparent error : %v", err)
		panic(err)
	}
}
