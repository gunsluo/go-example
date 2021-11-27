package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/mitchellh/mapstructure"
)

// Global koanf instance. Use . as the key path delimiter. This can be / or anything.
type ServeConfig struct {
	Version    string      `koanf:"version"`
	Cron       Cron        `koanf:"cron"`
	Connectors []Connector `koanf:"connectors"`
}

type Cron struct {
	Name     string        `koanf:"name"`
	NickName string        `koanf:"nick_name"`
	Period   time.Duration `koanf:"period"`
	Price    float64       `koanf:"price"`
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

func PrintEnvValuesHookFunc(paths map[string]interface{}, k, v string) (string, interface{}) {
	if strings.HasPrefix(k, "connectors.") {
		fmt.Println("---->", k, "=", v)
	}
	return "", nil
}

func ConnectorEnvValuesHookFunc(values map[string]interface{}, k, v string) (string, interface{}) {
	if !strings.HasPrefix(k, "connectors.") {
		return "", nil
	}

	items := strings.Split(k, ".")
	if len(items) < 3 {
		return "", nil
	}

	connectors := []Connector{}
	if val, ok := values["connectors"]; ok {
		switch vals := val.(type) {
		case []interface{}:
			decoder, err := mapstructure.NewDecoder(
				&mapstructure.DecoderConfig{
					DecodeHook: MapToConnectorHookFunc(),
					Result:     &connectors,
				})
			if err != nil {
				return "", nil
			}

			if err := decoder.Decode(vals); err != nil {
				return "", nil
			}
		case []Connector:
			connectors = vals
		default:
			return "", nil
		}
	}

	ty := items[1]
	var found bool
	for i, c := range connectors {
		if strings.ToLower(c.Type) == ty {
			fields := items[2:]
			if updater, ok := ConnectorsEnvUpdater[c.Type]; ok {
				updater(&c, v, fields...)
			}
			connectors[i] = c
			found = true
			break
		}
	}

	if !found {
		for k, updater := range ConnectorsEnvUpdater {
			if strings.ToLower(k) == ty {
				fields := items[2:]
				c := Connector{Type: k}
				updater(&c, v, fields...)
				connectors = append(connectors, c)
				break
			}
		}

	}

	values["connectors"] = connectors
	return "connectors", connectors
}

type EnvValuesHookFunc func(map[string]interface{}, string, string) (string, interface{})

func ComposeEnvValuesHookFunc(fs ...EnvValuesHookFunc) EnvValuesHookFunc {
	return func(values map[string]interface{}, k string, v string) (string, interface{}) {
		for _, fn := range fs {
			key, val := fn(values, k, v)
			if val != nil {
				return key, val
			}
		}
		return "", nil
	}
}

func main() {
	var delim = "."
	var prefix = ""
	var k = koanf.New(delim)
	envValuesHookFunc := ComposeEnvValuesHookFunc(PrintEnvValuesHookFunc, ConnectorEnvValuesHookFunc)

	// Load JSON config.
	if err := k.Load(file.Provider("mock/custom.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	//fmt.Println("->", k.String("cron.nick_name"))

	// environment variable
	flattenedPaths := k.All()
	k.Load(env.ProviderWithValue(prefix, delim, func(s string, v string) (string, interface{}) {
		// Strip out the MYVAR_ prefix and lowercase and get the key while also replacing
		// the _ character with . in the key (koanf delimeter).
		k := strings.Replace(strings.ToLower(strings.TrimPrefix(s, prefix)), "_", delim, -1)
		if _, ok := flattenedPaths[k]; ok {
			return k, v
		}
		for key := range flattenedPaths {
			normalized := strings.Replace(key, "_", delim, -1)
			if normalized == k {
				return key, v
			}
		}

		// If there is a space in the value, split the value into a slice by the space.
		// if strings.Contains(v, " ") {
		// 	return key, strings.Split(v, " ")
		// }

		// Otherwise, return the plain string.
		return envValuesHookFunc(flattenedPaths, k, v)
	}), nil)

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

	fmt.Println("->", o.Version, o.Cron, o.Cron.NickName)
	for _, c := range o.Connectors {
		fmt.Println("\t->", c.ID, c.Name, c.Type, c.Config)
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

var ConnectorsEnvUpdater = map[string]func(*Connector, string, ...string){
	"mockCallback": func(c *Connector, v string, fields ...string) {
		if len(fields) == 0 {
			return
		}
		field := fields[0]
		switch field {
		case "id":
			c.ID = v
		case "name":
			c.Name = v
		}

		if c.Config == nil {
			c.Config = &CallbackConfig{}
		}
	},
	"github": func(c *Connector, v string, fields ...string) {
		if len(fields) == 0 {
			return
		}

		field := fields[0]
		if field == "id" {
			c.ID = v
			return
		}

		if field == "name" {
			c.Name = v
			return
		}

		if field != "config" {
			return
		}

		if len(fields) < 2 {
			return
		}

		if c.Config == nil {
			c.Config = &Config{}
		}

		config := c.Config.(*Config)
		cfgField := fields[1]
		switch cfgField {
		case "clientid":
			config.ClientID = v
		case "clientsecret":
			config.ClientSecret = v
		case "redirecturi":
			config.RedirectURI = v
		}
	},
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
