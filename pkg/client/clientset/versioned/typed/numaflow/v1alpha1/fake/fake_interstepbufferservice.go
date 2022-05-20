/*


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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/numaproj/numaflow/pkg/apis/numaflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInterStepBufferServices implements InterStepBufferServiceInterface
type FakeInterStepBufferServices struct {
	Fake *FakeNumaflowV1alpha1
	ns   string
}

var interstepbufferservicesResource = schema.GroupVersionResource{Group: "numaflow.numaproj.io", Version: "v1alpha1", Resource: "interstepbufferservices"}

var interstepbufferservicesKind = schema.GroupVersionKind{Group: "numaflow.numaproj.io", Version: "v1alpha1", Kind: "InterStepBufferService"}

// Get takes name of the interStepBufferService, and returns the corresponding interStepBufferService object, and an error if there is any.
func (c *FakeInterStepBufferServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InterStepBufferService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(interstepbufferservicesResource, c.ns, name), &v1alpha1.InterStepBufferService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InterStepBufferService), err
}

// List takes label and field selectors, and returns the list of InterStepBufferServices that match those selectors.
func (c *FakeInterStepBufferServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InterStepBufferServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(interstepbufferservicesResource, interstepbufferservicesKind, c.ns, opts), &v1alpha1.InterStepBufferServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InterStepBufferServiceList{ListMeta: obj.(*v1alpha1.InterStepBufferServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.InterStepBufferServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested interStepBufferServices.
func (c *FakeInterStepBufferServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(interstepbufferservicesResource, c.ns, opts))

}

// Create takes the representation of a interStepBufferService and creates it.  Returns the server's representation of the interStepBufferService, and an error, if there is any.
func (c *FakeInterStepBufferServices) Create(ctx context.Context, interStepBufferService *v1alpha1.InterStepBufferService, opts v1.CreateOptions) (result *v1alpha1.InterStepBufferService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(interstepbufferservicesResource, c.ns, interStepBufferService), &v1alpha1.InterStepBufferService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InterStepBufferService), err
}

// Update takes the representation of a interStepBufferService and updates it. Returns the server's representation of the interStepBufferService, and an error, if there is any.
func (c *FakeInterStepBufferServices) Update(ctx context.Context, interStepBufferService *v1alpha1.InterStepBufferService, opts v1.UpdateOptions) (result *v1alpha1.InterStepBufferService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(interstepbufferservicesResource, c.ns, interStepBufferService), &v1alpha1.InterStepBufferService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InterStepBufferService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInterStepBufferServices) UpdateStatus(ctx context.Context, interStepBufferService *v1alpha1.InterStepBufferService, opts v1.UpdateOptions) (*v1alpha1.InterStepBufferService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(interstepbufferservicesResource, "status", c.ns, interStepBufferService), &v1alpha1.InterStepBufferService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InterStepBufferService), err
}

// Delete takes name of the interStepBufferService and deletes it. Returns an error if one occurs.
func (c *FakeInterStepBufferServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(interstepbufferservicesResource, c.ns, name, opts), &v1alpha1.InterStepBufferService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInterStepBufferServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(interstepbufferservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InterStepBufferServiceList{})
	return err
}

// Patch applies the patch and returns the patched interStepBufferService.
func (c *FakeInterStepBufferServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InterStepBufferService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(interstepbufferservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.InterStepBufferService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InterStepBufferService), err
}
