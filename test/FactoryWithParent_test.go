package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/chrisehlen/knex"
)

var _ = Describe("Factory", func() {

	Describe("has a parent", func() {

		var (
			impl TypeWithNoRequires
			err  error
		)

		Describe("type is not registered with child but is registered with the parent", func() {

			Context("when getting by type", func() {
				BeforeEach(func() {
					parent := knex.NewFactory()
					parent.Register(new(typeWithNoRequiresOneImpl))
					child := knex.NewFactory()
					child.AddParent(parent)
					impl, err = child.GetByType(new(TypeWithNoRequires))
				})

				It("should be successful", func() {
					Ω(err).Should(Succeed())
				})

				It("should return the parents' implementation", func() {
					Ω(impl).ShouldNot(BeNil())
				})
			})

			Context("when getting all of type", func() {
				BeforeEach(func() {
					parent := knex.NewFactory()
					parent.Register(new(typeWithNoRequiresOneImpl))
					parent.Register(new(TypeWithNoRequiresTwoImpl))
					child := knex.NewFactory()
					child.AddParent(parent)
					impl, err = child.GetAllOfType(new(TypeWithNoRequires))
				})

				It("should be successful", func() {
					Ω(err).Should(Succeed())
				})

				It("should return the parents' implementation", func() {
					Ω(impl).ShouldNot(BeNil())
					Ω(impl).Should(BeEquivalentTo([]TypeWithNoRequires{new(typeWithNoRequiresOneImpl), new(TypeWithNoRequiresTwoImpl)}))
				})
			})

			Context("when getting by id", func() {
				BeforeEach(func() {
					parent := knex.NewFactory()
					parent.Register(new(typeWithIDImpl))
					child := knex.NewFactory()
					child.AddParent(parent)
					impl, err = child.GetByID("testId")
				})

				It("should be successful", func() {
					Ω(err).Should(Succeed())
				})

				It("should return the parents' implementation", func() {
					Ω(impl).ShouldNot(BeNil())
				})
			})

			Context("when getting field by type", func() {
				BeforeEach(func() {
					parent := knex.NewFactory()
					parent.Register(new(typeWithNoRequiresOneImpl))
					child := knex.NewFactory()
					child.AddParent(parent)
					child.Register(new(TypeWithRequiresImpl))
					impl, err = child.GetByType(new(typeWithRequires))
				})

				It("should be successful", func() {
					Ω(err).Should(Succeed())
				})

				It("should return the parents' implementation", func() {
					Ω(impl.(*TypeWithRequiresImpl).InjectedType).ShouldNot(BeNil())
				})
			})

			Context("when getting field all of type", func() {
				BeforeEach(func() {
					parent := knex.NewFactory()
					parent.Register(new(typeWithNoRequiresOneImpl))
					parent.Register(new(TypeWithNoRequiresTwoImpl))
					child := knex.NewFactory()
					child.AddParent(parent)
					child.Register(new(TypeWithSliceRequiresImpl))
					impl, err = child.GetByType(new(typeWithRequires))
				})

				It("should be successful", func() {
					Ω(err).Should(Succeed())
				})

				It("should return the parents' implementation", func() {
					Ω(impl).ShouldNot(BeNil())
					Ω(impl.(*TypeWithSliceRequiresImpl).InjectedType).Should(BeEquivalentTo([]TypeWithNoRequires{new(typeWithNoRequiresOneImpl), new(TypeWithNoRequiresTwoImpl)}))
				})
			})

			Context("when getting field by id", func() {
				BeforeEach(func() {
					parent := knex.NewFactory()
					parent.Register(new(typeWithIDImpl))
					child := knex.NewFactory()
					child.AddParent(parent)
					child.Register(new(TypeWithRequiresWithIdImpl))
					impl, err = child.GetByType(new(TypeWithRequiresWithId))
				})

				It("should be successful", func() {
					Ω(err).Should(Succeed())
				})

				It("should return the parents' implementation", func() {
					Ω(impl.(*TypeWithRequiresWithIdImpl).InjectedType).ShouldNot(BeNil())
				})
			})
		})
	})

	Describe("adding a parent factory", func() {

		Context("causes a circular dependency", func() {

			It("should return a circular dependency error", func() {
				parent := knex.NewFactory()
				child := knex.NewFactory()
				_ = child.AddParent(parent)
				err := parent.AddParent(child)
				Ω(err.Error()).Should(HavePrefix("Circular dependency "))
			})
		})
	})
})
