package main

import (
  "fmt"
  "sync"
)

/*
Variable explanation for square struct
  row is the value (0-8) of the square
  col (column) is the value (0-8) of the square
  num is the value (1-9) contained in the square (the number you see on the board)
  possibly is a slice containing the possible values contained in this square
    if the squares value is 0 than
    if the squares value is not 0 than
*/

// num int = 0
// possibly = []bool{false, true, true, true, true, true, true, true, true, true}
// mutex = &sync.Mutex{}

type square struct {
  row, col, num int
  possibly []bool
  mutex sync.Mutex
}

/*
  Run through of checking a squares column, row and local box.
  Collects all of the numbers our square could possibly be, appends them to a
  slice and passes them to be checked to see if any that list can narrow down
  the possibilities of the squares value to 1.
*/
func full_sequence(s square) {
  if s.num != 0 {
    return
  }

  var nums []int

  // Later we can make each check its own thread to see what happens.
  nums = append(nums, s.check_column()...)
  nums = append(nums, s.check_row()...)
  nums = append(nums, s.check_local_box()...)

  s.update_possibly_list(nums)
}

/*
  Used for testing purposes.
  Prints the value of a square as well as its list of potential values
*/
func (s square) print_stuff() {
  fmt.Println(s.num)
  fmt.Println(s.possibly)
}

/*
  Checks to see if we need to check for the value of our square. If the value
  is not 0 then it should be correct therefore we do not have to waste effort
  on determining what it is.
*/
func (s square) keep_going() bool {
  if s.num != 0 {
    return false
  } else {
    return true
  }
}

/*
  Returns all of the numbers on the same column as that of our square s.
*/
func (s square) check_column() []int {
  var nums = make([]int, 9, 9)

  count := 0
  for i := 0; i < 9; i++ {
    nums[count] = big_board[i][s.col].num
    count++
  }
  return nums
}

/*
  Returns all of the numbers on the same row as that of our square s.
*/
func (s square) check_row() []int {
  var nums = make([]int, 9, 9)

  count := 0
  for i := 0; i < 9; i++ {
    nums[count] = big_board[s.row][i].num
    count++
  }
  return nums
}

/*
  Returns all of the values contained in a squares local area. There are 9 3x3
  local boxes on a standard sudoku board. The value returned includes the value
  of the square itself but as this is only called when the square equals 0 we
  know that we could not accidently remove the actual number of the square from
  the running for the squares actual value.

*/
func (s square) check_local_box() []int {
  var nums = make([]int, 9, 9)

  start_row, start_col := (s.row / 3) * 3, (s.col / 3) * 3

  count := 0
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      nums[count] = big_board[start_row + i][start_col + j].num
      count++
    }
  }
  return nums
}

/*
  Given a slice of numbers representing those from the row, column and local box
  we can determine those numbers which our s.val cannot be. Those numbers which
  our val cannot be are made false on the s.possibly list. If that list only
  contains 1 true value we can pass the square to the check_possibly_list method
*/
func (s square) update_possibly_list(numbers []int) {
  s.mutex.Lock()
  for _, n := range numbers {
    // fmt.Println(n)
    s.possibly[n] = false
  }
  // fmt.Println("++++++++++++++++++++")
  choices := 0
  for _, p := range s.possibly {
    if p == true {
      choices++
    }
  }
  s.mutex.Unlock()
  if choices == 1 {
    s.check_possibly_list()
  }
}

/*
  Check the list of possible square values for the true value, if there is only 1
  value then we can know that there can only be one value for that square. That
  one vlaue will be the only value for which the square can represent.

  When this method is called there would only ever be a single true value in the
  s.possibly slice
*/
func (s square) check_possibly_list() {
  for i, p := range s.possibly {
    if p == true {
      // Both of the below appear neccessary, one for logic, one for image
      // But im not entirely sure of the above statement. Check in the morning.
      big_board[s.row][s.col].num = i
      s.num = i
      return
    }
  }
}
