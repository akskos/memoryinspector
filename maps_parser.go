// Parses /proc/{pid}/maps

package main

import (
  "fmt"
)

func fileToLines(filename string) []string {
    result := []string{}
    for i := 0; i < 10; i++ {
      result = append(result, fmt.Sprintf("line %d", i))
    }
    return result
}
