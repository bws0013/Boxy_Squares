package main

import (
  "fmt"

)

/*
  This file is just for testing methods to go into board without cluttering it up.
*/

func main() {

  // print_board("unsolved/board_easy_1.csv")
  // fmt.Println()


  m := make_board("unsolved/board_easy_1.csv")

  print_matrix(m)

  fmt.Println()

  r := get_local_box_numbers(3, 5, m)
  fmt.Println(m[3][5])

  fmt.Println(r)

}


func get_local_box_numbers(row, col int, matrix [][]int) []int {
  var nums = make([]int, 9, 9)

  start_row, start_col := (row / 3) * 3, (col / 3) * 3

  count := 0
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      nums[count] = matrix[start_row + i][start_col + j]
      count++
    }
  }

  return nums
}
