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

func print_stuff(this square) {
  fmt.Println(this.num)
  fmt.Println(this.possibly)
}

func keep_going(this square) bool {
  if this.num != 0 {
    return false
  } else {
    return true
  }
}

func check_column(this square) {

}

func check_row(this square) {

}

func check_local_box(this square) {

}

func update_possibly_list(this square, numbers []int) {

  this.mutex.Lock()
  for n := range numbers {
      this.possibly[n] = true
  }
  choices := 0
  for _, p := range this.possibly {
    if p == true {
      choices++
    }
  }
  this.mutex.Unlock()
  if choices == 1 {
    check_possibly_list(this)
  }
}

func check_possibly_list(this square) {
  for i, p := range this.possibly {
    if p == true {
      this.num = i
      return
    }
  }
}
