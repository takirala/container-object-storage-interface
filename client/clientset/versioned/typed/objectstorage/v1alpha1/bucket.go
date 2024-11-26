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

package v1alpha1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	v1alpha1 "sigs.k8s.io/container-object-storage-interface-api/client/apis/objectstorage/v1alpha1"
	scheme "sigs.k8s.io/container-object-storage-interface-api/client/clientset/versioned/scheme"
)

// BucketsGetter has a method to return a BucketInterface.
// A group's client should implement this interface.
type BucketsGetter interface {
	Buckets() BucketInterface
}

// BucketInterface has methods to work with Bucket resources.
type BucketInterface interface {
	Create(ctx context.Context, bucket *v1alpha1.Bucket, opts v1.CreateOptions) (*v1alpha1.Bucket, error)
	Update(ctx context.Context, bucket *v1alpha1.Bucket, opts v1.UpdateOptions) (*v1alpha1.Bucket, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, bucket *v1alpha1.Bucket, opts v1.UpdateOptions) (*v1alpha1.Bucket, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Bucket, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BucketList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Bucket, err error)
	BucketExpansion
}

// buckets implements BucketInterface
type buckets struct {
	*gentype.ClientWithList[*v1alpha1.Bucket, *v1alpha1.BucketList]
}

// newBuckets returns a Buckets
func newBuckets(c *ObjectstorageV1alpha1Client) *buckets {
	return &buckets{
		gentype.NewClientWithList[*v1alpha1.Bucket, *v1alpha1.BucketList](
			"buckets",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.Bucket { return &v1alpha1.Bucket{} },
			func() *v1alpha1.BucketList { return &v1alpha1.BucketList{} }),
	}
}
