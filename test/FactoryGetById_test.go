package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/chrisehlen/knex"
)

var _ = Describe("Factory", func() {

	Describe("get an implementation, that does not have requires, by id", func() {

		var (
			impl TypeWithNoRequires
			err  error
		)

		Context("when an implementation has been registered", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithIDImpl))
				impl, err = factory.GetById("testId")
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				Ω(impl).ShouldNot(BeNil())
				Ω(impl).Should(BeEquivalentTo(new(typeWithIDImpl)))
			})
		})

		Context("when an implementation has not been registered", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithNoRequiresOneImpl))
				impl, err = factory.GetById("testId")
			})

			It("should return a 'Undeclared resource' error", func() {
				Ω(err.Error()).Should(HavePrefix("Undeclared resource "))
			})

			It("should not return an implementation", func() {
				Ω(impl).Should(BeNil())
			})
		})
	})

	Describe("get an implementation, that requires a reference by id", func() {

		Context("when the reference id has been registered", func() {

			var (
				impl TypeWithNoRequires
				err  error
			)

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithIDImpl))
				factory.Register(new(TypeWithRequiresWithIdImpl))
				impl, err = factory.GetByType(new(TypeWithRequiresWithId))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				compareValue, _ := NewTypeWithRequiresWithIdImpl(new(typeWithIDImpl))
				Ω(impl).ShouldNot(BeNil())
				Ω(impl).Should(BeEquivalentTo(compareValue))
			})

			It("should inject the approprite type(s)", func() {
				value := impl.(*TypeWithRequiresWithIdImpl).InjectedType
				Ω(value).ShouldNot(BeNil())
				Ω(value).Should(BeEquivalentTo(new(typeWithIDImpl)))
			})
		})

		Context("when the reference id has not been registered", func() {

			var (
				impl TypeWithNoRequires
				err  error
			)

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithRequiresWithIdImpl))
				impl, err = factory.GetByType(new(TypeWithRequiresWithId))
			})

			It("should return a 'Undeclared resource' error", func() {
				Ω(err.Error()).Should(HavePrefix("Undeclared resource "))
			})

			It("should not return an implementation", func() {
				Ω(impl).Should(BeNil())
			})
		})
	})

	Describe("get a provider, that is required", func() {

		var (
			impl TypeWithRequires
			err  error
		)

		Context("when a provider is registered with an id", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type: new(TypeWithNoRequires),
					Id:   "testId",
					Instance: func() (interface{}, error) {
						return &TypeWithNoRequiresOneImpl{}, nil
					},
				})
				impl, err = factory.GetById("testId")
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				Ω(impl).ShouldNot(BeNil())
				Ω(impl).Should(BeEquivalentTo(new(TypeWithNoRequiresOneImpl)))
			})
		})

		Context("when the provider is injected, by id", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.RegisterProvider(knex.Provider{
					Type: new(TypeWithNoRequires),
					Id:   "testId",
					Instance: func() (interface{}, error) {
						return &TypeWithNoRequiresOneImpl{}, nil
					},
				})
				factory.Register(new(TypeWithRequiresWithIdImpl))
				impl, err = factory.GetByType(new(TypeWithRequiresWithId))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				compareValue, _ := NewTypeWithRequiresWithIdImpl(new(TypeWithNoRequiresOneImpl))
				Ω(impl).ShouldNot(BeNil())
				Ω(impl).Should(BeEquivalentTo(compareValue))
			})

			It("should inject the approprite type(s)", func() {
				value := impl.(*TypeWithRequiresWithIdImpl).InjectedType
				Ω(value).ShouldNot(BeNil())
				Ω(value).Should(BeEquivalentTo(new(TypeWithNoRequiresOneImpl)))
			})
		})
	})
})
