// Copyright 2018 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package services

import "context"

// AdapterHandler is an adapter handler to handle multi adapters
// an alias for int32
type AdapterHandler = int32

// EnforcerHandler is an enforcer handler to handle multi enforcers
// an alias for int32
type EnforcerHandler = int32

// Server defines the backend services of casbinsvc
type Server interface {
	// implementations in enforcer.go
	NewEnforcer(ctx context.Context, modelText string, adapter AdapterHandler) (EnforcerHandler, error)
	Enforce(ctx context.Context, enforcer EnforcerHandler, params []string) (bool, error)
	LoadPolicy(ctx context.Context, enforcer EnforcerHandler) error
	SavePolicy(ctx context.Context, enforcer EnforcerHandler) error

	// implementations in adapter.go
	NewAdapter(ctx context.Context, driverName, connectString string, dbSpecified bool) (EnforcerHandler, error)

	// implementations in management_api.go
	AddPolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (bool, error)
	AddNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (bool, error)
	RemovePolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (bool, error)
	RemoveNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (bool, error)
	RemoveFilteredPolicy(ctx context.Context, enforcer EnforcerHandler, fieldIndex int, fieldValues []string) (bool, error)
	RemoveFilteredNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, fieldIndex int, fieldValues []string) (bool, error)
	GetPolicy(ctx context.Context, enforcer EnforcerHandler) ([][]string, error)
	GetNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string) ([][]string, error)
	GetFilteredPolicy(ctx context.Context, enforcer EnforcerHandler, fieldIndex int, fieldValues []string) ([][]string, error)
	GetFilteredNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, fieldIndex int, fieldValues []string) ([][]string, error)
	AddGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (bool, error)
	AddNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (bool, error)
	RemoveGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (bool, error)
	RemoveNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (bool, error)
	RemoveFilteredGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, fieldIndex int, fieldValues []string) (bool, error)
	RemoveFilteredNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, fieldIndex int, fieldValues []string) (bool, error)
	GetGroupingPolicy(ctx context.Context, enforcer EnforcerHandler) ([][]string, error)
	GetNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string) ([][]string, error)
	GetFilteredGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, fieldIndex int, fieldValues []string) ([][]string, error)
	GetFilteredNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, fieldIndex int, fieldValues []string) ([][]string, error)
	GetAllSubjects(ctx context.Context, enforcer EnforcerHandler) ([]string, error)
	GetAllNamedSubjects(ctx context.Context, enforcer EnforcerHandler, pType string) ([]string, error)
	GetAllObjects(ctx context.Context, enforcer EnforcerHandler) ([]string, error)
	GetAllNamedObjects(ctx context.Context, enforcer EnforcerHandler, pType string) ([]string, error)
	GetAllActions(ctx context.Context, enforcer EnforcerHandler) ([]string, error)
	GetAllNamedActions(ctx context.Context, enforcer EnforcerHandler, pType string) ([]string, error)
	GetAllRoles(ctx context.Context, enforcer EnforcerHandler) ([]string, error)
	GetAllNamedRoles(ctx context.Context, enforcer EnforcerHandler, pType string) ([]string, error)
	HasPolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (bool, error)
	HasNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (bool, error)
	HasGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (bool, error)
	HasNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (bool, error)

	// rbac api
	HasRoleForUser(ctx context.Context, enforcer EnforcerHandler, user, role string) (bool, error)
	AddRoleForUser(ctx context.Context, enforcer EnforcerHandler, user, role string) (bool, error)
	DeleteRoleForUser(ctx context.Context, enforcer EnforcerHandler, user, role string) (bool, error)
	DeleteRolesForUser(ctx context.Context, enforcer EnforcerHandler, user string) (bool, error)
	DeleteUser(ctx context.Context, enforcer EnforcerHandler, user string) (bool, error)
	DeleteRole(ctx context.Context, enforcer EnforcerHandler, role string) (bool, error)
	DeletePermission(ctx context.Context, enforcer EnforcerHandler, permission []string) (bool, error)
	GetRolesForUser(ctx context.Context, enforcer EnforcerHandler, user string) ([]string, error)
	GetImplicitRolesForUser(ctx context.Context, enforcer EnforcerHandler, user string) ([]string, error)
	GetUsersForRole(ctx context.Context, enforcer EnforcerHandler, role string) ([]string, error)
	AddPermissionForUser(ctx context.Context, enforcer EnforcerHandler, user string, permissions []string) (bool, error)
	DeletePermissionForUser(ctx context.Context, enforcer EnforcerHandler, user string, permissions []string) (bool, error)
	GetPermissionsForUser(ctx context.Context, enforcer EnforcerHandler, user string, permissions []string) ([][]string, error)
	GetImplicitPermissionsForUser(ctx context.Context, enforcer EnforcerHandler, user string, permissions []string) ([][]string, error)
	HasPermissionForUser(ctx context.Context, enforcer EnforcerHandler, user string, permissions []string) (bool, error)
}
