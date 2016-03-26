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

	Describe("registers an implementaion within graph scope", func() {

		Context("when making multiple calls to get by type", func() {

			Context("when there are no requires", func() {
				BeforeEach(func() {
					factory := knex.NewFactory()
					factory.Register(new(TypeWithGraphScopeImpl))
					implOne, errOne = factory.GetByType(new(TypeWithNoRequires))
					implTwo, errTwo = factory.GetByType(new(TypeWithNoRequires))
					implTwo.(*TypeWithGraphScopeImpl).Value = "Updated value"
				})

				It("should be successful", func() {
					Ω(errOne).Should(Succeed())
					Ω(errTwo).Should(Succeed())
				})

				It("should not return the same implementations", func() {
					Ω(implOne).ShouldNot(BeNil())
					Ω(implTwo).ShouldNot(BeNil())
					Ω(implOne).ShouldNot(Equal(implTwo))
				})
			})

			Context("when makeing one call with multiple requires", func() {
				BeforeEach(func() {
					factory := knex.NewFactory()
					factory.Register(new(TypeWithGraphScopeImpl))
					factory.Register(new(TypeWithMultipleRequiresImpl))
					implOne, errOne = factory.GetByType(new(typeWithRequires))
					implOne.(*TypeWithMultipleRequiresImpl).InjectedTypeOne.(*TypeWithGraphScopeImpl).Value = "Updated value"
				})

				It("should be successful", func() {
					Ω(errOne).Should(Succeed())
				})

				It("should return the same implementations for each get", func() {
					Ω(implOne).ShouldNot(BeNil())
					Ω(implOne.(*TypeWithMultipleRequiresImpl).InjectedTypeOne).Should(Equal(implOne.(*TypeWithMultipleRequiresImpl).InjectedTypeTwo))
				})
			})
		})
	})

	Describe("provider is registered within graph scope", func() {

		Context("when there are no requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type:  new(TypeWithNoRequires),
					Scope: "graph",
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

			It("should not return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implTwo).ShouldNot(BeNil())
				Ω(implOne).ShouldNot(Equal(implTwo))
			})
		})

		Context("when makeing one call with multiple requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type:  new(TypeWithNoRequires),
					Scope: "graph",
					Instance: func() (interface{}, error) {
						return &TypeWithValueImpl{Value: "Initial value"}, nil
					},
				})
				factory.Register(new(TypeWithMultipleRequiresImpl))
				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implOne.(*TypeWithMultipleRequiresImpl).InjectedTypeOne.(*TypeWithValueImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
			})

			It("should return the same implementations for each get", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implOne.(*TypeWithMultipleRequiresImpl).InjectedTypeOne).Should(Equal(implOne.(*TypeWithMultipleRequiresImpl).InjectedTypeTwo))
			})
		})
	})
})
