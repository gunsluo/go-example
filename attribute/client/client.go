package main

import (
	"context"

	"github.com/gogo/protobuf/types"
	"github.com/gunsluo/go-example/attribute/acpb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	acClient := acpb.NewAccessControlClient(conn)

	value := &acpb.StringAttributeValue{Value: "chengdu"}
	any, err := types.MarshalAny(value)
	if err != nil {
		panic(err)
	}

	_, err = acClient.UpsertPredefinedPolicies(context.Background(),
		&acpb.UpsertPredefinedPoliciesRequest{
			Policies: []*acpb.PredefinedPolicy{
				&acpb.PredefinedPolicy{
					Name:        "test pp",
					Description: "this is a test pp",
					Resources:   []string{"r1", "r2"},
					Actions:     []string{"a1", "a2"},
					Conditions: []*acpb.Condition{
						&acpb.Condition{
							Name: "region",
							Type: "StringEqualCondition",
							Options: &acpb.ConditionOption{
								Attributes: []*acpb.Attribute{
									&acpb.Attribute{
										Name:     "equals",
										Type:     acpb.ATTRIBUTE_TYPE_STRING,
										Required: true,
										Default:  any,
									},
								},
							},
						},
					},
				},
			},
		})
	if err != nil {
		logrus.WithError(err).Fatal("unable to upsert pp")
	}

	logrus.Info("reply: ok")
}
