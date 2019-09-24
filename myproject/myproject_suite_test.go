package myproject_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMyproject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Myproject Suite")
}
