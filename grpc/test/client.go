package main

import (
	"context"
	"log"
	"time"

	"gitlab.com/tesgo/kit/proto/ac/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	grpcServerAddr = "127.0.0.1:7001"
)

func main() {
	conn, err := grpc.Dial(
		grpcServerAddr,
		grpc.WithInsecure(),
		grpc.WithBackoffConfig(grpc.BackoffConfig{
			MaxDelay: time.Hour,
		}))
	if err != nil {
		panic("failed to dialing grpc server")
	}

	client := pb.NewAccessControlClient(conn)
	ctx := context.Background()

	// test

	testCreateRole(ctx, client)
	testCreateRole2(ctx, client)
	testCreateRole3(ctx, client)
	testListRoles(ctx, client)
	testBindRole(ctx, client)

	testBindRoles(ctx, client)
	testBindRoles2(ctx, client)
	testBindRoleGroup(ctx, client)
	testListSubjects(ctx, client)
	testRegisterResource(ctx, client)
	testRegisterResources(ctx, client)
	testListResources(ctx, client)
	testRegisterAction(ctx, client)
	testRegisterActions(ctx, client)
	testListActions(ctx, client)

	testCreatePolicy(ctx, client)
	testCreatePolicy2(ctx, client)
	testUpdatePolicy(ctx, client)
	testGetPolicy(ctx, client)
	testListPolicies(ctx, client)
	testListPolicies2(ctx, client)
	testVerify(ctx, client)
	testVerifyList(ctx, client)

	testRemoveAction(ctx, client)
	testRemoveResource(ctx, client)
	testRemovePolicy(ctx, client)

	testRemoveRole(ctx, client)
	testUnbindRoles(ctx, client)
	testUnbindRoles2(ctx, client)
	testUnbindRoleGroup(ctx, client)

	conn.Close()
}

func testCreateRole(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.CreateRole(ctx,
		&pb.CreateRoleRequest{
			Role: &pb.Role{Name: "target:reach:visitor", Description: "test"},
		})

	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK && errStatus.Code() != codes.AlreadyExists {
		log.Println("create role to access control grpc:", errStatus.Err())
		return
	}

	log.Println("create role to access control grpc success.")
}

func testCreateRole2(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.CreateRole(ctx,
		&pb.CreateRoleRequest{
			Role: &pb.Role{Name: "demo:test:role1", Description: "test"},
		})

	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK && errStatus.Code() != codes.AlreadyExists {
		log.Println("create role to access control grpc:", errStatus.Err())
		return
	}

	log.Println("create role to access control grpc success.")
}

func testCreateRole3(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.CreateRole(ctx,
		&pb.CreateRoleRequest{
			Role: &pb.Role{Name: "demo:test:role2", Description: "test"},
		})

	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK && errStatus.Code() != codes.AlreadyExists {
		log.Println("create role to access control grpc:", errStatus.Err())
		return
	}

	log.Println("create role to access control grpc success.")
}

