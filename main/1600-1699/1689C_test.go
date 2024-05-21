// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1689/C
// https://codeforces.com/problemset/status/1689/problem/C
func Test_cf1689C(t *testing.T) {
	testCases := [][2]string{
		{
			`4
2
1 2
4
1 2
2 3
2 4
7
1 2
1 5
2 3
2 4
5 6
5 7
15
1 2
2 3
3 4
4 5
4 6
3 7
2 8
1 9
9 10
9 11
10 12
10 13
11 14
11 15`,
			`0
2
2
10`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1689C)
}