package pok

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	EscapeBattle = false
	AttackValue  = 0
	Capt         = false
	Soin         = 0
	NextPok      = false
)

type InteBattle struct {
	img        *ebiten.Image
	x, y       float64
	bat, bag   bool
	att1, att2 int
}

func (i *InteBattle) SetPokAtt(att1, att2 int) {
	i.att1, i.att2 = att1, att2
}

func (i *InteBattle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(i.x, i.y)
	x, y := int(i.x), int(i.y)
	if !i.bag && !i.bat { //aff menu base
		screen.DrawImage(i.img, op)
		text.Draw(screen, "Attack", Fname, 15+x, 40+y, color.Black)
		text.Draw(screen, "Bag", Fname, 205+x, 40+y, color.Black)
		text.Draw(screen, "Pokemon", Fname, 15+x, 80+y, color.Black)
		text.Draw(screen, "Escape", Fname, 205+x, 80+y, color.Black)
	} else if i.bat { //aff att
		screen.DrawImage(i.img, op)
		text.Draw(screen, strconv.Itoa(i.att1), Fname, 15+x, 40+y, color.Black)
		text.Draw(screen, strconv.Itoa(i.att2), Fname, 205+x, 40+y, color.Black)
	} else if i.bag { //aff att
		screen.DrawImage(i.img, op)
		text.Draw(screen, "Pokeball", Fname, 15+x, 40+y, color.Black)
		text.Draw(screen, "Potion", Fname, 205+x, 40+y, color.Black)
	}
}

func (i *InteBattle) Event(count int) {
	if count%5 != 0 {
		return
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !i.bag && !i.bat {
		x, y := ebiten.CursorPosition()
		if int(i.x+10) <= x && x <= int(i.x+185) && int(i.y+10) <= y && y <= int(i.y+42) {
			fmt.Println("ATT")
			i.bat = true
		} else if int(i.x+210) <= x && x <= int(i.x+395) && int(i.y+10) <= y && y <= int(i.y+42) {
			fmt.Println("Bag")
			i.bag = true
		} else if int(i.x+10) <= x && x <= int(i.x+185) && int(i.y+52) <= y && y <= int(i.y+96) {
			NextPok = true
		} else if int(i.x+210) <= x && x <= int(i.x+395) && int(i.y+52) <= y && y <= int(i.y+96) {
			fmt.Println("Esc")
			EscapeBattle = true
		}
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && i.bat {
		x, y := ebiten.CursorPosition()
		if int(i.x+10) <= x && x <= int(i.x+185) && int(i.y+10) <= y && y <= int(i.y+42) {
			AttackValue = int(float64(i.att1)*RandForAtt1()) / 4
			i.bat = false
		} else if int(i.x+210) <= x && x <= int(i.x+395) && int(i.y+10) <= y && y <= int(i.y+42) {
			AttackValue = int(float64(i.att2)*RandForAtt2()) / 4
			i.bat = false
		}
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && i.bag {
		x, y := ebiten.CursorPosition()
		if int(i.x+10) <= x && x <= int(i.x+185) && int(i.y+10) <= y && y <= int(i.y+42) {
			Capt = true
			i.bag = false
		} else if int(i.x+210) <= x && x <= int(i.x+395) && int(i.y+10) <= y && y <= int(i.y+42) {
			Soin = 10
			i.bag = false
		}
	}
}
