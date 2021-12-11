package pok

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	PokManager     = false
	TempMan        = false
	PokManagerView = PokControl{}
)

func EventMainEnterPokMan(pok *[6]*Pokemon) {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		TempMan = true
		PokManagerView.Init(pok)
	} else if inpututil.IsKeyJustReleased(ebiten.KeyQ) && TempMan {
		PokManager = !PokManager
		TempMan = false
	}
}

func EventMainExitPokMan() {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		TempMan = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyQ) && TempMan {
		PokManager = !PokManager
		PokManagerView.Reset()
		TempMan = false
	}
}
