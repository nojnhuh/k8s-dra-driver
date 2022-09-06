//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatableDevices) DeepCopyInto(out *AllocatableDevices) {
	*out = *in
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(AllocatableGpu)
		**out = **in
	}
	if in.Mig != nil {
		in, out := &in.Mig, &out.Mig
		*out = new(AllocatableMigDevices)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatableDevices.
func (in *AllocatableDevices) DeepCopy() *AllocatableDevices {
	if in == nil {
		return nil
	}
	out := new(AllocatableDevices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatableGpu) DeepCopyInto(out *AllocatableGpu) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatableGpu.
func (in *AllocatableGpu) DeepCopy() *AllocatableGpu {
	if in == nil {
		return nil
	}
	out := new(AllocatableGpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatableMigDevices) DeepCopyInto(out *AllocatableMigDevices) {
	*out = *in
	if in.Placements != nil {
		in, out := &in.Placements, &out.Placements
		*out = make([]MigDevicePlacement, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatableMigDevices.
func (in *AllocatableMigDevices) DeepCopy() *AllocatableMigDevices {
	if in == nil {
		return nil
	}
	out := new(AllocatableMigDevices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedDevice) DeepCopyInto(out *AllocatedDevice) {
	*out = *in
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(AllocatedGpu)
		**out = **in
	}
	if in.Mig != nil {
		in, out := &in.Mig, &out.Mig
		*out = new(AllocatedMigDevice)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedDevice.
func (in *AllocatedDevice) DeepCopy() *AllocatedDevice {
	if in == nil {
		return nil
	}
	out := new(AllocatedDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AllocatedDevices) DeepCopyInto(out *AllocatedDevices) {
	{
		in := &in
		*out = make(AllocatedDevices, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedDevices.
func (in AllocatedDevices) DeepCopy() AllocatedDevices {
	if in == nil {
		return nil
	}
	out := new(AllocatedDevices)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedGpu) DeepCopyInto(out *AllocatedGpu) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedGpu.
func (in *AllocatedGpu) DeepCopy() *AllocatedGpu {
	if in == nil {
		return nil
	}
	out := new(AllocatedGpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedMigDevice) DeepCopyInto(out *AllocatedMigDevice) {
	*out = *in
	out.Placement = in.Placement
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedMigDevice.
func (in *AllocatedMigDevice) DeepCopy() *AllocatedMigDevice {
	if in == nil {
		return nil
	}
	out := new(AllocatedMigDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeInstanceClaimParameters) DeepCopyInto(out *ComputeInstanceClaimParameters) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeInstanceClaimParameters.
func (in *ComputeInstanceClaimParameters) DeepCopy() *ComputeInstanceClaimParameters {
	if in == nil {
		return nil
	}
	out := new(ComputeInstanceClaimParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeInstanceClaimParameters) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeInstanceClaimParametersList) DeepCopyInto(out *ComputeInstanceClaimParametersList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComputeInstanceClaimParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeInstanceClaimParametersList.
func (in *ComputeInstanceClaimParametersList) DeepCopy() *ComputeInstanceClaimParametersList {
	if in == nil {
		return nil
	}
	out := new(ComputeInstanceClaimParametersList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeInstanceClaimParametersList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeInstanceClaimParametersSpec) DeepCopyInto(out *ComputeInstanceClaimParametersSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeInstanceClaimParametersSpec.
func (in *ComputeInstanceClaimParametersSpec) DeepCopy() *ComputeInstanceClaimParametersSpec {
	if in == nil {
		return nil
	}
	out := new(ComputeInstanceClaimParametersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceClassParameters) DeepCopyInto(out *DeviceClassParameters) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceClassParameters.
func (in *DeviceClassParameters) DeepCopy() *DeviceClassParameters {
	if in == nil {
		return nil
	}
	out := new(DeviceClassParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceClassParameters) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceClassParametersList) DeepCopyInto(out *DeviceClassParametersList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeviceClassParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceClassParametersList.
func (in *DeviceClassParametersList) DeepCopy() *DeviceClassParametersList {
	if in == nil {
		return nil
	}
	out := new(DeviceClassParametersList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceClassParametersList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceClassParametersSpec) DeepCopyInto(out *DeviceClassParametersSpec) {
	*out = *in
	if in.DeviceSelector != nil {
		in, out := &in.DeviceSelector, &out.DeviceSelector
		*out = make([]DeviceSelector, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceClassParametersSpec.
func (in *DeviceClassParametersSpec) DeepCopy() *DeviceClassParametersSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceClassParametersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSelector) DeepCopyInto(out *DeviceSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSelector.
func (in *DeviceSelector) DeepCopy() *DeviceSelector {
	if in == nil {
		return nil
	}
	out := new(DeviceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuClaimParameters) DeepCopyInto(out *GpuClaimParameters) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuClaimParameters.
func (in *GpuClaimParameters) DeepCopy() *GpuClaimParameters {
	if in == nil {
		return nil
	}
	out := new(GpuClaimParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GpuClaimParameters) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuClaimParametersList) DeepCopyInto(out *GpuClaimParametersList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GpuClaimParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuClaimParametersList.
func (in *GpuClaimParametersList) DeepCopy() *GpuClaimParametersList {
	if in == nil {
		return nil
	}
	out := new(GpuClaimParametersList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GpuClaimParametersList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuClaimParametersSpec) DeepCopyInto(out *GpuClaimParametersSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuClaimParametersSpec.
func (in *GpuClaimParametersSpec) DeepCopy() *GpuClaimParametersSpec {
	if in == nil {
		return nil
	}
	out := new(GpuClaimParametersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDeviceClaimParameters) DeepCopyInto(out *MigDeviceClaimParameters) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDeviceClaimParameters.
func (in *MigDeviceClaimParameters) DeepCopy() *MigDeviceClaimParameters {
	if in == nil {
		return nil
	}
	out := new(MigDeviceClaimParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigDeviceClaimParameters) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDeviceClaimParametersList) DeepCopyInto(out *MigDeviceClaimParametersList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigDeviceClaimParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDeviceClaimParametersList.
func (in *MigDeviceClaimParametersList) DeepCopy() *MigDeviceClaimParametersList {
	if in == nil {
		return nil
	}
	out := new(MigDeviceClaimParametersList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigDeviceClaimParametersList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDeviceClaimParametersSpec) DeepCopyInto(out *MigDeviceClaimParametersSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDeviceClaimParametersSpec.
func (in *MigDeviceClaimParametersSpec) DeepCopy() *MigDeviceClaimParametersSpec {
	if in == nil {
		return nil
	}
	out := new(MigDeviceClaimParametersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDevicePlacement) DeepCopyInto(out *MigDevicePlacement) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDevicePlacement.
func (in *MigDevicePlacement) DeepCopy() *MigDevicePlacement {
	if in == nil {
		return nil
	}
	out := new(MigDevicePlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationState) DeepCopyInto(out *NodeAllocationState) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationState.
func (in *NodeAllocationState) DeepCopy() *NodeAllocationState {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeAllocationState) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationStateList) DeepCopyInto(out *NodeAllocationStateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeAllocationState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationStateList.
func (in *NodeAllocationStateList) DeepCopy() *NodeAllocationStateList {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationStateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeAllocationStateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationStateSpec) DeepCopyInto(out *NodeAllocationStateSpec) {
	*out = *in
	if in.AllocatableDevices != nil {
		in, out := &in.AllocatableDevices, &out.AllocatableDevices
		*out = make([]AllocatableDevices, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClaimRequests != nil {
		in, out := &in.ClaimRequests, &out.ClaimRequests
		*out = make(map[string]RequestedDevices, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ClaimAllocations != nil {
		in, out := &in.ClaimAllocations, &out.ClaimAllocations
		*out = make(map[string]AllocatedDevices, len(*in))
		for key, val := range *in {
			var outVal []AllocatedDevice
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(AllocatedDevices, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationStateSpec.
func (in *NodeAllocationStateSpec) DeepCopy() *NodeAllocationStateSpec {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationStateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedDevices) DeepCopyInto(out *RequestedDevices) {
	*out = *in
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(RequestedGpus)
		(*in).DeepCopyInto(*out)
	}
	if in.Mig != nil {
		in, out := &in.Mig, &out.Mig
		*out = new(RequestedMigDevices)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedDevices.
func (in *RequestedDevices) DeepCopy() *RequestedDevices {
	if in == nil {
		return nil
	}
	out := new(RequestedDevices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedGpu) DeepCopyInto(out *RequestedGpu) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedGpu.
func (in *RequestedGpu) DeepCopy() *RequestedGpu {
	if in == nil {
		return nil
	}
	out := new(RequestedGpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedGpus) DeepCopyInto(out *RequestedGpus) {
	*out = *in
	out.Spec = in.Spec
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]RequestedGpu, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedGpus.
func (in *RequestedGpus) DeepCopy() *RequestedGpus {
	if in == nil {
		return nil
	}
	out := new(RequestedGpus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedMigDevice) DeepCopyInto(out *RequestedMigDevice) {
	*out = *in
	out.Placement = in.Placement
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedMigDevice.
func (in *RequestedMigDevice) DeepCopy() *RequestedMigDevice {
	if in == nil {
		return nil
	}
	out := new(RequestedMigDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedMigDevices) DeepCopyInto(out *RequestedMigDevices) {
	*out = *in
	out.Spec = in.Spec
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]RequestedMigDevice, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedMigDevices.
func (in *RequestedMigDevices) DeepCopy() *RequestedMigDevices {
	if in == nil {
		return nil
	}
	out := new(RequestedMigDevices)
	in.DeepCopyInto(out)
	return out
}
