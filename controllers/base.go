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

package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/CloudNativeSDWAN/cnwan-operator/internal/types"
	sr "github.com/CloudNativeSDWAN/cnwan-operator/pkg/servregistry"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ktypes "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// BaseReconciler is the base controller/reconciler upon which all other
// controllers will be based.
type BaseReconciler struct {
	client.Client
	Log                logr.Logger
	Scheme             *runtime.Scheme
	ServRegBroker      *sr.Broker
	AllowedAnnotations map[string]bool
	CurrentNsPolicy    types.ListPolicy
}

// shouldWatchNs returns true a namespace should be watched according to the
// provided labels and the list policy currently implemented.
func (b *BaseReconciler) shouldWatchNs(labels map[string]string) (watch bool) {
	switch b.CurrentNsPolicy {
	case types.AllowList:
		if _, exists := labels[types.AllowedKey]; exists {
			watch = true
		}
	case types.BlockList:
		if _, exists := labels[types.BlockedKey]; !exists {
			watch = true
		}
	}

	return
}

func (b *BaseReconciler) shouldWatchSrv(srv *corev1.Service) bool {
	nsrv := ktypes.NamespacedName{Namespace: srv.Namespace, Name: srv.Name}
	l := b.Log.WithValues("service", nsrv)
	if srv.Spec.Type != corev1.ServiceTypeLoadBalancer {
		return false
	}

	if len(srv.Status.LoadBalancer.Ingress) == 0 {
		return false
	}

	filteredAnnotations := b.filterAnnotations(srv.Annotations)
	if len(filteredAnnotations) == 0 {
		return false
	}

	var ns corev1.Namespace
	if err := b.Get(context.Background(), ktypes.NamespacedName{Name: srv.Namespace}, &ns); err != nil {
		l.Error(err, "error while getting parent namespace from service")
		return false
	}

	return b.shouldWatchNs(ns.Labels)
}

// filterAnnotations takes a map of annotations and returnes a new one
// stripped from the ones that should not be registered on the service
// registry.
func (b *BaseReconciler) filterAnnotations(annotations map[string]string) map[string]string {
	if len(annotations) == 0 {
		return map[string]string{}
	}
	if len(b.AllowedAnnotations) == 0 {
		return map[string]string{}
	}

	if _, exists := b.AllowedAnnotations["*/*"]; exists {
		return annotations
	}

	filtered := map[string]string{}
	for key, val := range annotations {

		// Check this key specifically
		if _, exists := b.AllowedAnnotations[key]; exists {
			filtered[key] = val
			continue
		}

		prefixName := strings.Split(key, "/")
		if len(prefixName) != 2 {
			// This key is not in prefix/name format
			continue
		}

		prefixWildcard := fmt.Sprintf("%s/*", prefixName[0])
		if _, exists := b.AllowedAnnotations[prefixWildcard]; exists {
			filtered[key] = val
			continue
		}

		wildcardName := fmt.Sprintf("*/%s", prefixName[1])
		if _, exists := b.AllowedAnnotations[wildcardName]; exists {
			filtered[key] = val
		}
	}

	return filtered
}