package rbac

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// model represents the whole access control model.
type model map[string]assertionMap

// assertionMap is the collection of assertions, can be "r", "p", "g", "e", "m".
type assertionMap map[string]*Assertion

// Assertion represents an expression in a section of the model.
// For example: r = sub, obj, act
type Assertion struct {
	Key    string
	Value  string
	Tokens []string
	Rule   [][]string
	RL     *roleLinks
}

var sectionKVMap = map[string]string{
	"r": "sub, obj, act",
	"p": "sub, obj, act",
	"g": "_, _",
	"e": "some(where (p.eft == allow)) && !some(where (p.eft == deny))",
	"m": "g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)",
}

// newModel creates a model.
func newModel() (model, error) {
	m := make(model)

	//if err := loadSection(m, "p"); err != nil {
	//	return nil, err
	//}

	if err := loadSection(m, "g"); err != nil {
		return nil, err
	}

	return m, nil
}

func loadSection(md model, sec string) error {
	if loadAssertion(md, sec, sec) {
		return nil
	}

	return errors.Errorf("failed loadSection %s", sec)
}

func loadAssertion(md model, sec string, key string) bool {
	if value, ok := sectionKVMap[key]; ok {
		return md.addDef(sec, key, value)
	}

	return false
}

// addDef adds an assertion to the model.
func (md model) addDef(sec string, key string, value string) bool {
	ast := Assertion{}
	ast.Key = key
	ast.Value = value

	if ast.Value == "" {
		return false
	}

	if sec == "r" || sec == "p" {
		ast.Tokens = strings.Split(ast.Value, ", ")
		for i := range ast.Tokens {
			ast.Tokens[i] = key + "_" + ast.Tokens[i]
		}
	} else {
		ast.Value = RemoveComments(EscapeAssertion(ast.Value))
	}

	_, ok := md[sec]
	if !ok {
		md[sec] = make(assertionMap)
	}

	md[sec][key] = &ast
	return true
}

// Clear clears all current rule.
func (md model) Clear() {
	//for _, ast := range md["p"] {
	//	ast.Rule = nil
	//}

	for _, ast := range md["g"] {
		ast.Rule = nil
	}
}

// GetRule gets all rules by sec and ptype.
func (md model) GetRule(sec string, ptype string) [][]string {
	return md[sec][ptype].Rule
}

func (md model) Count(sec string, ptype string) int {
	return len(md[sec][ptype].Rule)
}

// GetFilteredRule gets rules based on field filters from a model.
func (md model) GetFilteredRule(sec string, ptype string, fieldIndex int, fieldValues ...string) [][]string {
	res := [][]string{}

	for _, rule := range md[sec][ptype].Rule {
		matched := true
		for i, fieldValue := range fieldValues {
			if fieldValue != "" && rule[fieldIndex+i] != fieldValue {
				matched = false
				break
			}
		}

		if matched {
			res = append(res, rule)
		}
	}

	return res
}

// HasRule determines whether a model has the specified rule.
func (md model) HasRule(sec string, ptype string, rule []string) bool {
	for _, r := range md[sec][ptype].Rule {
		if ArrayEquals(rule, r) {
			return true
		}
	}

	return false
}

// Add adds a rule to the model.
func (md model) Add(sec string, ptype string, rule []string) bool {
	if !md.HasRule(sec, ptype, rule) {
		md[sec][ptype].Rule = append(md[sec][ptype].Rule, rule)
		return true
	}
	return false
}

// Remove removes a rule from the model.
func (md model) Remove(sec string, ptype string, rule []string) bool {
	for i, r := range md[sec][ptype].Rule {
		if ArrayEquals(rule, r) {
			md[sec][ptype].Rule = append(md[sec][ptype].Rule[:i], md[sec][ptype].Rule[i+1:]...)
			return true
		}
	}

	return false
}

// RemoveFilteredRule removes rules based on field filters from the model.
func (md model) RemoveFilteredRule(sec string, ptype string, fieldIndex int, fieldValues ...string) bool {
	tmp := [][]string{}
	res := false
	for _, rule := range md[sec][ptype].Rule {
		matched := true
		for i, fieldValue := range fieldValues {
			if fieldValue != "" && rule[fieldIndex+i] != fieldValue {
				matched = false
				break
			}
		}

		if matched {
			res = true
		} else {
			tmp = append(tmp, rule)
		}
	}

	md[sec][ptype].Rule = tmp
	return res
}

// GetValuesForFieldInRule gets all values for a field for all rules in a model, duplicated values are removed.
func (md model) GetValuesForFieldInRule(sec string, ptype string, fieldIndex int) []string {
	values := []string{}

	for _, rule := range md[sec][ptype].Rule {
		values = append(values, rule[fieldIndex])
	}

	ArrayRemoveDuplicates(&values)
	// sort.Strings(values)

	return values
}

// buildRoleLinks build the rules to roleLinks.
func (md model) buildRoleLinks(rl *roleLinks) {
	for _, ast := range md["g"] {
		ast.buildRoleLinks(rl)
	}
}

func (ast *Assertion) buildRoleLinks(rl *roleLinks) {
	ast.RL = rl
	count := strings.Count(ast.Value, "_")

	for _, rule := range ast.Rule {
		if count < 2 {
			//fmt.Pritln("the number of \"_\" in role definition should be at least 2")
			continue
		}
		if len(rule) < count {
			//fmt.Pritln("grouping policy elements do not meet role definition")
			continue
		}

		if count == 2 {
			ast.RL.AddLink(rule[0], rule[1])
		} else if count == 3 {
			ast.RL.AddLink(rule[0], rule[1], rule[2])
		} else if count == 4 {
			ast.RL.AddLink(rule[0], rule[1], rule[2], rule[3])
		}
	}

	fmt.Println("Role links for: " + ast.Key)
	ast.RL.PrintRoles()
}

// loadRuleLine loads a text line as a rule to model.
func loadRuleLine(line string, md model) {
	if line == "" {
		return
	}

	if strings.HasPrefix(line, "#") {
		return
	}

	tokens := strings.Split(line, ", ")

	key := tokens[0]
	sec := key[:1]
	md[sec][key].Rule = append(md[sec][key].Rule, tokens[1:])
}
