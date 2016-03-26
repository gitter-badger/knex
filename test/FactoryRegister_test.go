package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/chrisehlen/knex"
)

var _ = Describe("Factory", func() {

	Describe("register an implementation", func() {

		var (
			err error
		)

		Context("when a propertly annotated implementation is  registered", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				err = factory.Register(new(typeWithNoRequiresOneImpl))
			})

			It("should be successful", func() {
				Ω(err).Should(Succeed())
			})
		})

		Context("when an implementaion's require filed has an invalid value", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				err = factory.Register(new(typeWithInvalidRequiresImpl))
			})

			It("should return a 'Invalid require value' error", func() {
				Ω(err.Error()).Should(HavePrefix("Invalid require value "))
			})
		})

		Context("when an implementaion's provide filed has an invalid value", func() {

			BeforeEach(func() {
				factory := knex.NewFactory()
				err = factory.Register(new(typeWithInvalidProvidesImpl))
			})

			It("should return a 'Invalid provide value' error", func() {
				Ω(err.Error()).Should(HavePrefix("Invalid provide value "))
			})
		})
	})
})
