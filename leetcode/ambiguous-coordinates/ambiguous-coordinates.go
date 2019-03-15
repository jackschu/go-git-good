// solving https://leetcode.com/problems/ambiguous-coordinates/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func insert(i int, val rune, slice []rune) []rune {
	slice = append(slice, rune(0))
	copy(slice[i+1:], slice[i:])
	slice[i] = val
	return slice
}

func isNoDecimalLegal(list []rune) bool {
	if len(list) == 1 && list[0] == '0' {
		return true
	}
	if len(list) == 0 {
		return false
	}
	return list[0] != '0'
}

func isValidDecimalPart(list []rune) bool {
	if len(list) == 0 {
		return false
	}
	if list[len(list)-1] == '0' {
		return false
	}
	return true
}

func validDecimals(list []rune) []string {
	var out []string
	if isNoDecimalLegal(list) {
		temp := make([]rune, len(list))
		copy(temp, list)
		out = append(out, string(temp))
	}

	for i := 1; i < len(list); i++ {
		front := list[:i]
		back := list[i:]
		if isNoDecimalLegal(front) && isValidDecimalPart(back) {
			out = append(out, string(front)+"."+string(back))
		}

	}

	return out
}

func ambiguousCoordinates(S string) []string {
	list := []rune(strings.Trim(S, "()"))
	var out []string
	had := make(map[string]bool)

	for i := 1; i < len(list); i++ {
		first := list[:i]
		second := list[i:]
		for _, pre := range validDecimals(first) {
			for _, post := range validDecimals(second) {
				temp := "(" + pre + ", " + post + ")"
				_, found := had[temp]
				if !found {
					out = append(out, temp)
					had[temp] = true
				}

			}
		}

	}

	return out
}

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	sc.Split(bufio.ScanWords)
	defer writer.Flush()
	var s string
	sc.Scan()
	s = sc.Text()
	fmt.Println(ambiguousCoordinates(s))
}

//  LocalWords:  ambiguousCoordinates AmbiguousCoordinates
