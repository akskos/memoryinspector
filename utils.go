// Common utilities

package main

import (
  "bytes"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

// Returns offsets/indexes from {haystack} where {needle} matches
func findMatchesInByteArray(haystack []byte, needle []byte) []int {
  offsets := []int{}
  for {
    index := bytes.Index(haystack, needle)
    if (index == -1) {
      break
    }
    if len(offsets) == 0 {
      offsets = append(offsets, index)
    } else {
      offsets = append(offsets, index + offsets[len(offsets)-1] + 1)
    }
    haystack = haystack[index+1:]
  }
  return offsets
}
