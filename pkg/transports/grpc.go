package transports

import (
	"casbinsvc/pb"
	"casbinsvc/pkg/endpoints"
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

func NewGRPCService(endpoints endpoints.Set) pb.CasbinServer {
	options := []grpctransport.ServerOption{}
	return &grpcService{
		newEnforcer: grpctransport.NewServer(
			endpoints.NewEnforcerEndpoint,
			decodeGRPCNewEnforcerRequest,
			encodeGRPCNewEnforcerResponse,
			options...),
		enforce: grpctransport.NewServer(
			endpoints.EnforceEndpoint,
			decodeGRPCEnforceRequest,
			encodeGRPCEnforceResponse,
			options...),
		loadPolicy: grpctransport.NewServer(
			endpoints.LoadPolicyEndpoint,
			decodeGRPCLoadPolicyRequest,
			encodeGRPCLoadPolicyResponse,
			options...),
		savePolicy: grpctransport.NewServer(
			endpoints.SavePolicyEndpoint,
			decodeGRPCSavePolicyRequest,
			encodeGRPCSavePolicyResponse,
			options...),
		newAdapter: grpctransport.NewServer(
			endpoints.NewAdapterEndpoint,
			decodeGRPCNewAdapterRequest,
			encodeGRPCNewAdapterResponse,
			options...),
		addPolicy: grpctransport.NewServer(
			endpoints.AddPolicyEndpoint,
			decodeGRPCAddPolicyRequest,
			encodeGRPCAddPolicyResponse,
			options...),
		addNamedPolicy: grpctransport.NewServer(
			endpoints.AddNamedPolicyEndpoint,
			decodeGRPCAddNamedPolicyRequest,
			encodeGRPCAddNamedPolicyResponse,
			options...),
		removePolicy: grpctransport.NewServer(
			endpoints.RemovePolicyEndpoint,
			decodeGRPCRemovePolicyRequest,
			encodeGRPCRemovePolicyResponse,
			options...),
		removeNamedPolicy: grpctransport.NewServer(
			endpoints.RemoveNamedPolicyEndpoint,
			decodeGRPCRemoveNamedPolicyRequest,
			encodeGRPCRemoveNamedPolicyResponse,
			options...),
		removeFilteredPolicy: grpctransport.NewServer(
			endpoints.RemoveFilteredPolicyEndpoint,
			decodeGRPCRemoveFilteredPolicyRequest,
			encodeGRPCRemoveFilteredPolicyResponse,
			options...),
		removeFilteredNamedPolicy: grpctransport.NewServer(
			endpoints.RemoveFilteredNamedPolicyEndpoint,
			decodeGRPCRemoveFilteredNamedPolicyRequest,
			encodeGRPCRemoveFilteredNamedPolicyResponse,
			options...),
		getPolicy: grpctransport.NewServer(
			endpoints.GetPolicyEndpoint,
			decodeGRPCGetPolicyRequest,
			encodeGRPCGetPolicyResponse,
			options...),
		getNamedPolicy: grpctransport.NewServer(
			endpoints.GetNamedPolicyEndpoint,
			decodeGRPCGetNamedPolicyRequest,
			encodeGRPCGetNamedPolicyResponse,
			options...),
		getFilteredPolicy: grpctransport.NewServer(
			endpoints.GetFilteredPolicyEndpoint,
			decodeGRPCGetFilteredPolicyRequest,
			encodeGRPCGetFilteredPolicyResponse,
			options...),
		getFilteredNamedPolicy: grpctransport.NewServer(
			endpoints.GetFilteredNamedPolicyEndpoint,
			decodeGRPCGetFilteredNamedPolicyRequest,
			encodeGRPCGetFilteredNamedPolicyResponse,
			options...),
		addGroupingPolicy: grpctransport.NewServer(
			endpoints.AddGroupingPolicyEndpoint,
			decodeGRPCAddGroupingPolicyRequest,
			encodeGRPCAddGroupingPolicyResponse,
			options...),
		addNamedGroupingPolicy: grpctransport.NewServer(
			endpoints.AddNamedGroupingPolicyEndpoint,
			decodeGRPCAddNamedGroupingPolicyRequest,
			encodeGRPCAddNamedGroupingPolicyResponse,
			options...),
		removeGroupingPolicy: grpctransport.NewServer(
			endpoints.RemoveGroupingPolicyEndpoint,
			decodeGRPCRemoveGroupingPolicyRequest,
			encodeGRPCRemoveGroupingPolicyResponse,
			options...),
		removeNamedGroupingPolicy: grpctransport.NewServer(
			endpoints.RemoveNamedGroupingPolicyEndpoint,
			decodeGRPCRemoveNamedGroupingPolicyRequest,
			encodeGRPCRemoveNamedGroupingPolicyResponse,
			options...),
		removeFilteredGroupingPolicy: grpctransport.NewServer(
			endpoints.RemoveFilteredGroupingPolicyEndpoint,
			decodeGRPCRemoveFilteredGroupingPolicyRequest,
			encodeGRPCRemoveFilteredGroupingPolicyResponse,
			options...),
		removeFilteredNamedGroupingPolicy: grpctransport.NewServer(
			endpoints.RemoveFilteredNamedGroupingPolicyEndpoint,
			decodeGRPCRemoveFilteredNamedGroupingPolicyRequest,
			encodeGRPCRemoveFilteredNamedGroupingPolicyResponse,
			options...),
		getGroupingPolicy: grpctransport.NewServer(
			endpoints.GetGroupingPolicyEndpoint,
			decodeGRPCGetGroupingPolicyRequest,
			encodeGRPCGetGroupingPolicyResponse,
			options...),
		getNamedGroupingPolicy: grpctransport.NewServer(
			endpoints.GetNamedGroupingPolicyEndpoint,
			decodeGRPCGetNamedGroupingPolicyRequest,
			encodeGRPCGetNamedGroupingPolicyResponse,
			options...),
		getFilteredGroupingPolicy: grpctransport.NewServer(
			endpoints.GetFilteredGroupingPolicyEndpoint,
			decodeGRPCGetFilteredGroupingPolicyRequest,
			encodeGRPCGetFilteredGroupingPolicyResponse,
			options...),
		getFilteredNamedGroupingPolicy: grpctransport.NewServer(
			endpoints.GetFilteredNamedGroupingPolicyEndpoint,
			decodeGRPCGetFilteredNamedGroupingPolicyRequest,
			encodeGRPCGetFilteredNamedGroupingPolicyResponse,
			options...),
		getAllSubjects: grpctransport.NewServer(
			endpoints.GetAllSubjectsEndpoint,
			decodeGRPCGetAllSubjectsRequest,
			encodeGRPCGetAllSubjectsResponse,
			options...),
		getAllNamedSubjects: grpctransport.NewServer(
			endpoints.GetAllNamedSubjectsEndpoint,
			decodeGRPCGetAllNamedSubjectsRequest,
			encodeGRPCGetAllNamedSubjectsResponse,
			options...),
		getAllObjects: grpctransport.NewServer(
			endpoints.GetAllObjectsEndpoint,
			decodeGRPCGetAllObjectsRequest,
			encodeGRPCGetAllObjectsResponse,
			options...),
		getAllNamedObjects: grpctransport.NewServer(
			endpoints.GetAllNamedObjectsEndpoint,
			decodeGRPCGetAllNamedObjectsRequest,
			encodeGRPCGetAllNamedObjectsResponse,
			options...),
		getAllActions: grpctransport.NewServer(
			endpoints.GetAllActionsEndpoint,
			decodeGRPCGetAllActionsRequest,
			encodeGRPCGetAllActionsResponse,
			options...),
		getAllNamedActions: grpctransport.NewServer(
			endpoints.GetAllNamedActionsEndpoint,
			decodeGRPCGetAllNamedActionsRequest,
			encodeGRPCGetAllNamedActionsResponse,
			options...),
		getAllRoles: grpctransport.NewServer(
			endpoints.GetAllRolesEndpoint,
			decodeGRPCGetAllRolesRequest,
			encodeGRPCGetAllRolesResponse,
			options...),
		getAllNamedRoles: grpctransport.NewServer(
			endpoints.GetAllNamedRolesEndpoint,
			decodeGRPCGetAllNamedRolesRequest,
			encodeGRPCGetAllNamedRolesResponse,
			options...),
		hasPolicy: grpctransport.NewServer(
			endpoints.HasPolicyEndpoint,
			decodeGRPCHasPolicyRequest,
			encodeGRPCHasPolicyResponse,
			options...),
		hasNamedPolicy: grpctransport.NewServer(
			endpoints.HasNamedPolicyEndpoint,
			decodeGRPCHasNamedPolicyRequest,
			encodeGRPCHasNamedPolicyResponse,
			options...),
		hasGroupingPolicy: grpctransport.NewServer(
			endpoints.HasGroupingPolicyEndpoint,
			decodeGRPCHasGroupingPolicyRequest,
			encodeGRPCHasGroupingPolicyResponse,
			options...),
		hasNamedGroupingPolicy: grpctransport.NewServer(
			endpoints.HasNamedGroupingPolicyEndpoint,
			decodeGRPCHasNamedGroupingPolicyRequest,
			encodeGRPCHasNamedGroupingPolicyResponse,
			options...),
		hasRoleForUser: grpctransport.NewServer(
			endpoints.HasRoleForUserEndpoint,
			decodeGRPCHasRoleForUserRequest,
			encodeGRPCHasRoleForUserResponse,
			options...),
		addRoleForUser: grpctransport.NewServer(
			endpoints.AddRoleForUserEndpoint,
			decodeGRPCAddRoleForUserRequest,
			encodeGRPCAddRoleForUserResponse,
			options...),
		deleteRoleForUser: grpctransport.NewServer(
			endpoints.DeleteRoleForUserEndpoint,
			decodeGRPCDeleteRoleForUserRequest,
			encodeGRPCDeleteRoleForUserResponse,
			options...),
		deleteRolesForUser: grpctransport.NewServer(
			endpoints.DeleteRolesForUserEndpoint,
			decodeGRPCDeleteRolesForUserRequest,
			encodeGRPCDeleteRolesForUserResponse,
			options...),
		deleteUser: grpctransport.NewServer(
			endpoints.DeleteUserEndpoint,
			decodeGRPCDeleteUserRequest,
			encodeGRPCDeleteUserResponse,
			options...),
		deleteRole: grpctransport.NewServer(
			endpoints.DeleteRoleEndpoint,
			decodeGRPCDeleteRoleRequest,
			encodeGRPCDeleteRoleResponse,
			options...),
		deletePermission: grpctransport.NewServer(
			endpoints.DeletePermissionEndpoint,
			decodeGRPCDeletePermissionRequest,
			encodeGRPCDeletePermissionResponse,
			options...),
		getRolesForUser: grpctransport.NewServer(
			endpoints.GetRolesForUserEndpoint,
			decodeGRPCGetRolesForUserRequest,
			encodeGRPCGetRolesForUserResponse,
			options...),
		getImplicitRolesForUser: grpctransport.NewServer(
			endpoints.GetImplicitRolesForUserEndpoint,
			decodeGRPCGetImplicitRolesForUserRequest,
			encodeGRPCGetImplicitRolesForUserResponse,
			options...),
		getUsersForRole: grpctransport.NewServer(
			endpoints.GetUsersForRoleEndpoint,
			decodeGRPCGetUsersForRoleRequest,
			encodeGRPCGetUsersForRoleResponse,
			options...),
		addPermissionForUser: grpctransport.NewServer(
			endpoints.AddPermissionForUserEndpoint,
			decodeGRPCAddPermissionForUserRequest,
			encodeGRPCAddPermissionForUserResponse,
			options...),
		deletePermissionForUser: grpctransport.NewServer(
			endpoints.DeletePermissionForUserEndpoint,
			decodeGRPCDeletePermissionForUserRequest,
			encodeGRPCDeletePermissionForUserResponse,
			options...),
		getPermissionsForUser: grpctransport.NewServer(
			endpoints.GetPermissionsForUserEndpoint,
			decodeGRPCGetPermissionsForUserRequest,
			encodeGRPCGetPermissionsForUserResponse,
			options...),
		getImplicitPermissionsForUser: grpctransport.NewServer(
			endpoints.GetImplicitPermissionsForUserEndpoint,
			decodeGRPCGetImplicitPermissionsForUserRequest,
			encodeGRPCGetImplicitPermissionsForUserResponse,
			options...),
		hasPermissionForUser: grpctransport.NewServer(
			endpoints.HasPermissionForUserEndpoint,
			decodeGRPCHasPermissionForUserRequest,
			encodeGRPCHasPermissionForUserResponse,
			options...),
	}
}

type grpcService struct {
	newEnforcer                       grpctransport.Handler
	enforce                           grpctransport.Handler
	loadPolicy                        grpctransport.Handler
	savePolicy                        grpctransport.Handler
	newAdapter                        grpctransport.Handler
	addPolicy                         grpctransport.Handler
	addNamedPolicy                    grpctransport.Handler
	removePolicy                      grpctransport.Handler
	removeNamedPolicy                 grpctransport.Handler
	removeFilteredPolicy              grpctransport.Handler
	removeFilteredNamedPolicy         grpctransport.Handler
	getPolicy                         grpctransport.Handler
	getNamedPolicy                    grpctransport.Handler
	getFilteredPolicy                 grpctransport.Handler
	getFilteredNamedPolicy            grpctransport.Handler
	addGroupingPolicy                 grpctransport.Handler
	addNamedGroupingPolicy            grpctransport.Handler
	removeGroupingPolicy              grpctransport.Handler
	removeNamedGroupingPolicy         grpctransport.Handler
	removeFilteredGroupingPolicy      grpctransport.Handler
	removeFilteredNamedGroupingPolicy grpctransport.Handler
	getGroupingPolicy                 grpctransport.Handler
	getNamedGroupingPolicy            grpctransport.Handler
	getFilteredGroupingPolicy         grpctransport.Handler
	getFilteredNamedGroupingPolicy    grpctransport.Handler
	getAllSubjects                    grpctransport.Handler
	getAllNamedSubjects               grpctransport.Handler
	getAllObjects                     grpctransport.Handler
	getAllNamedObjects                grpctransport.Handler
	getAllActions                     grpctransport.Handler
	getAllNamedActions                grpctransport.Handler
	getAllRoles                       grpctransport.Handler
	getAllNamedRoles                  grpctransport.Handler
	hasPolicy                         grpctransport.Handler
	hasNamedPolicy                    grpctransport.Handler
	hasGroupingPolicy                 grpctransport.Handler
	hasNamedGroupingPolicy            grpctransport.Handler
	hasRoleForUser                    grpctransport.Handler
	addRoleForUser                    grpctransport.Handler
	deleteRoleForUser                 grpctransport.Handler
	deleteRolesForUser                grpctransport.Handler
	deleteUser                        grpctransport.Handler
	deleteRole                        grpctransport.Handler
	deletePermission                  grpctransport.Handler
	getRolesForUser                   grpctransport.Handler
	getImplicitRolesForUser           grpctransport.Handler
	getUsersForRole                   grpctransport.Handler
	addPermissionForUser              grpctransport.Handler
	deletePermissionForUser           grpctransport.Handler
	getPermissionsForUser             grpctransport.Handler
	getImplicitPermissionsForUser     grpctransport.Handler
	hasPermissionForUser              grpctransport.Handler
}

// NewEnforcer implements grpc server interface
func (s *grpcService) NewEnforcer(ctx context.Context, req *pb.NewEnforcerRequest) (*pb.NewEnforcerResponse, error) {
	_, resp, err := s.newEnforcer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.NewEnforcerResponse), nil
}

