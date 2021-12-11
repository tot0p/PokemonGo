package pok

import (
	"encoding/json"
	"os"
)

func LoadJson(s string) map[string]interface{} {
	file, _ := os.Open(s)
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	cont := make([]byte, fileinfo.Size())
	_, err = file.Read(cont)
	if err != nil {
		panic(err)
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(cont, &jsonMap)
	if err != nil {
		panic(err)
	}
	return jsonMap
}

func WriteJson(s string, m map[string]interface{}) {
	cont, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(s, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(cont)
	if err != nil {
		panic(err)
	}
}
