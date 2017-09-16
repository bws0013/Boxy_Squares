package main

import (
  "fmt"

)

/*
  This file is just for testing methods to go into board without cluttering it up.
*/

func example_box_of_squares() {

  // print_board("unsolved/board_easy_1.csv")
  // fmt.Println()

  m := make_board("unsolved/board_easy_1.csv")

  print_matrix(m)

  fmt.Println()
  lb := get_local_box_numbers(3, 5, m)
  fmt.Println(lb)

  fmt.Println()
  r := get_row(3, m)
  fmt.Println(r)

  fmt.Println()
  c := get_col(3, m)
  fmt.Println(c)

}

// Test for getting the column numbers for a square
func get_col(col int, matrix [][]int) []int {
  var nums = make([]int, 9, 9)

  count := 0
  for i := 0; i < 9; i++ {
    nums[count] = matrix[i][col]
    count++
  }
  return nums
}

// Test for getting the row numbers for a square
func get_row(row int, matrix [][]int) []int {
  var nums = make([]int, 9, 9)

  count := 0
  for i := 0; i < 9; i++ {
    nums[count] = matrix[row][i]
    count++
  }
  return nums
}

// Test for getting local box numbers
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