// Enforce implements grpc server interface
func (s *grpcService) Enforce(ctx context.Context, req *pb.EnforceRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.enforce.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// LoadPolicy implements grpc server interface
func (s *grpcService) LoadPolicy(ctx context.Context, req *pb.EmptyRequest) (*pb.EmptyResponse, error) {
	_, resp, err := s.loadPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.EmptyResponse), nil
}

// SavePolicy implements grpc server interface
func (s *grpcService) SavePolicy(ctx context.Context, req *pb.EmptyRequest) (*pb.EmptyResponse, error) {
	_, resp, err := s.savePolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.EmptyResponse), nil
}

// NewAdapter implements grpc server interface
func (s *grpcService) NewAdapter(ctx context.Context, req *pb.NewAdapterRequest) (*pb.NewAdapterResponse, error) {
	_, resp, err := s.newAdapter.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.NewAdapterResponse), nil
}

// AddPolicy implements grpc server interface
func (s *grpcService) AddPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.addPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// AddNamedPolicy implements grpc server interface
func (s *grpcService) AddNamedPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.addNamedPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// RemovePolicy implements grpc server interface
func (s *grpcService) RemovePolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.removePolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// RemoveNamedPolicy implements grpc server interface
func (s *grpcService) RemoveNamedPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.removeNamedPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// RemoveFilteredPolicy implements grpc server interface
func (s *grpcService) RemoveFilteredPolicy(ctx context.Context, req *pb.FilteredPolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.removeFilteredPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// RemoveFilteredNamedPolicy implements grpc server interface
func (s *grpcService) RemoveFilteredNamedPolicy(ctx context.Context, req *pb.FilteredPolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.removeFilteredNamedPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// GetPolicy implements grpc server interface
func (s *grpcService) GetPolicy(ctx context.Context, req *pb.EmptyRequest) (*pb.Array2DResponse, error) {
	_, resp, err := s.getPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Array2DResponse), nil
}

// GetNamedPolicy implements grpc server interface
func (s *grpcService) GetNamedPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.Array2DResponse, error) {
	_, resp, err := s.getNamedPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Array2DResponse), nil
}

