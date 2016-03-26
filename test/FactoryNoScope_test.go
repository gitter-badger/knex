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

	Describe("registers an implementaion with no scope", func() {

		Context("when there are no requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithNoScopeImpl))
				implOne, errOne = factory.GetByType(new(TypeWithNoRequires))
				implTwo, errTwo = factory.GetByType(new(TypeWithNoRequires))
				implTwo.(*typeWithNoScopeImpl).Value = "Updated value"
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

		Context("when there are requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithNoScopeImpl))
				factory.Register(new(TypeWithRequiresImpl))
				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implTwo, errTwo = factory.GetByType(new(typeWithRequires))
				implTwo.(*TypeWithRequiresImpl).InjectedType.(*typeWithNoScopeImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
				Ω(errTwo).Should(Succeed())
			})

			It("should not return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implTwo).ShouldNot(BeNil())
				Ω(implOne.(*TypeWithRequiresImpl).InjectedType).ShouldNot(Equal(implTwo.(*TypeWithRequiresImpl).InjectedType))
			})
		})

		Context("when there are muliple requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithNoScopeImpl))
				factory.Register(new(TypeWithMultipleRequiresImpl))
				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implOne.(*TypeWithMultipleRequiresImpl).InjectedTypeOne.(*typeWithNoScopeImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
			})

			It("should not return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implOne.(*TypeWithMultipleRequiresImpl).InjectedTypeOne).ShouldNot(Equal(implOne.(*TypeWithMultipleRequiresImpl).InjectedTypeTwo))
			})
		})
	})

	Describe("provider is registered with no scope", func() {

		Context("when there are no requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type: new(TypeWithNoRequires),
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

		Context("when there are requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type: new(TypeWithNoRequires),
					Instance: func() (interface{}, error) {
						return &TypeWithValueImpl{Value: "Initial value"}, nil
					},
				})
				factory.Register(new(TypeWithRequiresImpl))
				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implTwo, errTwo = factory.GetByType(new(typeWithRequires))
				implTwo.(*TypeWithRequiresImpl).InjectedType.(*TypeWithValueImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
				Ω(errTwo).Should(Succeed())
			})

			It("should not return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implTwo).ShouldNot(BeNil())
				Ω(implOne.(*TypeWithRequiresImpl).InjectedType).ShouldNot(Equal(implTwo.(*TypeWithRequiresImpl).InjectedType))
			})
		})

		Context("when there are muliple requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type: new(TypeWithNoRequires),
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

			It("should not return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implOne.(*TypeWithMultipleRequiresImpl).InjectedTypeOne).ShouldNot(Equal(implOne.(*TypeWithMultipleRequiresImpl).InjectedTypeTwo))
			})
		})
	})
})
