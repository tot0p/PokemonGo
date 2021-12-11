package pok

import "math/rand"

func RandPourInt(n int) bool {
	return rand.Intn(100) <= n
}

func RandPourFloat(n float64) bool {
	return float64(rand.Intn(100)) <= n
}

func RandForAtt1() float64 {
	return rand.Float64()
}

func RandForAtt2() float64 {
	return rand.Float64() + rand.Float64()
}

func RandListString(g []string) string {
	return g[rand.Intn(len(g))]
}
