/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package docker10

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerConfig) DeepCopyInto(out *DockerConfig) {
	*out = *in
	if in.PortSpecs != nil {
		in, out := &in.PortSpecs, &out.PortSpecs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExposedPorts != nil {
		in, out := &in.ExposedPorts, &out.ExposedPorts
		*out = make(map[string]struct{}, len(*in))
		for key := range *in {
			(*out)[key] = struct{}{}
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Cmd != nil {
		in, out := &in.Cmd, &out.Cmd
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make(map[string]struct{}, len(*in))
		for key := range *in {
			(*out)[key] = struct{}{}
		}
	}
	if in.Entrypoint != nil {
		in, out := &in.Entrypoint, &out.Entrypoint
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityOpts != nil {
		in, out := &in.SecurityOpts, &out.SecurityOpts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OnBuild != nil {
		in, out := &in.OnBuild, &out.OnBuild
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerConfig.
func (in *DockerConfig) DeepCopy() *DockerConfig {
	if in == nil {
		return nil
	}
	out := new(DockerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerImage) DeepCopyInto(out *DockerImage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Created.DeepCopyInto(&out.Created)
	in.ContainerConfig.DeepCopyInto(&out.ContainerConfig)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		if *in == nil {
			*out = nil
		} else {
			*out = new(DockerConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerImage.
func (in *DockerImage) DeepCopy() *DockerImage {
	if in == nil {
		return nil
	}
	out := new(DockerImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerImage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}