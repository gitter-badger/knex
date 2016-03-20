package test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUtility(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Knex Suite")
}
