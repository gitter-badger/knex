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

	Describe("registers an implementaion with no scope", func() {

		Context("when there are no requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithNoScopeImpl))
				implOne, errOne = factory.GetByType(new(typeWithNoRequires))
				implTwo, errTwo = factory.GetByType(new(typeWithNoRequires))
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
				factory.Register(new(typeWithRequiresImpl))
				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implTwo, errTwo = factory.GetByType(new(typeWithRequires))
				implTwo.(*typeWithRequiresImpl).InjectedType.(*typeWithNoScopeImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
				Ω(errTwo).Should(Succeed())
			})

			It("should not return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implTwo).ShouldNot(BeNil())
				Ω(implOne.(*typeWithRequiresImpl).InjectedType).ShouldNot(Equal(implTwo.(*typeWithRequiresImpl).InjectedType))
			})
		})

		Context("when there are muliple requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithNoScopeImpl))
				factory.Register(new(typeWithMultipleRequiresImpl))
				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implOne.(*typeWithMultipleRequiresImpl).InjectedTypeOne.(*typeWithNoScopeImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
			})

			It("should not return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implOne.(*typeWithMultipleRequiresImpl).InjectedTypeOne).ShouldNot(Equal(implOne.(*typeWithMultipleRequiresImpl).InjectedTypeTwo))
			})
		})
	})

	Describe("provider is registered with no scope", func() {

		Context("when there are no requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type: new(typeWithNoRequires),
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

		Context("when there are requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type: new(typeWithNoRequires),
					Instance: func() (interface{}, error) {
						return &typeWithValueImpl{Value: "Initial value"}, nil
					},
				})
				factory.Register(new(typeWithRequiresImpl))
				implOne, errOne = factory.GetByType(new(typeWithRequires))
				implTwo, errTwo = factory.GetByType(new(typeWithRequires))
				implTwo.(*typeWithRequiresImpl).InjectedType.(*typeWithValueImpl).Value = "Updated value"
			})

			It("should be successful", func() {
				Ω(errOne).Should(Succeed())
				Ω(errTwo).Should(Succeed())
			})

			It("should not return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implTwo).ShouldNot(BeNil())
				Ω(implOne.(*typeWithRequiresImpl).InjectedType).ShouldNot(Equal(implTwo.(*typeWithRequiresImpl).InjectedType))
			})
		})

		Context("when there are muliple requires", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type: new(typeWithNoRequires),
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

			It("should not return the same implementations", func() {
				Ω(implOne).ShouldNot(BeNil())
				Ω(implOne.(*typeWithMultipleRequiresImpl).InjectedTypeOne).ShouldNot(Equal(implOne.(*typeWithMultipleRequiresImpl).InjectedTypeTwo))
			})
		})
	})
})
