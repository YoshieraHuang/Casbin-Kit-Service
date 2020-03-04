package endpoints

import (
	"casbinsvc/pkg/services"
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Set collects all of the endpoints that compose a service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	NewEnforcerEndpoint                       endpoint.Endpoint
	EnforceEndpoint                           endpoint.Endpoint
	LoadPolicyEndpoint                        endpoint.Endpoint
	SavePolicyEndpoint                        endpoint.Endpoint
	NewAdapterEndpoint                        endpoint.Endpoint
	AddPolicyEndpoint                         endpoint.Endpoint
	AddNamedPolicyEndpoint                    endpoint.Endpoint
	RemovePolicyEndpoint                      endpoint.Endpoint
	RemoveNamedPolicyEndpoint                 endpoint.Endpoint
	RemoveFilteredPolicyEndpoint              endpoint.Endpoint
	RemoveFilteredNamedPolicyEndpoint         endpoint.Endpoint
	GetPolicyEndpoint                         endpoint.Endpoint
	GetNamedPolicyEndpoint                    endpoint.Endpoint
	GetFilteredPolicyEndpoint                 endpoint.Endpoint
	GetFilteredNamedPolicyEndpoint            endpoint.Endpoint
	AddGroupingPolicyEndpoint                 endpoint.Endpoint
	AddNamedGroupingPolicyEndpoint            endpoint.Endpoint
	RemoveGroupingPolicyEndpoint              endpoint.Endpoint
	RemoveNamedGroupingPolicyEndpoint         endpoint.Endpoint
	RemoveFilteredGroupingPolicyEndpoint      endpoint.Endpoint
	RemoveFilteredNamedGroupingPolicyEndpoint endpoint.Endpoint
	GetGroupingPolicyEndpoint                 endpoint.Endpoint
	GetNamedGroupingPolicyEndpoint            endpoint.Endpoint
	GetFilteredGroupingPolicyEndpoint         endpoint.Endpoint
	GetFilteredNamedGroupingPolicyEndpoint    endpoint.Endpoint
	GetAllSubjectsEndpoint                    endpoint.Endpoint
	GetAllNamedSubjectsEndpoint               endpoint.Endpoint
	GetAllObjectsEndpoint                     endpoint.Endpoint
	GetAllNamedObjectsEndpoint                endpoint.Endpoint
	GetAllActionsEndpoint                     endpoint.Endpoint
	GetAllNamedActionsEndpoint                endpoint.Endpoint
	GetAllRolesEndpoint                       endpoint.Endpoint
	GetAllNamedRolesEndpoint                  endpoint.Endpoint
	HasPolicyEndpoint                         endpoint.Endpoint
	HasNamedPolicyEndpoint                    endpoint.Endpoint
	HasGroupingPolicyEndpoint                 endpoint.Endpoint
	HasNamedGroupingPolicyEndpoint            endpoint.Endpoint
	HasRoleForUserEndpoint                    endpoint.Endpoint
	AddRoleForUserEndpoint                    endpoint.Endpoint
	DeleteRoleForUserEndpoint                 endpoint.Endpoint
	DeleteRolesForUserEndpoint                endpoint.Endpoint
	DeleteUserEndpoint                        endpoint.Endpoint
	DeleteRoleEndpoint                        endpoint.Endpoint
	DeletePermissionEndpoint                  endpoint.Endpoint
	GetRolesForUserEndpoint                   endpoint.Endpoint
	GetImplicitRolesForUserEndpoint           endpoint.Endpoint
	GetUsersForRoleEndpoint                   endpoint.Endpoint
	AddPermissionForUserEndpoint              endpoint.Endpoint
	DeletePermissionForUserEndpoint           endpoint.Endpoint
	GetPermissionsForUserEndpoint             endpoint.Endpoint
	GetImplicitPermissionsForUserEndpoint     endpoint.Endpoint
	HasPermissionForUserEndpoint              endpoint.Endpoint
}

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(s services.Server) Set {
	var newEnforcer endpoint.Endpoint
	{
		newEnforcer = MakeNewEnforcerEndpoint(s)
	}
	var enforce endpoint.Endpoint
	{
		enforce = MakeEnforceEndpoint(s)
	}
	var loadPolicy endpoint.Endpoint
	{
		loadPolicy = MakeLoadPolicyEndpoint(s)
	}
	var savePolicy endpoint.Endpoint
	{
		savePolicy = MakeSavePolicyEndpoint(s)
	}
	var newAdapter endpoint.Endpoint
	{
		newAdapter = MakeNewAdapterEndpoint(s)
	}
	var addPolicy endpoint.Endpoint
	{
		addPolicy = MakeAddPolicyEndpoint(s)
	}
	var addNamedPolicy endpoint.Endpoint
	{
		addNamedPolicy = MakeAddNamedPolicyEndpoint(s)
	}
	var removePolicy endpoint.Endpoint
	{
		removePolicy = MakeRemovePolicyEndpoint(s)
	}
	var removeNamedPolicy endpoint.Endpoint
	{
		removeNamedPolicy = MakeRemoveNamedPolicyEndpoint(s)
	}
	var removeFilteredPolicy endpoint.Endpoint
	{
		removeFilteredPolicy = MakeRemoveFilteredPolicyEndpoint(s)
	}
	var removeFilteredNamedPolicy endpoint.Endpoint
	{
		removeFilteredNamedPolicy = MakeRemoveFilteredNamedPolicyEndpoint(s)
	}
	var getPolicy endpoint.Endpoint
	{
		getPolicy = MakeGetPolicyEndpoint(s)
	}
	var getNamedPolicy endpoint.Endpoint
	{
		getNamedPolicy = MakeGetNamedPolicyEndpoint(s)
	}
	var getFilteredPolicy endpoint.Endpoint
	{
		getFilteredPolicy = MakeGetFilteredPolicyEndpoint(s)
	}
	var getFilteredNamedPolicy endpoint.Endpoint
	{
		getFilteredNamedPolicy = MakeGetFilteredNamedPolicyEndpoint(s)
	}
	var addGroupingPolicy endpoint.Endpoint
	{
		addGroupingPolicy = MakeAddGroupingPolicyEndpoint(s)
	}
	var addNamedGroupingPolicy endpoint.Endpoint
	{
		addNamedGroupingPolicy = MakeAddNamedGroupingPolicyEndpoint(s)
	}
	var removeGroupingPolicy endpoint.Endpoint
	{
		removeGroupingPolicy = MakeRemoveGroupingPolicyEndpoint(s)
	}
	var removeNamedGroupingPolicy endpoint.Endpoint
	{
		removeNamedGroupingPolicy = MakeRemoveNamedGroupingPolicyEndpoint(s)
	}
	var removeFilteredGroupingPolicy endpoint.Endpoint
	{
		removeFilteredGroupingPolicy = MakeRemoveFilteredGroupingPolicyEndpoint(s)
	}
	var removeFilteredNamedGroupingPolicy endpoint.Endpoint
	{
		removeFilteredNamedGroupingPolicy = MakeRemoveFilteredNamedGroupingPolicyEndpoint(s)
	}
	var getGroupingPolicy endpoint.Endpoint
	{
		getGroupingPolicy = MakeGetGroupingPolicyEndpoint(s)
	}
	var getNamedGroupingPolicy endpoint.Endpoint
	{
		getNamedGroupingPolicy = MakeGetNamedGroupingPolicyEndpoint(s)
	}
	var getFilteredGroupingPolicy endpoint.Endpoint
	{
		getFilteredGroupingPolicy = MakeGetFilteredGroupingPolicyEndpoint(s)
	}
	var getFilteredNamedGroupingPolicy endpoint.Endpoint
	{
		getFilteredNamedGroupingPolicy = MakeGetFilteredNamedGroupingPolicyEndpoint(s)
	}
	var getAllSubjects endpoint.Endpoint
	{
		getAllSubjects = MakeGetAllSubjectsEndpoint(s)
	}
	var getAllNamedSubjects endpoint.Endpoint
	{
		getAllNamedSubjects = MakeGetAllNamedSubjectsEndpoint(s)
	}
	var getAllObjects endpoint.Endpoint
	{
		getAllObjects = MakeGetAllObjectsEndpoint(s)
	}
	var getAllNamedObjects endpoint.Endpoint
	{
		getAllNamedObjects = MakeGetAllNamedObjectsEndpoint(s)
	}
	var getAllActions endpoint.Endpoint
	{
		getAllActions = MakeGetAllActionsEndpoint(s)
	}
	var getAllNamedActions endpoint.Endpoint
	{
		getAllNamedActions = MakeGetAllNamedActionsEndpoint(s)
	}
	var getAllRoles endpoint.Endpoint
	{
		getAllRoles = MakeGetAllRolesEndpoint(s)
	}
	var getAllNamedRoles endpoint.Endpoint
	{
		getAllNamedRoles = MakeGetAllNamedRolesEndpoint(s)
	}
	var hasPolicy endpoint.Endpoint
	{
		hasPolicy = MakeHasPolicyEndpoint(s)
	}
	var hasNamedPolicy endpoint.Endpoint
	{
		hasNamedPolicy = MakeHasNamedPolicyEndpoint(s)
	}
	var hasGroupingPolicy endpoint.Endpoint
	{
		hasGroupingPolicy = MakeHasGroupingPolicyEndpoint(s)
	}
	var hasNamedGroupingPolicy endpoint.Endpoint
	{
		hasNamedGroupingPolicy = MakeHasNamedGroupingPolicyEndpoint(s)
	}
	var hasRoleForUser endpoint.Endpoint
	{
		hasRoleForUser = MakeHasRoleForUserEndpoint(s)
	}
	var addRoleForUser endpoint.Endpoint
	{
		addRoleForUser = MakeAddRoleForUserEndpoint(s)
	}
	var deleteRoleForUser endpoint.Endpoint
	{
		deleteRoleForUser = MakeDeleteRoleForUserEndpoint(s)
	}
	var deleteRolesForUser endpoint.Endpoint
	{
		deleteRolesForUser = MakeDeleteRolesForUserEndpoint(s)
	}
	var deleteUser endpoint.Endpoint
	{
		deleteUser = MakeDeleteUserEndpoint(s)
	}
	var deleteRole endpoint.Endpoint
	{
		deleteRole = MakeDeleteRoleEndpoint(s)
	}
	var deletePermission endpoint.Endpoint
	{
		deletePermission = MakeDeletePermissionEndpoint(s)
	}
	var getRolesForUser endpoint.Endpoint
	{
		getRolesForUser = MakeGetRolesForUserEndpoint(s)
	}
	var getImplicitRolesForUser endpoint.Endpoint
	{
		getImplicitRolesForUser = MakeGetImplicitRolesForUserEndpoint(s)
	}
	var getUsersForRole endpoint.Endpoint
	{
		getUsersForRole = MakeGetUsersForRoleEndpoint(s)
	}
	var addPermissionForUser endpoint.Endpoint
	{
		addPermissionForUser = MakeAddPermissionForUserEndpoint(s)
	}
	var deletePermissionForUser endpoint.Endpoint
	{
		deletePermissionForUser = MakeDeletePermissionForUserEndpoint(s)
	}
	var getPermissionsForUser endpoint.Endpoint
	{
		getPermissionsForUser = MakeGetPermissionsForUserEndpoint(s)
	}
	var getImplicitPermissionsForUser endpoint.Endpoint
	{
		getImplicitPermissionsForUser = MakeGetImplicitPermissionsForUserEndpoint(s)
	}
	var hasPermissionForUser endpoint.Endpoint
	{
		hasPermissionForUser = MakeHasPermissionForUserEndpoint(s)
	}
	return Set{NewEnforcerEndpoint: newEnforcer, EnforceEndpoint: enforce, LoadPolicyEndpoint: loadPolicy, SavePolicyEndpoint: savePolicy, NewAdapterEndpoint: newAdapter, AddPolicyEndpoint: addPolicy, AddNamedPolicyEndpoint: addNamedPolicy, RemovePolicyEndpoint: removePolicy, RemoveNamedPolicyEndpoint: removeNamedPolicy, RemoveFilteredPolicyEndpoint: removeFilteredPolicy, RemoveFilteredNamedPolicyEndpoint: removeFilteredNamedPolicy, GetPolicyEndpoint: getPolicy, GetNamedPolicyEndpoint: getNamedPolicy, GetFilteredPolicyEndpoint: getFilteredPolicy, GetFilteredNamedPolicyEndpoint: getFilteredNamedPolicy, AddGroupingPolicyEndpoint: addGroupingPolicy, AddNamedGroupingPolicyEndpoint: addNamedGroupingPolicy, RemoveGroupingPolicyEndpoint: removeGroupingPolicy, RemoveNamedGroupingPolicyEndpoint: removeNamedGroupingPolicy, RemoveFilteredGroupingPolicyEndpoint: removeFilteredGroupingPolicy, RemoveFilteredNamedGroupingPolicyEndpoint: removeFilteredNamedGroupingPolicy, GetGroupingPolicyEndpoint: getGroupingPolicy, GetNamedGroupingPolicyEndpoint: getNamedGroupingPolicy, GetFilteredGroupingPolicyEndpoint: getFilteredGroupingPolicy, GetFilteredNamedGroupingPolicyEndpoint: getFilteredNamedGroupingPolicy, GetAllSubjectsEndpoint: getAllSubjects, GetAllNamedSubjectsEndpoint: getAllNamedSubjects, GetAllObjectsEndpoint: getAllObjects, GetAllNamedObjectsEndpoint: getAllNamedObjects, GetAllActionsEndpoint: getAllActions, GetAllNamedActionsEndpoint: getAllNamedActions, GetAllRolesEndpoint: getAllRoles, GetAllNamedRolesEndpoint: getAllNamedRoles, HasPolicyEndpoint: hasPolicy, HasNamedPolicyEndpoint: hasNamedPolicy, HasGroupingPolicyEndpoint: hasGroupingPolicy, HasNamedGroupingPolicyEndpoint: hasNamedGroupingPolicy, HasRoleForUserEndpoint: hasRoleForUser, AddRoleForUserEndpoint: addRoleForUser, DeleteRoleForUserEndpoint: deleteRoleForUser, DeleteRolesForUserEndpoint: deleteRolesForUser, DeleteUserEndpoint: deleteUser, DeleteRoleEndpoint: deleteRole, DeletePermissionEndpoint: deletePermission, GetRolesForUserEndpoint: getRolesForUser, GetImplicitRolesForUserEndpoint: getImplicitRolesForUser, GetUsersForRoleEndpoint: getUsersForRole, AddPermissionForUserEndpoint: addPermissionForUser, DeletePermissionForUserEndpoint: deletePermissionForUser, GetPermissionsForUserEndpoint: getPermissionsForUser, GetImplicitPermissionsForUserEndpoint: getImplicitPermissionsForUser, HasPermissionForUserEndpoint: hasPermissionForUser}
}

