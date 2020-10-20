// Copyright 2018 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/jkljajic/mysql-operator/pkg/apis/mysql/v1alpha1"
	scheme "github.com/jkljajic/mysql-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackupSchedulesGetter has a method to return a BackupScheduleInterface.
// A group's client should implement this interface.
type BackupSchedulesGetter interface {
	BackupSchedules(namespace string) BackupScheduleInterface
}

// BackupScheduleInterface has methods to work with BackupSchedule resources.
type BackupScheduleInterface interface {
	Create(*v1alpha1.BackupSchedule) (*v1alpha1.BackupSchedule, error)
	Update(*v1alpha1.BackupSchedule) (*v1alpha1.BackupSchedule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BackupSchedule, error)
	List(opts v1.ListOptions) (*v1alpha1.BackupScheduleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BackupSchedule, err error)
	BackupScheduleExpansion
}

// backupSchedules implements BackupScheduleInterface
type backupSchedules struct {
	client rest.Interface
	ns     string
}

// newBackupSchedules returns a BackupSchedules
func newBackupSchedules(c *MySQLV1alpha1Client, namespace string) *backupSchedules {
	return &backupSchedules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backupSchedule, and returns the corresponding backupSchedule object, and an error if there is any.
func (c *backupSchedules) Get(name string, options v1.GetOptions) (result *v1alpha1.BackupSchedule, err error) {
	result = &v1alpha1.BackupSchedule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlbackupschedules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupSchedules that match those selectors.
func (c *backupSchedules) List(opts v1.ListOptions) (result *v1alpha1.BackupScheduleList, err error) {
	result = &v1alpha1.BackupScheduleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlbackupschedules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupSchedules.
func (c *backupSchedules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqlbackupschedules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}

// Create takes the representation of a backupSchedule and creates it.  Returns the server's representation of the backupSchedule, and an error, if there is any.
func (c *backupSchedules) Create(backupSchedule *v1alpha1.BackupSchedule) (result *v1alpha1.BackupSchedule, err error) {
	result = &v1alpha1.BackupSchedule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqlbackupschedules").
		Body(backupSchedule).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a backupSchedule and updates it. Returns the server's representation of the backupSchedule, and an error, if there is any.
func (c *backupSchedules) Update(backupSchedule *v1alpha1.BackupSchedule) (result *v1alpha1.BackupSchedule, err error) {
	result = &v1alpha1.BackupSchedule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlbackupschedules").
		Name(backupSchedule.Name).
		Body(backupSchedule).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the backupSchedule and deletes it. Returns an error if one occurs.
func (c *backupSchedules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlbackupschedules").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupSchedules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlbackupschedules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched backupSchedule.
func (c *backupSchedules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BackupSchedule, err error) {
	result = &v1alpha1.BackupSchedule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqlbackupschedules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}
