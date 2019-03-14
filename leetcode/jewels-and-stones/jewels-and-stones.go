package main

import (
	"bufio"
	"fmt"
	"os"
)

func numJewelsInStones(J string, S string) int {
	printf("%s ha %s\n", J, S)
	return 3 
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
	fmt.Println(j,s)
	numJewelsInStones(j,s)
}

//  LocalWords:  numJewelsInStones
