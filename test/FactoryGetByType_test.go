package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/chrisehlen/knex"
)

var _ = Describe("Factory", func() {

	Describe("get an implementation, that does not require, by type", func() {

		var (
			impl TypeWithNoRequires
			err  error
		)

		Context("when an implementation has been registered", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithNoRequiresOneImpl))
				impl, err = factory.GetByType(new(TypeWithNoRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				Ω(impl).ShouldNot(BeNil())
				Ω(impl).Should(BeEquivalentTo(new(TypeWithNoRequiresOneImpl)))
			})
		})

		Context("when an implementation has not been registered", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				impl, err = factory.GetByType(new(TypeWithNoRequires))
			})

			It("should return a 'Undeclared resource' error", func() {
				Ω(err.Error()).Should(HavePrefix("Undeclared resource "))
			})

			It("should not return an implementation", func() {
				Ω(impl).Should(BeNil())
			})
		})

		Context("when an implementation does not have a injector", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithNoInjectorImpl))
				impl, err = factory.GetByType(new(TypeWithNoRequires))
			})

			It("should return a 'missing injector' error", func() {
				Ω(err.Error()).Should(ContainSubstring("missing injector"))
			})

			It("should not return an implementation", func() {
				Ω(impl).Should(BeNil())
			})
		})

		Context("when an implementation's injector fails", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithErrorInjectorImpl))
				impl, err = factory.GetByType(new(TypeWithNoRequires))
			})

			It("should fail", func() {
				Ω(err).Should(HaveOccurred())
			})

			It("should not return an implementation", func() {
				Ω(impl).Should(BeNil())
			})
		})

		Context("when multiple implementations of the same type have been registered", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithNoRequiresOneImpl))
				factory.Register(new(TypeWithNoRequiresTwoImpl))
				impl, err = factory.GetByType(new(TypeWithNoRequires))
			})

			It("should return a 'Multiple implementations for type' error", func() {
				Ω(err.Error()).Should(HavePrefix("Multiple implementations for type "))
			})

			It("should not return an implementation", func() {
				Ω(impl).Should(BeNil())
			})
		})
	})

	Describe("get an implementation, that does require, by type", func() {

		var (
			impl TypeWithRequires
			err  error
		)

		Context("when the required value is injected successfully", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithNoRequiresOneImpl))
				factory.Register(new(TypeWithRequiresImpl))
				impl, err = factory.GetByType(new(TypeWithRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				compareValue, _ := NewTypeWithRequiresImpl(new(TypeWithNoRequiresOneImpl))
				Ω(impl).ShouldNot(BeNil())
				Ω(impl).Should(BeEquivalentTo(compareValue))
			})

			It("should inject the approprite type(s)", func() {
				value := impl.(*TypeWithRequiresImpl).InjectedType
				Ω(value).ShouldNot(BeNil())
				Ω(value).Should(BeEquivalentTo(new(TypeWithNoRequiresOneImpl)))
			})
		})

		Context("when the required field, is not a slice, and there are multiple implementations", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithNoRequiresOneImpl))
				factory.Register(new(TypeWithNoRequiresTwoImpl))
				factory.Register(new(TypeWithRequiresImpl))
				impl, err = factory.GetByType(new(TypeWithRequires))
			})

			It("should return a 'Multiple implementations for type' error", func() {
				Ω(err.Error()).Should(HavePrefix("Multiple implementations for type "))
			})

			It("should not return an implementation", func() {
				Ω(impl).Should(BeNil())
			})
		})

		Context("when the required value is not injected successfully", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithErrorInjectorImpl))
				factory.Register(new(TypeWithRequiresImpl))
				impl, err = factory.GetByType(new(TypeWithRequires))
			})

			It("should fail", func() {
				Ω(err).Should(HaveOccurred())
			})

			It("should not return an implementation", func() {
				Ω(impl).Should(BeNil())
			})
		})

		Context("when there is a circular dependency", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithCircularDependencyImpl))
				impl, err = factory.GetByType(new(TypeWithCircularDependency))
			})

			It("should return a 'Circular dependency' error", func() {
				Ω(err.Error()).Should(HavePrefix("Circular dependency "))
			})

			It("should not return an implementation", func() {
				Ω(impl).Should(BeNil())
			})
		})
	})

	Describe("get an implementation, that does require, but the require is optional", func() {

		var (
			impl TypeWithNoRequires
			err  error
		)

		Context("when the require dependency has been registered", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithNoRequiresOneImpl))
				factory.Register(new(TypeWithOptionalRequiresImpl))
				impl, err = factory.GetByType(new(TypeWithRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				compareValue, _ := NewTypeWithOptionalRequiresImpl(new(TypeWithNoRequiresOneImpl))
				Ω(impl).ShouldNot(BeNil())
				Ω(impl).Should(BeEquivalentTo(compareValue))
			})

			It("should inject the approprite type(s)", func() {
				value := impl.(*TypeWithOptionalRequiresImpl).InjectedType
				Ω(value).ShouldNot(BeNil())
				Ω(value).Should(BeEquivalentTo(new(TypeWithNoRequiresOneImpl)))
			})
		})

		Context("when the require dependency has not been registered", func() {
			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithOptionalRequiresImpl))
				impl, err = factory.GetByType(new(TypeWithRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				compareValue, _ := NewTypeWithOptionalRequiresImpl(nil)
				Ω(impl).ShouldNot(BeNil())
				Ω(impl).Should(BeEquivalentTo(compareValue))
			})

			It("should inject nil", func() {
				value := impl.(*TypeWithOptionalRequiresImpl).InjectedType
				Ω(value).Should(BeNil())
			})
		})
	})

	Describe("provider is registered", func() {

		var (
			impl TypeWithNoRequires
			err  error
		)

		Context("when getting a provider", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type: new(TypeWithNoRequires),
					Instance: func() (interface{}, error) {
						return &TypeWithNoRequiresOneImpl{}, nil
					},
				})
				impl, err = factory.GetByType(new(TypeWithNoRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				Ω(impl).ShouldNot(BeNil())
				Ω(impl).Should(BeEquivalentTo(new(TypeWithNoRequiresOneImpl)))
			})
		})

		Context("when the provider is injected, by type", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type: new(TypeWithNoRequires),
					Instance: func() (interface{}, error) {
						return &TypeWithNoRequiresOneImpl{}, nil
					},
				})
				factory.Register(new(TypeWithRequiresImpl))
				impl, err = factory.GetByType(new(TypeWithRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				compareValue, _ := NewTypeWithRequiresImpl(new(TypeWithNoRequiresOneImpl))
				Ω(impl).ShouldNot(BeNil())
				Ω(impl).Should(BeEquivalentTo(compareValue))
			})

			It("should inject the approprite type(s)", func() {
				value := impl.(*TypeWithRequiresImpl).InjectedType
				Ω(value).ShouldNot(BeNil())
				Ω(value).Should(BeEquivalentTo(new(TypeWithNoRequiresOneImpl)))
			})
		})
	})
})
