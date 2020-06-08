package kata

func DigitalRoot(n int) int {
   sum := 0
    for ;n > 0; {
        sum += n % 10 
        n /= 10
        
        if n <= 0 {
            if sum < 10 {
        break
        }
          n, sum = sum, 0
        }
        
    }
    
    return sum
    
}
