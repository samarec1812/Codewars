package kata


func sum_of_arithmetic_progression(count, number int) int {

      return (2*number + number*(count - 1))*count/ 2
}

func Multiple3And5(number int) int {
  var (
    count1, count2, count3, sum int
  )
  count1 = (number-1)/3
  count2 = (number-1)/5
  count3 = (number-1)/15
  
    sum += sum_of_arithmetic_progression(count1, 3)
    sum += sum_of_arithmetic_progression(count2, 5)
    sum -= sum_of_arithmetic_progression(count3, 15)

  return sum

}
