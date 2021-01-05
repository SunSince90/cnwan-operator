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

// GetEndp returns the endpoint, if it exists.
func (e *etcdServReg) GetEndp(nsName, servName, endpName string) (*sr.Endpoint, error) {
	key, err := KeyFromServiceRegistryObject(&sr.Endpoint{
		NsName: nsName, ServName: servName, Name: endpName,
	})
	if err != nil {
		return nil, err
	}

	ctx, canc := context.WithTimeout(e.mainCtx, defaultTimeout)
	defer canc()

	endp, err := e.getOne(ctx, key)
	if err != nil {
		return nil, err
	}

	return endp.(*sr.Endpoint), nil
}

// ListServ returns a list of services inside the provided namespace.
func (e *etcdServReg) ListEndp(nsName, servName string) (endpList []*sr.Endpoint, err error) {
	// TODO: implement me
	return nil, nil
}

// CreateEndp creates the endpoint.
func (e *etcdServReg) CreateEndp(endp *sr.Endpoint) (*sr.Endpoint, error) {
	// TODO: implement me
	return nil, nil
}

// UpdateEndp updates the endpoint.
func (e *etcdServReg) UpdateEndp(endp *sr.Endpoint) (*sr.Endpoint, error) {
	// TODO: implement me
	return nil, nil
}

// DeleteEndp deletes the endpoint.
func (e *etcdServReg) DeleteEndp(nsName, servName, endpName string) error {
	// TODO: implement me
	return nil
}