package main

import (
  "fmt"
)

func fibonacci() func() int {
  x, y := 0, 1
  return func() (r int) {
      r = x
      x, y = y, x + y
      return
  }
}

func main() {
  f := fibonacci()
  for counter := 0; counter < 10; counter++ {
    fmt.Println(f())
  }
}