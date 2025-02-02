// Copyright 2020 Google LLC
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

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeVPNGateways implements ComputeVPNGatewayInterface
type FakeComputeVPNGateways struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computevpngatewaysResource = schema.GroupVersionResource{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Resource: "computevpngateways"}

var computevpngatewaysKind = schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeVPNGateway"}

// Get takes name of the computeVPNGateway, and returns the corresponding computeVPNGateway object, and an error if there is any.
func (c *FakeComputeVPNGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computevpngatewaysResource, c.ns, name), &v1beta1.ComputeVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeVPNGateway), err
}

// List takes label and field selectors, and returns the list of ComputeVPNGateways that match those selectors.
func (c *FakeComputeVPNGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeVPNGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computevpngatewaysResource, computevpngatewaysKind, c.ns, opts), &v1beta1.ComputeVPNGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeVPNGatewayList{ListMeta: obj.(*v1beta1.ComputeVPNGatewayList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeVPNGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeVPNGateways.
func (c *FakeComputeVPNGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computevpngatewaysResource, c.ns, opts))

}

// Create takes the representation of a computeVPNGateway and creates it.  Returns the server's representation of the computeVPNGateway, and an error, if there is any.
func (c *FakeComputeVPNGateways) Create(ctx context.Context, computeVPNGateway *v1beta1.ComputeVPNGateway, opts v1.CreateOptions) (result *v1beta1.ComputeVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computevpngatewaysResource, c.ns, computeVPNGateway), &v1beta1.ComputeVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeVPNGateway), err
}

// Update takes the representation of a computeVPNGateway and updates it. Returns the server's representation of the computeVPNGateway, and an error, if there is any.
func (c *FakeComputeVPNGateways) Update(ctx context.Context, computeVPNGateway *v1beta1.ComputeVPNGateway, opts v1.UpdateOptions) (result *v1beta1.ComputeVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computevpngatewaysResource, c.ns, computeVPNGateway), &v1beta1.ComputeVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeVPNGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeVPNGateways) UpdateStatus(ctx context.Context, computeVPNGateway *v1beta1.ComputeVPNGateway, opts v1.UpdateOptions) (*v1beta1.ComputeVPNGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computevpngatewaysResource, "status", c.ns, computeVPNGateway), &v1beta1.ComputeVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeVPNGateway), err
}

// Delete takes name of the computeVPNGateway and deletes it. Returns an error if one occurs.
func (c *FakeComputeVPNGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computevpngatewaysResource, c.ns, name, opts), &v1beta1.ComputeVPNGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeVPNGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computevpngatewaysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeVPNGatewayList{})
	return err
}

// Patch applies the patch and returns the patched computeVPNGateway.
func (c *FakeComputeVPNGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computevpngatewaysResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeVPNGateway), err
}
