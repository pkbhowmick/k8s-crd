/*
Copyright Pulak Kanti Bhowmick(pulak@appscode.com).

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

	v1alpha1 "github.com/pkbhowmick/k8s-crd/pkg/apis/stable.example.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeApis implements KubeApiInterface
type FakeKubeApis struct {
	Fake *FakeStableV1alpha1
	ns   string
}

var kubeapisResource = schema.GroupVersionResource{Group: "stable.example.com", Version: "v1alpha1", Resource: "kubeapis"}

var kubeapisKind = schema.GroupVersionKind{Group: "stable.example.com", Version: "v1alpha1", Kind: "KubeApi"}

// Get takes name of the kubeApi, and returns the corresponding kubeApi object, and an error if there is any.
func (c *FakeKubeApis) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KubeApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubeapisResource, c.ns, name), &v1alpha1.KubeApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeApi), err
}

// List takes label and field selectors, and returns the list of KubeApis that match those selectors.
func (c *FakeKubeApis) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KubeApiList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubeapisResource, kubeapisKind, c.ns, opts), &v1alpha1.KubeApiList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubeApiList{ListMeta: obj.(*v1alpha1.KubeApiList).ListMeta}
	for _, item := range obj.(*v1alpha1.KubeApiList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeApis.
func (c *FakeKubeApis) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubeapisResource, c.ns, opts))

}

// Create takes the representation of a kubeApi and creates it.  Returns the server's representation of the kubeApi, and an error, if there is any.
func (c *FakeKubeApis) Create(ctx context.Context, kubeApi *v1alpha1.KubeApi, opts v1.CreateOptions) (result *v1alpha1.KubeApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubeapisResource, c.ns, kubeApi), &v1alpha1.KubeApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeApi), err
}

// Update takes the representation of a kubeApi and updates it. Returns the server's representation of the kubeApi, and an error, if there is any.
func (c *FakeKubeApis) Update(ctx context.Context, kubeApi *v1alpha1.KubeApi, opts v1.UpdateOptions) (result *v1alpha1.KubeApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubeapisResource, c.ns, kubeApi), &v1alpha1.KubeApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeApi), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeApis) UpdateStatus(ctx context.Context, kubeApi *v1alpha1.KubeApi, opts v1.UpdateOptions) (*v1alpha1.KubeApi, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubeapisResource, "status", c.ns, kubeApi), &v1alpha1.KubeApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeApi), err
}

// Delete takes name of the kubeApi and deletes it. Returns an error if one occurs.
func (c *FakeKubeApis) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubeapisResource, c.ns, name), &v1alpha1.KubeApi{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeApis) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubeapisResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubeApiList{})
	return err
}

// Patch applies the patch and returns the patched kubeApi.
func (c *FakeKubeApis) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KubeApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubeapisResource, c.ns, name, pt, data, subresources...), &v1alpha1.KubeApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeApi), err
}
