package services

import (
	"context"

	"github.com/go-kit/kit/log"
)

// Middleware describes s service middleware
type Middleware func(Server) Server

// LoggingMiddleware takes a logger as a dependency and returns a service Middleware
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Server) Server {
		return &loggingMiddleware{logger: logger, next: next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Server
}

// NewEnforcer implements Server interface
func (mw *loggingMiddleware) NewEnforcer(ctx context.Context, modelText string, adapter AdapterHandler) (handler EnforcerHandler, err error) {
	defer func() {
		mw.logger.Log("method", "NewEnforcer", "modelText", modelText, "adapter", adapter, "enforcer", handler, "error", err)
	}()
	return mw.next.NewEnforcer(ctx, modelText, adapter)
}

// Enforce implements Server interface
func (mw *loggingMiddleware) Enforce(ctx context.Context, enforcer EnforcerHandler, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "Enforce", "enforcer", enforcer, "params", params, "result", result, "error", err)
	}()
	return mw.next.Enforce(ctx, enforcer, params)
}

// LoadPolicy implements Server interface
func (mw *loggingMiddleware) LoadPolicy(ctx context.Context, enforcer EnforcerHandler) (err error) {
	defer func() {
		mw.logger.Log("method", "LoadPolicy", "enforcer", enforcer, "error", err)
	}()
	return mw.next.LoadPolicy(ctx, enforcer)
}

// SavePolicy implements Server interface
func (mw *loggingMiddleware) SavePolicy(ctx context.Context, enforcer EnforcerHandler) (err error) {
	defer func() {
		mw.logger.Log("method", "SavePolicy", "enforcer", enforcer, "error", err)
	}()
	return mw.next.SavePolicy(ctx, enforcer)
}

// NewAdapter implements Server interface
func (mw *loggingMiddleware) NewAdapter(ctx context.Context, driverName, connectString string, dbSpecified bool) (handler EnforcerHandler, err error) {
	defer func() {
		mw.logger.Log("method", "NewAdapter", "driverName", driverName, "connectString", connectString, "dbSpecified", dbSpecified, "enforcer", handler, "error", err)
	}()
	return mw.next.NewAdapter(ctx, driverName, connectString, dbSpecified)
}

// AddPolicy implements Server interface
func (mw *loggingMiddleware) AddPolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "AddPolicy", "enforcer", enforcer, "params", params, "result", result, "error", err)
	}()
	return mw.next.AddPolicy(ctx, enforcer, params)
}

// AddNamedPolicy implements Server interface
func (mw *loggingMiddleware) AddNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "AddNamedPolicy", "enforcer", enforcer, "pType", pType, "params", params, "result", result, "error", err)
	}()
	return mw.next.AddNamedPolicy(ctx, enforcer, pType, params)
}

// RemovePolicy implements Server interface
func (mw *loggingMiddleware) RemovePolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "RemovePolicy", "enforcer", enforcer, "params", params, "result", result, "error", err)
	}()
	return mw.next.RemovePolicy(ctx, enforcer, params)
}

// RemoveNamedPolicy implements Server interface
func (mw *loggingMiddleware) RemoveNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "RemoveNamedPolicy", "enforcer", enforcer, "pType", pType, "params", params, "result", result, "error", err)
	}()
	return mw.next.RemoveNamedPolicy(ctx, enforcer, pType, params)
}

// RemoveFilterePolicy implements Server interface
func (mw *loggingMiddleware) RemoveFilteredPolicy(ctx context.Context, enforcer EnforcerHandler, fieldIndex int, fieldValues []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "RemoveFilteredPolicy", "enforcer", enforcer, "fieldIndex", fieldIndex, "fieldValues", fieldValues, "result", result, "error", err)
	}()
	return mw.next.RemoveFilteredPolicy(ctx, enforcer, fieldIndex, fieldValues)
}

// RevmoeFilteredNamedPolicy implements Server interface
func (mw *loggingMiddleware) RemoveFilteredNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, fieldIndex int, fieldValues []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "RemoveFilteredNamedPolicy", "enforcer", enforcer, "pType", pType, "fieldIndex", fieldIndex, "fieldValues", fieldValues, "result", result, "error", err)
	}()
	return mw.next.RemoveFilteredNamedPolicy(ctx, enforcer, pType, fieldIndex, fieldValues)
}

