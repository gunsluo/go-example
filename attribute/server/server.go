package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

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

		conditions, err := attr.ConvertConditions(pp.Conditions)
		if err != nil {
			return nil, err
		}

		pps = append(pps, &attr.PredefinedPolicy{
			Name:        pp.Name,
			Description: pp.Description,
			Resources:   pp.Resources,
			Actions:     pp.Actions,
			Conditions:  conditions,
		})

		buffer, err := json.Marshal(pp)
		fmt.Printf("\npp json string: %s\n", string(buffer))

		buffer, err = conditions.MarshalJSON()
		if err != nil {
			return nil, err
		}

		fmt.Printf("\npp attribute to conditions string: %s\n", string(buffer))
	}

	return &acpb.UpsertPredefinedPoliciesReply{}, nil
}

func (s *Service) UpsertPoliciesUsingDTO(ctx context.Context, req *acpb.UpsertPoliciesUsingDTORequest) (*acpb.UpsertPoliciesUsingDTOReply, error) {
	return nil, nil
}

// protoc -I=. --gogoslick_out=plugins=grpc,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:. ac.proto
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
