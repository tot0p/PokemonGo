package pok

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

var InPokCenter = false

//Map
type MapPok struct {
	tileset      *ebiten.Image
	mp           []tile
	pokcenter    *Pokecenter
	xComp, yComp float64
	rect         [4]int
}

func (m *MapPok) Init() {
	m.tileset = LoadImg("data/Tilesets/Outside.png")
	temp := LoadFileMap("data/Map/map.txt")
	m.pokcenter = &Pokecenter{}
	m.pokcenter.Init(256, 64)
	m.rect[0], m.rect[1] = 0, 0
	for y, k := range temp {
		for x, i := range k {
			t := tile{}
			sx, sy := 32*(i%8), 32*(i/8)
			//fmt.Println(i%8, i/8, " ", i)
			img := m.tileset.SubImage(image.Rect(sx, sy, sx+32, sy+32)).(*ebiten.Image)
			t.Init(float64(x), float64(y), i, img)
			m.mp = append(m.mp, t)
			m.rect[0], m.rect[1] = x*32, y*32
			m.rect[2], m.rect[3] = x*32+32, y*32+32
			if x > 15 {
				m.rect[0] += 32 * (x - 14)
			}
		}
		if y > 11 {
			m.rect[1] += 32 * (y - 10)
		}
	}
}

func (m *MapPok) DrawInPokCenter(screen *ebiten.Image) {
	m.pokcenter.DrawIn(screen)
}

func (m *MapPok) Draw(screen *ebiten.Image) {
	for _, i := range m.mp {
		i.Draw(screen, m.xComp, m.yComp)
	}
	m.pokcenter.DrawOutside(screen, m.xComp, m.yComp)
}

func (m *MapPok) EnterPokeCenter(h *HitBox) bool {
	return m.pokcenter.Enter(h)
}

func (m *MapPok) checkCollPlayer(h *HitBox) bool {
	InPokCenter = m.EnterPokeCenter(h)
	return m.pokcenter.CheckCollideOut(h)
}

func (m *MapPok) GetComp() (float64, float64) {
	return m.xComp, m.yComp
}

func (m *MapPok) OnTallGrass(x, y float64) bool {
	for _, i := range m.mp {
		if i.PlayerOnIt(x, y) {
			if i.GetSpec() == 6 {
				return true
			}
		}
	}
	return false
}

//tile
type tile struct {
	img   *ebiten.Image
	x     float64
	y     float64
	spec  int
	solid bool
}

func (t *tile) PlayerOnIt(x, y float64) bool {
	if (t.x*32 <= x && x < t.x*32+32) && (t.y*32 <= y && y < t.y*32+32) {
		return true
	}
	return false
}

func (t *tile) SetSolidCond(cond func(int) bool) {
	if cond(t.spec) {
		t.solid = !t.solid
	}
}

func (t *tile) Init(x, y float64, spec int, img *ebiten.Image) {
	t.x = x
	t.y = y
	t.img = img
	t.spec = spec
	/*
		if spec == 6 {
			fmt.Println("x = ", x*32, "y = ", y*32)
		}
	*/
}

func (t *tile) GetSpec() int {
	return t.spec
}

func (t *tile) Draw(screen *ebiten.Image, xComp, yComp float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((t.x*32)-xComp, (t.y*32)-yComp)
	screen.DrawImage(t.img, op)
}

//loadFileMap
func LoadFileMap(filepath string) [][]int {
	var res [][]int
	var cursor int
	res = append(res, []int{})
	b, _ := ioutil.ReadFile(filepath)
	str2 := string(b)
	var str string
	for _, i := range str2 {
		if i != '\r' {
			str += string(i)
		}
	}
	t := strings.Split(str, "\n")
	for _, i := range t {
		for _, k := range strings.Split(i, ",") {
			t3, err := strconv.Atoi(k)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			res[cursor] = append(res[cursor], t3)
		}
		cursor++
		res = append(res, []int{})
	}
	return res[:len(res)-1]
}
