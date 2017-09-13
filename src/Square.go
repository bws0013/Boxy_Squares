package main

import (
  "fmt"
  "sync"
)

// num int = 0
// possibly = []bool{false, true, true, true, true, true, true, true, true, true}
// mutex = &sync.Mutex{}

type square struct {
  row, col, num int
  possibly []bool
  mutex sync.Mutex
}

func full_sequence(s square) {
  var nums []int

  nums = append(nums, s.check_column()...)
  nums = append(nums, s.check_row()...)
  nums = append(nums, s.check_local_box()...)

  s.update_possibly_list(nums)

  // col_nums := s.check_column()
  // row_nums := s.check_row()
  // local_box_nums := s.check_local_box()
}

func (s square) print_stuff() {
  fmt.Println(s.num)
  fmt.Println(s.possibly)
}

func (s square) keep_going() bool {
  if s.num != 0 {
    return false
  } else {
    return true
  }
}

func (s square) check_column() []int {
  var nums = make([]int, 9, 9)

  count := 0
  for i := 0; i < 9; i++ {
    nums[count] = big_board[i][s.col].num
    count++
  }
  return nums
}

func (s square) check_row() []int {
  var nums = make([]int, 9, 9)

  count := 0
  for i := 0; i < 9; i++ {
    nums[count] = big_board[s.row][i].num
    count++
  }
  return nums
}

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

func (s square) update_possibly_list(numbers []int) {
  s.mutex.Lock()
  for _, n := range numbers {
    //fmt.Println(n)
    s.possibly[n] = false
  }
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

func (s square) check_possibly_list() {
  for i, p := range s.possibly {
    if p == true {
      big_board[s.row][s.col].num = i
      return
    }
  }
}
