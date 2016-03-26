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

	Describe("registers an implementaion within graph scope", func() {

		Context("when making multiple calls to get by type", func() {

			Context("when there are no requires", func() {
				BeforeEach(func() {
					factory := knex.NewFactory()
					factory.Register(new(typeWithGraphScopeImpl))
					implOne, errOne = factory.GetByType(new(typeWithNoRequires))
					implTwo, errTwo = factory.GetByType(new(typeWithNoRequires))
					implTwo.(*typeWithGraphScopeImpl).Value = "Updated value"
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
					factory.Register(new(typeWithGraphScopeImpl))
					factory.Register(new(typeWithMultipleRequiresImpl))
					implOne, errOne = factory.GetByType(new(typeWithRequires))
					implOne.(*typeWithMultipleRequiresImpl).InjectedTypeOne.(*typeWithGraphScopeImpl).Value = "Updated value"
				})

				It("should be successful", func() {
					Ω(errOne).Should(Succeed())
				})

				It("should return the same implementations for each get", func() {
					Ω(implOne).ShouldNot(BeNil())
					Ω(implOne.(*typeWithMultipleRequiresImpl).InjectedTypeOne).Should(Equal(implOne.(*typeWithMultipleRequiresImpl).InjectedTypeTwo))
				})
			})
		})
	})

	Describe("provider is registered within graph scope", func() {

		Context("when there are no requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type:  new(typeWithNoRequires),
					Scope: "graph",
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
					Type:  new(typeWithNoRequires),
					Scope: "graph",
					Instance: func() (interface{}, error) {
						return &typeWithValueImpl{Value: "Initial value"}, nil
					},
				})
				factory.Register(new(typeWithMultipleRequiresImpl))
				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implOne.(*typeWithMultipleRequiresImpl).InjectedTypeOne.(*typeWithValueImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
			})

			It("should return the same implementations for each get", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implOne.(*typeWithMultipleRequiresImpl).InjectedTypeOne).Should(Equal(implOne.(*typeWithMultipleRequiresImpl).InjectedTypeTwo))
			})
		})
	})
})
