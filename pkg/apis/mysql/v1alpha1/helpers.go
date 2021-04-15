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
	"github.com/jkljajic/mysql-operator/pkg/constants"
	"github.com/jkljajic/mysql-operator/pkg/version"
)

const (
	// DefaultVersion is the MySQL version to use if not specified explicitly by user
	DefaultVersion      = "8.0.22"
	defaultMembers      = 3
	defaultBaseServerID = 1000
	defaultLogLevel     = 4
	// maxBaseServerID is the maximum safe value for BaseServerID calculated
	// as max MySQL server_id value - max Replication Group size.
	maxBaseServerID uint32 = 4294967295 - 9
	// MysqlServer is the image to use if no image is specified explicitly by the user.
	MysqlServer = "mysql/mysql-server"
	// ServiceAccountName is default value for service account name
	ServiceAccountName = "mysql-agent"

	PodManagementPolicy = "OrderedReady"

	ClusterName = "Cluster"
)

const (
	// MaxInnoDBClusterMembers is the maximum number of members supported by InnoDB
	// group replication.
	MaxInnoDBClusterMembers = 9

	// ClusterNameMaxLen is the maximum supported length of a
	// Cluster name.
	// See: https://bugs.mysql.com/bug.php?id=90601
	ClusterNameMaxLen = 28
)

// setOperatorVersionLabel sets the specified operator version label on the label map.
func setOperatorVersionLabel(labelMap map[string]string, label string) {
	labelMap[constants.MySQLOperatorVersionLabel] = label
}

// getOperatorVersionLabel get the specified operator version label on the label map.
func getOperatorVersionLabel(labelMap map[string]string) string {
	return labelMap[constants.MySQLOperatorVersionLabel]
}

// EnsureDefaults will ensure that if a user omits any fields in the
// spec that are required, we set some sensible defaults.
// For example a user can choose to omit the version and number of
// members.
func (c *Cluster) EnsureDefaults() *Cluster {
	if c.Spec.Members == 0 {
		c.Spec.Members = defaultMembers
	}

	if c.Spec.BaseServerID == 0 {
		c.Spec.BaseServerID = defaultBaseServerID
	}

	if c.Spec.Version == "" {
		c.Spec.Version = DefaultVersion
	}

	if c.Spec.ServiceAccountName == "" {
		c.Spec.ServiceAccountName = ServiceAccountName
	}
	if c.Spec.PodManagementPolicy == "" {
		c.Spec.PodManagementPolicy = PodManagementPolicy
	}

	if c.Spec.ClusterName == "" {
		c.Spec.ClusterName = ClusterName
	}

	return c
}

// Validate returns an error if a cluster is invalid
func (c *Cluster) Validate() error {
	return validateCluster(c).ToAggregate()
}

// RequiresConfigMount will return true if a user has specified a config map
// for configuring the cluster else false
func (c *Cluster) RequiresConfigMount() bool {
	return c.Spec.Config != nil
}

// RequiresSecret returns true if a secret should be generated
// for a MySQL cluster else false
func (c *Cluster) RequiresSecret() bool {
	return c.Spec.RootPasswordSecret == nil
}

// RequiresServicePrimary returns true if Multimaster is false
// for a MySQL cluster else false
func (c *Cluster) RequiresServicePrimary() bool {
	return !c.Spec.MultiMaster
}

// GetServiceHeadlessName returns name of service based on cluster name
// for a MySQL cluster
func (c *Cluster) GetServiceHeadlessName() string {
	return c.Name + constants.DefaultServiceSuffix
}

// GetServiceName returns name of service based on cluster name
// for a MySQL cluster
func (c *Cluster) GetServiceName() string {
	return c.Name
}

func (c *Cluster) RequiresReplicationGroupName() bool {
	return c.Spec.ReplicationGroupName != ""
}

// RequiresCustomSSLSetup returns true is the user has provided a secret
// that contains CA cert, server cert and server key for group replication
// SSL support
func (c *Cluster) RequiresCustomSSLSetup() bool {
	return c.Spec.SSLSecret != nil
}

// EnsureDefaults can be invoked to ensure the default values are present.
func (b Backup) EnsureDefaults() *Backup {
	buildVersion := version.GetBuildVersion()
	if buildVersion != "" {
		if b.Labels == nil {
			b.Labels = make(map[string]string)
		}
		_, hasKey := b.Labels[constants.MySQLOperatorVersionLabel]
		if !hasKey {
			setOperatorVersionLabel(b.Labels, buildVersion)
		}
	}
	return &b
}

// Validate checks if the resource spec is valid.
func (b Backup) Validate() error {
	return validateBackup(&b).ToAggregate()
}

// EnsureDefaults can be invoked to ensure the default values are present.
func (b BackupSchedule) EnsureDefaults() *BackupSchedule {
	buildVersion := version.GetBuildVersion()
	if buildVersion != "" {
		if b.Labels == nil {
			b.Labels = make(map[string]string)
		}
		_, hasKey := b.Labels[constants.MySQLOperatorVersionLabel]
		if !hasKey {
			setOperatorVersionLabel(b.Labels, buildVersion)
		}
	}
	return &b
}

// Validate checks if the resource spec is valid.
func (b BackupSchedule) Validate() error {
	return validateBackupSchedule(&b).ToAggregate()
}

// EnsureDefaults can be invoked to ensure the default values are present.
func (r Restore) EnsureDefaults() *Restore {
	buildVersion := version.GetBuildVersion()
	if buildVersion != "" {
		if r.Labels == nil {
			r.Labels = make(map[string]string)
		}
		_, hasKey := r.Labels[constants.MySQLOperatorVersionLabel]
		if !hasKey {
			setOperatorVersionLabel(r.Labels, buildVersion)
		}
	}
	return &r
}

// Validate checks if the resource spec is valid.
func (r Restore) Validate() error {
	return validateRestore(&r).ToAggregate()
}
