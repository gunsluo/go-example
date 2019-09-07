package main

import "github.com/pkg/errors"

// StringEqualCondition is a condition which is fulfilled if the given
// string value is the same as specified in StringEqualCondition
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
func (c *StringEqualCondition) Values(expression string, values map[string]interface{}) error {
	if len(values) != 1 {
		return errors.Errorf("required a single value but %d were found", len(values))
	}

	for _, v := range values {
		if s, ok := v.(string); ok {
			c.Equals = s
		} else {
			return errors.New("the value must be a string")
		}
	}

	return nil
}

// GetName returns the condition's name.
func (c *StringEqualCondition) GetName() string {
	return "StringEqualCondition"
}
