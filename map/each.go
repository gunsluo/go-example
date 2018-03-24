package main

import "fmt"

type Policy struct {
	Sources []string
}

var pmap = map[string]*Policy{
	"1": &Policy{
		Sources: []string{"abc", "dfg"},
	},
	"2": &Policy{
		Sources: []string{"dfg"},
	},
}

func main() {
	policies := getAll()

	sub := "dfg"
	for _, p := range policies {
		for i, s := range p.Sources {
			if s == sub {
				if i == 0 {
					p.Sources = p.Sources[1:]
				} else {
					p.Sources = append(p.Sources[:i-1], p.Sources[i:]...)
				}
				break
			}
		}
	}

	for _, p := range pmap {
		fmt.Println("->", p)
	}
}

func getAll() []*Policy {
	var ps []*Policy
	for _, p := range pmap {
		ps = append(ps, p)
	}

	return ps
}
