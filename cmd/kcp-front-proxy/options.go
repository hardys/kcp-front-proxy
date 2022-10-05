/*
Copyright 2022 The KCP Authors.

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

package main

import (
	"github.com/spf13/pflag"

	"k8s.io/component-base/config"
	"k8s.io/component-base/logs"

	proxyoptions "github.com/kcp-dev/kcp/pkg/proxy/options"
)

type Options struct {
	Proxy proxyoptions.Options
	Logs  *logs.Options
}

func NewOptions() *Options {
	o := &Options{
		Proxy: *proxyoptions.NewOptions(),
		Logs:  logs.NewOptions(),
	}

	// Default to -v=2
	o.Logs.Config.Verbosity = config.VerbosityLevel(2)
	return o
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	o.Proxy.AddFlags(fs)
	o.Logs.AddFlags(fs)
}

func (o *Options) Complete() error {
	if err := o.Proxy.Complete(); err != nil {
		return err
	}
	return nil
}

func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.Proxy.Validate()...)

	return errs
}
