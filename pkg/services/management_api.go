package services

import "context"

// GetAllNamedSubjects gets the list of subjects that show up in the current named policy
func (s *service) GetAllNamedSubjects(ctx context.Context, enforcerHandler EnforcerHandler, pType string) ([]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetModel().GetValuesForFieldInPolicy("p", pType, 0), nil
}

// GetAllSubjects gets the list of subjects that show up in the current policy
func (s *service) GetAllSubjects(ctx context.Context, enforcerHandler EnforcerHandler) ([]string, error) {
	return s.GetAllNamedSubjects(ctx, enforcerHandler, "p")
}

// GetAllNamedObjects gets the list of objects that show up in the current policy
func (s *service) GetAllNamedObjects(ctx context.Context, enforcerHandler EnforcerHandler, pType string) ([]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetModel().GetValuesForFieldInPolicy("p", pType, 1), nil
}

// GetAllObjects gets the list of objects that show up in the current policy
func (s *service) GetAllObjects(ctx context.Context, enforcerHandler EnforcerHandler) ([]string, error) {
	return s.GetAllNamedObjects(ctx, enforcerHandler, "p")
}

// GetAllNamedActions gets the list of actions that show up in the current named policy
func (s *service) GetAllNamedActions(ctx context.Context, enforcerHandler EnforcerHandler, pType string) ([]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetModel().GetValuesForFieldInPolicy("p", pType, 2), nil
}

// GetAllActions gets the list of actions that show up in the current policy
func (s *service) GetAllActions(ctx context.Context, enforcerHandler EnforcerHandler) ([]string, error) {
	return s.GetAllNamedActions(ctx, enforcerHandler, "p")
}

// GetNamedPolicy gets all the authorization rules in the named policy
func (s *service) GetNamedPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string) ([][]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetModel().GetPolicy("p", pType), nil
}

// GetAllRoles gets the list of roles that show up in the current policy.
func (s *service) GetAllRoles(ctx context.Context, enforcerHandler EnforcerHandler) ([]string, error) {
	return s.GetAllNamedRoles(ctx, enforcerHandler, "g")
}

// GetAllNamedRoles gets the list of roles that show up in the current named policy.
func (s *service) GetAllNamedRoles(ctx context.Context, enforcerHandler EnforcerHandler, pType string) ([]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetModel().GetValuesForFieldInPolicy("g", pType, 1), nil
}

// GetPolicy gets all the authorization rules in the policy
func (s *service) GetPolicy(ctx context.Context, enforcerHandler EnforcerHandler) ([][]string, error) {
	return s.GetNamedPolicy(ctx, enforcerHandler, "p")
}

// GetFilteredNamedPolicy gets all authorization in the named policy, field filtering can be specified
func (s *service) GetFilteredNamedPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string, fieldIndex int, fieldValues []string) ([][]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetModel().GetFilteredPolicy("p", pType, fieldIndex, fieldValues...), nil
}

// GetFilteredPolicy gets all authorization in the policy, field filtering can be specified
func (s *service) GetFilteredPolicy(ctx context.Context, enforcerHandler EnforcerHandler, fieldIndex int, fieldValues []string) ([][]string, error) {
	return s.GetFilteredNamedPolicy(ctx, enforcerHandler, "p", fieldIndex, fieldValues)
}

// GetNamedGroupingPolicy gets all the role inheritance rules in the policy
func (s *service) GetNamedGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string) ([][]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetModel().GetPolicy("g", pType), nil
}

// GetGroupingPolicy gets all the role inheritance rules in the policy
func (s *service) GetGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler) ([][]string, error) {
	return s.GetNamedGroupingPolicy(ctx, enforcerHandler, "g")
}

// GetFilteredNamedGroupingPolicy gets all the role inheritance rules in the policy, field filtering can be specified
func (s *service) GetFilteredNamedGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string, fieldIndex int, fieldValues []string) ([][]string, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return nil, err
	}

	return e.GetModel().GetFilteredPolicy("g", pType, fieldIndex, fieldValues...), nil
}

