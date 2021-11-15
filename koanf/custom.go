package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/mitchellh/mapstructure"
)

// Global koanf instance. Use . as the key path delimiter. This can be / or anything.
var k = koanf.New(".")

type ServeConfig struct {
	Version    string      `koanf:"version"`
	Connectors []Connector `koanf:"connectors"`
}

// Connector is a magical type that can unmarshal YAML dynamically. The
// Type field determines the connector type, which is then customized for Config.
type Connector struct {
	Type string `koanf:"type"`
	Name string `koanf:"name"`
	ID   string `koanf:"id"`

	Config ConnectorConfig `koanf:"config"`
}

// ConnectorConfig is a configuration that can open a connector.
type ConnectorConfig interface {
	Open(id string, logger log.Logger) error
}

func main() {
	// Load JSON config.
	if err := k.Load(file.Provider("mock/custom.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Unmarshal the whole root with FlatPaths: True.
	var o ServeConfig
	err := k.UnmarshalWithConf("", &o,
		koanf.UnmarshalConf{
			Tag: "koanf",
			DecoderConfig: &mapstructure.DecoderConfig{
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					MapToConnectorHookFunc(),
				),
				Metadata:         nil,
				Result:           &o,
				WeaklyTypedInput: true,
			},
		})
	if err != nil {
		panic(err)
	}
	for _, c := range o.Connectors {
		fmt.Println("->", c.ID, c.Name, c.Type, c.Config)
	}
}

// MapToConnectorHookFunc returns a DecodeHookFunc that converts
// strings to Connector.
func MapToConnectorHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.Map {
			return data, nil
		}

		c := Connector{}
		if t != reflect.TypeOf(c) {
			return data, nil
		}

		m, ok := data.(map[string]interface{})
		if !ok {
			return data, nil
		}
		if v, ok := m["id"]; !ok {
			return data, nil
		} else {
			c.ID = v.(string)
		}
		if v, ok := m["name"]; !ok {
			return data, nil
		} else {
			c.Name = v.(string)
		}
		if v, ok := m["type"]; !ok {
			return data, nil
		} else {
			c.Type = v.(string)
		}

		fn, ok := ConnectorsConfig[c.Type]
		if !ok {
			return c, fmt.Errorf("unknown connector type %q", c.Type)
		}
		c.Config = fn()
		if v, ok := m["config"]; ok && v != nil {
			if m, ok := v.(map[string]interface{}); ok {
				mapstructure.Decode(m, c.Config)
			}
		}

		return c, nil
	}
}

var ConnectorsConfig = map[string]func() ConnectorConfig{
	"mockCallback": func() ConnectorConfig { return new(CallbackConfig) },
	"github":       func() ConnectorConfig { return new(Config) },
}

// CallbackConfig holds the configuration parameters for a connector which requires no interaction.
type CallbackConfig struct{}

func (c *CallbackConfig) Open(id string, logger log.Logger) error {
	return nil
}

// Config holds configuration options for github logins.
type Config struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	RedirectURI  string `json:"redirectURI"`
}

func (c *Config) Open(id string, logger log.Logger) error {
	return nil
}
