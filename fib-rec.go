package main

import(
  "fmt"
)

func fibonacci(n int) int {
  if n <= 1 {
    return n
  }
  return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
  for counter := 0; counter < 10; counter ++ {
    fmt.Println(fibonacci(counter))
  }
}
