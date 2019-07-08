package main

import "github.com/pkg/errors"

type Key struct {
	Name      string `json:"name"`
	Required  bool   `json:"required"`
	Condition string `json:"condition"`
}

type Value interface{}

/*
type Value struct {
	Condition string `json:"condition"`
	Value     interface{}
}
*/

type Attributes map[Key]Value

func (as Attributes) ConvertConditions() (Conditions, error) {
	cs := Conditions{}
	for k, v := range as {
		fn, ok := ConditionFactories[k.Condition]
		if !ok {
			return nil, errors.Errorf("CCondition %s not found", k.Condition)
		}

		c := fn()
		c.Value(v)

		cs[k.Name] = c
	}

	return cs, nil
}

/*
import "github.com/pkg/errors"

type Attribute struct {
	Name      string
	Value     interface{}
	Required  bool
	Condition string
}

type Attributes []Attribute

func (a *Attribute) ConvertCondition() (Condition, error) {
	fn, ok := ConditionFactories[a.Condition]
	if !ok {
		return nil, errors.Errorf("CCondition %s not found", a.Condition)
	}

	c := fn()
	c.Value(a.Value)

	return c, nil
}

func (as Attributes) ConvertConditions() (Conditions, error) {
	cs := Conditions{}
	for _, a := range as {
		c, e := a.ConvertCondition()
		if e != nil {
			return nil, e
		}
		cs[a.Name] = c
	}

	return cs, nil
}
*/