// GetPolicy implements Server interface
func (mw *loggingMiddleware) GetPolicy(ctx context.Context, enforcer EnforcerHandler) (results [][]string, err error) {
	defer func() {
		mw.logger.Log("method", "GetPolicy", "enforcer", enforcer, "results", results, "error", err)
	}()
	return mw.next.GetPolicy(ctx, enforcer)
}

// GetNamedPolicy implements Server interface
func (mw *loggingMiddleware) GetNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string) (results [][]string, err error) {
	defer func() {
		mw.logger.Log("method", "GetNamedPolicy", "enforcer", enforcer, "pType", pType, "results", results, "error", err)
	}()
	return mw.next.GetNamedPolicy(ctx, enforcer, pType)
}

// GetFilteredPolicy implements Server interface
func (mw *loggingMiddleware) GetFilteredPolicy(ctx context.Context, enforcer EnforcerHandler, fieldIndex int, fieldValues []string) (results [][]string, err error) {
	defer func() {
		mw.logger.Log("method", "GetFilteredPolicy", "enforcer", enforcer, "fieldIndex", fieldIndex, "fieldValues", fieldValues, "results", results, "error", err)
	}()
	return mw.next.GetFilteredPolicy(ctx, enforcer, fieldIndex, fieldValues)
}

// GetFilteredNamedPolicy implements Server interface
func (mw *loggingMiddleware) GetFilteredNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, fieldIndex int, fieldValues []string) (results [][]string, err error) {
	defer func() {
		mw.logger.Log("method", "GetFilteredNamedPolicy", "enforcer", enforcer, "pType", pType, "fieldIndex", fieldIndex, "fieldValues", fieldValues, "results", results, "error", err)
	}()
	return mw.next.GetFilteredNamedPolicy(ctx, enforcer, pType, fieldIndex, fieldValues)
}

// AddGroupingPolicy implements Server interface
func (mw *loggingMiddleware) AddGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "AddGroupingPolicy", "enforcer", enforcer, "params", params, "result", result, "error", err)
	}()
	return mw.next.AddGroupingPolicy(ctx, enforcer, params)
}

// AddNamedGroupingPolicy implements Server interface
func (mw *loggingMiddleware) AddNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "AddNamedGroupingPolicy", "enforcer", enforcer, "pType", pType, "params", params, "result", result, "error", err)
	}()
	return mw.next.AddNamedGroupingPolicy(ctx, enforcer, pType, params)
}

// RemoveGroupingPolicy implements Server interface
func (mw *loggingMiddleware) RemoveGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "RemoveGroupingPolicy", "enforcer", enforcer, "params", params, "result", result, "error", err)
	}()
	return mw.next.RemoveGroupingPolicy(ctx, enforcer, params)
}

// RemoveNamedGroupingPolicy implements Server interface
func (mw *loggingMiddleware) RemoveNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "RemoveNamedGroupingPolicy", "enforcer", enforcer, "pType", pType, "params", params, "result", result, "error", err)
	}()
	return mw.next.RemoveNamedGroupingPolicy(ctx, enforcer, pType, params)
}

// RemoveFilteredGroupingPolicy implements Server interface
func (mw *loggingMiddleware) RemoveFilteredGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, fieldIndex int, fieldValues []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "RemoveFilteredGroupingPolicy", "enforcer", enforcer, "fieldIndex", fieldIndex, "fieldValues", fieldValues, "result", result, "error", err)
	}()
	return mw.next.RemoveFilteredGroupingPolicy(ctx, enforcer, fieldIndex, fieldValues)
}

// RemoveFilteredNamedGroupingPolicy implements Server interface
func (mw *loggingMiddleware) RemoveFilteredNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, fieldIndex int, fieldValues []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "RemoveFilteredNamedGroupingPolicy", "enforcer", enforcer, "pType", pType, "fieldIndex", fieldIndex, "fieldValues", fieldValues, "result", result, "error", err)
	}()
	return mw.next.RemoveFilteredNamedGroupingPolicy(ctx, enforcer, pType, fieldIndex, fieldValues)
}

