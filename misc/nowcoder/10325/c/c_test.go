// Code generated by copypasta/template/nowcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`2`,`[[0,0, 2, 0], [4, 0, -2, 0]]`,
			`0.000`,
		},
		{
			`3`,`[[0, 0, 2, 0],[ 4, 0, -2, 0],[ -5, 0, 2, 0]]`,
			`5.000`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, solve, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://ac.nowcoder.com/acm/contest/10325/c
