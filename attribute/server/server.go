package main

import (
	"context"
	"fmt"
	"net"

	"github.com/golang/protobuf/ptypes"
	"github.com/gunsluo/go-example/attribute/acpb"
	"github.com/gunsluo/go-example/attribute/attr"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	sAddress = "0.0.0.0:19000"
)

type Service struct {
}

func (s *Service) UpsertPredefinedPolicies(ctx context.Context, req *acpb.UpsertPredefinedPoliciesRequest) (*acpb.UpsertPredefinedPoliciesReply, error) {
	if len(req.Policies) == 0 {
		return nil, errors.New("Missing Policies")
	}

	var pps []*attr.PredefinedPolicy
	for _, pp := range req.Policies {
		fmt.Printf("pp: %s %s %v %v", pp.Name, pp.Description, pp.Resources, pp.Actions)
		for _, a := range pp.Attributes {
			fmt.Printf("pp attribute: %s %v", a.Name, a.Required)
			var attributes attr.Attributes
			if a.Value.Condition == "StringEqualCondition" {
				value := &acpb.StringEqualConditionAttributeValue{}
				err := ptypes.UnmarshalAny(a.Value.Attribute, value)
				if err != nil {
					return nil, err
				}
				fmt.Printf("pp attribute: %s %s", a.Value.Condition, value.Value)

				attributes = append(attributes, attr.Attribute{
					Name:      a.Name,
					Value:     value.Value,
					Required:  a.Required,
					Condition: a.Value.Condition,
				})
			}

			pps = append(pps, &attr.PredefinedPolicy{
				Name:        pp.Name,
				Description: pp.Description,
				Resources:   pp.Resources,
				Actions:     pp.Actions,
				Attributes:  attributes,
			})

			conditions, err := attributes.ConvertConditions()
			if err != nil {
				return nil, err
			}

			buffer, err := conditions.MarshalJSON()
			if err != nil {
				return nil, err
			}

			fmt.Printf("\npp attribute to conditions string: %s\n", string(buffer))
		}
	}

	return &acpb.UpsertPredefinedPoliciesReply{}, nil
}

func main() {
	server := grpc.NewServer()
	acpb.RegisterAccessControlServer(server, &Service{})

	listener, err := net.Listen("tcp", sAddress)
	if err != nil {
		panic(err)
	}

	logrus.WithField("addr", sAddress).Println("Starting server")
	server.Serve(listener)
}
