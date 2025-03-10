/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fake

import (
	infrav1 "sigs.k8s.io/cluster-api-provider-vsphere/apis/vmware/v1beta1"
	"sigs.k8s.io/cluster-api-provider-vsphere/pkg/context"
	"sigs.k8s.io/cluster-api-provider-vsphere/pkg/context/vmware"
)

// NewVmwareClusterContext returns a fake ClusterContext for unit testing
// reconcilers with a fake client.
func NewVmwareClusterContext(ctx *context.ControllerContext, namespace string, vsphereCluster *infrav1.VSphereCluster) *vmware.ClusterContext {
	// Create the cluster resources.
	cluster := newClusterV1()
	if vsphereCluster == nil {
		v := NewVSphereCluster(namespace)
		vsphereCluster = &v
	}

	// Add the cluster resources to the fake cluster client.
	if err := ctx.Client.Create(ctx, &cluster); err != nil {
		panic(err)
	}
	if err := ctx.Client.Create(ctx, vsphereCluster); err != nil {
		panic(err)
	}

	return &vmware.ClusterContext{
		ControllerContext: ctx,
		VSphereCluster:    vsphereCluster,
		Logger:            ctx.Logger.WithName(vsphereCluster.Name),
	}
}