// MakeNewEnforcerEndpoint constructs a NewEnforcer endpoint wrapping the service.
func MakeNewEnforcerEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(NewEnforcerRequest)
		enforcer, err := s.NewEnforcer(ctx, req.ModelText, req.Adapter)
		return NewEnforcerResponse{Enforcer: enforcer, Err: err}, nil
	}
}

// MakeEnforceEndpoint constructs a Enforce endpoint wrapping the service.
func MakeEnforceEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EnforceRequest)
		result, err := s.Enforce(ctx, req.Enforcer, req.Params)
		return EnforceResponse{Result: result, Err: err}, nil
	}
}

// MakeLoadPolicyEndpoint constructs a LoadPolicy endpoint wrapping the service.
func MakeLoadPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoadPolicyRequest)
		err := s.LoadPolicy(ctx, req.
			Enforcer)
		return LoadPolicyResponse{Err: err}, nil
	}
}

// MakeSavePolicyEndpoint constructs a SavePolicy endpoint wrapping the service.
func MakeSavePolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SavePolicyRequest)
		err := s.SavePolicy(ctx, req.
			Enforcer)
		return SavePolicyResponse{Err: err}, nil
	}
}

// MakeNewAdapterEndpoint constructs a NewAdapter endpoint wrapping the service.
func MakeNewAdapterEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(NewAdapterRequest)
		adapter, err := s.NewAdapter(ctx, req.DriverName, req.ConnectString, req.DbSpecified)
		return NewAdapterResponse{Adapter: adapter, Err: err}, nil
	}
}

