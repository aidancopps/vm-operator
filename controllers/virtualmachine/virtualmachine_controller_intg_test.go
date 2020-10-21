// +build integration

// Copyright (c) 2019-2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package virtualmachine

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"

	"github.com/vmware-tanzu/vm-operator/pkg/vmprovider"
	"github.com/vmware-tanzu/vm-operator/pkg/vmprovider/providers/vsphere"
	"github.com/vmware-tanzu/vm-operator/test/integration"
)

func TestVirtualMachineOldIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "VirtualMachine Suite", []Reporter{printer.NewlineReporter{}})
}

var (
	cfg           *rest.Config
	vcSim         *integration.VcSimInstance
	testEnv       *envtest.Environment
	vSphereConfig *vsphere.VSphereVmProviderConfig
	session       *vsphere.Session
	vmProvider    vmprovider.VirtualMachineProviderInterface
)

var _ = BeforeSuite(func() {
	testEnv, vSphereConfig, cfg, vcSim, session, vmProvider = integration.SetupIntegrationEnv([]string{integration.DefaultNamespace})
})

var _ = AfterSuite(func() {
	integration.TeardownIntegrationEnv(testEnv, vcSim)
})