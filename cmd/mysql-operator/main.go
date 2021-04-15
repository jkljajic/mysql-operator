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

package main

import (
	goflag "flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	utilflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog/v2"

	"github.com/jkljajic/mysql-operator/cmd/mysql-operator/app"
	operatoropts "github.com/jkljajic/mysql-operator/pkg/options/operator"
	"github.com/jkljajic/mysql-operator/pkg/version"
)

const (
	configPath      = "/etc/mysql-operator/mysql-operator-config.yaml"
	metricsEndpoint = "0.0.0.0:9183"
)

func main() {
	klog.InitFlags(nil)
	fmt.Fprintf(os.Stderr, "Starting mysql-operator version '%s'\n", version.GetBuildVersion())
	opts, err := operatoropts.NewMySQLOperatorOpts(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading config: %v\n", err)
		os.Exit(1)
	}
	defer klog.Flush()

	opts.AddFlags(pflag.CommandLine)
	pflag.CommandLine.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()
	goflag.CommandLine.Parse([]string{})

	pflag.VisitAll(func(flag *pflag.Flag) {
		klog.V(6).Infof("FLAG: --%s=%q", flag.Name, flag.Value)
	})

	if err := app.Run(opts); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