// MakeAddPolicyEndpoint constructs a AddPolicy endpoint wrapping the service.
func MakeAddPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddPolicyRequest)
		result, err := s.AddPolicy(ctx, req.Enforcer, req.Params)
		return AddPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeAddNamedPolicyEndpoint constructs a AddNamedPolicy endpoint wrapping the service.
func MakeAddNamedPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddNamedPolicyRequest)
		result, err := s.AddNamedPolicy(ctx, req.Enforcer, req.PType, req.Params)
		return AddNamedPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeRemovePolicyEndpoint constructs a RemovePolicy endpoint wrapping the service.
func MakeRemovePolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemovePolicyRequest)
		result, err := s.RemovePolicy(ctx, req.Enforcer, req.Params)
		return RemovePolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeRemoveNamedPolicyEndpoint constructs a RemoveNamedPolicy endpoint wrapping the service.
func MakeRemoveNamedPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveNamedPolicyRequest)
		result, err := s.RemoveNamedPolicy(ctx, req.
			Enforcer, req.PType, req.Params)
		return RemoveNamedPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeRemoveFilteredPolicyEndpoint constructs a RemoveFilteredPolicy endpoint wrapping the service.
func MakeRemoveFilteredPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveFilteredPolicyRequest)
		result, err := s.RemoveFilteredPolicy(ctx, req.
			Enforcer, req.FieldIndex, req.FieldValues)
		return RemoveFilteredPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeRemoveFilteredNamedPolicyEndpoint constructs a RemoveFilteredNamedPolicy endpoint wrapping the service.
func MakeRemoveFilteredNamedPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveFilteredNamedPolicyRequest)
		result, err := s.RemoveFilteredNamedPolicy(ctx, req.
			Enforcer, req.PType, req.FieldIndex, req.FieldValues)
		return RemoveFilteredNamedPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeGetPolicyEndpoint constructs a GetPolicy endpoint wrapping the service.
func MakeGetPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPolicyRequest)
		results, err := s.GetPolicy(ctx, req.Enforcer)
		return GetPolicyResponse{Results: results, Err: err}, nil
	}
}

// MakeGetNamedPolicyEndpoint constructs a GetNamedPolicy endpoint wrapping the service.
func MakeGetNamedPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNamedPolicyRequest)
		results, err := s.GetNamedPolicy(ctx, req.Enforcer, req.PType)
		return GetNamedPolicyResponse{Results: results, Err: err}, nil
	}
}

// MakeGetFilteredPolicyEndpoint constructs a GetFilteredPolicy endpoint wrapping the service.
func MakeGetFilteredPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFilteredPolicyRequest)
		results, err := s.GetFilteredPolicy(ctx, req.Enforcer, req.FieldIndex, req.FieldValues)
		return GetFilteredPolicyResponse{Results: results, Err: err}, nil
	}
}

// MakeGetFilteredNamedPolicyEndpoint constructs a GetFilteredNamedPolicy endpoint wrapping the service.
func MakeGetFilteredNamedPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFilteredNamedPolicyRequest)
		results, err := s.GetFilteredNamedPolicy(ctx, req.Enforcer, req.PType, req.FieldIndex, req.FieldValues)
		return GetFilteredNamedPolicyResponse{Results: results, Err: err}, nil
	}
}

// MakeAddGroupingPolicyEndpoint constructs a AddGroupingPolicy endpoint wrapping the service.
func MakeAddGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddGroupingPolicyRequest)
		result, err := s.AddGroupingPolicy(ctx, req.Enforcer, req.Params)
		return AddGroupingPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeAddNamedGroupingPolicyEndpoint constructs a AddNamedGroupingPolicy endpoint wrapping the service.
func MakeAddNamedGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddNamedGroupingPolicyRequest)
		result, err := s.AddNamedGroupingPolicy(ctx, req.Enforcer, req.PType, req.Params)
		return AddNamedGroupingPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeRemoveGroupingPolicyEndpoint constructs a RemoveGroupingPolicy endpoint wrapping the service.
func MakeRemoveGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveGroupingPolicyRequest)
		result, err := s.RemoveGroupingPolicy(ctx, req.Enforcer, req.Params)
		return RemoveGroupingPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeRemoveNamedGroupingPolicyEndpoint constructs a RemoveNamedGroupingPolicy endpoint wrapping the service.
func MakeRemoveNamedGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveNamedGroupingPolicyRequest)
		result, err := s.RemoveNamedGroupingPolicy(ctx, req.Enforcer, req.PType, req.Params)
		return RemoveNamedGroupingPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeRemoveFilteredGroupingPolicyEndpoint constructs a RemoveFilteredGroupingPolicy endpoint wrapping the service.
func MakeRemoveFilteredGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveFilteredGroupingPolicyRequest)
		result, err := s.RemoveFilteredGroupingPolicy(ctx, req.Enforcer, req.FieldIndex, req.FieldValues)
		return RemoveFilteredGroupingPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeRemoveFilteredNamedGroupingPolicyEndpoint constructs a RemoveFilteredNamedGroupingPolicy endpoint wrapping the service.
func MakeRemoveFilteredNamedGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveFilteredNamedGroupingPolicyRequest)
		result, err := s.RemoveFilteredNamedGroupingPolicy(ctx, req.Enforcer, req.PType, req.FieldIndex, req.FieldValues)
		return RemoveFilteredNamedGroupingPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeGetGroupingPolicyEndpoint constructs a GetGroupingPolicy endpoint wrapping the service.
func MakeGetGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupingPolicyRequest)
		results, err := s.GetGroupingPolicy(ctx, req.Enforcer)
		return GetGroupingPolicyResponse{Results: results, Err: err}, nil
	}
}

// MakeGetNamedGroupingPolicyEndpoint constructs a GetNamedGroupingPolicy endpoint wrapping the service.
func MakeGetNamedGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNamedGroupingPolicyRequest)
		results, err := s.GetNamedGroupingPolicy(ctx, req.Enforcer, req.PType)
		return GetNamedGroupingPolicyResponse{Results: results, Err: err}, nil
	}
}

// MakeGetFilteredGroupingPolicyEndpoint constructs a GetFilteredGroupingPolicy endpoint wrapping the service.
func MakeGetFilteredGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFilteredGroupingPolicyRequest)
		results, err := s.GetFilteredGroupingPolicy(ctx, req.Enforcer, req.FieldIndex, req.FieldValues)
		return GetFilteredGroupingPolicyResponse{Results: results, Err: err}, nil
	}
}

// MakeGetFilteredNamedGroupingPolicyEndpoint constructs a GetFilteredNamedGroupingPolicy endpoint wrapping the service.
func MakeGetFilteredNamedGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFilteredNamedGroupingPolicyRequest)
		results, err := s.GetFilteredNamedGroupingPolicy(ctx, req.Enforcer, req.PType, req.FieldIndex, req.FieldValues)
		return GetFilteredNamedGroupingPolicyResponse{Results: results, Err: err}, nil
	}
}

// MakeGetAllSubjectsEndpoint constructs a GetAllSubjects endpoint wrapping the service.
func MakeGetAllSubjectsEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllSubjectsRequest)
		results, err := s.GetAllSubjects(ctx, req.Enforcer)
		return GetAllSubjectsResponse{Results: results, Err: err}, nil
	}
}

// MakeGetAllNamedSubjectsEndpoint constructs a GetAllNamedSubjects endpoint wrapping the service.
func MakeGetAllNamedSubjectsEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllNamedSubjectsRequest)
		results, err := s.GetAllNamedSubjects(ctx, req.Enforcer, req.PType)
		return GetAllNamedSubjectsResponse{Results: results, Err: err}, nil
	}
}

// MakeGetAllObjectsEndpoint constructs a GetAllObjects endpoint wrapping the service.
func MakeGetAllObjectsEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllObjectsRequest)
		results, err := s.GetAllObjects(ctx, req.Enforcer)
		return GetAllObjectsResponse{Results: results, Err: err}, nil
	}
}

// MakeGetAllNamedObjectsEndpoint constructs a GetAllNamedObjects endpoint wrapping the service.
func MakeGetAllNamedObjectsEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllNamedObjectsRequest)
		results, err := s.GetAllNamedObjects(ctx, req.Enforcer, req.PType)
		return GetAllNamedObjectsResponse{Results: results, Err: err}, nil
	}
}

// MakeGetAllActionsEndpoint constructs a GetAllActions endpoint wrapping the service.
func MakeGetAllActionsEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllActionsRequest)
		results, err := s.GetAllActions(ctx, req.Enforcer)
		return GetAllActionsResponse{Results: results, Err: err}, nil
	}
}

// MakeGetAllNamedActionsEndpoint constructs a GetAllNamedActions endpoint wrapping the service.
func MakeGetAllNamedActionsEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllNamedActionsRequest)
		results, err := s.GetAllNamedActions(ctx, req.Enforcer, req.PType)
		return GetAllNamedActionsResponse{Results: results, Err: err}, nil
	}
}

// MakeGetAllRolesEndpoint constructs a GetAllRoles endpoint wrapping the service.
func MakeGetAllRolesEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllRolesRequest)
		results, err := s.GetAllRoles(ctx, req.Enforcer)
		return GetAllRolesResponse{Results: results, Err: err}, nil
	}
}

// MakeGetAllNamedRolesEndpoint constructs a GetAllNamedRoles endpoint wrapping the service.
func MakeGetAllNamedRolesEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllNamedRolesRequest)
		results, err := s.GetAllNamedRoles(ctx, req.Enforcer, req.PType)
		return GetAllNamedRolesResponse{Results: results, Err: err}, nil
	}
}

// MakeHasPolicyEndpoint constructs a HasPolicy endpoint wrapping the service.
func MakeHasPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HasPolicyRequest)
		result, err := s.HasPolicy(ctx, req.Enforcer, req.Params)
		return HasPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeHasNamedPolicyEndpoint constructs a HasNamedPolicy endpoint wrapping the service.
func MakeHasNamedPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HasNamedPolicyRequest)
		result, err := s.HasNamedPolicy(ctx, req.Enforcer, req.PType, req.Params)
		return HasNamedPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeHasGroupingPolicyEndpoint constructs a HasGroupingPolicy endpoint wrapping the service.
func MakeHasGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HasGroupingPolicyRequest)
		result, err := s.HasGroupingPolicy(ctx, req.Enforcer, req.Params)
		return HasGroupingPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeHasNamedGroupingPolicyEndpoint constructs a HasNamedGroupingPolicy endpoint wrapping the service.
func MakeHasNamedGroupingPolicyEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HasNamedGroupingPolicyRequest)
		result, err := s.HasNamedGroupingPolicy(ctx, req.Enforcer, req.PType, req.Params)
		return HasNamedGroupingPolicyResponse{Result: result, Err: err}, nil
	}
}

// MakeHasRoleForUserEndpoint constructs a HasRoleForUser endpoint wrapping the service.
func MakeHasRoleForUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HasRoleForUserRequest)
		result, err := s.HasRoleForUser(ctx, req.Enforcer, req.User, req.Role)
		return HasRoleForUserResponse{Result: result, Err: err}, nil
	}
}

// MakeAddRoleForUserEndpoint constructs a AddRoleForUser endpoint wrapping the service.
func MakeAddRoleForUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRoleForUserRequest)
		result, err := s.AddRoleForUser(ctx, req.Enforcer, req.User, req.Role)
		return AddRoleForUserResponse{Result: result, Err: err}, nil
	}
}