// GetFilteredPolicy implements grpc server interface
func (s *grpcService) GetFilteredPolicy(ctx context.Context, req *pb.FilteredPolicyRequest) (*pb.Array2DResponse, error) {
	_, resp, err := s.getFilteredPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Array2DResponse), nil
}

// GetFilterdNamedPolicy implements grpc server interface
func (s *grpcService) GetFilteredNamedPolicy(ctx context.Context, req *pb.FilteredPolicyRequest) (*pb.Array2DResponse, error) {
	_, resp, err := s.getFilteredNamedPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Array2DResponse), nil
}

// AddGroupingPolicy implements grpc server interface
func (s *grpcService) AddGroupingPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.addGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// AddNamedGroupingPolicy implements grpc server interface
func (s *grpcService) AddNamedGroupingPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.addNamedGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// RemoveGroupingPolicy implements grpc server interface
func (s *grpcService) RemoveGroupingPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.removeGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// RemoveNamedGroupingPolicy implements grpc server interface
func (s *grpcService) RemoveNamedGroupingPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.removeNamedGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// RemoveFilteredGroupingPolicy implements grpc server interface
func (s *grpcService) RemoveFilteredGroupingPolicy(ctx context.Context, req *pb.FilteredPolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.removeFilteredGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// RemoveFilteredNamedGroupingPolicy implements grpc server interface
func (s *grpcService) RemoveFilteredNamedGroupingPolicy(ctx context.Context, req *pb.FilteredPolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.removeFilteredNamedGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// GetGroupingPolicy implements grpc server interface
func (s *grpcService) GetGroupingPolicy(ctx context.Context, req *pb.EmptyRequest) (*pb.Array2DResponse, error) {
	_, resp, err := s.getGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Array2DResponse), nil
}

// GetNamedGroupingPolicy implements grpc server interface
func (s *grpcService) GetNamedGroupingPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.Array2DResponse, error) {
	_, resp, err := s.getNamedGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Array2DResponse), nil
}

