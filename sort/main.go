package main

import (
	"fmt"
	"sort"
)

type connectorInfo struct {
	ID   string
	Name string
}

func main() {
	connectors := []connectorInfo{
		connectorInfo{ID: "", Name: "GitHub"},
		connectorInfo{ID: "", Name: "Google"},
		connectorInfo{ID: "", Name: "LinkedIn"},
		connectorInfo{ID: "", Name: "Microsoft"},
		connectorInfo{ID: "", Name: "GitLab"},
	}
	sort.Sort(byNamePriority(connectors))
	for _, c := range connectors {
		fmt.Println("==>", c.ID, c.Name)
	}
}

type byName []connectorInfo

func (n byName) Len() int           { return len(n) }
func (n byName) Less(i, j int) bool { return n[i].Name < n[j].Name }
func (n byName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

var connectorInfoPriority = map[string]int{
	"LinkedIn":  0,
	"Google":    1,
	"Microsoft": 2,
	"GitHub":    3,
	"GitLab":    4,
}

type byNamePriority []connectorInfo

func (n byNamePriority) Len() int      { return len(n) }
func (n byNamePriority) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n byNamePriority) Less(i, j int) bool {
	in := n[i].Name
	jn := n[j].Name

	var ip, jp int
	if p, ok := connectorInfoPriority[in]; ok {
		ip = p
	} else {
		ip = 99
	}

	if p, ok := connectorInfoPriority[jn]; ok {
		jp = p
	} else {
		jp = 99
	}

	return ip < jp
}
