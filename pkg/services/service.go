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
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"github.com/go-kit/kit/log"
)

type service struct {
	enforcerMap map[EnforcerHandler]*casbin.Enforcer
	adapterMap  map[AdapterHandler]persist.Adapter
}

// NewServer returns a basic struct that implements Server interface
func NewServer() Server {
	s := service{}

	s.enforcerMap = map[EnforcerHandler]*casbin.Enforcer{}
	s.adapterMap = map[AdapterHandler]persist.Adapter{}

	return &s
}

// New returns go-kit service Server
func New(logger log.Logger) Server {
	s := NewServer()
	s = LoggingMiddleware(logger)(s)
	return s
}
