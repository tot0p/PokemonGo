package pok

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//Chance Apparition Pokemon
func ChanceAp(x float64) float64 {
	return x / 187.5
}

type Pokemon struct {
	pv       int
	pvmax    int
	att1     int
	att2     int
	typ1     string
	typ2     string
	name     string
	Evl      int
	Evn      string
	tc       int
	imgName  string
	lv       int
	exp      int
	SavePath string
}

func (p *Pokemon) Init(name string) {
	p.name = name
	b := LoadJson("data/pokemon/info/" + name + ".json")
	p.typ1, p.typ2 = b["type1"].(string), b["type2"].(string)
	p.pv = int(b["pv"].(float64))
	p.pvmax = p.pv
	p.att1 = int(b["att1"].(float64))
	p.att2 = int(b["att2"].(float64))
	p.Evl = int(b["EvLv"].(float64))
	p.Evn = b["EvName"].(string)
	p.tc = int(b["tc"].(float64))
	p.imgName = b["id"].(string)
	p.lv = 1
	p.exp = 0
	p.SavePath = "data/save/pok/"
}

func (p *Pokemon) AddExp(h int) {
	p.exp += h
	for p.exp >= 100 {
		p.exp -= 100
		p.lv += 1
	}
}

func (p *Pokemon) GetId() string {
	return p.imgName
}

func (p *Pokemon) GetPv() int {
	return p.pv
}

func (p *Pokemon) SetPv(pv int) {
	p.pv = pv
}

func (p *Pokemon) SetName(name string) {
	p.SavePath = "data/save/pok/"
	p.name = name
}

func (p *Pokemon) Load(fileName string) {
	var name string
	for i := range fileName {
		if fileName[i] > 64 {
			name += string(fileName[i])
		}
	}
	p.Init(name)
	b, _ := ioutil.ReadFile(p.SavePath + fileName + ".txt")
	str := string(b)
	t := strings.Split(str, "\n")
	p.pv, _ = strconv.Atoi(t[0])
	p.exp, _ = strconv.Atoi(t[1])
	p.lv, _ = strconv.Atoi(t[2])
}

func (p *Pokemon) Debug() {
	fmt.Println("Name ", p.name)
	fmt.Println("Type ", p.typ1, "/", p.typ2)
	fmt.Println("Att ", p.att1, "/", p.att2)
	fmt.Println("Pv ", p.pv, "/", p.pvmax)
	fmt.Println("Lv ", p.lv, "/", p.exp)
	fmt.Println("Ev ", p.Evl, "/", p.Evn)
	fmt.Println("tc ", p.tc)
	fmt.Println("Img ", p.imgName)
}

func (p *Pokemon) Save() {
	var cursor int
	var cont []byte
	var file *os.File
	_, t := os.Stat(p.SavePath + p.name + ".txt")
	for t == nil {
		_, t = os.Stat(p.SavePath + p.name + strconv.Itoa(cursor) + ".txt")
		cursor++
	}
	if cursor != 0 {
		file, _ = os.Create(p.SavePath + p.name + strconv.Itoa(cursor-1) + ".txt")
	} else {
		file, _ = os.Create(p.SavePath + p.name + ".txt")
	}
	cont = []byte(strconv.Itoa(p.pv) + "\n" + strconv.Itoa(p.exp) + "\n" + strconv.Itoa(p.lv))
	file.Write(cont)
	file.Close()
}
