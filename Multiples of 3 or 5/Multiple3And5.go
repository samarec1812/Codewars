package kata

func Multiple3And5(number int) int {
  primeSlice := 0

  for i := 3; i < number; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      primeSlice += i
    }
  }
  return primeSlice
}
