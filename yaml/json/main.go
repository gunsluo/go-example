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

// ProfileConfig sync user name from profile service
type ProfileConfig struct {
	Enable                 bool     `json:"enable"`
	EnableLoginSync        bool     `json:"enableLoginSync"`
	AllowedConnectorIds    []string `json:"allowedConnectorIds"`
	AllowedConnectorIdsAll bool     `json:"allowedConnectorIdsAll"`
}

// RegionConfig is for region configuration
type RegionConfig struct {
	Name               string        `json:"name"`
	DefaultRedirectURI string        `json:"defaultRedirectURI"`
	Profile            ProfileConfig `json:"profile"`
}

// RegionsConfig is for multi region configuration
type RegionsConfig struct {
	SyncAddr   string         `json:"syncAddr"`
	ServerName string         `json:"serverName"`
	TLSCA      string         `json:"tlsCA"`
	TLSCert    string         `json:"tlsCert"`
	TLSKey     string         `json:"tlsKey"`
	Regions    []RegionConfig `json:"regions"`
}

type Config struct {
	Issuer  string  `json:"issuer"`
	Storage Storage `json:"storage"`

	RegionManager RegionsConfig `json:"regions"`
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

	fmt.Println("-->", c.Issuer, c.Storage, c.RegionManager.SyncAddr)

	for _, region := range c.RegionManager.Regions {
		fmt.Println("-->", region.Name, region.Profile.AllowedConnectorIds)
	}
}
