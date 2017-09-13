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

func (s square) check_column(board [][]int) []int {
  var nums = make([]int, 9, 9)

  count := 0
  for i := 0; i < 9; i++ {
    nums[count] = board[i][s.col]
    count++
  }
  return nums
}

func (s square) check_row(board [][]int) []int {
  var nums = make([]int, 9, 9)

  count := 0
  for i := 0; i < 9; i++ {
    nums[count] = matrix[s.row][i]
    count++
  }
  return nums
}

func (s square) check_local_box(board [][]int) []int {
  var nums = make([]int, 9, 9)

  start_row, start_col := (s.row / 3) * 3, (s.col / 3) * 3

  count := 0
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      nums[count] = board[start_row + i][start_col + j]
      count++
    }
  }
  return nums
}

func (s square) update_possibly_list(numbers []int) {
  s.mutex.Lock()
  for n := range numbers {
      s.possibly[n] = true
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
      s.num = i
      return
    }
  }
}
