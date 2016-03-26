package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/chrisehlen/knex"
)

var _ = Describe("Factory", func() {

	var (
		implOne TypeWithNoRequires
		implTwo TypeWithNoRequires
		errOne  error
		errTwo  error
	)

	Describe("registers an implementaion within factory scope", func() {

		Context("when there are no requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithFactoryScopeImpl))
				implOne, errOne = factory.GetByType(new(TypeWithNoRequires))
				implTwo, errTwo = factory.GetByType(new(TypeWithNoRequires))
				implTwo.(*TypeWithFactoryScopeImpl).Value = "Updated value"
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
				factory.Register(new(TypeWithFactoryScopeImpl))
				factory.Register(new(TypeWithRequiresImpl))
				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implTwo, errTwo = factory.GetByType(new(typeWithRequires))
				implTwo.(*TypeWithRequiresImpl).InjectedType.(*TypeWithFactoryScopeImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
				Ω(errTwo).Should(Succeed())
			})

			It("should return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implTwo).ShouldNot(BeNil())
				Ω(implOne.(*TypeWithRequiresImpl).InjectedType).Should(Equal(implTwo.(*TypeWithRequiresImpl).InjectedType))
			})
		})
	})

	Describe("provider is registered within factory scope", func() {

		Context("when there are no requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type:  new(TypeWithNoRequires),
					Scope: "factory",
					Instance: func() (interface{}, error) {
						return &TypeWithValueImpl{Value: "Initial value"}, nil
					},
				})
				implOne, errOne = factory.GetByType(new(TypeWithNoRequires))
				implTwo, errTwo = factory.GetByType(new(TypeWithNoRequires))
				implTwo.(*TypeWithValueImpl).Value = "Updated value"
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
				factory.Register(new(TypeWithRequiresImpl))
				factory.RegisterProvider(knex.Provider{
					Type:  new(TypeWithNoRequires),
					Scope: "factory",
					Instance: func() (interface{}, error) {
						return &TypeWithValueImpl{Value: "Initial value"}, nil
					},
				})

				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implTwo, errTwo = factory.GetByType(new(typeWithRequires))
				implTwo.(*TypeWithRequiresImpl).InjectedType.(*TypeWithValueImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
				Ω(errTwo).Should(Succeed())
			})

			It("should return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implTwo).ShouldNot(BeNil())
				Ω(implOne.(*TypeWithRequiresImpl).InjectedType).Should(Equal(implTwo.(*TypeWithRequiresImpl).InjectedType))
			})
		})
	})
})