// MakeDeleteRoleForUserEndpoint constructs a DeleteRoleForUser endpoint wrapping the service.
func MakeDeleteRoleForUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRoleForUserRequest)
		result, err := s.DeleteRoleForUser(ctx, req.Enforcer, req.User, req.Role)
		return DeleteRoleForUserResponse{Result: result, Err: err}, nil
	}
}

// MakeDeleteRolesForUserEndpoint constructs a DeleteRolesForUser endpoint wrapping the service.
func MakeDeleteRolesForUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRolesForUserRequest)
		result, err := s.DeleteRolesForUser(ctx, req.Enforcer, req.User)
		return DeleteRolesForUserResponse{Result: result, Err: err}, nil
	}
}

// MakeDeleteUserEndpoint constructs a DeleteUser endpoint wrapping the service.
func MakeDeleteUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		result, err := s.DeleteUser(ctx, req.Enforcer, req.User)
		return DeleteUserResponse{Result: result, Err: err}, nil
	}
}

// MakeDeleteRoleEndpoint constructs a DeleteRole endpoint wrapping the service.
func MakeDeleteRoleEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRoleRequest)
		result, err := s.DeleteRole(ctx, req.Enforcer, req.Role)
		return DeleteRoleResponse{Result: result, Err: err}, nil
	}
}

// MakeDeletePermissionEndpoint constructs a DeletePermission endpoint wrapping the service.
func MakeDeletePermissionEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePermissionRequest)
		result, err := s.DeletePermission(ctx, req.Enforcer, req.Permission)
		return DeletePermissionResponse{Result: result, Err: err}, nil
	}
}

// MakeGetRolesForUserEndpoint constructs a GetRolesForUser endpoint wrapping the service.
func MakeGetRolesForUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRolesForUserRequest)
		results, err := s.GetRolesForUser(ctx, req.Enforcer, req.User)
		return GetRolesForUserResponse{Results: results, Err: err}, nil
	}
}

// MakeGetImplicitRolesForUserEndpoint constructs a GetImplicitRolesForUser endpoint wrapping the service.
func MakeGetImplicitRolesForUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetImplicitRolesForUserRequest)
		results, err := s.GetImplicitRolesForUser(ctx, req.Enforcer, req.User)
		return GetImplicitRolesForUserResponse{Results: results, Err: err}, nil
	}
}

// MakeGetUsersForRoleEndpoint constructs a GetUsersForRole endpoint wrapping the service.
func MakeGetUsersForRoleEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUsersForRoleRequest)
		results, err := s.GetUsersForRole(ctx, req.Enforcer, req.Role)
		return GetUsersForRoleResponse{Results: results, Err: err}, nil
	}
}

// MakeAddPermissionForUserEndpoint constructs a AddPermissionForUser endpoint wrapping the service.
func MakeAddPermissionForUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddPermissionForUserRequest)
		result, err := s.AddPermissionForUser(ctx, req.Enforcer, req.User, req.Permissions)
		return AddPermissionForUserResponse{Result: result, Err: err}, nil
	}
}

// MakeDeletePermissionForUserEndpoint constructs a DeletePermissionForUser endpoint wrapping the service.
func MakeDeletePermissionForUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePermissionForUserRequest)
		result, err := s.DeletePermissionForUser(ctx, req.Enforcer, req.User, req.Permissions)
		return DeletePermissionForUserResponse{Result: result, Err: err}, nil
	}
}

// MakeGetPermissionsForUserEndpoint constructs a GetPermissionsForUser endpoint wrapping the service.
func MakeGetPermissionsForUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPermissionsForUserRequest)
		results, err := s.GetPermissionsForUser(ctx, req.Enforcer, req.User, req.Permissions)
		return GetPermissionsForUserResponse{Results: results, Err: err}, nil
	}
}

// MakeGetImplicitPermissionsForUserEndpoint constructs a GetImplicitPermissionsForUser endpoint wrapping the service.
func MakeGetImplicitPermissionsForUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetImplicitPermissionsForUserRequest)
		results, err := s.GetImplicitPermissionsForUser(ctx, req.Enforcer, req.User, req.Permissions)
		return GetImplicitPermissionsForUserResponse{Results: results, Err: err}, nil
	}
}

// MakeHasPermissionForUserEndpoint constructs a HasPermissionForUser endpoint wrapping the service.
func MakeHasPermissionForUserEndpoint(s services.Server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HasPermissionForUserRequest)
		result, err := s.HasPermissionForUser(ctx, req.Enforcer, req.User, req.Permissions)
		return HasPermissionForUserResponse{Result: result, Err: err}, nil
	}
}

// NewEnforcerRequest collects the request parameters for the NewEnforcer method.
type NewEnforcerRequest struct {
	ModelText string
	Adapter   services.AdapterHandler
}

// EnforceRequest collects the request parameters for the Enforce method.
type EnforceRequest struct {
	Enforcer services.EnforcerHandler
	Params   []string
}

// LoadPolicyRequest collects the request parameters for the LoadPolicy method.
type LoadPolicyRequest struct {
	Enforcer services.EnforcerHandler
}

// SavePolicyRequest collects the request parameters for the SavePolicy method.
type SavePolicyRequest struct {
	Enforcer services.EnforcerHandler
}

// NewAdapterRequest collects the request parameters for the NewAdapter method.
type NewAdapterRequest struct {
	DriverName, ConnectString string
	DbSpecified               bool
}

// AddPolicyRequest collects the request parameters for the AddPolicy method.
type AddPolicyRequest struct {
	Enforcer services.EnforcerHandler
	Params   []string
}

// AddNamedPolicyRequest collects the request parameters for the AddNamedPolicy method.
type AddNamedPolicyRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
	Params   []string
}

// RemovePolicyRequest collects the request parameters for the RemovePolicy method.
type RemovePolicyRequest struct {
	Enforcer services.EnforcerHandler
	Params   []string
}

// RemoveNamedPolicyRequest collects the request parameters for the RemoveNamedPolicy method.
type RemoveNamedPolicyRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
	Params   []string
}

// RemoveFilteredPolicyRequest collects the request parameters for the RemoveFilteredPolicy method.
type RemoveFilteredPolicyRequest struct {
	Enforcer    services.EnforcerHandler
	FieldIndex  int
	FieldValues []string
}

// RemoveFilteredNamedPolicyRequest collects the request parameters for the RemoveFilteredNamedPolicy method.
type RemoveFilteredNamedPolicyRequest struct {
	Enforcer    services.EnforcerHandler
	PType       string
	FieldIndex  int
	FieldValues []string
}

// GetPolicyRequest collects the request parameters for the GetPolicy method.
type GetPolicyRequest struct {
	Enforcer services.EnforcerHandler
}

// GetNamedPolicyRequest collects the request parameters for the GetNamedPolicy method.
type GetNamedPolicyRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
}

