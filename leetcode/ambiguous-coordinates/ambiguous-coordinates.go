// solving https://leetcode.com/problems/ambiguous-coordinates/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ambiguousCoordinates(S string) []string {
	r := strings.Split(strings.Trim(S,"()"),"")
	fmt.Println(r)
	return []string{"hi"}
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
