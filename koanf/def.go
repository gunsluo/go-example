package main

import (
	"fmt"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
var k = koanf.New(".")

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
	// Load default values using the structs provider.
	// We provide a struct along with the struct tag `koanf` to the
	// provider.
	k.Load(structs.Provider(sampleStruct{
		Type:  "json",
		Empty: make(map[string]string),
		Parent1: parentStruct{
			Name: "parent1",
			ID:   1234,
			Child1: childStruct{
				Name:  "child",
				Type:  "json",
				Empty: make(map[string]string),
				Grandchild1: grandchildStruct{
					Ids: []int{1, 2, 3},
					On:  true,
				},
			},
		},
	}, "koanf"), nil)

	fmt.Printf("name is = `%s`\n", k.String("parent1.child1.name"))

	k.Load(file.Provider("mock/mock.yaml"), yaml.Parser())

	fmt.Printf("name is = `%s`\n", k.String("parent1.child1.name"))
}
