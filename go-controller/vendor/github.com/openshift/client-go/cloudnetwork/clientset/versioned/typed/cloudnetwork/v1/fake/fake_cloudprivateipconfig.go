// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	cloudnetworkv1 "github.com/openshift/api/cloudnetwork/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudPrivateIPConfigs implements CloudPrivateIPConfigInterface
type FakeCloudPrivateIPConfigs struct {
	Fake *FakeCloudV1
}

var cloudprivateipconfigsResource = schema.GroupVersionResource{Group: "cloud.network.openshift.io", Version: "v1", Resource: "cloudprivateipconfigs"}

var cloudprivateipconfigsKind = schema.GroupVersionKind{Group: "cloud.network.openshift.io", Version: "v1", Kind: "CloudPrivateIPConfig"}

// Get takes name of the cloudPrivateIPConfig, and returns the corresponding cloudPrivateIPConfig object, and an error if there is any.
func (c *FakeCloudPrivateIPConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *cloudnetworkv1.CloudPrivateIPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cloudprivateipconfigsResource, name), &cloudnetworkv1.CloudPrivateIPConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudnetworkv1.CloudPrivateIPConfig), err
}

// List takes label and field selectors, and returns the list of CloudPrivateIPConfigs that match those selectors.
func (c *FakeCloudPrivateIPConfigs) List(ctx context.Context, opts v1.ListOptions) (result *cloudnetworkv1.CloudPrivateIPConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cloudprivateipconfigsResource, cloudprivateipconfigsKind, opts), &cloudnetworkv1.CloudPrivateIPConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cloudnetworkv1.CloudPrivateIPConfigList{ListMeta: obj.(*cloudnetworkv1.CloudPrivateIPConfigList).ListMeta}
	for _, item := range obj.(*cloudnetworkv1.CloudPrivateIPConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudPrivateIPConfigs.
func (c *FakeCloudPrivateIPConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cloudprivateipconfigsResource, opts))
}

// Create takes the representation of a cloudPrivateIPConfig and creates it.  Returns the server's representation of the cloudPrivateIPConfig, and an error, if there is any.
func (c *FakeCloudPrivateIPConfigs) Create(ctx context.Context, cloudPrivateIPConfig *cloudnetworkv1.CloudPrivateIPConfig, opts v1.CreateOptions) (result *cloudnetworkv1.CloudPrivateIPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cloudprivateipconfigsResource, cloudPrivateIPConfig), &cloudnetworkv1.CloudPrivateIPConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudnetworkv1.CloudPrivateIPConfig), err
}

// Update takes the representation of a cloudPrivateIPConfig and updates it. Returns the server's representation of the cloudPrivateIPConfig, and an error, if there is any.
func (c *FakeCloudPrivateIPConfigs) Update(ctx context.Context, cloudPrivateIPConfig *cloudnetworkv1.CloudPrivateIPConfig, opts v1.UpdateOptions) (result *cloudnetworkv1.CloudPrivateIPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cloudprivateipconfigsResource, cloudPrivateIPConfig), &cloudnetworkv1.CloudPrivateIPConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudnetworkv1.CloudPrivateIPConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudPrivateIPConfigs) UpdateStatus(ctx context.Context, cloudPrivateIPConfig *cloudnetworkv1.CloudPrivateIPConfig, opts v1.UpdateOptions) (*cloudnetworkv1.CloudPrivateIPConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(cloudprivateipconfigsResource, "status", cloudPrivateIPConfig), &cloudnetworkv1.CloudPrivateIPConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudnetworkv1.CloudPrivateIPConfig), err
}

// Delete takes name of the cloudPrivateIPConfig and deletes it. Returns an error if one occurs.
func (c *FakeCloudPrivateIPConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cloudprivateipconfigsResource, name), &cloudnetworkv1.CloudPrivateIPConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudPrivateIPConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cloudprivateipconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &cloudnetworkv1.CloudPrivateIPConfigList{})
	return err
}

// Patch applies the patch and returns the patched cloudPrivateIPConfig.
func (c *FakeCloudPrivateIPConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cloudnetworkv1.CloudPrivateIPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cloudprivateipconfigsResource, name, pt, data, subresources...), &cloudnetworkv1.CloudPrivateIPConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudnetworkv1.CloudPrivateIPConfig), err
}