// GetGroupingPolicy implements Server interface
func (mw *loggingMiddleware) GetGroupingPolicy(ctx context.Context, enforcer EnforcerHandler) (results [][]string, err error) {
	defer func() {
		mw.logger.Log("method", "GetGroupingPolicy", "enforcer", enforcer, "results", results, "error", err)
	}()
	return mw.next.GetGroupingPolicy(ctx, enforcer)
}

// GetNamedPolicy implements Server interface
func (mw *loggingMiddleware) GetNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string) (results [][]string, err error) {
	defer func() {
		mw.logger.Log("method", "GetNamedPolicy", "enforcer", enforcer, "pType", pType, "results", results, "error", err)
	}()
	return mw.next.GetNamedPolicy(ctx, enforcer, pType)
}

// GetFilteredGroupingPolicy implements Server interface
func (mw *loggingMiddleware) GetFilteredGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, fieldIndex int, fieldValues []string) (results [][]string, err error) {
	defer func() {
		mw.logger.Log("method", "GetFilteredGroupingPolicy", "enforcer", enforcer, "fieldIndex", fieldIndex, "fieldValues", fieldValues, "results", results, "error", err)
	}()
	return mw.next.GetFilteredGroupingPolicy(ctx, enforcer, fieldIndex, fieldValues)
}

// GetFilteredNamedGroupingPolicy implements Server interface
func (mw *loggingMiddleware) GetFilteredNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, fieldIndex int, fieldValues []string) (results [][]string, err error) {
	defer func() {
		mw.logger.Log("method", "GetFilteredNamedGroupingPolicy", "enforcer", enforcer, "pType", pType, "fieldIndex", fieldIndex, "fieldValues", fieldValues, "results", results, "error", err)
	}()
	return mw.next.GetFilteredNamedGroupingPolicy(ctx, enforcer, pType, fieldIndex, fieldValues)
}

// GetAllSubjects implements Server interface
func (mw *loggingMiddleware) GetAllSubjects(ctx context.Context, enforcer EnforcerHandler) (results []string, err error) {
	defer func() {
		mw.logger.Log("method", "GetAllSubjects", "enforcer", enforcer, "resutls", results, "error", err)
	}()
	return mw.next.GetAllSubjects(ctx, enforcer)
}

// GetAllNamedSubjects implements Server interface
func (mw *loggingMiddleware) GetAllNamedSubjects(ctx context.Context, enforcer EnforcerHandler, pType string) (results []string, err error) {
	defer func() {
		mw.logger.Log("method", "GetAllNamedSubjects", "enforcer", enforcer, "pType", pType, "results", results, "error", err)
	}()
	return mw.next.GetAllNamedSubjects(ctx, enforcer, pType)
}

// GetAllObjects implements Server interface
func (mw *loggingMiddleware) GetAllObjects(ctx context.Context, enforcer EnforcerHandler) (results []string, err error) {
	defer func() {
		mw.logger.Log("method", "GetAllObjects", "enforcer", enforcer, "results", results, "error", err)
	}()
	return mw.next.GetAllObjects(ctx, enforcer)
}

// GetAllNamedObjects implements Server interface
func (mw *loggingMiddleware) GetAllNamedObjects(ctx context.Context, enforcer EnforcerHandler, pType string) (results []string, err error) {
	defer func() {
		mw.logger.Log("method", "GetAllNamedObjects", "enforcer", enforcer, "pType", pType)
	}()
	return mw.next.GetAllNamedObjects(ctx, enforcer, pType)
}

// GetAllActions implements Server interface
func (mw *loggingMiddleware) GetAllActions(ctx context.Context, enforcer EnforcerHandler) (results []string, err error) {
	defer func() {
		mw.logger.Log("method", "GetAllActions", "enforcer", enforcer, "results", results, "error", err)
	}()
	return mw.next.GetAllActions(ctx, enforcer)
}

