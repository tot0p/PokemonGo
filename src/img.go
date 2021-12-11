package pok

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImg(s string) *ebiten.Image {
	file, _ := os.Open(s)
	defer file.Close()
	img, _, _ := image.Decode(file)
	return ebiten.NewImageFromImage(img)
}

func LoadImgImage(s string) image.Image {
	file, _ := os.Open(s)
	defer file.Close()
	img, _, _ := image.Decode(file)
	return img
}
