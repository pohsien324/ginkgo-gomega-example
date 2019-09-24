package myproject_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/pohsienshih/ginkgo-gomega-example/myproject"
)

var _ = Describe("Mytest", func() {
        
    var name string= "test"
    var number int= 99

    Describe("Test greeting function", func() {
        Context("When receive the name", func() {
            It("should greeting", func() {
                Expect(Greeting(name)).To(Equal("Hello "+name+"."))
            })
        })
   })

    Describe("Test calc function", func() {
        Context("Give the number", func() {
            It("should get the correct result", func() {
                Expect(Calc(number)).To(Equal(number*10))
            })
        })
   })
})
