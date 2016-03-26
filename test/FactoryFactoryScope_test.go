package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/chrisehlen/knex"
)

var _ = Describe("Factory", func() {

	var (
		implOne typeWithNoRequires
		implTwo typeWithNoRequires
		errOne  error
		errTwo  error
	)

	Describe("registers an implementaion within factory scope", func() {

		Context("when there are no requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithFactoryScopeImpl))
				implOne, errOne = factory.GetByType(new(typeWithNoRequires))
				implTwo, errTwo = factory.GetByType(new(typeWithNoRequires))
				implTwo.(*typeWithFactoryScopeImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
				Ω(errTwo).Should(Succeed())
			})

			It("should return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implTwo).ShouldNot(BeNil())
				Ω(implOne).Should(Equal(implTwo))
			})
		})

		Context("when there are requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithFactoryScopeImpl))
				factory.Register(new(typeWithRequiresImpl))
				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implTwo, errTwo = factory.GetByType(new(typeWithRequires))
				implTwo.(*typeWithRequiresImpl).InjectedType.(*typeWithFactoryScopeImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
				Ω(errTwo).Should(Succeed())
			})

			It("should return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implTwo).ShouldNot(BeNil())
				Ω(implOne.(*typeWithRequiresImpl).InjectedType).Should(Equal(implTwo.(*typeWithRequiresImpl).InjectedType))
			})
		})
	})

	Describe("provider is registered within factory scope", func() {

		Context("when there are no requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type:  new(typeWithNoRequires),
					Scope: "factory",
					Instance: func() (interface{}, error) {
						return &typeWithValueImpl{Value: "Initial value"}, nil
					},
				})
				implOne, errOne = factory.GetByType(new(typeWithNoRequires))
				implTwo, errTwo = factory.GetByType(new(typeWithNoRequires))
				implTwo.(*typeWithValueImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
				Ω(errTwo).Should(Succeed())
			})

			It("should return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implTwo).ShouldNot(BeNil())
				Ω(implOne).Should(Equal(implTwo))
			})
		})

		Context("when there are requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithRequiresImpl))
				factory.RegisterProvider(knex.Provider{
					Type:  new(typeWithNoRequires),
					Scope: "factory",
					Instance: func() (interface{}, error) {
						return &typeWithValueImpl{Value: "Initial value"}, nil
					},
				})

				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implTwo, errTwo = factory.GetByType(new(typeWithRequires))
				implTwo.(*typeWithRequiresImpl).InjectedType.(*typeWithValueImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
				Ω(errTwo).Should(Succeed())
			})

			It("should return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implTwo).ShouldNot(BeNil())
				Ω(implOne.(*typeWithRequiresImpl).InjectedType).Should(Equal(implTwo.(*typeWithRequiresImpl).InjectedType))
			})
		})
	})
})
