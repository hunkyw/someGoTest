package main

import (
	"testing"
)

func TestSubstr(t *testing.T){
	tests := [] struct{
		s string
		ans int
 	}{
 		//normal cases
 		{"abcabcbb",3},
 		{"pwwkew",3},
 		// Edge cases
 		{"",0},
 		{"b",1},
 		{"bbbbbbbbb",1},
 		{"abcabcabcabcd",4},
 	//chinene support
 		{"这里是慕课网",6},
 		{"一二三三二一",3},
 		{"黑化肥会挥发发灰会花飞灰化肥挥发发黑会飞花",8},
	}

	for _,tt := range tests{
		actual := lengthOfNonrepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s ;" + "expected %d ", actual , tt.s , tt.ans)
		}
	}
}
