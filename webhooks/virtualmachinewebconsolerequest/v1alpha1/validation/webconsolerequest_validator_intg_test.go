// Copyright (c) 2022-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package validation_test

import (
	"crypto/rsa"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	vmopv1a1 "github.com/vmware-tanzu/vm-operator/api/v1alpha1"
	"github.com/vmware-tanzu/vm-operator/pkg/constants/testlabels"
	"github.com/vmware-tanzu/vm-operator/test/builder"
)

func intgTests() {
	Describe(
		"Create",
		Label(
			testlabels.Create,
			testlabels.EnvTest,
			testlabels.V1Alpha1,
			testlabels.Validation,
			testlabels.Webhook,
		),
		intgTestsValidateCreate,
	)
	Describe(
		"Update",
		Label(
			testlabels.Update,
			testlabels.EnvTest,
			testlabels.V1Alpha1,
			testlabels.Validation,
			testlabels.Webhook,
		),
		intgTestsValidateUpdate,
	)
	Describe(
		"Delete",
		Label(
			testlabels.Delete,
			testlabels.EnvTest,
			testlabels.V1Alpha1,
			testlabels.Validation,
			testlabels.Webhook,
		),
		intgTestsValidateDelete,
	)
}

type intgValidatingWebhookContext struct {
	builder.IntegrationTestContext
	wcr        *vmopv1a1.WebConsoleRequest
	privateKey *rsa.PrivateKey
}

func newIntgValidatingWebhookContext() *intgValidatingWebhookContext {
	privateKey, publicKeyPem := builder.WebConsoleRequestKeyPair()

	ctx := &intgValidatingWebhookContext{
		IntegrationTestContext: *suite.NewIntegrationTestContext(),
	}

	ctx.wcr = builder.DummyWebConsoleRequest(ctx.Namespace, "some-name", "some-vm-name", publicKeyPem)
	ctx.privateKey = privateKey
	return ctx
}

func intgTestsValidateCreate() {
	var (
		err error
		ctx *intgValidatingWebhookContext
	)
	BeforeEach(func() {
		ctx = newIntgValidatingWebhookContext()
	})
	AfterEach(func() {
		err = nil
		ctx = nil
	})

	When("create is performed", func() {
		BeforeEach(func() {
			err = ctx.Client.Create(ctx, ctx.wcr)
		})
		It("should allow the request", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})
}

func intgTestsValidateUpdate() {
	var (
		err error
		ctx *intgValidatingWebhookContext
	)

	BeforeEach(func() {
		ctx = newIntgValidatingWebhookContext()
		err = ctx.Client.Create(ctx, ctx.wcr)
		Expect(err).ToNot(HaveOccurred())
	})
	JustBeforeEach(func() {
		err = ctx.Client.Update(suite, ctx.wcr)
	})
	AfterEach(func() {

		err = nil
		ctx = nil
	})

	When("update is performed with changed vm name", func() {
		BeforeEach(func() {
			ctx.wcr.Spec.VirtualMachineName = "alternate-vm-name"
		})
		It("should deny the request", func() {
			Expect(err).To(HaveOccurred())
		})
	})
}

func intgTestsValidateDelete() {
	var (
		err error
		ctx *intgValidatingWebhookContext
	)

	BeforeEach(func() {
		ctx = newIntgValidatingWebhookContext()
		err = ctx.Client.Create(ctx, ctx.wcr)
		Expect(err).ToNot(HaveOccurred())
	})
	JustBeforeEach(func() {
		err = ctx.Client.Delete(suite, ctx.wcr)
	})
	AfterEach(func() {

		err = nil
		ctx = nil
	})

	When("delete is performed", func() {
		It("should allow the request", func() {
			Expect(ctx.Namespace).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})
	})
}
