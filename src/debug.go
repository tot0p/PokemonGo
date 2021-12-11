package pok

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var debug bool

func DebugMod(screen *ebiten.Image) {
	if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		debug = !debug
	}
	if debug {
		tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
		FDebug, _ := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    22,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		text.Draw(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()), FDebug, 5, 25, color.Black)
		text.Draw(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()), FDebug, 5, 50, color.Black)
	}

}
