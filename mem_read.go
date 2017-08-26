// Funtions for reading process memory

package main

import (
  "os"
  "fmt"
)

// Reads memory space specified in {space} and returns a byte array of the
// contents
func readMemorySpace(pid int, space [2]int64) []byte {
  file, err := os.Open(fmt.Sprintf("/proc/%d/mem", pid))
  check(err)
  defer file.Close()
  file.Seek(space[0], 0)
  buffer := make([]byte, 100)
  n, err := file.Read(buffer)
  check(err)
  fmt.Printf("%d bytes read\n", n)
  return buffer
}
