package main

import "fmt"

type P struct {
	Name string
	Subs []string
}

func main() {

	ps := []P{
		P{
			Name: "tom",
			Subs: []string{"tom"},
		},
		P{
			Name: "peter",
			Subs: []string{"peter"},
		},
		P{
			Name: "jack",
			Subs: []string{"jack"},
		},
		P{
			Name: "tome",
			Subs: []string{"jerry"},
		},
	}
	roles := []string{"peter", "jack"}

	var ret []P
	var isMatch bool
	for _, p := range ps {
		isMatch = false
	LOOP:
		for _, sub := range p.Subs {
			for _, r := range roles {
				if r == sub {
					isMatch = true
					break LOOP
				}
			}
		}
		if isMatch {
			ret = append(ret, p)
		}
	}

	fmt.Println("=>", ret)
}
