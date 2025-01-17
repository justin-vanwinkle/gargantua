/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	scheme "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProgressesGetter has a method to return a ProgressInterface.
// A group's client should implement this interface.
type ProgressesGetter interface {
	Progresses(namespace string) ProgressInterface
}

// ProgressInterface has methods to work with Progress resources.
type ProgressInterface interface {
	Create(ctx context.Context, progress *v1.Progress, opts metav1.CreateOptions) (*v1.Progress, error)
	Update(ctx context.Context, progress *v1.Progress, opts metav1.UpdateOptions) (*v1.Progress, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Progress, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ProgressList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Progress, err error)
	ProgressExpansion
}

// progresses implements ProgressInterface
type progresses struct {
	client rest.Interface
	ns     string
}

// newProgresses returns a Progresses
func newProgresses(c *HobbyfarmV1Client, namespace string) *progresses {
	return &progresses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the progress, and returns the corresponding progress object, and an error if there is any.
func (c *progresses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Progress, err error) {
	result = &v1.Progress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("progresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Progresses that match those selectors.
func (c *progresses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ProgressList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ProgressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("progresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested progresses.
func (c *progresses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("progresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a progress and creates it.  Returns the server's representation of the progress, and an error, if there is any.
func (c *progresses) Create(ctx context.Context, progress *v1.Progress, opts metav1.CreateOptions) (result *v1.Progress, err error) {
	result = &v1.Progress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("progresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(progress).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a progress and updates it. Returns the server's representation of the progress, and an error, if there is any.
func (c *progresses) Update(ctx context.Context, progress *v1.Progress, opts metav1.UpdateOptions) (result *v1.Progress, err error) {
	result = &v1.Progress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("progresses").
		Name(progress.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(progress).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the progress and deletes it. Returns an error if one occurs.
func (c *progresses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("progresses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *progresses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("progresses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched progress.
func (c *progresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Progress, err error) {
	result = &v1.Progress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("progresses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
