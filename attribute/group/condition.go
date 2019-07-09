package group

import (
	"github.com/gogo/protobuf/types"
	"github.com/gunsluo/go-example/attribute/acpb"
	"github.com/gunsluo/go-example/attribute/attr"
	"github.com/pkg/errors"
)

type PredefinedPolicy struct {
	Name        string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Resources   []string   `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty"`
	Actions     []string   `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	Conditions  Conditions `protobuf:"bytes,5,opt,name=conditions,proto3" json:"conditions,omitempty"`
}

type Attribute struct {
	Name     string      `json:"name,omitempty"`
	Type     string      `json:"type,omitempty"`
	Required bool        `json:"required,omitempty"`
	Value    interface{} `json:"default,omitempty"`
}

type ConditionOption struct {
	Expression string       `json:"expression,omitempty"`
	Attributes []*Attribute `json:"attributes,omitempty"`
}

type Condition struct {
	Name    string           `json:"name,omitempty"`
	Type    string           `json:"type,omitempty"`
	Options *ConditionOption `json:"options,omitempty"`
}

type Conditions []*Condition

func ConvertCondition(c *acpb.Condition) (*Condition, error) {
	nc := &Condition{
		Name: c.Name,
		Type: c.Type,
	}

	if c.Options != nil {
		nc.Options = &ConditionOption{
			Expression: c.Options.Expression,
		}

		attributes, err := convertAttributes(c.Options.Attributes)
		if err != nil {
			return nil, err
		}
		nc.Options.Attributes = attributes
		/*
			for _, a := range c.Options.Attributes {
				na := &Attribute{
					Name:     a.Name,
					Required: a.Required,
				}

				switch a.Type {
				case acpb.ATTRIBUTE_TYPE_STRING:
					if a.Value != nil {
						v := &acpb.StringAttributeValue{}
						err := types.UnmarshalAny(a.Value, v)
						if err != nil {
							return nil, err
						}
						na.Value = v.Value
					}
					na.Type = "string"
					nc.Options.Attributes = append(nc.Options.Attributes, na)
				case acpb.ATTRIBUTE_TYPE_NUMBER:
					if a.Value != nil {
						v := &acpb.NumberAttributeValue{}
						err := types.UnmarshalAny(a.Value, v)
						if err != nil {
							return nil, err
						}
						na.Value = v.Value
					}
					na.Type = "number"
					nc.Options.Attributes = append(nc.Options.Attributes, na)
				case acpb.ATTRIBUTE_TYPE_BOOLEAN:
					if a.Value != nil {
						v := &acpb.BooleanAttributeValue{}
						err := types.UnmarshalAny(a.Value, v)
						if err != nil {
							return nil, err
						}
						na.Value = v.Value
					}
					na.Type = "boolean"
					nc.Options.Attributes = append(nc.Options.Attributes, na)
				}
			}
		*/

	}

	return nc, nil
}

func ConvertConditions(cs []*acpb.Condition) (Conditions, error) {
	var ncs Conditions
	for _, c := range cs {
		nc, err := ConvertCondition(c)
		if err != nil {
			return nil, err
		}
		ncs = append(ncs, nc)
	}

	return ncs, nil
}

func (c *Condition) ConvertCondition(values map[string]interface{}) (attr.Condition, error) {
	fn, ok := attr.ConditionFactories[c.Type]
	if !ok {
		return nil, errors.Errorf("Condition %s not found", c.Type)
	}

	nc := fn()
	if c.Options != nil {
		nvs := make(map[string]interface{})
		for _, a := range c.Options.Attributes {
			if v, ok := values[a.Name]; ok {
				nvs[a.Name] = v
			} else {
				if a.Required {
					return nil, errors.Errorf("attribute %s is required", a.Name)
				}
				if a.Value != nil {
					nvs[a.Name] = a.Value
				}
			}
		}

		nc.Values(c.Options.Expression, values)
	}

	return nc, nil
}

func (cs Conditions) ConvertConditions(all map[string]map[string]interface{}) (attr.Conditions, error) {
	ncs := attr.Conditions{}
	for _, c := range cs {
		var values map[string]interface{}
		if v, ok := all[c.Name]; ok {
			values = v
		} else {
			values = make(map[string]interface{})
		}

		nc, err := c.ConvertCondition(values)
		if err != nil {
			return nil, err
		}
		ncs[c.Name] = nc
	}

	return ncs, nil
}

func ConvertAttributes(attributesTable map[string]*acpb.PolicyDTO_Attributes) (map[string]map[string]interface{}, error) {
	all := make(map[string]map[string]interface{})
	for k, v := range attributesTable {
		attributes, err := convertAttributes(v.Attributes)
		if err != nil {
			return nil, err
		}

		values := make(map[string]interface{})
		for _, a := range attributes {
			values[a.Name] = a.Value
		}
		all[k] = values
	}

	return all, nil
}

func convertAttributes(attributes []*acpb.Attribute) ([]*Attribute, error) {
	var attrs []*Attribute
	for _, a := range attributes {
		na := &Attribute{
			Name:     a.Name,
			Required: a.Required,
		}

		switch a.Type {
		case acpb.ATTRIBUTE_TYPE_STRING:
			if a.Value != nil {
				v := &acpb.StringAttributeValue{}
				err := types.UnmarshalAny(a.Value, v)
				if err != nil {
					return nil, err
				}
				na.Value = v.Value
			}
			na.Type = "string"
			attrs = append(attrs, na)
		case acpb.ATTRIBUTE_TYPE_NUMBER:
			if a.Value != nil {
				v := &acpb.NumberAttributeValue{}
				err := types.UnmarshalAny(a.Value, v)
				if err != nil {
					return nil, err
				}
				na.Value = v.Value
			}
			na.Type = "number"
			attrs = append(attrs, na)
		case acpb.ATTRIBUTE_TYPE_BOOLEAN:
			if a.Value != nil {
				v := &acpb.BooleanAttributeValue{}
				err := types.UnmarshalAny(a.Value, v)
				if err != nil {
					return nil, err
				}
				na.Value = v.Value
			}
			na.Type = "boolean"
			attrs = append(attrs, na)
		}
	}

	return attrs, nil
}

//AttributeValues map[string]*PolicyDTO_Attributes `protobuf:"bytes,8,rep,name=attribute_values,json=attributeValues,proto3" json:"attribute_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
