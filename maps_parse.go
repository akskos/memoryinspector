// Functions for parsing /proc/{pid}/maps

package main

import (
  "os"
  "bufio"
  "strings"
  "strconv"
)

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
func getAddressSpaceForLabel(lines []string, label string) [2]int64 {
  addressSpace := [2]int64{}
  for _, line := range lines {
    if strings.Contains(line, label) {
      memoryColumn := strings.Split(line, " ")[0]
      addressStrings := strings.Split(memoryColumn, "-")
      var err error
      addressSpace[0], err = strconv.ParseInt(addressStrings[0], 16, 64)
      check(err)
      addressSpace[1], err = strconv.ParseInt(addressStrings[1], 16, 64)
      check(err)
      break
    }
  }
  return addressSpace
}
