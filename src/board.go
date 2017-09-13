package main

import (
  "io"
  "os"
  "fmt"
  "log"
  "sync"
  "bufio"
  "strconv"
  "encoding/csv"
)

// func main() {
//
//   print_board("unsolved/board_easy_1.csv")
//   fmt.Println()
//
//
//   m := make_board("unsolved/board_easy_1.csv")
//
//
//   print_matrix(m)
//
// }

func main() {

  board_name := "unsolved/board_easy_1.csv"

  print_board(board_name)

  fmt.Println()
  m := make_smart_board(board_name)

  print_smart_board(m)

}

func make_smart_board(filename string) [][]square {
  row, col := 9, 9

  basic_board := make_board(filename)

  smart_board := make([][]square, row)
  for i := range smart_board {
      smart_board[i] = make([]square, col)
  }

  // num int = 0
  // possibly = []bool{false, true, true, true, true, true, true, true, true, true}
  // mutex = &sync.Mutex{}

  for i := 0; i < row; i++ {
    for j := 0; j < col; j++ {
      var row = i
      var col = j
      var num = basic_board[i][j]
      var possibly = []bool{false, true, true, true, true, true, true, true, true, true}
      var mutex = sync.Mutex{}
      possibly[num] = false
      smart_board[i][j] = square{row, col, num, possibly, mutex}
    }
  }

  return smart_board
}


func print_smart_board(board [][]square) {
  row, col := 9, 9

  for i := 0; i < row; i++ {
    for j := 0; j < col; j++ {
      fmt.Print(board[i][j].num, " ")
    }
    fmt.Println()
  }
}



// Prints a board given a filename
func print_board(filename string) {
  f, err := os.Open("../storage/boards/" + filename)
  if err != nil {
    log.Fatal(err)
  }
    // Create a new reader.
  r := csv.NewReader(bufio.NewReader(f))

  // This could be a for-ever loop but we can assume 9x9
  for i := 0; i < 9; i++ {
    record, err := r.Read()
    // Stop at EOF.
    if err == io.EOF {
        break
    }
    for value := range record {
        fmt.Printf("%v ", record[value])
    }
    fmt.Println()
  }
}

// Prints a 9x9 matrix, fixed dimensions
func print_matrix(matrix [][]int) {
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      fmt.Print(matrix[i][j], " ")
    }
    fmt.Println()
  }
}

// Creates a board given a file name
func make_board(filename string) [][]int {
  var row, col int
  row = 9
  col = 9

  f, err := os.Open("../storage/boards/" + filename)
  if err != nil {
    log.Fatal(err)
  }
    // Create a new reader.
  r := csv.NewReader(bufio.NewReader(f))

  board := make([][]int, row)
  for i := range board {
      board[i] = make([]int, col)
  }

  // This could be a for-ever loop but we can assume 9x9
  for i := 0; i < 9; i++ {
    record, err := r.Read()
    // Stop at EOF.
    if err == io.EOF {
        break
    }
    j := 0
    for value := range record {
      val, err := strconv.Atoi(record[value])
      if err == nil {
        board[i][j] = val
      }
      j++
    }
  }

  return board
}
