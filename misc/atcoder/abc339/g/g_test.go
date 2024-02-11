// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc339/tasks/abc339_g
// 提交：https://atcoder.jp/contests/abc339/submit?taskScreenName=abc339_g
// 对拍：https://atcoder.jp/contests/abc339/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc339_g&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc339/submissions?f.Status=AC&f.Task=abc339_g&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [g]")
	testCases := [][2]string{
		{
			`8
2 0 2 4 0 2 0 3
5
1 8 3
10 12 11
3 3 2
3 6 5
12 0 11`,
			`9
2
0
8
5`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}