// GetFilteredPolicyRequest collects the request parameters for the GetFilteredPolicy method.
type GetFilteredPolicyRequest struct {
	Enforcer    services.EnforcerHandler
	FieldIndex  int
	FieldValues []string
}

// GetFilteredNamedPolicyRequest collects the request parameters for the GetFilteredNamedPolicy method.
type GetFilteredNamedPolicyRequest struct {
	Enforcer    services.EnforcerHandler
	PType       string
	FieldIndex  int
	FieldValues []string
}

// AddGroupingPolicyRequest collects the request parameters for the AddGroupingPolicy method.
type AddGroupingPolicyRequest struct {
	Enforcer services.EnforcerHandler
	Params   []string
}

// AddNamedGroupingPolicyRequest collects the request parameters for the AddNamedGroupingPolicy method.
type AddNamedGroupingPolicyRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
	Params   []string
}

// RemoveGroupingPolicyRequest collects the request parameters for the RemoveGroupingPolicy method.
type RemoveGroupingPolicyRequest struct {
	Enforcer services.EnforcerHandler
	Params   []string
}

// RemoveNamedGroupingPolicyRequest collects the request parameters for the RemoveNamedGroupingPolicy method.
type RemoveNamedGroupingPolicyRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
	Params   []string
}

// RemoveFilteredGroupingPolicyRequest collects the request parameters for the RemoveFilteredGroupingPolicy method.
type RemoveFilteredGroupingPolicyRequest struct {
	Enforcer    services.EnforcerHandler
	FieldIndex  int
	FieldValues []string
}

// RemoveFilteredNamedGroupingPolicyRequest collects the request parameters for the RemoveFilteredNamedGroupingPolicy method.
type RemoveFilteredNamedGroupingPolicyRequest struct {
	Enforcer    services.EnforcerHandler
	PType       string
	FieldIndex  int
	FieldValues []string
}

// GetGroupingPolicyRequest collects the request parameters for the GetGroupingPolicy method.
type GetGroupingPolicyRequest struct {
	Enforcer services.EnforcerHandler
}

// GetNamedGroupingPolicyRequest collects the request parameters for the GetNamedGroupingPolicy method.
type GetNamedGroupingPolicyRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
}

// GetFilteredGroupingPolicyRequest collects the request parameters for the GetFilteredGroupingPolicy method.
type GetFilteredGroupingPolicyRequest struct {
	Enforcer    services.EnforcerHandler
	FieldIndex  int
	FieldValues []string
}

// GetFilteredNamedGroupingPolicyRequest collects the request parameters for the GetFilteredNamedGroupingPolicy method.
type GetFilteredNamedGroupingPolicyRequest struct {
	Enforcer    services.EnforcerHandler
	PType       string
	FieldIndex  int
	FieldValues []string
}

// GetAllSubjectsRequest collects the request parameters for the GetAllSubjects method.
type GetAllSubjectsRequest struct {
	Enforcer services.EnforcerHandler
}

// GetAllNamedSubjectsRequest collects the request parameters for the GetAllNamedSubjects method.
type GetAllNamedSubjectsRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
}

// GetAllObjectsRequest collects the request parameters for the GetAllObjects method.
type GetAllObjectsRequest struct {
	Enforcer services.EnforcerHandler
}

// GetAllNamedObjectsRequest collects the request parameters for the GetAllNamedObjects method.
type GetAllNamedObjectsRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
}

// GetAllActionsRequest collects the request parameters for the GetAllActions method.
type GetAllActionsRequest struct {
	Enforcer services.EnforcerHandler
}

// GetAllNamedActionsRequest collects the request parameters for the GetAllNamedActions method.
type GetAllNamedActionsRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
}

// GetAllRolesRequest collects the request parameters for the GetAllRoles method.
type GetAllRolesRequest struct {
	Enforcer services.EnforcerHandler
}

// GetAllNamedRolesRequest collects the request parameters for the GetAllNamedRoles method.
type GetAllNamedRolesRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
}

// HasPolicyRequest collects the request parameters for the HasPolicy method.
type HasPolicyRequest struct {
	Enforcer services.EnforcerHandler
	Params   []string
}

// HasNamedPolicyRequest collects the request parameters for the HasNamedPolicy method.
type HasNamedPolicyRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
	Params   []string
}

// HasGroupingPolicyRequest collects the request parameters for the HasGroupingPolicy method.
type HasGroupingPolicyRequest struct {
	Enforcer services.EnforcerHandler
	Params   []string
}

// HasNamedGroupingPolicyRequest collects the request parameters for the HasNamedGroupingPolicy method.
type HasNamedGroupingPolicyRequest struct {
	Enforcer services.EnforcerHandler
	PType    string
	Params   []string
}

// HasRoleForUserRequest collects the request parameters for the HasRoleForUser method.
type HasRoleForUserRequest struct {
	Enforcer   services.EnforcerHandler
	User, Role string
}

// AddRoleForUserRequest collects the request parameters for the AddRoleForUser method.
type AddRoleForUserRequest struct {
	Enforcer   services.EnforcerHandler
	User, Role string
}

// DeleteRoleForUserRequest collects the request parameters for the DeleteRoleForUser method.
type DeleteRoleForUserRequest struct {
	Enforcer   services.EnforcerHandler
	User, Role string
}

// DeleteRolesForUserRequest collects the request parameters for the DeleteRolesForUser method.
type DeleteRolesForUserRequest struct {
	Enforcer services.EnforcerHandler
	User     string
}

// DeleteUserRequest collects the request parameters for the DeleteUser method.
type DeleteUserRequest struct {
	Enforcer services.EnforcerHandler
	User     string
}

// DeleteRoleRequest collects the request parameters for the DeleteRole method.
type DeleteRoleRequest struct {
	Enforcer services.EnforcerHandler
	Role     string
}

// DeletePermissionRequest collects the request parameters for the DeletePermission method.
type DeletePermissionRequest struct {
	Enforcer   services.EnforcerHandler
	Permission []string
}

// GetRolesForUserRequest collects the request parameters for the GetRolesForUser method.
type GetRolesForUserRequest struct {
	Enforcer services.EnforcerHandler
	User     string
}

// GetImplicitRolesForUserRequest collects the request parameters for the GetImplicitRolesForUser method.
type GetImplicitRolesForUserRequest struct {
	Enforcer services.EnforcerHandler
	User     string
}

// GetUsersForRoleRequest collects the request parameters for the GetUsersForRole method.
type GetUsersForRoleRequest struct {
	Enforcer services.EnforcerHandler
	Role     string
}

// AddPermissionForUserRequest collects the request parameters for the AddPermissionForUser method.
type AddPermissionForUserRequest struct {
	Enforcer    services.EnforcerHandler
	User        string
	Permissions []string
}

// DeletePermissionForUserRequest collects the request parameters for the DeletePermissionForUser method.
type DeletePermissionForUserRequest struct {
	Enforcer    services.EnforcerHandler
	User        string
	Permissions []string
}

// GetPermissionsForUserRequest collects the request parameters for the GetPermissionsForUser method.
type GetPermissionsForUserRequest struct {
	Enforcer    services.EnforcerHandler
	User        string
	Permissions []string
}

