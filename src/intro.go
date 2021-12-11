package pok

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Intro struct {
	name  string
	imgB  *ebiten.Image
	profI *ebiten.Image
}

func (i *Intro) Init() {
	i.name = ""
	i.imgB = LoadImg("data/introbg.png")
	i.profI = LoadImg("data/introOak.png")
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	Fname, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

func (i *Intro) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(i.imgB, op)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(201, 100)
	screen.DrawImage(i.profI, op)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(25, 140)
	screen.DrawImage(LoadImg("data/pokemon/img/001.png"), op)
	op.GeoM.Translate(150, 0)
	screen.DrawImage(LoadImg("data/pokemon/img/004.png"), op)
	op.GeoM.Translate(150, 0)
	screen.DrawImage(LoadImg("data/pokemon/img/007.png"), op)
	text.Draw(screen, "Click on Pokemon For Choose", Fname, 175, 25, color.Black)
}

func (i *Intro) Choice(pl *[6]*Pokemon) bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		var t bool
		var str string
		if 69 < x && x < 138 && 235 < y && y < 300 {
			str, t = "Bulbizarre", true
		} else if 218 < y && y < 300 && 221 < x && x < 278 {
			str, t = "Salameche", true
		} else if 206 < y && y < 300 && 367 < x && x < 448 {
			str, t = "Carapuce", true
		}
		if t {
			p := Pokemon{}
			p.Init(str)
			pl[0] = &p
			return true
		}
	}
	return false
}