// GetFilterdGroupingPolicy implements grpc server interface
func (s *grpcService) GetFilteredGroupingPolicy(ctx context.Context, req *pb.FilteredPolicyRequest) (*pb.Array2DResponse, error) {
	_, resp, err := s.getFilteredGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Array2DResponse), nil
}

// GetFilteredNamedGroupingPolicy implements grpc server interface
func (s *grpcService) GetFilteredNamedGroupingPolicy(ctx context.Context, req *pb.FilteredPolicyRequest) (*pb.Array2DResponse, error) {
	_, resp, err := s.getFilteredNamedGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Array2DResponse), nil
}

// GetAllSubjects implemnets grpc server interface
func (s *grpcService) GetAllSubjects(ctx context.Context, req *pb.EmptyRequest) (*pb.ArrayResponse, error) {
	_, resp, err := s.getAllSubjects.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArrayResponse), nil
}

// GetAllNamedSubjects implements grpc server interface
func (s *grpcService) GetAllNamedSubjects(ctx context.Context, req *pb.SimpleGetRequest) (*pb.ArrayResponse, error) {
	_, resp, err := s.getAllNamedSubjects.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArrayResponse), nil
}

// GetAllObjects implements grpc server interface
func (s *grpcService) GetAllObjects(ctx context.Context, req *pb.EmptyRequest) (*pb.ArrayResponse, error) {
	_, resp, err := s.getAllObjects.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArrayResponse), nil
}

// GetAllNamedObjects implements grpc server interface
func (s *grpcService) GetAllNamedObjects(ctx context.Context, req *pb.SimpleGetRequest) (*pb.ArrayResponse, error) {
	_, resp, err := s.getAllNamedObjects.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArrayResponse), nil
}

// GetAllActions implements grpc server interface
func (s *grpcService) GetAllActions(ctx context.Context, req *pb.EmptyRequest) (*pb.ArrayResponse, error) {
	_, resp, err := s.getAllActions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArrayResponse), nil
}

