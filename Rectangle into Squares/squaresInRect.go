package kata

import (
    "math"
)


func SquaresInRect(lng int, wdth int) []int {
    var arr [] int    
    if lng == wdth {
     return arr
    }
    for lng > 0 && wdth > 0   {
    newLength := int(math.Min(float64(lng), float64(wdth)))
    arr = append(arr, newLength)
    if newLength == lng {
         wdth -= newLength
     } else if newLength == wdth {
        lng -= newLength
    }
  

    }
    return arr
}
