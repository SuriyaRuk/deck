// +build !ignore_autogenerated

/*
Copyright 2021 Kong Inc.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package konnect

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlane) DeepCopyInto(out *ControlPlane) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ControlPlaneType)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlane.
func (in *ControlPlane) DeepCopy() *ControlPlane {
	if in == nil {
		return nil
	}
	out := new(ControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneServiceRelation) DeepCopyInto(out *ControlPlaneServiceRelation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ControlPlaneEntityID != nil {
		in, out := &in.ControlPlaneEntityID, &out.ControlPlaneEntityID
		*out = new(string)
		**out = **in
	}
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(ControlPlane)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneServiceRelation.
func (in *ControlPlaneServiceRelation) DeepCopy() *ControlPlaneServiceRelation {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneServiceRelation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneType) DeepCopyInto(out *ControlPlaneType) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneType.
func (in *ControlPlaneType) DeepCopy() *ControlPlaneType {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePackage) DeepCopyInto(out *ServicePackage) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]ServiceVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePackage.
func (in *ServicePackage) DeepCopy() *ServicePackage {
	if in == nil {
		return nil
	}
	out := new(ServicePackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceVersion) DeepCopyInto(out *ServiceVersion) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.ServicePackage != nil {
		in, out := &in.ServicePackage, &out.ServicePackage
		*out = new(ServicePackage)
		(*in).DeepCopyInto(*out)
	}
	if in.ControlPlaneServiceRelation != nil {
		in, out := &in.ControlPlaneServiceRelation, &out.ControlPlaneServiceRelation
		*out = new(ControlPlaneServiceRelation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceVersion.
func (in *ServiceVersion) DeepCopy() *ServiceVersion {
	if in == nil {
		return nil
	}
	out := new(ServiceVersion)
	in.DeepCopyInto(out)
	return out
}
