// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2_test

import (
	"testing"

	. "github.com/onsi/ginkgo"

	ctrlmgr "sigs.k8s.io/controller-runtime/pkg/manager"

	virtualmachinesetresourcepolicy "github.com/vmware-tanzu/vm-operator/controllers/virtualmachinesetresourcepolicy/v1alpha2"
	ctrlContext "github.com/vmware-tanzu/vm-operator/pkg/context"
	providerfake "github.com/vmware-tanzu/vm-operator/pkg/vmprovider/fake"
	"github.com/vmware-tanzu/vm-operator/test/builder"
)

var intgFakeVMProvider = providerfake.NewVMProviderA2()

var suite = builder.NewTestSuiteForController(
	virtualmachinesetresourcepolicy.AddToManager,
	func(ctx *ctrlContext.ControllerManagerContext, _ ctrlmgr.Manager) error {
		ctx.VMProviderA2 = intgFakeVMProvider
		return nil
	},
)

func TestVirtualMachineSetResourcePolicy(t *testing.T) {
	_ = intgTests
	suite.Register(t, "VirtualMachineSetResourcePolicy controller suite", nil /*intgTests*/, unitTests)
}

var _ = BeforeSuite(suite.BeforeSuite)

var _ = AfterSuite(suite.AfterSuite)