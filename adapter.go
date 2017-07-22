// Copyright 2017 The casbin Authors. All Rights Reserved.
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

package protobufadapter

import (
	"strings"

	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/util"
	"github.com/golang/protobuf/proto"
)

// Adapter represents the Protocol Buffers adapter for policy persistence.
// It can load policy from protobuf bytes or save policy to protobuf bytes.
type Adapter struct {
	source *[]byte
	policy *Policy
}

// NewAdapter is the constructor for Adapter.
func NewAdapter(source *[]byte) *Adapter {
	a := Adapter{}
	a.source = source
	a.policy = &Policy{}
	return &a
}

func (a *Adapter) saveToBuffer() error {
	data, err := proto.Marshal(a.policy)
	if err == nil {
		*a.source = data
	}
	return err
}

func (a *Adapter) loadFromBuffer() error {
	policy := &Policy{}
	err := proto.Unmarshal(*a.source, policy)
	if err == nil {
		a.policy = policy
	}
	return err
}

func loadPolicyLine(line string, model model.Model) {
	if line == "" {
		return
	}

	tokens := strings.Split(line, ", ")

	key := tokens[0]
	sec := key[:1]
	model[sec][key].Policy = append(model[sec][key].Policy, tokens[1:])
}

// LoadPolicy loads policy from protobuf bytes.
func (a *Adapter) LoadPolicy(model model.Model) error {
	err := a.loadFromBuffer()
	if err != nil {
		return err
	}

	for _, line := range a.policy.Rules {
		loadPolicyLine(line, model)
	}
	return nil
}

// SavePolicy saves policy to protobuf bytes.
func (a *Adapter) SavePolicy(model model.Model) error {
	a.policy.Reset()

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			tmp := ptype + ", " + util.ArrayToString(rule)
			a.policy.Rules = append(a.policy.Rules, tmp)
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			tmp := ptype + ", " + util.ArrayToString(rule)
			a.policy.Rules = append(a.policy.Rules, tmp)
		}
	}

	err := a.saveToBuffer()
	return err
}
