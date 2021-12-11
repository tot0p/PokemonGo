package pok

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	ExitIn = false
)

func (p *Player) Control() {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		p.g = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyA) || inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) {
		p.g = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		p.d = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyD) || inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		p.d = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		p.h = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyW) || inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) {
		p.h = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		p.b = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyS) || inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
		p.b = false
	}
	if p.g && p.x >= 0 {
		p.x -= 2
		p.hb.Update(p.x, p.y)
		if p.mp.checkCollPlayer(p.hb) {
			p.x += 2
			p.g = false
			p.hb.Update(p.x, p.y)
		}
	}
	if p.d && p.x <= float64(p.mp.rect[2])-32 {
		p.x += 2
		p.hb.Update(p.x, p.y)
		if p.mp.checkCollPlayer(p.hb) {
			p.x -= 2
			p.d = false
			p.hb.Update(p.x, p.y)
		}
	}
	if p.h && p.y >= 0 {
		p.y -= 2
		p.hb.Update(p.x, p.y)
		if p.mp.checkCollPlayer(p.hb) {
			p.y += 2
			p.h = false
			p.hb.Update(p.x, p.y)
		}
	}
	if p.b && p.y+16 <= float64(p.mp.rect[3])-32 {
		p.y += 2
		p.hb.Update(p.x, p.y)
		if p.mp.checkCollPlayer(p.hb) {
			p.y -= 2
			p.b = false
			p.hb.Update(p.x, p.y)
		}
	}
	if p.x >= 256 && p.mp.xComp+p.x < float64(p.mp.rect[0])-256 {
		if p.g {
			p.mp.xComp -= 2
		}
		if p.d {
			p.mp.xComp += 2
		}
	}
	if p.y >= 176 && p.mp.yComp+p.y+16 < float64(p.mp.rect[1])-176 {
		if p.h {
			p.mp.yComp -= 2
		}
		if p.b {
			p.mp.yComp += 2
		}
	}
	//Reset
	if p.y < 192 {
		p.mp.yComp = 0
	}
	if p.x < 256 {
		p.mp.xComp = 0
	}
}

func (p *Player) Debug() {
	fmt.Println("x = ", p.x, " xComp = ", p.mp.xComp, " some = ", p.mp.xComp+p.x)
	fmt.Println("y = ", p.y, " yComp = ", p.mp.yComp, " some = ", p.mp.yComp+p.y)
}

func (p *Player) ControlIn() {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		p.g = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyA) || inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) {
		p.g = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		p.d = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyD) || inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		p.d = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		p.h = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyW) || inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) {
		p.h = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		p.b = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyS) || inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
		p.b = false
	}
	if p.g && p.xIn > 0 {
		p.xIn -= 2
	}
	if p.d && p.xIn < 448 {
		p.xIn += 2
	}
	if p.h && p.yIn > 0 {
		p.yIn -= 2
	}
	if p.b && p.yIn < 308 {
		p.yIn += 2
	}
	p.hb.Update(p.xIn, p.yIn)
	ExitIn = p.mp.pokcenter.EventExit(p.hb)
	if ExitIn {
		p.y += 10
		if p.y >= 176 {
			p.mp.yComp += 10
		}
	}
}

func (p *Player) ResetDep() {
	p.h, p.b, p.d, p.g = false, false, false, false
}
