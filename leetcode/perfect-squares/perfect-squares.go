// solving https://leetcode.com/problems/perfect-squares/
package main

import (
	"strconv"
	"bufio"
	"fmt"
	"math"
	"os"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func numSquares(n int) int {
	return int(math.Sqrt(9))
}

func NumSquares(n int) int {
	return numSquares(n)
}

func main() {
	sc.Split(bufio.ScanLines)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	numSquares(n)
}
