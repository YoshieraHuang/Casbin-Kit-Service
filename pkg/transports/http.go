package transports

import (
	"casbinsvc/pkg/endpoints"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints endpoints.Set) http.Handler {
	options := []httptransport.ServerOption{}
	m := http.NewServeMux()
	m.Handle("/newEnforcer", httptransport.NewServer(endpoints.NewEnforcerEndpoint, decodeHTTPNewEnforcerRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/enforce", httptransport.NewServer(endpoints.EnforceEndpoint, decodeHTTPEnforceRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/loadPolicy", httptransport.NewServer(endpoints.LoadPolicyEndpoint, decodeHTTPLoadPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/savePolicy", httptransport.NewServer(endpoints.SavePolicyEndpoint, decodeHTTPSavePolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/newAdapter", httptransport.NewServer(endpoints.NewAdapterEndpoint, decodeHTTPNewAdapterRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/addPolicy", httptransport.NewServer(endpoints.AddPolicyEndpoint, decodeHTTPAddPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/addNamedPolicy", httptransport.NewServer(endpoints.AddNamedPolicyEndpoint, decodeHTTPAddNamedPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/removePolicy", httptransport.NewServer(endpoints.RemovePolicyEndpoint, decodeHTTPRemovePolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/removeNamedPolicy", httptransport.NewServer(endpoints.RemoveNamedPolicyEndpoint, decodeHTTPRemoveNamedPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/removeFilteredPolicy", httptransport.NewServer(endpoints.RemoveFilteredPolicyEndpoint, decodeHTTPRemoveFilteredPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/removeFilteredNamedPolicy", httptransport.NewServer(endpoints.RemoveFilteredNamedPolicyEndpoint, decodeHTTPRemoveFilteredNamedPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getPolicy", httptransport.NewServer(endpoints.GetPolicyEndpoint, decodeHTTPGetPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getNamedPolicy", httptransport.NewServer(endpoints.GetNamedPolicyEndpoint, decodeHTTPGetNamedPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getFilteredPolicy", httptransport.NewServer(endpoints.GetFilteredPolicyEndpoint, decodeHTTPGetFilteredPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getFilteredNamedPolicy", httptransport.NewServer(endpoints.GetFilteredNamedPolicyEndpoint, decodeHTTPGetFilteredNamedPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/addGroupingPolicy", httptransport.NewServer(endpoints.AddGroupingPolicyEndpoint, decodeHTTPAddGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/addNamedGroupingPolicy", httptransport.NewServer(endpoints.AddNamedGroupingPolicyEndpoint, decodeHTTPAddNamedGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/removeGroupingPolicy", httptransport.NewServer(endpoints.RemoveGroupingPolicyEndpoint, decodeHTTPRemoveGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/removeNamedGroupingPolicy", httptransport.NewServer(endpoints.RemoveNamedGroupingPolicyEndpoint, decodeHTTPRemoveNamedGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/removeFilteredGroupingPolicy", httptransport.NewServer(endpoints.RemoveFilteredGroupingPolicyEndpoint, decodeHTTPRemoveFilteredGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/removeFilteredNamedGroupingPolicy", httptransport.NewServer(endpoints.RemoveFilteredNamedGroupingPolicyEndpoint, decodeHTTPRemoveFilteredNamedGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getGroupingPolicy", httptransport.NewServer(endpoints.GetGroupingPolicyEndpoint, decodeHTTPGetGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getNamedGroupingPolicy", httptransport.NewServer(endpoints.GetNamedGroupingPolicyEndpoint, decodeHTTPGetNamedGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getFilteredGroupingPolicy", httptransport.NewServer(endpoints.GetFilteredGroupingPolicyEndpoint, decodeHTTPGetFilteredGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getFilteredNamedGroupingPolicy", httptransport.NewServer(endpoints.GetFilteredNamedGroupingPolicyEndpoint, decodeHTTPGetFilteredNamedGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getAllSubjects", httptransport.NewServer(endpoints.GetAllSubjectsEndpoint, decodeHTTPGetAllSubjectsRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getAllNamedSubjects", httptransport.NewServer(endpoints.GetAllNamedSubjectsEndpoint, decodeHTTPGetAllNamedSubjectsRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getAllObjects", httptransport.NewServer(endpoints.GetAllObjectsEndpoint, decodeHTTPGetAllObjectsRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getAllNamedObjects", httptransport.NewServer(endpoints.GetAllNamedObjectsEndpoint, decodeHTTPGetAllNamedObjectsRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getAllActions", httptransport.NewServer(endpoints.GetAllActionsEndpoint, decodeHTTPGetAllActionsRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getAllNamedActions", httptransport.NewServer(endpoints.GetAllNamedActionsEndpoint, decodeHTTPGetAllNamedActionsRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getAllRoles", httptransport.NewServer(endpoints.GetAllRolesEndpoint, decodeHTTPGetAllRolesRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getAllNamedRoles", httptransport.NewServer(endpoints.GetAllNamedRolesEndpoint, decodeHTTPGetAllNamedRolesRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/hasPolicy", httptransport.NewServer(endpoints.HasPolicyEndpoint, decodeHTTPHasPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/hasNamedPolicy", httptransport.NewServer(endpoints.HasNamedPolicyEndpoint, decodeHTTPHasNamedPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/hasGroupingPolicy", httptransport.NewServer(endpoints.HasGroupingPolicyEndpoint, decodeHTTPHasGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/hasNamedGroupingPolicy", httptransport.NewServer(endpoints.HasNamedGroupingPolicyEndpoint, decodeHTTPHasNamedGroupingPolicyRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/hasRoleForUser", httptransport.NewServer(endpoints.HasRoleForUserEndpoint, decodeHTTPHasRoleForUserRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/addRoleForUser", httptransport.NewServer(endpoints.AddRoleForUserEndpoint, decodeHTTPAddRoleForUserRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/deleteRoleForUser", httptransport.NewServer(endpoints.DeleteRoleForUserEndpoint, decodeHTTPDeleteRoleForUserRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/deleteRolesForUser", httptransport.NewServer(endpoints.DeleteRolesForUserEndpoint, decodeHTTPDeleteRolesForUserRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/deleteUser", httptransport.NewServer(endpoints.DeleteUserEndpoint, decodeHTTPDeleteUserRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/deleteRole", httptransport.NewServer(endpoints.DeleteRoleEndpoint, decodeHTTPDeleteRoleRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/deletePermission", httptransport.NewServer(endpoints.DeletePermissionEndpoint, decodeHTTPDeletePermissionRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getRolesForUser", httptransport.NewServer(endpoints.GetRolesForUserEndpoint, decodeHTTPGetRolesForUserRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getImplicitRolesForUser", httptransport.NewServer(endpoints.GetImplicitRolesForUserEndpoint, decodeHTTPGetImplicitRolesForUserRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getUsersForRole", httptransport.NewServer(endpoints.GetUsersForRoleEndpoint, decodeHTTPGetUsersForRoleRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/addPermissionForUser", httptransport.NewServer(endpoints.AddPermissionForUserEndpoint, decodeHTTPAddPermissionForUserRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/deletePermissionForUser", httptransport.NewServer(endpoints.DeletePermissionForUserEndpoint, decodeHTTPDeletePermissionForUserRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getPermissionsForUser", httptransport.NewServer(endpoints.GetPermissionsForUserEndpoint, decodeHTTPGetPermissionsForUserRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/getImplicitPermissionsForUser", httptransport.NewServer(endpoints.GetImplicitPermissionsForUserEndpoint, decodeHTTPGetImplicitPermissionsForUserRequest, encodeHTTPGenericResponse, options...))
	m.Handle("/hasPermissionForUser", httptransport.NewServer(endpoints.HasPermissionForUserEndpoint, decodeHTTPHasPermissionForUserRequest, encodeHTTPGenericResponse, options...))
	return m
}

// decodeHTTPNewEnforcerRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded newEnforcer request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPNewEnforcerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.NewEnforcerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPEnforceRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded enforce request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPEnforceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.EnforceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPLoadPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded loadPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPLoadPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.LoadPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPSavePolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded savePolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPSavePolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.SavePolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPNewAdapterRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded newAdapter request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPNewAdapterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.NewAdapterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPAddPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded addPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPAddPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.AddPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPAddNamedPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded addNamedPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPAddNamedPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.AddNamedPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPRemovePolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded removePolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPRemovePolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.RemovePolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPRemoveNamedPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded removeNamedPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPRemoveNamedPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.RemoveNamedPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPRemoveFilteredPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded removeFilteredPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPRemoveFilteredPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.RemoveFilteredPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPRemoveFilteredNamedPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded removeFilteredNamedPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPRemoveFilteredNamedPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.RemoveFilteredNamedPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetNamedPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getNamedPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetNamedPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetNamedPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetFilteredPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getFilteredPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetFilteredPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetFilteredPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetFilteredNamedPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getFilteredNamedPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetFilteredNamedPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetFilteredNamedPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPAddGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded addGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPAddGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.AddGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPAddNamedGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded addNamedGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPAddNamedGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.AddNamedGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPRemoveGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded removeGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPRemoveGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.RemoveGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPRemoveNamedGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded removeNamedGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPRemoveNamedGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.RemoveNamedGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPRemoveFilteredGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded removeFilteredGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPRemoveFilteredGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.RemoveFilteredGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPRemoveFilteredNamedGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded removeFilteredNamedGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPRemoveFilteredNamedGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.RemoveFilteredNamedGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetNamedGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getNamedGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetNamedGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetNamedGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetFilteredGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getFilteredGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetFilteredGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetFilteredGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetFilteredNamedGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getFilteredNamedGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetFilteredNamedGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetFilteredNamedGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetAllSubjectsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getAllSubjects request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetAllSubjectsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetAllSubjectsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetAllNamedSubjectsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getAllNamedSubjects request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetAllNamedSubjectsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetAllNamedSubjectsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetAllObjectsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getAllObjects request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetAllObjectsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetAllObjectsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetAllNamedObjectsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getAllNamedObjects request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetAllNamedObjectsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetAllNamedObjectsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetAllActionsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getAllActions request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetAllActionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetAllActionsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetAllNamedActionsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getAllNamedActions request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetAllNamedActionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetAllNamedActionsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetAllRolesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getAllRoles request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetAllRolesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetAllRolesRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetAllNamedRolesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getAllNamedRoles request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetAllNamedRolesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetAllNamedRolesRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPHasPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded hasPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPHasPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.HasPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPHasNamedPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded hasNamedPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPHasNamedPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.HasNamedPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPHasGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded hasGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPHasGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.HasGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPHasNamedGroupingPolicyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded hasNamedGroupingPolicy request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPHasNamedGroupingPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.HasNamedGroupingPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPHasRoleForUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded hasRoleForUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPHasRoleForUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.HasRoleForUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPAddRoleForUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded addRoleForUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPAddRoleForUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.AddRoleForUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPDeleteRoleForUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded deleteRoleForUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPDeleteRoleForUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.DeleteRoleForUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPDeleteRolesForUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded deleteRolesForUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPDeleteRolesForUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.DeleteRolesForUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPDeleteUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded deleteUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.DeleteUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPDeleteRoleRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded deleteRole request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPDeleteRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.DeleteRoleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPDeletePermissionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded deletePermission request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPDeletePermissionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.DeletePermissionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetRolesForUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getRolesForUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetRolesForUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetRolesForUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetImplicitRolesForUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getImplicitRolesForUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetImplicitRolesForUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetImplicitRolesForUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetUsersForRoleRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getUsersForRole request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetUsersForRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetUsersForRoleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPAddPermissionForUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded addPermissionForUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPAddPermissionForUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.AddPermissionForUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPDeletePermissionForUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded deletePermissionForUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPDeletePermissionForUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.DeletePermissionForUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetPermissionsForUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getPermissionsForUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetPermissionsForUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetPermissionsForUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPGetImplicitPermissionsForUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded getImplicitPermissionsForUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPGetImplicitPermissionsForUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetImplicitPermissionsForUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPHasPermissionForUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded hasPermissionForUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPHasPermissionForUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.HasPermissionForUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

type errorWrapper struct {
	Error string `json:"error"`
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

// encodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
