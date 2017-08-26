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

  lines := fileToLines(fmt.Sprintf("/proc/%d/maps", *pidPtr))
  addrSpace := getAddressSpaceForLabel(lines, "[heap]")
  heapData := readMemorySpace(*pidPtr, addrSpace)

  fmt.Printf("heap begins at %d\n", addrSpace[0])

  matchString := []byte("hello")
  fmt.Println(findMatchesInByteArray(heapData, matchString))
}
