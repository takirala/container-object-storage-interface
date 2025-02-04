/*
Copyright 2024 The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/container-object-storage-interface-api/client/apis/objectstorage/v1alpha1"
)

// FakeBucketAccessClasses implements BucketAccessClassInterface
type FakeBucketAccessClasses struct {
	Fake *FakeObjectstorageV1alpha1
}

var bucketaccessclassesResource = v1alpha1.SchemeGroupVersion.WithResource("bucketaccessclasses")

var bucketaccessclassesKind = v1alpha1.SchemeGroupVersion.WithKind("BucketAccessClass")

// Get takes name of the bucketAccessClass, and returns the corresponding bucketAccessClass object, and an error if there is any.
func (c *FakeBucketAccessClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BucketAccessClass, err error) {
	emptyResult := &v1alpha1.BucketAccessClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(bucketaccessclassesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.BucketAccessClass), err
}

// List takes label and field selectors, and returns the list of BucketAccessClasses that match those selectors.
func (c *FakeBucketAccessClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BucketAccessClassList, err error) {
	emptyResult := &v1alpha1.BucketAccessClassList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(bucketaccessclassesResource, bucketaccessclassesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BucketAccessClassList{ListMeta: obj.(*v1alpha1.BucketAccessClassList).ListMeta}
	for _, item := range obj.(*v1alpha1.BucketAccessClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bucketAccessClasses.
func (c *FakeBucketAccessClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(bucketaccessclassesResource, opts))
}

// Create takes the representation of a bucketAccessClass and creates it.  Returns the server's representation of the bucketAccessClass, and an error, if there is any.
func (c *FakeBucketAccessClasses) Create(ctx context.Context, bucketAccessClass *v1alpha1.BucketAccessClass, opts v1.CreateOptions) (result *v1alpha1.BucketAccessClass, err error) {
	emptyResult := &v1alpha1.BucketAccessClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(bucketaccessclassesResource, bucketAccessClass, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.BucketAccessClass), err
}

// Update takes the representation of a bucketAccessClass and updates it. Returns the server's representation of the bucketAccessClass, and an error, if there is any.
func (c *FakeBucketAccessClasses) Update(ctx context.Context, bucketAccessClass *v1alpha1.BucketAccessClass, opts v1.UpdateOptions) (result *v1alpha1.BucketAccessClass, err error) {
	emptyResult := &v1alpha1.BucketAccessClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(bucketaccessclassesResource, bucketAccessClass, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.BucketAccessClass), err
}

// Delete takes name of the bucketAccessClass and deletes it. Returns an error if one occurs.
func (c *FakeBucketAccessClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(bucketaccessclassesResource, name, opts), &v1alpha1.BucketAccessClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBucketAccessClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(bucketaccessclassesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BucketAccessClassList{})
	return err
}

// Patch applies the patch and returns the patched bucketAccessClass.
func (c *FakeBucketAccessClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BucketAccessClass, err error) {
	emptyResult := &v1alpha1.BucketAccessClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(bucketaccessclassesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.BucketAccessClass), err
}
