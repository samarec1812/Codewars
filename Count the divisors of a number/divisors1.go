package kata

func Divisors(n int)int{
   if(n == 1) {
      return n
   }
   count := 2
   for i := 2; i <= n/2; i++ {
        if n % i == 0 {
            count += 1
        }
   }
   return count 
  
}
