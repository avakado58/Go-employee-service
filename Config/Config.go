package Config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadCfg() *Settings {
	filePath := "./.Config/config.json"
	fileByte, err := ioutil.ReadFile(filePath)
	if err != nil{
		log.Fatal(err)
		panic(err)
	}

	var cfg Settings
	err = json.NewDecoder(bytes.NewBuffer(fileByte)).Decode(&cfg)
	if err != nil{
		log.Fatal(err)
		panic(err)
	}

	return &cfg
}