func testRegisterResource(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.RegisterResource(ctx,
		&pb.RegisterResourceRequest{
			&pb.Resource{
				Name:        "target:reach:/r1",
				Description: "test",
			},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK && errStatus.Code() != codes.AlreadyExists {
		log.Println("register resource to access control grpc:", errStatus.Err())
		return
	}

	log.Println("register resource to access control grpc success.")
}

func testRegisterResources(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.RegisterResources(ctx,
		&pb.RegisterResourcesRequest{
			Resources: []*pb.Resource{
				&pb.Resource{
					Name:        "target:reach:/r2",
					Description: "test",
				},
				&pb.Resource{
					Name:        "target:reach:/<.+>",
					Description: "test",
				},
			},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK && errStatus.Code() != codes.AlreadyExists {
		log.Println("register resources to access control grpc:", errStatus.Err())
		return
	}

	log.Println("register resources to access control grpc success.")
}

func testRegisterAction(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.RegisterAction(context.Background(),
		&pb.RegisterActionRequest{
			Action: &pb.Action{
				Name:        "reach:get",
				Description: "test",
			},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK && errStatus.Code() != codes.AlreadyExists {
		log.Println("register action to access control grpc:", errStatus.Err())
		return
	}

	log.Println("register action to access control grpc success.")
}

func testRegisterActions(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.RegisterActions(ctx,
		&pb.RegisterActionsRequest{
			Actions: []*pb.Action{
				&pb.Action{
					Name:        "reach:create",
					Description: "test",
				},
				&pb.Action{
					Name:        "reach:update",
					Description: "test",
				},
			},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK && errStatus.Code() != codes.AlreadyExists {
		log.Println("register actions to access control grpc:", errStatus.Err())
		return
	}

	log.Println("register actions to access control grpc success.")
}

func testListResources(ctx context.Context, client pb.AccessControlClient) {
	reply, err := client.ListResources(ctx,
		&pb.ListResourcesRequest{
			LikeResource: "target:reach",
			Limit:        10,
			Offset:       0,
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("register resource to access control grpc:", errStatus.Err())
		return
	}

	log.Println("register resource to access control grpc success.", len(reply.Resources))
}

func testListActions(ctx context.Context, client pb.AccessControlClient) {
	reply, err := client.ListActions(ctx,
		&pb.ListActionsRequest{
			LikeAction: "reach",
			Limit:      10,
			Offset:     0,
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("get actions to access control grpc:", errStatus.Err())
		return
	}

	log.Println("get actions to access control grpc success.", len(reply.Actions))
}

func testListRoles(ctx context.Context, client pb.AccessControlClient) {
	reply, err := client.ListRoles(ctx,
		&pb.ListRolesRequest{})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("get roles to access control grpc:", errStatus.Err())
		return
	}

	log.Println("get roles to access control grpc success.", len(reply.Roles))
}

func testBindRole(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.BindRole(ctx,
		&pb.BindRoleRequest{
			User: &pb.User{Type: pb.SUBJECT, Name: "tom"},
			Role: "target:reach:visitor",
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK && errStatus.Code() != codes.AlreadyExists {
		log.Println("role binding to access control grpc:", errStatus.Err())
		return
	}

	log.Println("role binding to access control grpc success.")
}

func testBindRoles(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.BindRoles(ctx,
		&pb.BindRolesRequest{
			User:  &pb.User{Type: pb.SUBJECT, Name: "luoji"},
			Roles: []string{"demo:test:role1", "demo:test:role2"},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("roles binding to access control grpc:", errStatus.Err())
		return
	}

	log.Println("roles binding to access control grpc success.")
}

func testBindRoles2(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.BindRoles(ctx,
		&pb.BindRolesRequest{
			User:  &pb.User{Type: pb.ROLE, Name: "target:reach:visitor"},
			Roles: []string{"demo:test:role1", "demo:test:role2"},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("roles binding to access control grpc:", errStatus.Err())
		return
	}

	log.Println("roles binding to access control grpc success.")
}

func testBindRoleGroup(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.BindRoleGroup(ctx,
		&pb.BindRoleGroupRequest{
			RoleGroup: "target:reach:visitor",
			Roles:     []string{"demo:test:role1", "demo:test:role2"},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("roles binding to access control grpc:", errStatus.Err())
		return
	}

	log.Println("roles binding to access control grpc success.")
}

func testListSubjects(ctx context.Context, client pb.AccessControlClient) {
	reply, err := client.ListUsersByRole(ctx,
		&pb.ListUsersByRoleRequest{
			LikeRole: "target:reach:visitor",
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("get subjects to access control grpc:", errStatus.Err())
		return
	}

	log.Println("get subjects to access control grpc success.", len(reply.Users))
}

func testCreatePolicy(ctx context.Context, client pb.AccessControlClient) {
	id, err := client.CreatePolicy(ctx,
		&pb.CreatePolicyRequest{
			Policy: &pb.Policy{
				Id:          "reach-policy1",
				Description: "policy of visitor",
				Effect:      pb.ALLOW,
				Subjects:    []string{"target:reach:visitor"},
				Resources:   []string{"target:reach:/r1"},
				Actions:     []string{"reach:get"},
			},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK && errStatus.Code() != codes.AlreadyExists {
		log.Println("create policy to access control grpc:", errStatus.Err())
		return
	}

	log.Println("create policy to access control grpc success. id:", id)
}

func testCreatePolicy2(ctx context.Context, client pb.AccessControlClient) {
	reply, err := client.CreatePolicy(ctx,
		&pb.CreatePolicyRequest{
			Policy: &pb.Policy{
				Id:          "reach-policy2",
				Description: "policy of visitor",
				Effect:      pb.ALLOW,
				Subjects:    []string{"target:reach:admin"},
				Resources: []string{
					"target:reach:/r1",
					"target:reach:/r2",
					"target:reach:/<.+>",
				},
				Actions:    []string{"reach:update"},
				Conditions: []byte(`{"clientIP":{"type":"CIDRCondition","options":{"cidr":"127.0.0.1/32"}}}`),
			},
		})

	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK && errStatus.Code() != codes.AlreadyExists {
		log.Println("create policy to access control grpc:", errStatus.Err())
		return
	}

	log.Println("create policy to access control grpc success. id:", reply.Id)
}

func testUpdatePolicy(ctx context.Context, client pb.AccessControlClient) {
	reply, err := client.UpdatePolicy(ctx,
		&pb.UpdatePolicyRequest{
			Policy: &pb.Policy{
				Id:          "reach-policy2",
				Description: "policy of visitor",
				Effect:      pb.ALLOW,
				Subjects:    []string{"target:reach:admin"},
				Resources: []string{
					"target:reach:/r1",
					"target:reach:/r2",
					"target:reach:/<.+>",
				},
				Actions:    []string{"reach:create"},
				Conditions: []byte(`{"clientIP":{"type":"CIDRCondition","options":{"cidr":"127.0.0.1/24"}}}`),
			},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("update policy to access control grpc:", errStatus.Err())
		return
	}

	log.Println("update policy to access control grpc success. id:", reply.Id)
}

func testGetPolicy(ctx context.Context, client pb.AccessControlClient) {
	reply, err := client.GetPolicy(ctx,
		&pb.GetPolicyRequest{
			Id: "reach-policy1",
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("get policy to access control grpc:", errStatus.Err())
		return
	}

	log.Println("get policy to access control grpc success.", reply.Policy)
}

func testListPolicies(ctx context.Context, client pb.AccessControlClient) {
	reply, err := client.ListPolicies(ctx,
		&pb.ListPoliciesRequest{
			Subject: "tom",
			Limit:   10,
			Offset:  0,
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("get policies to access control grpc:", errStatus.Err())
		return
	}

	log.Println("get policies to access control grpc success.", len(reply.Policies))
}

func testListPolicies2(ctx context.Context, client pb.AccessControlClient) {
	reply, err := client.ListPolicies(ctx,
		&pb.ListPoliciesRequest{
			Role:   "target:reach:admin",
			Limit:  10,
			Offset: 0,
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("get policies to access control grpc:", errStatus.Err())
		return
	}

	log.Println("get policies to access control grpc success.", len(reply.Policies))
}

func testVerify(ctx context.Context, client pb.AccessControlClient) {
	reply, err := client.Verify(ctx,
		&pb.VerifyRequest{
			Subject:  "tom",
			Resource: "target:reach:/r1",
			Action:   "reach:get",
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("verify subject permissions to access control grpc:", errStatus.Err())
		return
	}

	log.Println("verify subject permissions to access control grpc result:", reply.Allowed)
}

func testVerifyList(ctx context.Context, client pb.AccessControlClient) {
	reply, err := client.VerifyList(ctx,
		&pb.VerifyListRequest{
			List: []*pb.VerifyRequest{
				&pb.VerifyRequest{
					Subject:  "tom",
					Resource: "target:reach:/r1",
					Action:   "reach:get",
				},
				&pb.VerifyRequest{
					Subject:  "target:reach:admin",
					Resource: "target:reach:/r2",
					Action:   "reach:create",
					Context:  []byte(`{"clientIP": "127.0.0.5"}`),
				},
			},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("verify list subject permissions to access control grpc:", errStatus.Err())
		return
	}

	for _, res := range reply.Results {
		log.Println("verify list subject permissions to access control grpc result:", res.Allowed)
	}
}

func testRemoveAction(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.RemoveAction(ctx,
		&pb.RemoveActionRequest{
			Action: "reach:get",
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("remove action to access control grpc:", errStatus.Err())
		return
	}

	_, err = client.RemoveAction(ctx,
		&pb.RemoveActionRequest{
			Action: "reach:create",
		})
	errStatus = status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("remove action to access control grpc:", errStatus.Err())
		return
	}

	_, err = client.RemoveAction(ctx,
		&pb.RemoveActionRequest{
			Action: "reach:update",
		})
	errStatus = status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("remove action to access control grpc:", errStatus.Err())
		return
	}

	log.Println("remove action to access control grpc success.")
}

func testRemoveResource(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.RemoveResource(ctx,
		&pb.RemoveResourceRequest{
			Resource: "target:reach:/r1",
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("remove resource to access control grpc:", errStatus.Err())
		return
	}

	_, err = client.RemoveResource(ctx,
		&pb.RemoveResourceRequest{
			Resource: "target:reach:/r2",
		})
	errStatus = status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("remove resource to access control grpc:", errStatus.Err())
		return
	}

	_, err = client.RemoveResource(ctx,
		&pb.RemoveResourceRequest{
			Resource: "target:reach:/<.+>",
		})
	errStatus = status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("remove resource to access control grpc:", errStatus.Err())
		return
	}

	log.Println("remove resource to access control grpc success.")
}

func testRemovePolicy(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.RemovePolicy(ctx,
		&pb.RemovePolicyRequest{
			Id: "reach-policy1",
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("remove policy to access control grpc:", errStatus.Err())
		return
	}

	_, err = client.RemovePolicy(ctx,
		&pb.RemovePolicyRequest{
			Id: "reach-policy2",
		})
	errStatus = status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("remove policy to access control grpc:", errStatus.Err())
		return
	}

	log.Println("remove policy to access control grpc success.")
}

func testRemoveRole(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.RemoveRole(ctx,
		&pb.RemoveRoleRequest{
			Role: "target:reach:visitor",
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("remove role to access control grpc:", errStatus.Err())
		return
	}

	log.Println("remove role to access control grpc success.")
}

func testUnbindRoles(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.UnbindRoles(ctx,
		&pb.UnbindRolesRequest{
			User:  &pb.User{Type: pb.SUBJECT, Name: "luoji"},
			Roles: []string{"demo:test:role1", "demo:test:role2"},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("remove role to access control grpc:", errStatus.Err())
		return
	}

	log.Println("remove role to access control grpc success.")
}

func testUnbindRoles2(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.UnbindRoles(ctx,
		&pb.UnbindRolesRequest{
			User:  &pb.User{Type: pb.ROLE, Name: "target:reach:visitor"},
			Roles: []string{"demo:test:role1", "demo:test:role2"},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("remove role to access control grpc:", errStatus.Err())
		return
	}

	log.Println("remove role to access control grpc success.")
}

func testUnbindRoleGroup(ctx context.Context, client pb.AccessControlClient) {
	_, err := client.UnbindRoleGroup(ctx,
		&pb.UnbindRoleGroupRequest{
			RoleGroup: "target:reach:visitor",
			Roles:     []string{"demo:test:role1", "demo:test:role2"},
		})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Println("roles unbinding to access control grpc:", errStatus.Err())
		return
	}

	log.Println("roles unbinding to access control grpc success.")
}
