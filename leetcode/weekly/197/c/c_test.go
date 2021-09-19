// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`3`, `[[0,1],[1,2],[0,2]]`, `[0.5,0.5,0.2]`, `0`, `2`, 
			`0.25000`,
		},
		{
			`3`, `[[0,1],[1,2],[0,2]]`, `[0.5,0.5,0.3]`, `0`, `2`, 
			`0.30000`,
		},
		{
			`3`, `[[0,1]]`, `[0.5]`, `0`, `2`, 
			`0.00000`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxProbability, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-197/problems/path-with-maximum-probability/