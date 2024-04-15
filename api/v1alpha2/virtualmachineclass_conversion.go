// Copyright (c) 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/vmware-tanzu/vm-operator/api/v1alpha3"
)

// ConvertTo converts this VirtualMachineClass to the Hub version.
func (src *VirtualMachineClass) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha3.VirtualMachineClass)
	return Convert_v1alpha2_VirtualMachineClass_To_v1alpha3_VirtualMachineClass(src, dst, nil)
}

// ConvertFrom converts the hub version to this VirtualMachineClass.
func (dst *VirtualMachineClass) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha3.VirtualMachineClass)
	return Convert_v1alpha3_VirtualMachineClass_To_v1alpha2_VirtualMachineClass(src, dst, nil)
}

// ConvertTo converts this VirtualMachineClassList to the Hub version.
func (src *VirtualMachineClassList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha3.VirtualMachineClassList)
	return Convert_v1alpha2_VirtualMachineClassList_To_v1alpha3_VirtualMachineClassList(src, dst, nil)
}

// ConvertFrom converts the hub version to this VirtualMachineClassList.
func (dst *VirtualMachineClassList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha3.VirtualMachineClassList)
	return Convert_v1alpha3_VirtualMachineClassList_To_v1alpha2_VirtualMachineClassList(src, dst, nil)
}
