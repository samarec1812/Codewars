package kata_test
import (
  "fmt"
  "math/rand"
  "time"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

var _ = Describe("Full Test Suite",func() {
  var RR = func(x,y int) int {return rand.Intn(y-x) + x}
  
  var shuffle = func(r []int) []int {
    for i := 0; i < len(r); i++ {
      n := RR(0,len(r))
      r[n],r[i] = r[i],r[n]}
    return r
  }
  
  // initiate random seed
  rand.Seed(time.Now().UTC().UnixNano())
  
  Context("Example tests",func() {
    arr := [...][]int{
      {20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5},
      {1,1,2,-2,5,2,4,4,-1,-2,5},
      {20,1,1,2,2,3,3,5,5,4,20,4,5},
      {10},
      {1,1,1,1,1,1,10,1,1,1,1},
      {5,4,3,2,1,5,4,3,2,10,10},
    }
    sol := [...]int{5,-1,5,10,10,1}
    for i,v := range arr {
      It(fmt.Sprintf("Testing input %v",v),func() {Expect(FindOdd(v)).To(Equal(sol[i]))})}})
  
  Context("Random tests",func() {
    for i := 0; i < 40; i++ {
      size := RR(50,1001)
      if size%2 == 0 {size++}
      nr := make([]int,size)
      for c := 0; c < size-1; c += 2 {
        rn := RR(0,1000)
        nr[c] = rn
        nr[c+1] = rn
      }
      qn := RR(0,1000)
      nr[size-1] = qn
      shuffle(nr)
      It("",func() {Expect(FindOdd(nr)).To(Equal(qn))})}})
})
