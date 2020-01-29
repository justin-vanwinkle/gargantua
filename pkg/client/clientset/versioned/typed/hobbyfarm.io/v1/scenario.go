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
	"time"

	v1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	scheme "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ScenariosGetter has a method to return a ScenarioInterface.
// A group's client should implement this interface.
type ScenariosGetter interface {
	Scenarios() ScenarioInterface
}

// ScenarioInterface has methods to work with Scenario resources.
type ScenarioInterface interface {
	Create(*v1.Scenario) (*v1.Scenario, error)
	Update(*v1.Scenario) (*v1.Scenario, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Scenario, error)
	List(opts metav1.ListOptions) (*v1.ScenarioList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Scenario, err error)
	ScenarioExpansion
}

// scenarios implements ScenarioInterface
type scenarios struct {
	client rest.Interface
}

// newScenarios returns a Scenarios
func newScenarios(c *HobbyfarmV1Client) *scenarios {
	return &scenarios{
		client: c.RESTClient(),
	}
}

// Get takes name of the scenario, and returns the corresponding scenario object, and an error if there is any.
func (c *scenarios) Get(name string, options metav1.GetOptions) (result *v1.Scenario, err error) {
	result = &v1.Scenario{}
	err = c.client.Get().
		Resource("scenarios").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Scenarios that match those selectors.
func (c *scenarios) List(opts metav1.ListOptions) (result *v1.ScenarioList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ScenarioList{}
	err = c.client.Get().
		Resource("scenarios").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested scenarios.
func (c *scenarios) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("scenarios").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a scenario and creates it.  Returns the server's representation of the scenario, and an error, if there is any.
func (c *scenarios) Create(scenario *v1.Scenario) (result *v1.Scenario, err error) {
	result = &v1.Scenario{}
	err = c.client.Post().
		Resource("scenarios").
		Body(scenario).
		Do().
		Into(result)
	return
}

// Update takes the representation of a scenario and updates it. Returns the server's representation of the scenario, and an error, if there is any.
func (c *scenarios) Update(scenario *v1.Scenario) (result *v1.Scenario, err error) {
	result = &v1.Scenario{}
	err = c.client.Put().
		Resource("scenarios").
		Name(scenario.Name).
		Body(scenario).
		Do().
		Into(result)
	return
}

// Delete takes name of the scenario and deletes it. Returns an error if one occurs.
func (c *scenarios) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("scenarios").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *scenarios) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("scenarios").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched scenario.
func (c *scenarios) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Scenario, err error) {
	result = &v1.Scenario{}
	err = c.client.Patch(pt).
		Resource("scenarios").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
