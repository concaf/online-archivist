package internalversion

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	scheme "github.com/openshift/origin/pkg/authorization/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RoleBindingsGetter has a method to return a RoleBindingInterface.
// A group's client should implement this interface.
type RoleBindingsGetter interface {
	RoleBindings(namespace string) RoleBindingInterface
}

// RoleBindingInterface has methods to work with RoleBinding resources.
type RoleBindingInterface interface {
	Create(*authorization.RoleBinding) (*authorization.RoleBinding, error)
	Update(*authorization.RoleBinding) (*authorization.RoleBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*authorization.RoleBinding, error)
	List(opts v1.ListOptions) (*authorization.RoleBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authorization.RoleBinding, err error)
	RoleBindingExpansion
}

// roleBindings implements RoleBindingInterface
type roleBindings struct {
	client rest.Interface
	ns     string
}

// newRoleBindings returns a RoleBindings
func newRoleBindings(c *AuthorizationClient, namespace string) *roleBindings {
	return &roleBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a roleBinding and creates it.  Returns the server's representation of the roleBinding, and an error, if there is any.
func (c *roleBindings) Create(roleBinding *authorization.RoleBinding) (result *authorization.RoleBinding, err error) {
	result = &authorization.RoleBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rolebindings").
		Body(roleBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a roleBinding and updates it. Returns the server's representation of the roleBinding, and an error, if there is any.
func (c *roleBindings) Update(roleBinding *authorization.RoleBinding) (result *authorization.RoleBinding, err error) {
	result = &authorization.RoleBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rolebindings").
		Name(roleBinding.Name).
		Body(roleBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the roleBinding and deletes it. Returns an error if one occurs.
func (c *roleBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rolebindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *roleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rolebindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the roleBinding, and returns the corresponding roleBinding object, and an error if there is any.
func (c *roleBindings) Get(name string, options v1.GetOptions) (result *authorization.RoleBinding, err error) {
	result = &authorization.RoleBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rolebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RoleBindings that match those selectors.
func (c *roleBindings) List(opts v1.ListOptions) (result *authorization.RoleBindingList, err error) {
	result = &authorization.RoleBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested roleBindings.
func (c *roleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched roleBinding.
func (c *roleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authorization.RoleBinding, err error) {
	result = &authorization.RoleBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rolebindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
