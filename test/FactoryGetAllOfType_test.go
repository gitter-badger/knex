package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/chrisehlen/knex"
)

var _ = Describe("Factory", func() {

	Describe("get all implementations, that does not require, by type", func() {
		var (
			allValues interface{}
			implSlice []TypeWithNoRequires
			err       error
		)

		Context("when an implementation has been registered", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithNoRequiresOneImpl))
				allValues, err = factory.GetAllOfType(new(TypeWithNoRequires))
				implSlice = allValues.([]TypeWithNoRequires)
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				Ω(implSlice).ShouldNot(BeNil())
				Ω(implSlice).Should(BeEquivalentTo([]TypeWithNoRequires{new(TypeWithNoRequiresOneImpl)}))
			})
		})

		Context("when multiple implementations have been registered with the same type", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithNoRequiresOneImpl))
				factory.Register(new(TypeWithNoRequiresTwoImpl))
				allValues, err = factory.GetAllOfType(new(TypeWithNoRequires))
				implSlice = allValues.([]TypeWithNoRequires)
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return implementaions of the correct type", func() {
				Ω(implSlice).ShouldNot(BeNil())
				Ω(implSlice).Should(BeEquivalentTo([]TypeWithNoRequires{new(TypeWithNoRequiresOneImpl), new(TypeWithNoRequiresTwoImpl)}))
			})
		})

		Context("when no implementations have been registered with the given type", func() {

			var allvalues interface{}

			BeforeEach(func() {
				factory := knex.NewFactory()
				allvalues, err = factory.GetAllOfType(new(TypeWithNoRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an empty slice", func() {
				Ω(allvalues).Should(BeEquivalentTo([]TypeWithNoRequires{}))
			})
		})

		Context("when one of the required fields is a slice", func() {

			var allvalues interface{}

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithNoRequiresOneImpl))
				factory.Register(new(TypeWithNoRequiresTwoImpl))
				factory.Register(new(TypeWithSliceRequiresImpl))
				allvalues, err = factory.GetByType(new(TypeWithRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return implementaions of the correct type", func() {
				Ω(allvalues).ShouldNot(BeNil())
				impl, _ := NewTypeWithSliceRequiresImpl([]TypeWithNoRequires{new(TypeWithNoRequiresOneImpl), new(TypeWithNoRequiresTwoImpl)})
				Ω(allvalues).Should(BeEquivalentTo(impl))
			})
		})

		Context("when one of the required fields is a slice and no types have been registered for the slice", func() {

			var allvalues interface{}

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithSliceRequiresImpl))
				allvalues, err = factory.GetByType(new(TypeWithRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return implementaion with slice with legth of 0", func() {
				Ω(allvalues).ShouldNot(BeNil())
				impl, _ := NewTypeWithSliceRequiresImpl([]TypeWithNoRequires{})
				Ω(allvalues).Should(BeEquivalentTo(impl))
			})
		})

		Context("when multiple implementations have been registered and one of the implementations fail", func() {

			var allvalues interface{}

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(TypeWithNoRequiresOneImpl))
				factory.Register(new(TypeWithErrorInjectorImpl))
				allvalues, err = factory.GetAllOfType(new(TypeWithNoRequires))
			})

			It("should fail", func() {
				Ω(err).Should(HaveOccurred())
			})

			It("should not return an implementation", func() {
				Ω(allvalues).Should(BeNil())
			})
		})
	})
})
