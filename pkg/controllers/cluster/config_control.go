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

package cluster

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

// ConfigMapControlInterface defines the interface that the
// ClusterController uses to create Configmaps. It is implemented as an
// interface to enable testing.
type ConfigMapControlInterface interface {
	CreateConfigMap(c *v1.ConfigMap) error
}

type realConfigMapControl struct {
	client kubernetes.Interface
}

// NewRealConfigMapControl creates a concrete implementation of the
// ConfigMapControlInterface.
func NewRealConfigMapControl(client kubernetes.Interface) ConfigMapControlInterface {
	return &realConfigMapControl{client: client}
}

func (rsc *realConfigMapControl) CreateConfigMap(c *v1.ConfigMap) error {
	_, err := rsc.client.CoreV1().ConfigMaps(c.Namespace).Create(context.Background(), c, metav1.CreateOptions{})
	return err
}