// GetImplicitPermissionsForUserRequest collects the request parameters for the GetImplicitPermissionsForUser method.
type GetImplicitPermissionsForUserRequest struct {
	Enforcer    services.EnforcerHandler
	User        string
	Permissions []string
}

// HasPermissionForUserRequest collects the request parameters for the HasPermissionForUser method.
type HasPermissionForUserRequest struct {
	Enforcer    services.EnforcerHandler
	User        string
	Permissions []string
}

// NewEnforcerResponse collects the response parameters for the NewEnforcer method.
type NewEnforcerResponse struct {
	Enforcer services.EnforcerHandler `json:"enforcer"`
	Err      error                    `json:"-"`
}

// EnforceResponse collects the response parameters for the Enforce method.
type EnforceResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// LoadPolicyResponse collects the response parameters for the LoadPolicy method.
type LoadPolicyResponse struct {
	Err error `json:"-"`
}

// SavePolicyResponse collects the response parameters for the SavePolicy method.
type SavePolicyResponse struct {
	Err error `json:"-"`
}

// NewAdapterResponse collects the response parameters for the NewAdapter method.
type NewAdapterResponse struct {
	Adapter services.AdapterHandler `json:"adapter"`
	Err     error                   `json:"-"`
}

// AddPolicyResponse collects the response parameters for the AddPolicy method.
type AddPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// AddNamedPolicyResponse collects the response parameters for the AddNamedPolicy method.
type AddNamedPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// RemovePolicyResponse collects the response parameters for the RemovePolicy method.
type RemovePolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// RemoveNamedPolicyResponse collects the response parameters for the RemoveNamedPolicy method.
type RemoveNamedPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// RemoveFilteredPolicyResponse collects the response parameters for the RemoveFilteredPolicy method.
type RemoveFilteredPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// RemoveFilteredNamedPolicyResponse collects the response parameters for the RemoveFilteredNamedPolicy method.
type RemoveFilteredNamedPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// GetPolicyResponse collects the response parameters for the GetPolicy method.
type GetPolicyResponse struct {
	Results [][]string `json:"results"`
	Err     error      `json:"-"`
}

// GetNamedPolicyResponse collects the response parameters for the GetNamedPolicy method.
type GetNamedPolicyResponse struct {
	Results [][]string `json:"results"`
	Err     error      `json:"-"`
}

// GetFilteredPolicyResponse collects the response parameters for the GetFilteredPolicy method.
type GetFilteredPolicyResponse struct {
	Results [][]string `json:"results"`
	Err     error      `json:"-"`
}

// GetFilteredNamedPolicyResponse collects the response parameters for the GetFilteredNamedPolicy method.
type GetFilteredNamedPolicyResponse struct {
	Results [][]string `json:"results"`
	Err     error      `json:"-"`
}

// AddGroupingPolicyResponse collects the response parameters for the AddGroupingPolicy method.
type AddGroupingPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// AddNamedGroupingPolicyResponse collects the response parameters for the AddNamedGroupingPolicy method.
type AddNamedGroupingPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// RemoveGroupingPolicyResponse collects the response parameters for the RemoveGroupingPolicy method.
type RemoveGroupingPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// RemoveNamedGroupingPolicyResponse collects the response parameters for the RemoveNamedGroupingPolicy method.
type RemoveNamedGroupingPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// RemoveFilteredGroupingPolicyResponse collects the response parameters for the RemoveFilteredGroupingPolicy method.
type RemoveFilteredGroupingPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// RemoveFilteredNamedGroupingPolicyResponse collects the response parameters for the RemoveFilteredNamedGroupingPolicy method.
type RemoveFilteredNamedGroupingPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// GetGroupingPolicyResponse collects the response parameters for the GetGroupingPolicy method.
type GetGroupingPolicyResponse struct {
	Results [][]string `json:"results"`
	Err     error      `json:"-"`
}

// GetNamedGroupingPolicyResponse collects the response parameters for the GetNamedGroupingPolicy method.
type GetNamedGroupingPolicyResponse struct {
	Results [][]string `json:"results"`
	Err     error      `json:"-"`
}

// GetFilteredGroupingPolicyResponse collects the response parameters for the GetFilteredGroupingPolicy method.
type GetFilteredGroupingPolicyResponse struct {
	Results [][]string `json:"results"`
	Err     error      `json:"-"`
}

// GetFilteredNamedGroupingPolicyResponse collects the response parameters for the GetFilteredNamedGroupingPolicy method.
type GetFilteredNamedGroupingPolicyResponse struct {
	Results [][]string `json:"results"`
	Err     error      `json:"-"`
}

// GetAllSubjectsResponse collects the response parameters for the GetAllSubjects method.
type GetAllSubjectsResponse struct {
	Results []string `json:"results"`
	Err     error    `json:"-"`
}

// GetAllNamedSubjectsResponse collects the response parameters for the GetAllNamedSubjects method.
type GetAllNamedSubjectsResponse struct {
	Results []string `json:"results"`
	Err     error    `json:"-"`
}

// GetAllObjectsResponse collects the response parameters for the GetAllObjects method.
type GetAllObjectsResponse struct {
	Results []string `json:"results"`
	Err     error    `json:"-"`
}

// GetAllNamedObjectsResponse collects the response parameters for the GetAllNamedObjects method.
type GetAllNamedObjectsResponse struct {
	Results []string `json:"results"`
	Err     error    `json:"-"`
}

// GetAllActionsResponse collects the response parameters for the GetAllActions method.
type GetAllActionsResponse struct {
	Results []string `json:"results"`
	Err     error    `json:"-"`
}

// GetAllNamedActionsResponse collects the response parameters for the GetAllNamedActions method.
type GetAllNamedActionsResponse struct {
	Results []string `json:"results"`
	Err     error    `json:"-"`
}

// GetAllRolesResponse collects the response parameters for the GetAllRoles method.
type GetAllRolesResponse struct {
	Results []string `json:"results"`
	Err     error    `json:"-"`
}

// GetAllNamedRolesResponse collects the response parameters for the GetAllNamedRoles method.
type GetAllNamedRolesResponse struct {
	Results []string `json:"results"`
	Err     error    `json:"-"`
}

// HasPolicyResponse collects the response parameters for the HasPolicy method.
type HasPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// HasNamedPolicyResponse collects the response parameters for the HasNamedPolicy method.
type HasNamedPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// HasGroupingPolicyResponse collects the response parameters for the HasGroupingPolicy method.
type HasGroupingPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// HasNamedGroupingPolicyResponse collects the response parameters for the HasNamedGroupingPolicy method.
type HasNamedGroupingPolicyResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// HasRoleForUserResponse collects the response parameters for the HasRoleForUser method.
type HasRoleForUserResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// AddRoleForUserResponse collects the response parameters for the AddRoleForUser method.
type AddRoleForUserResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// DeleteRoleForUserResponse collects the response parameters for the DeleteRoleForUser method.
type DeleteRoleForUserResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// DeleteRolesForUserResponse collects the response parameters for the DeleteRolesForUser method.
type DeleteRolesForUserResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// DeleteUserResponse collects the response parameters for the DeleteUser method.
type DeleteUserResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// DeleteRoleResponse collects the response parameters for the DeleteRole method.
type DeleteRoleResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// DeletePermissionResponse collects the response parameters for the DeletePermission method.
type DeletePermissionResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// GetRolesForUserResponse collects the response parameters for the GetRolesForUser method.
type GetRolesForUserResponse struct {
	Results []string `json:"results"`
	Err     error    `json:"-"`
}

