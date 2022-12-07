package main

import (
	"fmt"
)

func i_swap(arr[]int) []int{
  out := []int{};
  added := 0;
  none_index := 0
  for i := 0; i < len(arr); i++{
    for j := 0; j < len(arr); j++{ 
      if i != arr[j]{
        added = -1
        continue;
      } else{
        added = 1
        none_index = j
        break;
      }
    }
    if added == -1{
      out = append(out, added);
    } 
    if added == 1{
      out = append(out, arr[none_index]);
    }
  }
  return out
}

func reverseString(str string) (result string){
  for _, v := range str{
    result = string(v) + result;
  }
  return 
}
func main(){
  arr := []int{19, 7, 0, 3, 18, 15, 12, 6, 1, 8, 11, 10, 9, 5, 13, 16, 2, 14, 17, 4}
  out := i_swap(arr)
  fmt.Println(out)
  fmt.Println(reverseString("Hello"))
}