// GetAllNamedActions implements grpc server interface
func (s *grpcService) GetAllNamedActions(ctx context.Context, req *pb.SimpleGetRequest) (*pb.ArrayResponse, error) {
	_, resp, err := s.getAllNamedActions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArrayResponse), nil
}

// GetAllRoles implements grpc server interface
func (s *grpcService) GetAllRoles(ctx context.Context, req *pb.EmptyRequest) (*pb.ArrayResponse, error) {
	_, resp, err := s.getAllRoles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArrayResponse), nil
}

// GetAllNamedRoles implements grpc server interface
func (s *grpcService) GetAllNamedRoles(ctx context.Context, req *pb.SimpleGetRequest) (*pb.ArrayResponse, error) {
	_, resp, err := s.getAllNamedRoles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArrayResponse), nil
}

// HasPolicy implements grpc server interface
func (s *grpcService) HasPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.hasPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// HasNamedPolicy implements grpc server interface
func (s *grpcService) HasNamedPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.hasNamedPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// HasGroupingPolicy implements grpc server interface
func (s *grpcService) HasGroupingPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.hasGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// HasNamedGroupingPolicy implements grpc server interface
func (s *grpcService) HasNamedGroupingPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.hasNamedGroupingPolicy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// HasRoleForUser implements grpc server interface
func (s *grpcService) HasRoleForUser(ctx context.Context, req *pb.UserRoleRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.hasRoleForUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// AddRoleForUser implements grpc server interface
func (s *grpcService) AddRoleForUser(ctx context.Context, req *pb.UserRoleRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.addRoleForUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// DeleteRoleForUser implements grpc server interface
func (s *grpcService) DeleteRoleForUser(ctx context.Context, req *pb.UserRoleRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.deleteRoleForUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// DeleteRolesForUser implements grpc server interface
func (s *grpcService) DeleteRolesForUser(ctx context.Context, req *pb.UserRoleRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.deleteRolesForUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// DeleteUser implements grpc server interface
func (s *grpcService) DeleteUser(ctx context.Context, req *pb.UserRoleRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// DeleteRole implements grpc server interface
func (s *grpcService) DeleteRole(ctx context.Context, req *pb.UserRoleRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.deleteRole.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// DeletePermission implements grpc server interface
func (s *grpcService) DeletePermission(ctx context.Context, req *pb.PermissionRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.deletePermission.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// GetRolesForUser implements grpc server interface
func (s *grpcService) GetRolesForUser(ctx context.Context, req *pb.UserRoleRequest) (*pb.ArrayResponse, error) {
	_, resp, err := s.getRolesForUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArrayResponse), nil
}

// GetImplicitRolesForUser implements grpc server interface
func (s *grpcService) GetImplicitRolesForUser(ctx context.Context, req *pb.UserRoleRequest) (*pb.ArrayResponse, error) {
	_, resp, err := s.getImplicitRolesForUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArrayResponse), nil
}

// GetUserForRole implements grpc server interface
func (s *grpcService) GetUsersForRole(ctx context.Context, req *pb.UserRoleRequest) (*pb.ArrayResponse, error) {
	_, resp, err := s.getUsersForRole.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArrayResponse), nil
}

// AddPermissionForUser implements grpc server interface
func (s *grpcService) AddPermissionForUser(ctx context.Context, req *pb.PermissionRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.addPermissionForUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// DeletePermissionForUser implements grpc server interface
func (s *grpcService) DeletePermissionForUser(ctx context.Context, req *pb.PermissionRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.deletePermissionForUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// GetPermissionsForUser implements grpc server interface
func (s *grpcService) GetPermissionsForUser(ctx context.Context, req *pb.PermissionRequest) (*pb.Array2DResponse, error) {
	_, resp, err := s.getPermissionsForUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Array2DResponse), nil
}

// GetImplicitPermissionsForUser implements grpc server interface
func (s *grpcService) GetImplicitPermissionsForUser(ctx context.Context, req *pb.PermissionRequest) (*pb.Array2DResponse, error) {
	_, resp, err := s.getImplicitPermissionsForUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Array2DResponse), nil
}

// HasPermissionForUser implements grpc server interface
func (s *grpcService) HasPermissionForUser(ctx context.Context, req *pb.PermissionRequest) (*pb.BoolResponse, error) {
	_, resp, err := s.hasPermissionForUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BoolResponse), nil
}

// decodeGRPCNewEnforcerRequest decodes grpc NewEnforcerRequest to endpoint NewEnforcerRequest for NewEnforcer method
func decodeGRPCNewEnforcerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.NewEnforcerRequest)
	return endpoints.NewEnforcerRequest{ModelText: req.ModelText, Adapter: req.AdapterHandler}, nil
}

// encodeGRPCNewEnforcerResponse encodes endpoint NewEnforcerResponse to grpc NewEnforcerResponse
func encodeGRPCNewEnforcerResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*endpoints.NewEnforcerResponse)
	return pb.NewEnforcerRequest{Handler: int32(resp.Handler), Err: err2str(resp.Err)}, nil
}
func decodeGRPCEnforceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.EnforceRequest)
	return endpoints.EnforceRequest{Enforcer: req.Enforcer, Params: req.Params}, nil
}
func encodeGRPCEnforceResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.EnforceResponse)
	return endpoints.EnforceRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCLoadPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.LoadPolicyRequest)
	return endpoints.LoadPolicyRequest{Enforcer: req.Enforcer}, nil
}
func encodeGRPCLoadPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.LoadPolicyResponse)
	return endpoints.LoadPolicyRequest{Err: err2str(resp.Err)}, nil
}
func decodeGRPCSavePolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.SavePolicyRequest)
	return endpoints.SavePolicyRequest{Enforcer: req.Enforcer}, nil
}
func encodeGRPCSavePolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.SavePolicyResponse)
	return endpoints.SavePolicyRequest{Err: err2str(resp.Err)}, nil
}
func decodeGRPCNewAdapterRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.NewAdapterRequest)
	return endpoints.NewAdapterRequest{DriverName: req.DriverName, ConnectString: req.ConnectString, DbSpecified: req.DbSpecified}, nil
}
func encodeGRPCNewAdapterResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.NewAdapterResponse)
	return endpoints.NewAdapterRequest{E: resp.E, Err: err2str(resp.Err)}, nil
}
func decodeGRPCAddPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.AddPolicyRequest)
	return endpoints.AddPolicyRequest{Enforcer: req.Enforcer, Params: req.Params}, nil
}
func encodeGRPCAddPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.AddPolicyResponse)
	return endpoints.AddPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCAddNamedPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.AddNamedPolicyRequest)
	return endpoints.AddNamedPolicyRequest{Enforcer: req.Enforcer, PType: req.PType, Params: req.Params}, nil
}
func encodeGRPCAddNamedPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.AddNamedPolicyResponse)
	return endpoints.AddNamedPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCRemovePolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.RemovePolicyRequest)
	return endpoints.RemovePolicyRequest{Enforcer: req.Enforcer, Params: req.Params}, nil
}
func encodeGRPCRemovePolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.RemovePolicyResponse)
	return endpoints.RemovePolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCRemoveNamedPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.RemoveNamedPolicyRequest)
	return endpoints.RemoveNamedPolicyRequest{Enforcer: req.Enforcer, PType: req.PType, Params: req.Params}, nil
}
func encodeGRPCRemoveNamedPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.RemoveNamedPolicyResponse)
	return endpoints.RemoveNamedPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCRemoveFilteredPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.RemoveFilteredPolicyRequest)
	return endpoints.RemoveFilteredPolicyRequest{Enforcer: req.Enforcer, FieldIndex: req.FieldIndex, FieldValues: req.FieldValues}, nil
}
func encodeGRPCRemoveFilteredPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.RemoveFilteredPolicyResponse)
	return endpoints.RemoveFilteredPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCRemoveFilteredNamedPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.RemoveFilteredNamedPolicyRequest)
	return endpoints.RemoveFilteredNamedPolicyRequest{Enforcer: req.Enforcer, PType: req.PType, FieldIndex: req.FieldIndex, FieldValues: req.FieldValues}, nil
}
func encodeGRPCRemoveFilteredNamedPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.RemoveFilteredNamedPolicyResponse)
	return endpoints.RemoveFilteredNamedPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetPolicyRequest)
	return endpoints.GetPolicyRequest{Enforcer: req.Enforcer}, nil
}
func encodeGRPCGetPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetPolicyResponse)
	return endpoints.GetPolicyRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetNamedPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetNamedPolicyRequest)
	return endpoints.GetNamedPolicyRequest{Enforcer: req.Enforcer, PType: req.PType}, nil
}
func encodeGRPCGetNamedPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetNamedPolicyResponse)
	return endpoints.GetNamedPolicyRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetFilteredPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetFilteredPolicyRequest)
	return endpoints.GetFilteredPolicyRequest{Enforcer: req.Enforcer, FieldIndex: req.FieldIndex, FieldValues: req.FieldValues}, nil
}
func encodeGRPCGetFilteredPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetFilteredPolicyResponse)
	return endpoints.GetFilteredPolicyRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetFilteredNamedPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetFilteredNamedPolicyRequest)
	return endpoints.GetFilteredNamedPolicyRequest{Enforcer: req.Enforcer, PType: req.PType, FieldIndex: req.FieldIndex, FieldValues: req.FieldValues}, nil
}
func encodeGRPCGetFilteredNamedPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetFilteredNamedPolicyResponse)
	return endpoints.GetFilteredNamedPolicyRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCAddGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.AddGroupingPolicyRequest)
	return endpoints.AddGroupingPolicyRequest{Enforcer: req.Enforcer, Params: req.Params}, nil
}
func encodeGRPCAddGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.AddGroupingPolicyResponse)
	return endpoints.AddGroupingPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCAddNamedGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.AddNamedGroupingPolicyRequest)
	return endpoints.AddNamedGroupingPolicyRequest{Enforcer: req.Enforcer, PType: req.PType, Params: req.Params}, nil
}
func encodeGRPCAddNamedGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.AddNamedGroupingPolicyResponse)
	return endpoints.AddNamedGroupingPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCRemoveGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.RemoveGroupingPolicyRequest)
	return endpoints.RemoveGroupingPolicyRequest{Enforcer: req.Enforcer, Params: req.Params}, nil
}
func encodeGRPCRemoveGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.RemoveGroupingPolicyResponse)
	return endpoints.RemoveGroupingPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCRemoveNamedGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.RemoveNamedGroupingPolicyRequest)
	return endpoints.RemoveNamedGroupingPolicyRequest{Enforcer: req.Enforcer, PType: req.PType, Params: req.Params}, nil
}
func encodeGRPCRemoveNamedGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.RemoveNamedGroupingPolicyResponse)
	return endpoints.RemoveNamedGroupingPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCRemoveFilteredGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.RemoveFilteredGroupingPolicyRequest)
	return endpoints.RemoveFilteredGroupingPolicyRequest{Enforcer: req.Enforcer, FieldIndex: req.FieldIndex, FieldValues: req.FieldValues}, nil
}
func encodeGRPCRemoveFilteredGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.RemoveFilteredGroupingPolicyResponse)
	return endpoints.RemoveFilteredGroupingPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCRemoveFilteredNamedGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.RemoveFilteredNamedGroupingPolicyRequest)
	return endpoints.RemoveFilteredNamedGroupingPolicyRequest{Enforcer: req.Enforcer, PType: req.PType, FieldIndex: req.FieldIndex, FieldValues: req.FieldValues}, nil
}
func encodeGRPCRemoveFilteredNamedGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.RemoveFilteredNamedGroupingPolicyResponse)
	return endpoints.RemoveFilteredNamedGroupingPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetGroupingPolicyRequest)
	return endpoints.GetGroupingPolicyRequest{Enforcer: req.Enforcer}, nil
}
func encodeGRPCGetGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetGroupingPolicyResponse)
	return endpoints.GetGroupingPolicyRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetNamedGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetNamedGroupingPolicyRequest)
	return endpoints.GetNamedGroupingPolicyRequest{Enforcer: req.Enforcer, PType: req.PType}, nil
}
func encodeGRPCGetNamedGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetNamedGroupingPolicyResponse)
	return endpoints.GetNamedGroupingPolicyRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetFilteredGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetFilteredGroupingPolicyRequest)
	return endpoints.GetFilteredGroupingPolicyRequest{Enforcer: req.Enforcer, FieldIndex: req.FieldIndex, FieldValues: req.FieldValues}, nil
}
func encodeGRPCGetFilteredGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetFilteredGroupingPolicyResponse)
	return endpoints.GetFilteredGroupingPolicyRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetFilteredNamedGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetFilteredNamedGroupingPolicyRequest)
	return endpoints.GetFilteredNamedGroupingPolicyRequest{Enforcer: req.Enforcer, PType: req.PType, FieldIndex: req.FieldIndex, FieldValues: req.FieldValues}, nil
}
func encodeGRPCGetFilteredNamedGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetFilteredNamedGroupingPolicyResponse)
	return endpoints.GetFilteredNamedGroupingPolicyRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetAllSubjectsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetAllSubjectsRequest)
	return endpoints.GetAllSubjectsRequest{Enforcer: req.Enforcer}, nil
}
func encodeGRPCGetAllSubjectsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetAllSubjectsResponse)
	return endpoints.GetAllSubjectsRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetAllNamedSubjectsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetAllNamedSubjectsRequest)
	return endpoints.GetAllNamedSubjectsRequest{Enforcer: req.Enforcer, PType: req.PType}, nil
}
func encodeGRPCGetAllNamedSubjectsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetAllNamedSubjectsResponse)
	return endpoints.GetAllNamedSubjectsRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetAllObjectsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetAllObjectsRequest)
	return endpoints.GetAllObjectsRequest{Enforcer: req.Enforcer}, nil
}
func encodeGRPCGetAllObjectsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetAllObjectsResponse)
	return endpoints.GetAllObjectsRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetAllNamedObjectsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetAllNamedObjectsRequest)
	return endpoints.GetAllNamedObjectsRequest{Enforcer: req.Enforcer, PType: req.PType}, nil
}
func encodeGRPCGetAllNamedObjectsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetAllNamedObjectsResponse)
	return endpoints.GetAllNamedObjectsRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetAllActionsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetAllActionsRequest)
	return endpoints.GetAllActionsRequest{Enforcer: req.Enforcer}, nil
}
func encodeGRPCGetAllActionsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetAllActionsResponse)
	return endpoints.GetAllActionsRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetAllNamedActionsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetAllNamedActionsRequest)
	return endpoints.GetAllNamedActionsRequest{Enforcer: req.Enforcer, PType: req.PType}, nil
}
func encodeGRPCGetAllNamedActionsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetAllNamedActionsResponse)
	return endpoints.GetAllNamedActionsRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetAllRolesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetAllRolesRequest)
	return endpoints.GetAllRolesRequest{Enforcer: req.Enforcer}, nil
}
func encodeGRPCGetAllRolesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetAllRolesResponse)
	return endpoints.GetAllRolesRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetAllNamedRolesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetAllNamedRolesRequest)
	return endpoints.GetAllNamedRolesRequest{Enforcer: req.Enforcer, PType: req.PType}, nil
}
func encodeGRPCGetAllNamedRolesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetAllNamedRolesResponse)
	return endpoints.GetAllNamedRolesRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCHasPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.HasPolicyRequest)
	return endpoints.HasPolicyRequest{Enforcer: req.Enforcer, Params: req.Params}, nil
}
func encodeGRPCHasPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.HasPolicyResponse)
	return endpoints.HasPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCHasNamedPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.HasNamedPolicyRequest)
	return endpoints.HasNamedPolicyRequest{Enforcer: req.Enforcer, PType: req.PType, Params: req.Params}, nil
}
func encodeGRPCHasNamedPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.HasNamedPolicyResponse)
	return endpoints.HasNamedPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCHasGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.HasGroupingPolicyRequest)
	return endpoints.HasGroupingPolicyRequest{Enforcer: req.Enforcer, Params: req.Params}, nil
}
func encodeGRPCHasGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.HasGroupingPolicyResponse)
	return endpoints.HasGroupingPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCHasNamedGroupingPolicyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.HasNamedGroupingPolicyRequest)
	return endpoints.HasNamedGroupingPolicyRequest{Enforcer: req.Enforcer, PType: req.PType, Params: req.Params}, nil
}
func encodeGRPCHasNamedGroupingPolicyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.HasNamedGroupingPolicyResponse)
	return endpoints.HasNamedGroupingPolicyRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCHasRoleForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.HasRoleForUserRequest)
	return endpoints.HasRoleForUserRequest{Enforcer: req.Enforcer, User: req.User, Role: req.Role}, nil
}
func encodeGRPCHasRoleForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.HasRoleForUserResponse)
	return endpoints.HasRoleForUserRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCAddRoleForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.AddRoleForUserRequest)
	return endpoints.AddRoleForUserRequest{Enforcer: req.Enforcer, User: req.User, Role: req.Role}, nil
}
func encodeGRPCAddRoleForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.AddRoleForUserResponse)
	return endpoints.AddRoleForUserRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCDeleteRoleForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.DeleteRoleForUserRequest)
	return endpoints.DeleteRoleForUserRequest{Enforcer: req.Enforcer, User: req.User, Role: req.Role}, nil
}
func encodeGRPCDeleteRoleForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.DeleteRoleForUserResponse)
	return endpoints.DeleteRoleForUserRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCDeleteRolesForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.DeleteRolesForUserRequest)
	return endpoints.DeleteRolesForUserRequest{Enforcer: req.Enforcer, User: req.User}, nil
}
func encodeGRPCDeleteRolesForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.DeleteRolesForUserResponse)
	return endpoints.DeleteRolesForUserRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCDeleteUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.DeleteUserRequest)
	return endpoints.DeleteUserRequest{Enforcer: req.Enforcer, User: req.User}, nil
}
func encodeGRPCDeleteUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.DeleteUserResponse)
	return endpoints.DeleteUserRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCDeleteRoleRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.DeleteRoleRequest)
	return endpoints.DeleteRoleRequest{Enforcer: req.Enforcer, Role: req.Role}, nil
}
func encodeGRPCDeleteRoleResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.DeleteRoleResponse)
	return endpoints.DeleteRoleRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCDeletePermissionRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.DeletePermissionRequest)
	return endpoints.DeletePermissionRequest{Enforcer: req.Enforcer, Permission: req.Permission}, nil
}
func encodeGRPCDeletePermissionResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.DeletePermissionResponse)
	return endpoints.DeletePermissionRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetRolesForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetRolesForUserRequest)
	return endpoints.GetRolesForUserRequest{Enforcer: req.Enforcer, User: req.User}, nil
}
func encodeGRPCGetRolesForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetRolesForUserResponse)
	return endpoints.GetRolesForUserRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetImplicitRolesForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetImplicitRolesForUserRequest)
	return endpoints.GetImplicitRolesForUserRequest{Enforcer: req.Enforcer, User: req.User}, nil
}
func encodeGRPCGetImplicitRolesForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetImplicitRolesForUserResponse)
	return endpoints.GetImplicitRolesForUserRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetUsersForRoleRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetUsersForRoleRequest)
	return endpoints.GetUsersForRoleRequest{Enforcer: req.Enforcer, Role: req.Role}, nil
}
func encodeGRPCGetUsersForRoleResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetUsersForRoleResponse)
	return endpoints.GetUsersForRoleRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCAddPermissionForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.AddPermissionForUserRequest)
	return endpoints.AddPermissionForUserRequest{Enforcer: req.Enforcer, User: req.User, Permissions: req.Permissions}, nil
}
func encodeGRPCAddPermissionForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.AddPermissionForUserResponse)
	return endpoints.AddPermissionForUserRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCDeletePermissionForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.DeletePermissionForUserRequest)
	return endpoints.DeletePermissionForUserRequest{Enforcer: req.Enforcer, User: req.User, Permissions: req.Permissions}, nil
}
func encodeGRPCDeletePermissionForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.DeletePermissionForUserResponse)
	return endpoints.DeletePermissionForUserRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetPermissionsForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetPermissionsForUserRequest)
	return endpoints.GetPermissionsForUserRequest{Enforcer: req.Enforcer, User: req.User, Permissions: req.Permissions}, nil
}
func encodeGRPCGetPermissionsForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetPermissionsForUserResponse)
	return endpoints.GetPermissionsForUserRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCGetImplicitPermissionsForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.GetImplicitPermissionsForUserRequest)
	return endpoints.GetImplicitPermissionsForUserRequest{Enforcer: req.Enforcer, User: req.User, Permissions: req.Permissions}, nil
}
func encodeGRPCGetImplicitPermissionsForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.GetImplicitPermissionsForUserResponse)
	return endpoints.GetImplicitPermissionsForUserRequest{S: resp.S, Err: err2str(resp.Err)}, nil
}
func decodeGRPCHasPermissionForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req = grpcReq.(*pb.HasPermissionForUserRequest)
	return endpoints.HasPermissionForUserRequest{Enforcer: req.Enforcer, User: req.User, Permissions: req.Permissions}, nil
}
func encodeGRPCHasPermissionForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp = response.(*pb.HasPermissionForUserResponse)
	return endpoints.HasPermissionForUserRequest{B: resp.B, Err: err2str(resp.Err)}, nil
}
func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
