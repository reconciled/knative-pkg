//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Knative Authors

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

package testing

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InlinedPtrStruct) DeepCopyInto(out *InlinedPtrStruct) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InlinedPtrStruct.
func (in *InlinedPtrStruct) DeepCopy() *InlinedPtrStruct {
	if in == nil {
		return nil
	}
	out := new(InlinedPtrStruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InlinedStruct) DeepCopyInto(out *InlinedStruct) {
	*out = *in
	if in.InlinedPtrStruct != nil {
		in, out := &in.InlinedPtrStruct, &out.InlinedPtrStruct
		*out = new(InlinedPtrStruct)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InlinedStruct.
func (in *InlinedStruct) DeepCopy() *InlinedStruct {
	if in == nil {
		return nil
	}
	out := new(InlinedStruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InnerDefaultResource) DeepCopyInto(out *InnerDefaultResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InnerDefaultResource.
func (in *InnerDefaultResource) DeepCopy() *InnerDefaultResource {
	if in == nil {
		return nil
	}
	out := new(InnerDefaultResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InnerDefaultResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InnerDefaultSpec) DeepCopyInto(out *InnerDefaultSpec) {
	*out = *in
	if in.SubFields != nil {
		in, out := &in.SubFields, &out.SubFields
		*out = new(InnerDefaultSubSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InnerDefaultSpec.
func (in *InnerDefaultSpec) DeepCopy() *InnerDefaultSpec {
	if in == nil {
		return nil
	}
	out := new(InnerDefaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InnerDefaultStatus) DeepCopyInto(out *InnerDefaultStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InnerDefaultStatus.
func (in *InnerDefaultStatus) DeepCopy() *InnerDefaultStatus {
	if in == nil {
		return nil
	}
	out := new(InnerDefaultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InnerDefaultStruct) DeepCopyInto(out *InnerDefaultStruct) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InnerDefaultStruct.
func (in *InnerDefaultStruct) DeepCopy() *InnerDefaultStruct {
	if in == nil {
		return nil
	}
	out := new(InnerDefaultStruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InnerDefaultSubSpec) DeepCopyInto(out *InnerDefaultSubSpec) {
	*out = *in
	if in.DeprecatedStringPtr != nil {
		in, out := &in.DeprecatedStringPtr, &out.DeprecatedStringPtr
		*out = new(string)
		**out = **in
	}
	if in.DeprecatedIntPtr != nil {
		in, out := &in.DeprecatedIntPtr, &out.DeprecatedIntPtr
		*out = new(int64)
		**out = **in
	}
	if in.DeprecatedMap != nil {
		in, out := &in.DeprecatedMap, &out.DeprecatedMap
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeprecatedSlice != nil {
		in, out := &in.DeprecatedSlice, &out.DeprecatedSlice
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.DeprecatedStruct = in.DeprecatedStruct
	if in.DeprecatedStructPtr != nil {
		in, out := &in.DeprecatedStructPtr, &out.DeprecatedStructPtr
		*out = new(InnerDefaultStruct)
		**out = **in
	}
	in.InlinedStruct.DeepCopyInto(&out.InlinedStruct)
	if in.InlinedPtrStruct != nil {
		in, out := &in.InlinedPtrStruct, &out.InlinedPtrStruct
		*out = new(InlinedPtrStruct)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InnerDefaultSubSpec.
func (in *InnerDefaultSubSpec) DeepCopy() *InnerDefaultSubSpec {
	if in == nil {
		return nil
	}
	out := new(InnerDefaultSubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnContext) DeepCopyInto(out *OnContext) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnContext.
func (in *OnContext) DeepCopy() *OnContext {
	if in == nil {
		return nil
	}
	out := new(OnContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Resource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceList) DeepCopyInto(out *ResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Resource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceList.
func (in *ResourceList) DeepCopy() *ResourceList {
	if in == nil {
		return nil
	}
	out := new(ResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
	return out
}
