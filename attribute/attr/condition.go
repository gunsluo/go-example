package attr

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Context map[string]interface{}

// Request is the warden's request object.
type Request struct {
	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// Action is the action that is requested on the resource.
	Action string `json:"action"`

	// Subejct is the subject that is requesting access.
	Subject string `json:"subject"`

	// Context is the request's environmental context.
	Context Context `json:"context"`
}

// Condition either do or do not fulfill an access request.
type Condition interface {
	// GetName returns the condition's name.
	GetName() string

	// Fulfills returns true if the request is fulfilled by the condition.
	Fulfills(interface{}, *Request) bool

	// Value set value to the condition.
	Values(string, map[string]interface{}) bool
}

// Conditions is a collection of conditions.
type Conditions map[string]Condition

// AddCondition adds a condition to the collection.
func (cs Conditions) AddCondition(key string, c Condition) {
	cs[key] = c
}

// MarshalJSON marshals a list of conditions to json.
func (cs Conditions) MarshalJSON() ([]byte, error) {
	out := make(map[string]*jsonCondition, len(cs))
	for k, c := range cs {
		raw, err := json.Marshal(c)
		if err != nil {
			return []byte{}, errors.WithStack(err)
		}

		out[k] = &jsonCondition{
			Type:    c.GetName(),
			Options: json.RawMessage(raw),
		}
	}

	return json.Marshal(out)
}

// UnmarshalJSON unmarshals a list of conditions from json.
func (cs Conditions) UnmarshalJSON(data []byte) error {
	if cs == nil {
		return errors.New("Can not be nil")
	}

	var jcs map[string]jsonCondition
	var dc Condition

	if err := json.Unmarshal(data, &jcs); err != nil {
		return errors.WithStack(err)
	}

	for k, jc := range jcs {
		var found bool
		for name, c := range ConditionFactories {
			if name == jc.Type {
				found = true
				dc = c()

				if len(jc.Options) == 0 {
					cs[k] = dc
					break
				}

				if err := json.Unmarshal(jc.Options, dc); err != nil {
					return errors.WithStack(err)
				}

				cs[k] = dc
				break
			}
		}

		if !found {
			return errors.Errorf("Could not find condition type %s", jc.Type)
		}
	}

	return nil
}

type jsonCondition struct {
	Type    string          `json:"type"`
	Options json.RawMessage `json:"options"`
}

// ConditionFactories is where you can add custom conditions
var ConditionFactories = map[string]func() Condition{
	new(StringEqualCondition).GetName(): func() Condition {
		return new(StringEqualCondition)
	},
}

type StringEqualCondition struct {
	Equals string `json:"equals"`
}

// Fulfills returns true if the given value is a string and is the
// same as in StringEqualCondition.Equals
func (c *StringEqualCondition) Fulfills(value interface{}, _ *Request) bool {
	s, ok := value.(string)

	return ok && s == c.Equals
}

// Value set value to the condition's Equals.
func (c *StringEqualCondition) Values(expression string, values map[string]interface{}) bool {
	s, ok := values["equals"].(string)
	if ok {
		c.Equals = s
	}

	return ok
}

// GetName returns the condition's name.
func (c *StringEqualCondition) GetName() string {
	return "StringEqualCondition"
}
