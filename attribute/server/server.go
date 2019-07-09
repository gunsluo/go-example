package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/gunsluo/go-example/attribute/acpb"
	"github.com/gunsluo/go-example/attribute/group"
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

	var pps []*group.PredefinedPolicy
	for _, pp := range req.Policies {
		fmt.Printf("pp: %s %s %v %v", pp.Name, pp.Description, pp.Resources, pp.Actions)

		conditions, err := group.ConvertConditions(pp.Conditions)
		if err != nil {
			return nil, err
		}

		npp := &group.PredefinedPolicy{
			Name:        pp.Name,
			Description: pp.Description,
			Resources:   pp.Resources,
			Actions:     pp.Actions,
			Conditions:  conditions,
		}

		pps = append(pps, npp)

		buffer, err := json.Marshal(npp)
		if err != nil {
			return nil, err
		}
		fmt.Printf("\npp json string: %s\n", string(buffer))
	}

	return &acpb.UpsertPredefinedPoliciesReply{}, nil
}

func (s *Service) UpsertPoliciesUsingDTO(ctx context.Context, req *acpb.UpsertPoliciesUsingDTORequest) (*acpb.UpsertPoliciesUsingDTOReply, error) {
	if len(req.Dtos) == 0 {
		return nil, errors.New("Missing Dtos")
	}

	for _, dto := range req.Dtos {
		// get pp by name
		pp, err := mockGetPPByName(dto.PpName)
		if err != nil {
			return nil, err
		}

		if len(pp.Conditions) > 0 {
			all, err := group.ConvertAttributes(dto.AttributeValues)
			if err != nil {
				return nil, err
			}

			conditions, err := pp.Conditions.ConvertConditions(all)
			if err != nil {
				return nil, err
			}

			buffer, err := conditions.MarshalJSON()
			if err != nil {
				return nil, err
			}

			fmt.Printf("\npp conditions json string: %s\n", string(buffer))
		}
	}

	return &acpb.UpsertPoliciesUsingDTOReply{}, nil
}

func mockGetPPByName(name string) (*group.PredefinedPolicy, error) {
	jsonstr := `{"name":"test pp","description":"this is a test pp","resources":["r1","r2"],"actions":["a1","a2"],"conditions":[{"name":"region","type":"StringEqualCondition","options":{"attributes":[{"name":"equals","type":"string","required":true,"default":"chengdu"}]}}]}`
	pp := &group.PredefinedPolicy{}
	err := json.Unmarshal([]byte(jsonstr), pp)
	if err != nil {
		return nil, err
	}

	pp.Name = name
	return pp, err
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
