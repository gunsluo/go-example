package srv

import "github.com/gunsluo/go-example/ory/identity-ui/swagger/identityclient"

type Group struct {
	Name     string
	Includes []string
}

type From struct {
	Group      string       `json:"group"`
	GroupNodes []GroupNodes `json:"group_nodes"`
}

type GroupNodes struct {
	Name  string
	Nodes []identityclient.UiNode `json:"nodes"`
}

type Froms struct {
	Action   string                  `json:"action"`
	Method   string                  `json:"method"`
	Messages []identityclient.UiText `json:"messages,omitempty"`
	Froms    []From                  `json:"froms"`
}

func groupLoginUi(ui identityclient.UiContainer) Froms {
	groups := []Group{
		{
			Name:     "oidc",
			Includes: []string{"oidc", "default"},
		},
		{
			Name:     "password",
			Includes: []string{"password", "default"},
		},
		{
			Name:     "totp",
			Includes: []string{"totp", "default"},
		},
	}

	return groupUi(ui, groups)
}

func groupUi(ui identityclient.UiContainer, groups []Group) Froms {
	froms := Froms{
		Action:   ui.Action,
		Method:   ui.Method,
		Messages: ui.Messages,
		Froms:    make([]From, len(groups)),
	}

	for _, n := range ui.Nodes {
		for i, g := range groups {
			for j, item := range g.Includes {
				if n.Group == item {
					froms.Froms[i].Group = g.Name
					if len(froms.Froms[i].GroupNodes) == 0 {
						froms.Froms[i].GroupNodes = make([]GroupNodes, len(g.Includes))
					}

					froms.Froms[i].GroupNodes[j].Name = n.Group
					froms.Froms[i].GroupNodes[j].Nodes = append(froms.Froms[i].GroupNodes[j].Nodes, n)
					break
				}
			}
		}
	}

	return froms
}
