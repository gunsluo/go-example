package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

const (
	configFile = "config.yaml"
)

type SQLite3 struct {
	File string `json:"file"`
}

type Storage struct {
	Type   string  `json:"type"`
	Config SQLite3 `json:"config"`
}

type Config struct {
	Issuer  string  `json:"issuer"`
	Storage Storage `json:"storage"`
}

func main() {
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	var c Config
	if err := yaml.Unmarshal(configData, &c); err != nil {
		panic(err)
	}

	fmt.Println("-->", c.Issuer, c.Storage)
}
