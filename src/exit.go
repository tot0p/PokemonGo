package pok

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Exit() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}


