package fake

import (
	user_v1 "github.com/openshift/api/user/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeUserIdentityMappings implements UserIdentityMappingInterface
type FakeUserIdentityMappings struct {
	Fake *FakeUserV1
}

var useridentitymappingsResource = schema.GroupVersionResource{Group: "user.openshift.io", Version: "v1", Resource: "useridentitymappings"}

var useridentitymappingsKind = schema.GroupVersionKind{Group: "user.openshift.io", Version: "v1", Kind: "UserIdentityMapping"}

// Get takes name of the userIdentityMapping, and returns the corresponding userIdentityMapping object, and an error if there is any.
func (c *FakeUserIdentityMappings) Get(name string, options v1.GetOptions) (result *user_v1.UserIdentityMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(useridentitymappingsResource, name), &user_v1.UserIdentityMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user_v1.UserIdentityMapping), err
}

// Create takes the representation of a userIdentityMapping and creates it.  Returns the server's representation of the userIdentityMapping, and an error, if there is any.
func (c *FakeUserIdentityMappings) Create(userIdentityMapping *user_v1.UserIdentityMapping) (result *user_v1.UserIdentityMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(useridentitymappingsResource, userIdentityMapping), &user_v1.UserIdentityMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user_v1.UserIdentityMapping), err
}

// Update takes the representation of a userIdentityMapping and updates it. Returns the server's representation of the userIdentityMapping, and an error, if there is any.
func (c *FakeUserIdentityMappings) Update(userIdentityMapping *user_v1.UserIdentityMapping) (result *user_v1.UserIdentityMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(useridentitymappingsResource, userIdentityMapping), &user_v1.UserIdentityMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user_v1.UserIdentityMapping), err
}

// Delete takes name of the userIdentityMapping and deletes it. Returns an error if one occurs.
func (c *FakeUserIdentityMappings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(useridentitymappingsResource, name), &user_v1.UserIdentityMapping{})
	return err
}
