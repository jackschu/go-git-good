// tests for https://leetcode.com/problems/perfect-squares/
package main

import "testing"

func TestSquares(t *testing.T) {
	ins := [...]int{12, 13, 18}
	outs := [...]int{3, 2, 2}
	for i, n := range ins {
		got := NumSquares(n)
		if  got != outs[i] {
			t.Errorf("incorrect on input %d, got %d expected %d", n, got, outs[i])
		}
	}
}

//  LocalWords:  TestSquares NumSquares
