package main

import "fmt"

type Group struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Label       string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
}

type PredefinedPolicy struct {
	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Resources   []string `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty"`
	Actions     []string `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	Attributes  []byte   `protobuf:"bytes,5,opt,name=conditions,proto3" json:"conditions,omitempty"`
	//Conditions  []byte   `protobuf:"bytes,5,opt,name=conditions,proto3" json:"conditions,omitempty"`
}

func main() {
	/*
		group := &Group{
			Name:        "region-manager",
			Label:       "Region Manager",
			Description: "this is a test group",
		}

		pp := &PredefinedPolicy{
			Name:        "region-manager:read-access",
			Description: "test pp",
			Resources:   []string{"rs1", "rs2"},
			Actions:     []string{"read"},
			//Conditions:  []string{},
			//Conditions: Conditions{
			//	"region": &StringEqualCondition{
			//		Equals: "chengdu",
			//	},
			//},
		}

		// group + pp -> p
	*/

	/*
		attr := Attribute{
			Name:      "region",
			Value:     "chengdu",
			Required:  true,
			Condition: "StringEqualCondition",
		}

		var attrs Attributes
		attrs = append(attrs, attr)
	*/

	attrs := Attributes{
		Key{
			Name:      "region",
			Required:  true,
			Condition: "StringEqualCondition",
		}: "chengdu",
	}

	cs, err := attrs.ConvertConditions()
	if err != nil {
		panic(err)
	}

	buffer, err := cs.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("------------>", cs, string(buffer))

	/*
		pp := &PredefinedPolicy{
			Name:        "region-manager:read-access",
			Description: "test pp",
			Resources:   []string{"rs1", "rs2"},
			Actions:     []string{"read"},
			Attributes:  attrs,
		}
	*/
}
