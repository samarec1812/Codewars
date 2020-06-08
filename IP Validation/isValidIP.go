package kata

import (
      "strings"
      "strconv"
)

func Is_valid_ip(ip string) bool {

  slice := strings.Split(ip, ".")
    if len(slice) < 4 {
      return false
  }

  flag := true
  for i := 0; i < len(slice); i++ {

    if slice[i][0] == '0' && len(slice[i]) > 1 {
    flag = false
    return flag
  }
    for j:= 0; j < len(slice[i]); j++ {
      if slice[i][j] == ' ' || slice[i][j] > '9'{
        flag = false
        return flag
    }
    }
    number, errorSyntax := strconv.ParseInt(slice[i], 10, 32)
    if number >= 0 && number < 256 {
      continue
    } else if number < 0 || number > 255 || errorSyntax == strconv.ErrRange {
      flag = false
      return flag
    }

  }


  return true
}
