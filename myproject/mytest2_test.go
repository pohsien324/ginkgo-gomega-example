package myproject_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
        . "github.com/pohsienshih/ginkgo_gomega_example/myproject"
)

var _ = Describe("Mytest2", func() {
     var number2 int 

     // If you want to test BeforeEach function, comment the following assignment.
     //number2  = 99


     BeforeEach(func(){
       // If you want to test BeforeEach function, uncomment the following assignment.
       number2  = 99
     }) 

     Describe("Test Calc function with BeforeEach", func() {
        Context("Recieve correct number", func() {
            It("should return 990", func() {
                Expect(Calc(number2)).To(Equal(990))

                // If you don't use BeforeEach function, the value of number2 will be overridden after execution.
                number2 = Calc(number2)
            })
            It("should return 990 too", func() {
                Expect(Calc(number2)).To(Equal(990))
            })
        })
      })
})