// GetAllNamedActions implements Server interface
func (mw *loggingMiddleware) GetAllNamedActions(ctx context.Context, enforcer EnforcerHandler, pType string) (results []string, err error) {
	defer func() {
		mw.logger.Log("method", "GetAllNamedActions", "enforcer", enforcer, "pType", pType, "results", results, "error", err)
	}()
	return mw.next.GetAllNamedActions(ctx, enforcer, pType)
}

// GetAllRoles implements Server interface
func (mw *loggingMiddleware) GetAllRoles(ctx context.Context, enforcer EnforcerHandler) (results []string, err error) {
	defer func() {
		mw.logger.Log("method", "GetAllRoles", "enforcer", enforcer, "results", results, "error", err)
	}()
	return mw.next.GetAllRoles(ctx, enforcer)
}

// GetAllNamedRoles implements Server interface
func (mw *loggingMiddleware) GetAllNamedRoles(ctx context.Context, enforcer EnforcerHandler, pType string) (results []string, err error) {
	defer func() {
		mw.logger.Log("method", "GetAllNamedRoles", "enforcer", enforcer, "pType", pType, "results", results, "error", err)
	}()
	return mw.next.GetAllNamedRoles(ctx, enforcer, pType)
}

// HasPolicy implements Server interface
func (mw *loggingMiddleware) HasPolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "HasPolicy", "enforcer", enforcer, "params", params, "result", result, "error", err)
	}()
	return mw.next.HasPolicy(ctx, enforcer, params)
}

// HasNamedPolicy implements Server interface
func (mw *loggingMiddleware) HasNamedPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "HasNamedPolicy", "enforcer", enforcer, "pType", pType, "params", params, "result", result, "error", err)
	}()
	return mw.next.HasNamedPolicy(ctx, enforcer, pType, params)
}

// HasGroupingPolicy implements Server interface
func (mw *loggingMiddleware) HasGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "HasGroupingPolicy", "enforcer", enforcer, "params", params, "result", result, "error", err)
	}()
	return mw.next.HasGroupingPolicy(ctx, enforcer, params)
}

// HasNamedGroupingPolicy implements Server interface
func (mw *loggingMiddleware) HasNamedGroupingPolicy(ctx context.Context, enforcer EnforcerHandler, pType string, params []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "HasNamedGroupingPolicy", "enforcer", enforcer, "pType", pType, "params", params, "result", result, "error", err)
	}()
	return mw.next.HasNamedGroupingPolicy(ctx, enforcer, pType, params)
}

// HasRoleForUser implements Server interface
func (mw *loggingMiddleware) HasRoleForUser(ctx context.Context, enforcer EnforcerHandler, user, role string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "HasRoleForUser", "enforcer", enforcer, "user", user, "role", role, "result", result, "error", err)
	}()
	return mw.next.HasRoleForUser(ctx, enforcer, user, role)
}

// AddRoleForUser implements Server interface
func (mw *loggingMiddleware) AddRoleForUser(ctx context.Context, enforcer EnforcerHandler, user, role string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "AddRoleForUser", "enforcer", enforcer, "user", user, "role", role, "result", result, "error", err)
	}()
	return mw.next.AddRoleForUser(ctx, enforcer, user, role)
}

// DeleteRoleForUser implements Server interface
func (mw *loggingMiddleware) DeleteRoleForUser(ctx context.Context, enforcer EnforcerHandler, user, role string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "DeleteRoleForUser", "enforcer", enforcer, "user", user, "role", role, "result", result, "error", err)
	}()
	return mw.next.DeleteRoleForUser(ctx, enforcer, user, role)
}

// DeleteRolesForUser implements Server interface
func (mw *loggingMiddleware) DeleteRolesForUser(ctx context.Context, enforcer EnforcerHandler, user string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "DeleteRolesForUser", "enforcer", enforcer, "user", user, "result", result, "error", err)
	}()
	return mw.next.DeleteRolesForUser(ctx, enforcer, user)
}

// DeleteUser implements Server interface
func (mw *loggingMiddleware) DeleteUser(ctx context.Context, enforcer EnforcerHandler, user string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "DeleteUser", "enforcer", enforcer, "user", user, "result", result, "error", err)
	}()
	return mw.next.DeleteUser(ctx, enforcer, user)
}

