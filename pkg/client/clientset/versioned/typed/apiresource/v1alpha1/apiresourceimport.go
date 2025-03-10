/*
Copyright 2021 The Authors

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
	"time"

	v1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apiresource/v1alpha1"
	scheme "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// APIResourceImportsGetter has a method to return a APIResourceImportInterface.
// A group's client should implement this interface.
type APIResourceImportsGetter interface {
	APIResourceImports() APIResourceImportInterface
}

// APIResourceImportInterface has methods to work with APIResourceImport resources.
type APIResourceImportInterface interface {
	Create(ctx context.Context, aPIResourceImport *v1alpha1.APIResourceImport, opts v1.CreateOptions) (*v1alpha1.APIResourceImport, error)
	Update(ctx context.Context, aPIResourceImport *v1alpha1.APIResourceImport, opts v1.UpdateOptions) (*v1alpha1.APIResourceImport, error)
	UpdateStatus(ctx context.Context, aPIResourceImport *v1alpha1.APIResourceImport, opts v1.UpdateOptions) (*v1alpha1.APIResourceImport, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.APIResourceImport, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.APIResourceImportList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIResourceImport, err error)
	APIResourceImportExpansion
}

// aPIResourceImports implements APIResourceImportInterface
type aPIResourceImports struct {
	client rest.Interface
}

// newAPIResourceImports returns a APIResourceImports
func newAPIResourceImports(c *ApiresourceV1alpha1Client) *aPIResourceImports {
	return &aPIResourceImports{
		client: c.RESTClient(),
	}
}

// Get takes name of the aPIResourceImport, and returns the corresponding aPIResourceImport object, and an error if there is any.
func (c *aPIResourceImports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIResourceImport, err error) {
	result = &v1alpha1.APIResourceImport{}
	err = c.client.Get().
		Resource("apiresourceimports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of APIResourceImports that match those selectors.
func (c *aPIResourceImports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIResourceImportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.APIResourceImportList{}
	err = c.client.Get().
		Resource("apiresourceimports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aPIResourceImports.
func (c *aPIResourceImports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apiresourceimports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aPIResourceImport and creates it.  Returns the server's representation of the aPIResourceImport, and an error, if there is any.
func (c *aPIResourceImports) Create(ctx context.Context, aPIResourceImport *v1alpha1.APIResourceImport, opts v1.CreateOptions) (result *v1alpha1.APIResourceImport, err error) {
	result = &v1alpha1.APIResourceImport{}
	err = c.client.Post().
		Resource("apiresourceimports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIResourceImport).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aPIResourceImport and updates it. Returns the server's representation of the aPIResourceImport, and an error, if there is any.
func (c *aPIResourceImports) Update(ctx context.Context, aPIResourceImport *v1alpha1.APIResourceImport, opts v1.UpdateOptions) (result *v1alpha1.APIResourceImport, err error) {
	result = &v1alpha1.APIResourceImport{}
	err = c.client.Put().
		Resource("apiresourceimports").
		Name(aPIResourceImport.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIResourceImport).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aPIResourceImports) UpdateStatus(ctx context.Context, aPIResourceImport *v1alpha1.APIResourceImport, opts v1.UpdateOptions) (result *v1alpha1.APIResourceImport, err error) {
	result = &v1alpha1.APIResourceImport{}
	err = c.client.Put().
		Resource("apiresourceimports").
		Name(aPIResourceImport.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIResourceImport).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aPIResourceImport and deletes it. Returns an error if one occurs.
func (c *aPIResourceImports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apiresourceimports").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aPIResourceImports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apiresourceimports").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aPIResourceImport.
func (c *aPIResourceImports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIResourceImport, err error) {
	result = &v1alpha1.APIResourceImport{}
	err = c.client.Patch(pt).
		Resource("apiresourceimports").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
