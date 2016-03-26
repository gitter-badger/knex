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
			implSlice []typeWithNoRequires
			err       error
		)

		Context("when an implementation has been registered", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithNoRequiresOneImpl))
				allValues, err = factory.GetAllOfType(new(typeWithNoRequires))
				implSlice = allValues.([]typeWithNoRequires)
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an implementaion of the correct type", func() {
				Ω(implSlice).ShouldNot(BeNil())
				Ω(implSlice).Should(BeEquivalentTo([]typeWithNoRequires{new(typeWithNoRequiresOneImpl)}))
			})
		})

		Context("when multiple implementations have been registered with the same type", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithNoRequiresOneImpl))
				factory.Register(new(typeWithNoRequiresTwoImpl))
				allValues, err = factory.GetAllOfType(new(typeWithNoRequires))
				implSlice = allValues.([]typeWithNoRequires)
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return implementaions of the correct type", func() {
				Ω(implSlice).ShouldNot(BeNil())
				Ω(implSlice).Should(BeEquivalentTo([]typeWithNoRequires{new(typeWithNoRequiresOneImpl), new(typeWithNoRequiresTwoImpl)}))
			})
		})

		Context("when no implementations have been registered with the given type", func() {

			var allvalues interface{}

			BeforeEach(func() {
				factory := knex.NewFactory()
				allvalues, err = factory.GetAllOfType(new(typeWithNoRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return an empty slice", func() {
				Ω(allvalues).Should(BeEquivalentTo([]typeWithNoRequires{}))
			})
		})

		Context("when one of the required fields is a slice", func() {

			var allvalues interface{}

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithNoRequiresOneImpl))
				factory.Register(new(typeWithNoRequiresTwoImpl))
				factory.Register(new(typeWithSliceRequiresImpl))
				allvalues, err = factory.GetByType(new(typeWithRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return implementaions of the correct type", func() {
				Ω(allvalues).ShouldNot(BeNil())
				impl, _ := newTypeWithSliceRequiresImpl([]typeWithNoRequires{new(typeWithNoRequiresOneImpl), new(typeWithNoRequiresTwoImpl)})
				Ω(allvalues).Should(BeEquivalentTo(impl))
			})
		})

		Context("when one of the required fields is a slice and no types have been registered for the slice", func() {

			var allvalues interface{}

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithSliceRequiresImpl))
				allvalues, err = factory.GetByType(new(typeWithRequires))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})

			It("should return implementaion with slice with legth of 0", func() {
				Ω(allvalues).ShouldNot(BeNil())
				impl, _ := newTypeWithSliceRequiresImpl([]typeWithNoRequires{})
				Ω(allvalues).Should(BeEquivalentTo(impl))
			})
		})

		Context("when multiple implementations have been registered and one of the implementations fail", func() {

			var allvalues interface{}

			BeforeEach(func() {
				factory := knex.NewFactory()
				factory.Register(new(typeWithNoRequiresOneImpl))
				factory.Register(new(typeWithErrorInjectorImpl))
				allvalues, err = factory.GetAllOfType(new(typeWithNoRequires))
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
