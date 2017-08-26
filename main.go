package main

import (
  "fmt"
  "flag"
)

func main() {
  fmt.Println("This is meminspector")

  pidPtr := flag.Int("pid", 0, "process id")
  flag.Parse()
  fmt.Printf("Inspecting memory for process: %d\n", *pidPtr)
}