// GetFilteredGroupingPolicy gets all the role inheritance rules in the policy, field filtering can be specified
func (s *service) GetFilteredGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler, fieldIndex int, fieldValues []string) ([][]string, error) {
	return s.GetFilteredNamedGroupingPolicy(ctx, enforcerHandler, "g", fieldIndex, fieldValues)
}

// HasNamedPolicy determines whether a named role inheritance rule exists
func (s *service) HasNamedPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string, params []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.GetModel().HasPolicy("p", pType, params), nil
}

// HasPolicy determines whether a named role inheritance rule exists
func (s *service) HasPolicy(ctx context.Context, enforcerHandler EnforcerHandler, params []string) (bool, error) {
	return s.HasNamedPolicy(ctx, enforcerHandler, "p", params)
}

// HasNamedGroupingPolicy determines whether a named role inheritance rule exists
func (s *service) HasNamedGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string, params []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.GetModel().HasPolicy("g", pType, params), nil
}

// HasGroupingPolicy determines whether a role inheritance rule exists
func (s *service) HasGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler, params []string) (bool, error) {
	return s.HasNamedGroupingPolicy(ctx, enforcerHandler, "g", params)
}

// AddNamedPolicy adds a named policy to the current policy
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (s *service) AddNamedPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string, params []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.AddNamedPolicy(pType, params)
}

// AddPolicy adds a authorization rule to the current policy
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (s *service) AddPolicy(ctx context.Context, enforcerHandler EnforcerHandler, params []string) (bool, error) {
	return s.AddNamedPolicy(ctx, enforcerHandler, "p", params)
}

// RemoveNamedPolicy removes a named authorization rule from the current policy
func (s *service) RemoveNamedPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string, params []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.RemoveNamedPolicy(pType, params)
}

// RemovePolicy removes an authorization rule from the current policy
func (s *service) RemovePolicy(ctx context.Context, enforcerHandler EnforcerHandler, params []string) (bool, error) {
	return s.RemoveNamedPolicy(ctx, enforcerHandler, "p", params)
}

// RemoveFilteredPolicy removes an authorization rule from the current policy, field filters can be specified.
func (s *service) RemoveFilteredPolicy(ctx context.Context, enforcerHandler EnforcerHandler, fieldIndex int, fieldValues []string) (bool, error) {
	return s.RemoveFilteredNamedPolicy(ctx, enforcerHandler, "p", fieldIndex, fieldValues)
}

// RemoveFilteredNamedPolicy removes an authorization rule from the current named policy, field filters can be specified.
func (s *service) RemoveFilteredNamedPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string, fieldIndex int, fieldValues []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.RemoveFilteredNamedPolicy(pType, fieldIndex, fieldValues...)
}

// AddNamedGroupingPolicy adds a named role inheritance rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (s *service) AddNamedGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string, params []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.AddNamedGroupingPolicy(pType, params)
}

// AddGroupingPolicy adds a role inheritance rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (s *service) AddGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler, params []string) (bool, error) {
	return s.AddNamedGroupingPolicy(ctx, enforcerHandler, "p", params)
}

// RemoveGroupingPolicy removes a role inheritance rule from the current policy.
func (s *service) RemoveGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler, params []string) (bool, error) {
	return s.RemoveNamedGroupingPolicy(ctx, enforcerHandler, "g", params)
}

// RemoveNamedGroupingPolicy removes a role inheritance rule from the current named policy.
func (s *service) RemoveNamedGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string, params []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.RemoveNamedGroupingPolicy(pType, params)
}

// RemoveFilteredGroupingPolicy removes a role inheritance rule from the current policy, field filters can be specified.
func (s *service) RemoveFilteredGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler, fieldIndex int, fieldValues []string) (bool, error) {
	return s.RemoveFilteredNamedGroupingPolicy(ctx, enforcerHandler, "g", fieldIndex, fieldValues)
}

// RemoveFilteredNamedGroupingPolicy removes a role inheritance rule from the current named policy, field filters can be specified.
func (s *service) RemoveFilteredNamedGroupingPolicy(ctx context.Context, enforcerHandler EnforcerHandler, pType string, fieldIndex int, fieldValues []string) (bool, error) {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	return e.RemoveFilteredNamedGroupingPolicy(pType, fieldIndex, fieldValues...)
}
