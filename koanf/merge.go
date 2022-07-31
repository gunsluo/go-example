package main

import (
	"fmt"
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

var conf = koanf.Conf{
	Delim: ".",
	//StrictMerge: true,
}
var k = koanf.NewWithConf(conf)

type parentStruct struct {
	Name   string      `koanf:"name"`
	ID     int         `koanf:"id"`
	Child1 childStruct `koanf:"child1"`
}
type childStruct struct {
	Name        string            `koanf:"name"`
	Type        string            `koanf:"type"`
	Empty       map[string]string `koanf:"empty"`
	Grandchild1 grandchildStruct  `koanf:"grandchild1"`
}
type grandchildStruct struct {
	Ids []int `koanf:"ids"`
	On  bool  `koanf:"on"`
}
type sampleStruct struct {
	Type    string            `koanf:"type"`
	Empty   map[string]string `koanf:"empty"`
	Parent1 parentStruct      `koanf:"parent1"`
}

func main() {
	k.Load(structs.Provider(sampleStruct{
		Type:  "json",
		Empty: make(map[string]string),
		Parent1: parentStruct{
			Name: "parent1",
			ID:   1234,
			Child1: childStruct{
				Name:  "a",
				Type:  "json",
				Empty: make(map[string]string),
				Grandchild1: grandchildStruct{
					Ids: []int{1, 2, 3},
					On:  true,
				},
			},
		},
	}, "koanf"), nil)

	jsonPath := "mock/mock.json"
	if err := k.Load(file.Provider(jsonPath), json.Parser(), koanf.WithMergeFunc(notMergeNullValues)); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Printf("name is = `%s`\n", k.String("parent1.child1.name"))
}

func notMergeNullValues(src, dest map[string]interface{}) error {
	for k, v := range src {
		dv, ok := dest[k]
		if !ok {
			dest[k] = v
			continue
		}

		if v == nil {
			continue
		}

		var overwrite bool
		switch vv := v.(type) {
		case string:
			if vv != "" {
				overwrite = true
			}
		case []interface{}:
			if len(vv) > 0 {
				overwrite = true
			}
		case map[string]interface{}:
			if dvv, ok := dv.(map[string]interface{}); ok {
				if err := notMergeNullValues(vv, dvv); err != nil {
					return err
				}
				continue
			}

			overwrite = true
		default:
			overwrite = true
		}

		if overwrite {
			dest[k] = v
		}
	}

	return nil
}
