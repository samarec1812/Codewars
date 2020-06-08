function hasPair(array, sum){
  
 let first = 0;
  let last = array.length - 1;
  while (first < last) {
    let s = array[first] + array[last];
    if (s == sum) {
      return true; 
      first++;
      last--;
    } else {
      if (s < sum) first++;
      else last--;
    }
  }
  return false; 
}
