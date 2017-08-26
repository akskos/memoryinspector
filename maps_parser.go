// Parses /proc/{pid}/maps

package main

import (
  "fmt"
  "os"
  "bufio"
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
