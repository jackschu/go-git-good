// solving https://leetcode.com/problems/jewels-and-stones/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func numJewelsInStones(J string, S string) int {
	out := 0
	jewelSet := make(map[rune]bool)
	for _, rune := range J {		
		jewelSet[rune] = true
	}
	
	for _, rune := range S {
		_, ok := jewelSet[rune]
		if ok {
			out++
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
	var j, s string
	sc.Scan()
	j = sc.Text()
	sc.Scan()
	s = sc.Text()
	fmt.Println(numJewelsInStones(j, s))
}

//  LocalWords:  numJewelsInStones jewelSet ok
