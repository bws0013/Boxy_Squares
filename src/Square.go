package main

import (
  "fmt"
  "sync"
)

var (
  num int = 0
  possibly = []bool{false, true, true, true, true, true, true, true, true, true}
  mutex = &sync.Mutex{}
)

func main() {
  fmt.Println(num)
  fmt.Println(possibly)
}

func check_column() {

}

func check_row() {

}

func check_local_box() {

}

func update_possibly_list(numbers []int) {
  mutex.Lock()
  for n := range numbers {
      possibly[n] = true
  }
  int choices := 0
  for _, p := range possibly {
    if p == true {
      choices++
    }
  }
  mutex.Unlock()
  if p == 1 {
    check_possibly_list()
  }
}

func check_possibly_list() {
  for i, _ := range possibly {
    if p == true {
      num = i
      return
    }
  }
}
