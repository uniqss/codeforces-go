// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/dp/tasks/dp_v
// 最短：https://atcoder.jp/contests/dp/submissions?f.Status=AC&f.Task=dp_v&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [v]")
	testCases := [][2]string{
		{
			`3 100
1 2
2 3`,
			`3
4
3`,
		},
		{
			`4 100
1 2
1 3
1 4`,
			`8
5
5
5`,
		},
		{
			`1 100`,
			`1`,
		},
		{
			`10 2
8 5
10 8
6 5
1 5
4 8
2 10
3 6
9 2
1 7`,
			`0
0
1
1
1
0
1
0
1
1`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}