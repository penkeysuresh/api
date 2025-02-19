// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	scheme "github.com/tigera/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PacketCapturesGetter has a method to return a PacketCaptureInterface.
// A group's client should implement this interface.
type PacketCapturesGetter interface {
	PacketCaptures(namespace string) PacketCaptureInterface
}

// PacketCaptureInterface has methods to work with PacketCapture resources.
type PacketCaptureInterface interface {
	Create(ctx context.Context, packetCapture *v3.PacketCapture, opts v1.CreateOptions) (*v3.PacketCapture, error)
	Update(ctx context.Context, packetCapture *v3.PacketCapture, opts v1.UpdateOptions) (*v3.PacketCapture, error)
	UpdateStatus(ctx context.Context, packetCapture *v3.PacketCapture, opts v1.UpdateOptions) (*v3.PacketCapture, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.PacketCapture, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.PacketCaptureList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.PacketCapture, err error)
	PacketCaptureExpansion
}

// packetCaptures implements PacketCaptureInterface
type packetCaptures struct {
	client rest.Interface
	ns     string
}

// newPacketCaptures returns a PacketCaptures
func newPacketCaptures(c *ProjectcalicoV3Client, namespace string) *packetCaptures {
	return &packetCaptures{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the packetCapture, and returns the corresponding packetCapture object, and an error if there is any.
func (c *packetCaptures) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.PacketCapture, err error) {
	result = &v3.PacketCapture{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("packetcaptures").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PacketCaptures that match those selectors.
func (c *packetCaptures) List(ctx context.Context, opts v1.ListOptions) (result *v3.PacketCaptureList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.PacketCaptureList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("packetcaptures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested packetCaptures.
func (c *packetCaptures) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("packetcaptures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a packetCapture and creates it.  Returns the server's representation of the packetCapture, and an error, if there is any.
func (c *packetCaptures) Create(ctx context.Context, packetCapture *v3.PacketCapture, opts v1.CreateOptions) (result *v3.PacketCapture, err error) {
	result = &v3.PacketCapture{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("packetcaptures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(packetCapture).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a packetCapture and updates it. Returns the server's representation of the packetCapture, and an error, if there is any.
func (c *packetCaptures) Update(ctx context.Context, packetCapture *v3.PacketCapture, opts v1.UpdateOptions) (result *v3.PacketCapture, err error) {
	result = &v3.PacketCapture{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("packetcaptures").
		Name(packetCapture.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(packetCapture).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *packetCaptures) UpdateStatus(ctx context.Context, packetCapture *v3.PacketCapture, opts v1.UpdateOptions) (result *v3.PacketCapture, err error) {
	result = &v3.PacketCapture{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("packetcaptures").
		Name(packetCapture.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(packetCapture).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the packetCapture and deletes it. Returns an error if one occurs.
func (c *packetCaptures) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("packetcaptures").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *packetCaptures) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("packetcaptures").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched packetCapture.
func (c *packetCaptures) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.PacketCapture, err error) {
	result = &v3.PacketCapture{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("packetcaptures").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