// DeleteRole implements Server interface
func (mw *loggingMiddleware) DeleteRole(ctx context.Context, enforcer EnforcerHandler, role string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "DeleteRole", "enforcer", enforcer, "role", role, "result", result, "error", err)
	}()
	return mw.next.DeleteRole(ctx, enforcer, role)
}

// DeletePermission implements Server interface
func (mw *loggingMiddleware) DeletePermission(ctx context.Context, enforcer EnforcerHandler, permission []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "DeletePermission", "enforcer", enforcer, "permission", permission, "result", result, "error", err)
	}()
	return mw.next.DeletePermission(ctx, enforcer, permission)
}

// GetRolesForUser implements Server interface
func (mw *loggingMiddleware) GetRolesForUser(ctx context.Context, enforcer EnforcerHandler, user string) (results []string, err error) {
	defer func() {
		mw.logger.Log("method", "GetRolesForUser", "enforcer", enforcer, "user", user, "results", results, "error", err)
	}()
	return mw.next.GetRolesForUser(ctx, enforcer, user)
}

// GetImplicitRolesForUser implements Server interface
func (mw *loggingMiddleware) GetImplicitRolesForUser(ctx context.Context, enforcer EnforcerHandler, user string) (results []string, err error) {
	defer func() {
		mw.logger.Log("method", "GetImplicitRolesForUser", "enforcer", enforcer, "user", user, "results", results, "error", err)
	}()
	return mw.next.GetImplicitRolesForUser(ctx, enforcer, user)
}

// GetUsersForRole implements Server interface
func (mw *loggingMiddleware) GetUsersForRole(ctx context.Context, enforcer EnforcerHandler, role string) (results []string, err error) {
	defer func() {
		mw.logger.Log("method", "GetUsersForRole", "enforcer", enforcer, "role", role, "results", results, "error", err)
	}()
	return mw.next.GetUsersForRole(ctx, enforcer, role)
}

// AddPermissionForUser implements Server interface
func (mw *loggingMiddleware) AddPermissionForUser(ctx context.Context, enforcer EnforcerHandler, user string, permissions []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "AddPermissionForUser", "enforcer", enforcer, "user", user, "permissions", permissions, "result", result, "error", err)
	}()
	return mw.next.AddPermissionForUser(ctx, enforcer, user, permissions)
}

// DeletePermissionForUser implements Server interface
func (mw *loggingMiddleware) DeletePermissionForUser(ctx context.Context, enforcer EnforcerHandler, user string, permissions []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "DeletePermissionForUser", "enforcer", enforcer, "user", user, "permissions", permissions, "result", result, "error", err)
	}()
	return mw.next.DeletePermissionForUser(ctx, enforcer, user, permissions)
}

// GetPermissionsForUser implements Server interface
func (mw *loggingMiddleware) GetPermissionsForUser(ctx context.Context, enforcer EnforcerHandler, user string, permissions []string) (results [][]string, err error) {
	defer func() {
		mw.logger.Log("method", "GetPermissionsForUser", "enforcer", enforcer, "user", user, "permissions", permissions, "results", results, "error", err)
	}()
	return mw.next.GetPermissionsForUser(ctx, enforcer, user, permissions)
}

// GetImplicitPermissionsForUser implements Server interface
func (mw *loggingMiddleware) GetImplicitPermissionsForUser(ctx context.Context, enforcer EnforcerHandler, user string, permissions []string) (results [][]string, err error) {
	defer func() {
		mw.logger.Log("method", "GetImplicitPermissionsForUser", "enforcer", enforcer, "user", user, "permissions", permissions, "results", results, "error", err)
	}()
	return mw.next.GetImplicitPermissionsForUser(ctx, enforcer, user, permissions)
}

// HasPermissionForUser implements Server interface
func (mw *loggingMiddleware) HasPermissionForUser(ctx context.Context, enforcer EnforcerHandler, user string, permissions []string) (result bool, err error) {
	defer func() {
		mw.logger.Log("method", "HasPermissionForUser", "enforcer", enforcer, "user", user, "permissions", permissions, "result", result, "error", err)
	}()
	return mw.next.HasPermissionForUser(ctx, enforcer, user, permissions)
}
