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

import (
	"context"
	"errors"
	"strings"

	"github.com/casbin/casbin/v2/model"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
)

var (
	// ErrEnforcerNotFound is an error indicating that enforcer is not found
	ErrEnforcerNotFound = errors.New("enforcer not found")
	// ErrAdapterNotFound is an error indicating that adapter is not found
	ErrAdapterNotFound = errors.New("adapter not found")
)

func (s *service) getEnforcer(handle EnforcerHandler) (*casbin.Enforcer, error) {
	if _, ok := s.enforcerMap[handle]; ok {
		return s.enforcerMap[handle], nil
	}
	return nil, ErrEnforcerNotFound
}

func (s *service) addEnforcer(e *casbin.Enforcer) EnforcerHandler {
	cnt := EnforcerHandler(len(s.enforcerMap))
	s.enforcerMap[cnt] = e
	return cnt
}

// NewEnforcer construct a new enforcer and stored it
func (s *service) NewEnforcer(ctx context.Context, modelText string, adapterHandler AdapterHandler) (EnforcerHandler, error) {
	var a persist.Adapter
	var e *casbin.Enforcer

	// Fetch adapter and policies if adapterHandler is given
	if adapterHandler != -1 {
		var err error
		a, err = s.getAdapter(adapterHandler)
		if err != nil {
			return -1, err
		}
	}

	// TODO: load default model
	if modelText == "" {
		return -1, errors.New("Empty Model")
	}

	if a == nil {
		// init an enforcer without policies
		m, err := model.NewModelFromString(modelText)
		if err != nil {
			return -1, err
		}

		e, err = casbin.NewEnforcer(m)
		if err != nil {
			return -1, err
		}
	} else {
		// init an enforcer with policies
		m, err := model.NewModelFromString(modelText)
		if err != nil {
			return -1, err
		}

		e, err = casbin.NewEnforcer(m, a)
		if err != nil {
			return -1, err
		}
	}

	// Add enforcer to map
	return s.addEnforcer(e), nil
}

func (s *service) parseParams(params []string, m string) ([]interface{}, string) {
	// init a []interface{}, cap set as len(params) to minimize unnecessary allocations
	ps := make([]interface{}, 0, len(params))

	for _, p := range params {

		if strings.HasPrefix(p, "ABAC::") {
			// ABAC case, parse param
			attrList, err := resolveABAC(p)
			if err != nil {
				panic(err)
			}
			for k, v := range attrList.nameMap {
				old := "." + k
				if strings.Contains(m, old) {
					m = strings.Replace(m, old, "."+v, -1)
				}
			}
			// add attrList to ps
			ps = append(ps, attrList)
		} else {
			// common case

			// add params to ps
			ps = append(ps, p)
		}
	}

	return ps, m
}

// Enforce do the actual authorization, which decides whether a "subject" can access a "object" with the operation "action",
// input parameters are usually: (matcher, sub, obj, act), use model matcher by default when matcher is "".
func (s *service) Enforce(ctx context.Context, enforcerHandler EnforcerHandler, params []string) (bool, error) {
	// get enforcer by enforcerHandler
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return false, err
	}

	// get initial matcher from model
	m := e.GetModel()["m"]["m"].Value
	// parse []string params to []interface{} with matcher
	parsedParams, m := s.parseParams(params, m)

	// Authorization
	res, err := e.EnforceWithMatcher(m, parsedParams...)
	if err != nil {
		return false, err
	}

	return res, nil
}

// LoadPolicy loads policies in an enforcer from its associated adapter
func (s *service) LoadPolicy(ctx context.Context, enforcerHandler EnforcerHandler) error {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return err
	}

	return e.LoadPolicy()
}

// SavePolicy save changed policies in an enforcer to its associated adapter
func (s *service) SavePolicy(ctx context.Context, enforcerHandler EnforcerHandler) error {
	e, err := s.getEnforcer(enforcerHandler)
	if err != nil {
		return err
	}

	return e.SavePolicy()
}
