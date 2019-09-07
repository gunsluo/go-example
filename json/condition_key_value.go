package main

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// KeyValueCondition is a condition which is fulfilled if the given
// string value is the same as specified in KeyValueCondition
type KeyValueCondition struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Fulfills returns true if the given value is a string and is the
// same as in KeyValueCondition.Value
func (c *KeyValueCondition) Fulfills(value interface{}, _ *Request) bool {
	s, ok := value.(string)

	return ok && s == c.Value
}

// Value set value to the condition's Value.
func (c *KeyValueCondition) Values(expression string, values map[string]interface{}) error {
	if len(values) != 1 {
		return errors.Errorf("required a single value but %d were found", len(values))
	}

	for k, v := range values {
		c.Key = k
		if s, ok := v.(string); ok {
			c.Value = s
		}
	}

	return nil
}

// GetName returns the condition's name.
func (c *KeyValueCondition) GetName() string {
	return "KeyValueCondition"
}

// MarshalJSON marshals KeyValueCondition to json.
func (c *KeyValueCondition) MarshalJSON() ([]byte, error) {
	rawTable := map[string]string{c.Key: c.Value}
	return json.Marshal(rawTable)
}

// UnmarshalJSON unmarshals KeyValueCondition from json.
func (c *KeyValueCondition) UnmarshalJSON(data []byte) error {
	kv := map[string]string{}
	if err := json.Unmarshal(data, &kv); err != nil {
		return errors.WithStack(err)
	}

	for k, v := range kv {
		c.Key = k
		c.Value = v
	}

	return nil
}
