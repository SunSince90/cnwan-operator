// Copyright © 2021 Cisco
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// All rights reserved.

package etcd

import (
	"context"

	sr "github.com/CloudNativeSDWAN/cnwan-operator/pkg/servregistry"
)

// GetNs returns the namespace if exists.
func (e *etcdServReg) GetNs(name string) (*sr.Namespace, error) {
	key := KeyFromNames(name)
	if !key.IsValid() {
		return nil, sr.ErrNsNameNotProvided
	}

	ctx, canc := context.WithTimeout(e.mainCtx, defaultTimeout)
	defer canc()

	ns, err := e.getOne(ctx, key)
	if err != nil {
		return nil, err
	}

	return ns.(*sr.Namespace), nil
}

// ListNs returns a list of all namespaces.
func (e *etcdServReg) ListNs() ([]*sr.Namespace, error) {
	// TODO: implement me
	return nil, nil
}

// CreateNs creates the namespace.
func (e *etcdServReg) CreateNs(ns *sr.Namespace) (*sr.Namespace, error) {
	// TODO: implement me
	return nil, nil
}

// UpdateNs updates the namespace.
func (e *etcdServReg) UpdateNs(ns *sr.Namespace) (*sr.Namespace, error) {
	// TODO: implement me
	return nil, nil
}

// DeleteNs deletes the namespace.
func (e *etcdServReg) DeleteNs(name string) error {
	// TODO: implement me
	return nil
}