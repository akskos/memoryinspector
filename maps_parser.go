// Parses /proc/{pid}/maps

package main

import (
  "os"
  "bufio"
  "strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

// Reads {filename} and creates a slice of its lines
func fileToLines(filename string) []string {
  result := []string{}
  file, err := os.Open(filename)
  check(err)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    result = append(result, scanner.Text())
  }
  return result
}

// Parses {lines} and finds address space defined for {label}
func getAddressSpaceForLabel(lines []string, label string) [2]int {
  addressSpace := [2]int{}
  for _, line := range lines {
    if strings.Contains(line, label) {
      addressSpace[0] = 0
      addressSpace[1] = 88
      break
    }
  }
  return addressSpace
}
