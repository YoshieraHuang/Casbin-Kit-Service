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

// GetRolesForUser gets the roles that a user has.
func (s *service) GetRolesForUser(ctx context.Context, enforcerHandler EnforcerHandler, user string) ([]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetModel()["g"]["g"].RM.GetRoles(user)
}

// GetImplicitPermissionsForUser gets all permissions(including children) for a user or role.
func (s *service) GetImplicitRolesForUser(ctx context.Context, enforcerHandler EnforcerHandler, user string) ([]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}
	return e.GetImplicitRolesForUser(user)
}

// GetUsersForRole gets the users that has a role.
func (s *service) GetUsersForRole(ctx context.Context, enforcerHandler EnforcerHandler, role string) ([]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetModel()["g"]["g"].RM.GetUsers(role)
}

// HasRoleForUser determines whether a user has a role.
func (s *service) HasRoleForUser(ctx context.Context, enforcerHandler EnforcerHandler, user, role string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	roles, err := e.GetRolesForUser(user)
	if err != nil {
		return false, err
	}

	// determine whether given role in roles
	for _, r := range roles {
		if r == role {
			return true, nil
		}
	}

	return false, nil
}

// AddRoleForUser adds a role for a user.
// Returns false if the user already has the role (aka not affected).
func (s *service) AddRoleForUser(ctx context.Context, enforcerHandler EnforcerHandler, user, role string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.AddGroupingPolicy(user, role)
}

// DeleteRoleForUser deletes a role for a user.
// Returns false if the user does not have the role (aka not affected).
func (s *service) DeleteRoleForUser(ctx context.Context, enforcerHandler EnforcerHandler, user, role string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.RemoveGroupingPolicy(user, role)
}

// DeleteRolesForUser deletes all roles for a user.
// Returns false if the user does not have any roles (aka not affected).
func (s *service) DeleteRolesForUser(ctx context.Context, enforcerHandler EnforcerHandler, user string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.RemoveFilteredGroupingPolicy(0, user)
}

// DeleteUser deletes a user.
// Returns false if the user does not exist (aka not affected).
func (s *service) DeleteUser(ctx context.Context, enforcerHandler EnforcerHandler, user string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.RemoveFilteredGroupingPolicy(0, user)
}

// DeleteRole deletes a role.
// Returns false if the role does not exist (aka not affected)
func (s *service) DeleteRole(ctx context.Context, enforcerHandler EnforcerHandler, role string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.DeleteRole(role)
}

// DeletePermission deletes a permission.
// Returns false if the permission does not exist (aka not affected).
func (s *service) DeletePermission(ctx context.Context, enforcerHandler EnforcerHandler, permissions []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.RemoveFilteredPolicy(1, permissions...)
}

// AddPermissionForUser adds a permission for a user or role.
// Returns false if the user or role already has the permission (aka not affected).
func (s *service) AddPermissionForUser(ctx context.Context, enforcerHandler EnforcerHandler, user string, permissions []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.AddPolicy(s.convertPermissions(user, permissions...)...)
}

// DeletePermissionForUser deletes a permission for a user or role.
// Returns false if the user or role does not have the permission (aka not affected).
func (s *service) DeletePermissionForUser(ctx context.Context, enforcerHandler EnforcerHandler, user string, permissions []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.RemovePolicy(s.convertPermissions(user, permissions...)...)
}

// DeletePermissionsForUser deletes permissions for a user or role.
// Returns false if the user or role does not have any permissions (aka not affected).
func (s *service) DeletePermissionsForUser(ctx context.Context, enforcerHandler EnforcerHandler, user string, permissions []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.RemoveFilteredPolicy(0, user)
}

// GetPermissionsForUser gets permissions for a user or role.
func (s *service) GetPermissionsForUser(ctx context.Context, enforcerHandler EnforcerHandler, user string, permissions []string) ([][]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetFilteredPolicy(0, user), nil
}

// GetImplicitPermissionsForUser gets all permissions(including children) for a user or role.
func (s *service) GetImplicitPermissionsForUser(ctx context.Context, enforcerHandler EnforcerHandler, user string, permissions []string) ([][]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetImplicitPermissionsForUser(user)
}

// HasPermissionForUser determines whether a user has a permission.
func (s *service) HasPermissionForUser(ctx context.Context, enforcerHandler EnforcerHandler, user string, permissions []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.HasPolicy(s.convertPermissions(user, permissions...)...), nil
}

func (s *service) convertPermissions(user string, permissions ...string) []interface{} {
	params := make([]interface{}, 0, len(permissions)+1)
	params = append(params, user)
	for _, perm := range permissions {
		params = append(params, perm)
	}

	return params
}
