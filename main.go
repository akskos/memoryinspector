package main

import (
  "fmt"
  "flag"
  "os"
)

func main() {
  pidPtr := flag.Int("pid", 0, "process id")
  matchStringPtr := flag.String("m", "", "will find addresses for matches of this string in process memory data")
  overwritePtr := flag.String("o", "", "string used to overwrite matches")
  flag.Parse()
  fmt.Printf("Inspecting memory for process: %d\n", *pidPtr)

  lines := fileToLines(fmt.Sprintf("/proc/%d/maps", *pidPtr))
  addrSpace := getAddressSpaceForLabel(lines, "[heap]")
  heapData := readMemorySpace(*pidPtr, addrSpace)

  fmt.Printf("Heap begins at %d\n", addrSpace[0])

  matchBytes := []byte(*matchStringPtr)
  matches := findMatchesInByteArray(heapData, matchBytes)
  fmt.Println(matches)
  for _, match := range matches {
    fmt.Printf("Match at mem address: %d\n", addrSpace[0] + int64(match))
  }

  // Overwrite matches if specified
  if len(*overwritePtr) > 0 {
    file, err := os.OpenFile(fmt.Sprintf("/proc/%d/mem", *pidPtr), os.O_WRONLY, 0600)
    check(err)
    defer file.Close()
    for _, match := range matches {
      _, err := file.WriteAt([]byte(*overwritePtr), addrSpace[0] + int64(match))
      check(err)
      fmt.Printf("Overwritten match at mem address: %d\n", addrSpace[0] + int64(match))
    }
  }
}
