package kata

func Divisors(n int)int{
    kol := 1
    for i:= 2; i <= n; i++ {
      if n % i == 0 {
       kol += 1
      }
    }
    return kol
 
}
