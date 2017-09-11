package main

import (
  "io"
  "os"
  "fmt"
  "log"
  "bufio"
  "encoding/csv"
)

func main() {

  print_board("unsolved/board_easy_1.csv")

}

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
