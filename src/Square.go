package main

import (
  "fmt"
  "sync"
)

// num int = 0
// possibly = []bool{false, true, true, true, true, true, true, true, true, true}
// mutex = &sync.Mutex{}

type square struct {
  num int
  possibly []bool
  mutex sync.Mutex
}

// func main() {
//   fmt.Println(num)
//   fmt.Println(possibly)
// }

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

func (s square) check_column() {

}

func (s square) check_row() {

}

func (s square) check_local_box() {

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
