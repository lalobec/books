package main

import (
  "bytes"
  "testing"
)

func TestCountWords(t *testing.T) {
  b := bytes.NewBufferString("word1 word2 word3 word4\n")
  expected := 4
  result := count(b, false)

  if result != expected {
    t.Errorf("------------FAILED TEST-----------------")
    t.Errorf("Expected %d, got %d instead.\n", expected, result)
    t.Errorf("----------------------------------------")
  }
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 \nline2\nline3 word1")
	expected := 3
	result := count(b, true)
	
	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}

}
