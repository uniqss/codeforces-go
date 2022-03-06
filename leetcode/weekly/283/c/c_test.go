// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_c(t *testing.T) {
	examples := [][]string{
		{
			`[[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]`, 
			`[50,20,80,15,17,19]`,
		},
		{
			`[[1,2,1],[2,3,0],[3,4,1]]`, 
			`[1,2,null,null,3,4]`,
		},
		
	}
	targetCaseNum :=  -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, createBinaryTree, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-283/problems/create-binary-tree-from-descriptions/
