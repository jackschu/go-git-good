// solving https://leetcode.com/problems/perfect-squares/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

//maps from n  to number required
var memo map[int]int = make(map[int]int)

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	if memo[n] != 0 {
		return memo[n]
	}
	highestSquare := int(math.Sqrt(float64(n)))
	min := n + 1
	for i := highestSquare; i > 0; i-- {
		cur := numSquares(n - i*i)
		if cur < min {
			min = cur
		}
	}
	memo[n] = min + 1
	return memo[n]
}

func NumSquares(n int) int {
	return numSquares(n)
}

func main() {
	sc.Split(bufio.ScanLines)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	fmt.Println("result", numSquares(n))
}

//  LocalWords:  numSquares NumSquares