// GetImplicitRolesForUserResponse collects the response parameters for the GetImplicitRolesForUser method.
type GetImplicitRolesForUserResponse struct {
	Results []string `json:"results"`
	Err     error    `json:"-"`
}

// GetUsersForRoleResponse collects the response parameters for the GetUsersForRole method.
type GetUsersForRoleResponse struct {
	Results []string `json:"results"`
	Err     error    `json:"-"`
}

// AddPermissionForUserResponse collects the response parameters for the AddPermissionForUser method.
type AddPermissionForUserResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// DeletePermissionForUserResponse collects the response parameters for the DeletePermissionForUser method.
type DeletePermissionForUserResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// GetPermissionsForUserResponse collects the response parameters for the GetPermissionsForUser method.
type GetPermissionsForUserResponse struct {
	Results [][]string `json:"results"`
	Err     error      `json:"-"`
}

// GetImplicitPermissionsForUserResponse collects the response parameters for the GetImplicitPermissionsForUser method.
type GetImplicitPermissionsForUserResponse struct {
	Results [][]string `json:"results"`
	Err     error      `json:"-"`
}

// HasPermissionForUserResponse collects the response parameters for the HasPermissionForUser method.
type HasPermissionForUserResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"-"`
}

// compile time assertions for our response types implementing endpoint.Failer.
var (
	_ endpoint.Failer = NewEnforcerResponse{}
	_ endpoint.Failer = EnforceResponse{}
	_ endpoint.Failer = LoadPolicyResponse{}
	_ endpoint.Failer = SavePolicyResponse{}
	_ endpoint.Failer = NewAdapterResponse{}
	_ endpoint.Failer = AddPolicyResponse{}
	_ endpoint.Failer = AddNamedPolicyResponse{}
	_ endpoint.Failer = RemovePolicyResponse{}
	_ endpoint.Failer = RemoveNamedPolicyResponse{}
	_ endpoint.Failer = RemoveFilteredPolicyResponse{}
	_ endpoint.Failer = RemoveFilteredNamedPolicyResponse{}
	_ endpoint.Failer = GetPolicyResponse{}
	_ endpoint.Failer = GetNamedPolicyResponse{}
	_ endpoint.Failer = GetFilteredPolicyResponse{}
	_ endpoint.Failer = GetFilteredNamedPolicyResponse{}
	_ endpoint.Failer = AddGroupingPolicyResponse{}
	_ endpoint.Failer = AddNamedGroupingPolicyResponse{}
	_ endpoint.Failer = RemoveGroupingPolicyResponse{}
	_ endpoint.Failer = RemoveNamedGroupingPolicyResponse{}
	_ endpoint.Failer = RemoveFilteredGroupingPolicyResponse{}
	_ endpoint.Failer = RemoveFilteredNamedGroupingPolicyResponse{}
	_ endpoint.Failer = GetGroupingPolicyResponse{}
	_ endpoint.Failer = GetNamedGroupingPolicyResponse{}
	_ endpoint.Failer = GetFilteredGroupingPolicyResponse{}
	_ endpoint.Failer = GetFilteredNamedGroupingPolicyResponse{}
	_ endpoint.Failer = GetAllSubjectsResponse{}
	_ endpoint.Failer = GetAllNamedSubjectsResponse{}
	_ endpoint.Failer = GetAllObjectsResponse{}
	_ endpoint.Failer = GetAllNamedObjectsResponse{}
	_ endpoint.Failer = GetAllActionsResponse{}
	_ endpoint.Failer = GetAllNamedActionsResponse{}
	_ endpoint.Failer = GetAllRolesResponse{}
	_ endpoint.Failer = GetAllNamedRolesResponse{}
	_ endpoint.Failer = HasPolicyResponse{}
	_ endpoint.Failer = HasNamedPolicyResponse{}
	_ endpoint.Failer = HasGroupingPolicyResponse{}
	_ endpoint.Failer = HasNamedGroupingPolicyResponse{}
	_ endpoint.Failer = HasRoleForUserResponse{}
	_ endpoint.Failer = AddRoleForUserResponse{}
	_ endpoint.Failer = DeleteRoleForUserResponse{}
	_ endpoint.Failer = DeleteRolesForUserResponse{}
	_ endpoint.Failer = DeleteUserResponse{}
	_ endpoint.Failer = DeleteRoleResponse{}
	_ endpoint.Failer = DeletePermissionResponse{}
	_ endpoint.Failer = GetRolesForUserResponse{}
	_ endpoint.Failer = GetImplicitRolesForUserResponse{}
	_ endpoint.Failer = GetUsersForRoleResponse{}
	_ endpoint.Failer = AddPermissionForUserResponse{}
	_ endpoint.Failer = DeletePermissionForUserResponse{}
	_ endpoint.Failer = GetPermissionsForUserResponse{}
	_ endpoint.Failer = GetImplicitPermissionsForUserResponse{}
	_ endpoint.Failer = HasPermissionForUserResponse{}
)

// Failed implements endpoint.Failer
func (r NewEnforcerResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r EnforceResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r LoadPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r SavePolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r NewAdapterResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r AddPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r AddNamedPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r RemovePolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r RemoveNamedPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r RemoveFilteredPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r RemoveFilteredNamedPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetNamedPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetFilteredPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetFilteredNamedPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r AddGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r AddNamedGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r RemoveGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r RemoveNamedGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r RemoveFilteredGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r RemoveFilteredNamedGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetNamedGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetFilteredGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetFilteredNamedGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetAllSubjectsResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetAllNamedSubjectsResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetAllObjectsResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetAllNamedObjectsResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetAllActionsResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetAllNamedActionsResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetAllRolesResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetAllNamedRolesResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r HasPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r HasNamedPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r HasGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r HasNamedGroupingPolicyResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r HasRoleForUserResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r AddRoleForUserResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r DeleteRoleForUserResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r DeleteRolesForUserResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r DeleteUserResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r DeleteRoleResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r DeletePermissionResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetRolesForUserResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetImplicitRolesForUserResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetUsersForRoleResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r AddPermissionForUserResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r DeletePermissionForUserResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetPermissionsForUserResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r GetImplicitPermissionsForUserResponse) Failed() error {
	return r.Err
}

// Failed implements endpoint.Failer
func (r HasPermissionForUserResponse) Failed() error {
	return r.Err
}
