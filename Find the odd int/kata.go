package kata

func FindOdd(seq []int) int {
    var k int =  seq[0]
  for i := 1; i < len(seq); i++ {
    k ^= seq[i]
  }
   return k
}
