package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/mitchellh/mapstructure"
)

func main() {
	var delim = "."
	var prefix = ""
	var k = koanf.New(delim)
	envValuesHookFunc := ComposeEnvValuesHookFunc(PrintEnvValuesHookFunc, MethodsEnvValuesHookFunc)

	os.Setenv("SELFSERVICE_METHODS_PASSWORD_CONFIG_MIN_LENGTH", "8")

	// Load JSON config.
	if err := k.Load(file.Provider("mock/custom2.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// environment variable
	flattenedPaths := k.All()
	k.Load(env.ProviderWithValue(prefix, delim, func(s string, v string) (string, interface{}) {
		// Strip out the MYVAR_ prefix and lowercase and get the key while also replacing
		// the _ character with . in the key (koanf delimeter).
		k := strings.Replace(strings.ToLower(strings.TrimPrefix(s, prefix)), "_", delim, -1)
		if _, ok := flattenedPaths[k]; ok {
			fmt.Println("1.---->", s, k, v)
			return k, v
		}
		for key := range flattenedPaths {
			normalized := strings.Replace(key, "_", delim, -1)
			if normalized == k {
				fmt.Println("2.---->", s, key, v)
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
	v := k.Get("selfservice.methods.password.config.min_length")

	fmt.Printf("===>%T, %v, %d\n", v, v, k.Int("selfservice.methods.password.config.min_length"))

	// Unmarshal the whole root with FlatPaths: True.
	var o Conf
	err := k.UnmarshalWithConf("", &o,
		koanf.UnmarshalConf{
			Tag: "koanf",
			DecoderConfig: &mapstructure.DecoderConfig{
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					MapToMethodsHookFunc(),
				),
				Metadata:         nil,
				Result:           &o,
				WeaklyTypedInput: true,
			},
		})
	if err != nil {
		panic(err)
	}

	fmt.Println("->", o.Version)
	m, exists := o.SelfSrv.Methods.Get("password")
	if exists {
		pc := &PasswordMethodConfig{}
		if err := m.Decode(pc); err != nil {
			panic(err)
		}
		fmt.Println("->", pc)
	}

	for k, v := range o.SelfSrv.Methods {
		fmt.Println("\t->", k, v.Enabled, v.Config)
	}
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

func PrintEnvValuesHookFunc(paths map[string]interface{}, k, v string) (string, interface{}) {
	if strings.HasPrefix(k, "selfservice.") {
		fmt.Println("---->", k, "=", v)
	}
	return "", nil
}

func MethodsEnvValuesHookFunc(values map[string]interface{}, k, v string) (string, interface{}) {
	// if !strings.HasPrefix(k, "selfservice.") {
	// 	return "", nil
	// }
	fmt.Println("----------->", k, v)
	return "", nil
	// return k, values
}

// MapToConnectorHookFunc returns a DecodeHookFunc that converts
// strings to Connector.
func MapToMethodsHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.Map {
			return data, nil
		}

		ms := Methods{}
		if t != reflect.TypeOf(ms) {
			return data, nil
		}

		mapping, ok := data.(map[string]interface{})
		if !ok {
			return data, nil
		}

		for k, v := range mapping {
			fn, ok := methodsConfigMappging[k]
			if !ok {
				continue
			}

			method, ok := v.(map[string]interface{})
			if !ok {
				continue
			}

			mc := fn()
			if c, ok := method["config"]; ok && c != nil {
				if d, ok := c.(map[string]interface{}); ok {
					mapstructure.Decode(d, mc)
				}
			}
			m := Method{Config: mc}
			if c, ok := method["enabled"]; ok && c != nil {
				if d, ok := c.(bool); ok {
					m.Enabled = d
				}
			}
			ms[k] = m
		}

		return ms, nil
	}
}

var methodsConfigMappging = map[string]func() MethodConfig{
	"password": func() MethodConfig { return new(PasswordMethodConfig) },
	"oidc":     func() MethodConfig { return new(OidcMethodConfig) },
	"totp":     func() MethodConfig { return new(TotpMethodConfig) },
}

// Conf is the entry of the configuration file
type Conf struct {
	Version string      `koanf:"version"`
	SelfSrv SelfService `koanf:"selfservice"`
}

type SelfService struct {
	DefaultBrowserReturnUrl *url.URL   `koanf:"default_browser_return_url"`
	AllowedReturnUrls       []*url.URL `koanf:"allowed_return_urls"`
	Methods                 Methods    `koanf:"methods"`
}

type Methods map[string]Method

func (ms Methods) Get(k string) (Method, bool) {
	m, ok := ms[k]
	return m, ok
}

type Method struct {
	Enabled bool         `koanf:"enabled"`
	Config  MethodConfig `koanf:"config"`
}

func (m Method) Decode(out interface{}) error {
	// if reflect.TypeOf(out).Kind() != reflect.TypeOf(m.Config).Kind() {
	// 	return fmt.Errorf("expected type '%s', got unconvertible type '%s'", reflect.TypeOf(m.Config).Kind(), reflect.TypeOf(out).Kind())
	// }

	// if reflect.TypeOf(out).String() != reflect.TypeOf(m.Config).String() {
	// 	return fmt.Errorf("expected type '%s', got unconvertible type '%s'", reflect.TypeOf(m.Config).String(), reflect.TypeOf(out).String())
	// }

	// reflect.ValueOf(out).Elem().Set(reflect.ValueOf(m.Config).Elem())
	return nil
}

type MethodConfig interface{}

type PasswordMethodConfig struct {
	MinLength          int     `koanf:"min_length" mapstructure:"min_length"`
	SimilarityCheck    bool    `koanf:"similarity_check" mapstructure:"similarity_check"`
	MaxSubstrThreshold float64 `koanf:"max_substr_threshold" mapstructure:"max_substr_threshold"`
	MinSimilarityDist  int     `koanf:"min_similarity_dist" mapstructure:"min_similarity_dist"`
	RequireCapital     bool    `koanf:"require_capital" mapstructure:"require_capital"`
	RequireLower       bool    `koanf:"require_lower" mapstructure:"require_lower"`
	RequireNumber      bool    `koanf:"require_number" mapstructure:"require_number"`
	RequireSpecial     bool    `koanf:"require_special" mapstructure:"require_special"`
}

type OidcMethodConfig struct {
}

type TotpMethodConfig struct {
	Issuer string `koanf:"issuer"`
	Digits int    `koanf:"digits"`
	Period int    `koanf:"period"`
	// sha1 sha256 sha512 md5
	Algorithm string `koanf:"algorithm"`
}
