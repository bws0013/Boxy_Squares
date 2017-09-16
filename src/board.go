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

// TODO Add more comments.

// Global variable zone, all of them should go here, no exceptions
var (
  // The name of the file to be turned into the board
  // board_name = "almost/board_unsolved_1.csv"
  // board_name = "unsolved/board_medium_1.csv"
  board_name = "unsolved/board_hard_1.csv"

  // The board every square interacts with
  big_board = make_smart_board(board_name)

  // row and col are the dimensions of the board, they should be fixed at 9 each
  row int = 9
  col int = 9
)

func main() {

  print_big_board()
  fmt.Println()
  print_num_0s_big_board()
  fmt.Println("\n=================")
  fmt.Println()

  for i := 0; i < 10; i++ {
    run_each_square()
    print_num_0s_big_board()
  }
  fmt.Println()
  print_big_board()
}

// Run through the full list of operations on a particular square on the board.
func run_each_square() {
  for i := 0; i < row; i++ {
    for j := 0; j < col; j++ {
      full_sequence(big_board[i][j])
    }
  }
}

// Create a board of square objects given a filename
func make_smart_board(filename string) [][]square {
  basic_board := make_board(filename)

  smart_board := make([][]square, row)
  for i := range smart_board {
      smart_board[i] = make([]square, col)
  }

  // I use var instead of :=  in the loop because I feel it looks clearer in this case
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

// Print the number of 0s that appear on the board
func print_num_0s_big_board() {

  total_0s := 0

  for i := 0; i < row; i++ {
    for j := 0; j < col; j++ {
      if big_board[i][j].num == 0 {
        total_0s++
      }
    }
  }
  fmt.Println(total_0s)
}

/*
  Prints the vlaues of the squares of the global board being used.
*/
func print_big_board() {
  for i := 0; i < row; i++ {
    for j := 0; j < col; j++ {
      fmt.Print(big_board[i][j].num, " ")
    }
    fmt.Println()
  }
}

/*
  A smart_board is what I am calling a board of square structs. This method
  prints the value of those square structs.
*/
func print_smart_board(board [][]square) {

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
  for i := 0; i < row; i++ {
    for j := 0; j < col; j++ {
      fmt.Print(matrix[i][j], " ")
    }
    fmt.Println()
  }
}

// Creates a board given a file name
func make_board(filename string) [][]int {

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
