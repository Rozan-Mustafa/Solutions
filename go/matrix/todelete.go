package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

// New creates a new matrix from a string
func main() {
	m := Matrix{}
	s := "1 2 3\n4 5 6"
	
