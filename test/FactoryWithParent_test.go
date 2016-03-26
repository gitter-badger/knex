package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/chrisehlen/knex"
)

var _ = Describe("Factory", func() {

	Describe("has a parent", func() {

		var (
			impl typeWithNoRequires
			err  error
		)

		Describe("type is not registered with child but is registered with the parent", func() {

			Context("when getting by type", func() {
				BeforeEach(func() {
					parent := knex.NewFactory()
					parent.Register(new(typeWithNoRequiresOneImpl))
					child := knex.NewFactory()
					child.AddParent(parent)
					impl, err = child.GetByType(new(typeWithNoRequires))
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
					parent.Register(new(typeWithNoRequiresTwoImpl))
					child := knex.NewFactory()
					child.AddParent(parent)
					impl, err = child.GetAllOfType(new(typeWithNoRequires))
				})

				It("should be successful", func() {
					Ω(err).Should(Succeed())
				})

				It("should return the parents' implementation", func() {
					Ω(impl).ShouldNot(BeNil())
					Ω(impl).Should(BeEquivalentTo([]typeWithNoRequires{new(typeWithNoRequiresOneImpl), new(typeWithNoRequiresTwoImpl)}))
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
					child.Register(new(typeWithRequiresImpl))
					impl, err = child.GetByType(new(typeWithRequires))
				})

				It("should be successful", func() {
					Ω(err).Should(Succeed())
				})

				It("should return the parents' implementation", func() {
					Ω(impl.(*typeWithRequiresImpl).InjectedType).ShouldNot(BeNil())
				})
			})

			Context("when getting field all of type", func() {
				BeforeEach(func() {
					parent := knex.NewFactory()
					parent.Register(new(typeWithNoRequiresOneImpl))
					parent.Register(new(typeWithNoRequiresTwoImpl))
					child := knex.NewFactory()
					child.AddParent(parent)
					child.Register(new(typeWithSliceRequiresImpl))
					impl, err = child.GetByType(new(typeWithRequires))
				})

				It("should be successful", func() {
					Ω(err).Should(Succeed())
				})

				It("should return the parents' implementation", func() {
					Ω(impl).ShouldNot(BeNil())
					Ω(impl.(*typeWithSliceRequiresImpl).InjectedType).Should(BeEquivalentTo([]typeWithNoRequires{new(typeWithNoRequiresOneImpl), new(typeWithNoRequiresTwoImpl)}))
				})
			})

			Context("when getting field by id", func() {
				BeforeEach(func() {
					parent := knex.NewFactory()
					parent.Register(new(typeWithIDImpl))
					child := knex.NewFactory()
					child.AddParent(parent)
					child.Register(new(typeWithRequiresWithIDImpl))
					impl, err = child.GetByType(new(typeWithRequiresWithID))
				})

				It("should be successful", func() {
					Ω(err).Should(Succeed())
				})

				It("should return the parents' implementation", func() {
					Ω(impl.(*typeWithRequiresWithIDImpl).InjectedType).ShouldNot(BeNil())
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
