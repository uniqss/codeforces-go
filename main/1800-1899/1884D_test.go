// Generated by copypasta/template/generator_test.go
package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1884/D
// https://codeforces.com/problemset/status/1884/problem/D
func Test_cf1884D(t *testing.T) {
	testCases := [][2]string{
		{
			`6
4
2 4 4 4
4
2 3 4 4
9
6 8 9 4 6 8 9 4 9
9
7 7 4 4 9 9 6 2 9
18
10 18 18 15 14 4 5 6 8 9 10 12 15 16 18 17 13 11
21
12 19 19 18 18 12 2 18 19 12 12 3 12 12 12 18 19 16 18 19 12`,
			`0
3
26
26
124
82`,
		},
		{
			`1
6
3 5 3 6 3 4`,
			`9`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1884D)
}

func TestCompare_cf1884D(_t *testing.T) {
	return
	testutil.DebugTLE = 0
	rg := testutil.NewRandGenerator()
	inputGenerator := func() string {
		//return ``
		rg.Clear()
		rg.One()
		n := rg.Int(1, 6)
		rg.NewLine()
		rg.IntSlice(n, 1, n)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		solve := func(Case int) {
			var n int
			Fscan(in, &n)
			a := make([]int, n)
			for i := range a {
				Fscan(in, &a[i])
			}
			ans := 0
			for i, v := range a {
				o:for j := i + 1; j < len(a); j++ {
					w := a[j]
					for _, k := range a {
						if v%k==0 && w%k==0 {
							continue o
						}
					}
					ans++
				}
			}
			Fprintln(out, ans)
		}

		T := 1
		Fscan(in, &T)
		for Case := 1; Case <= T; Case++ {
			solve(Case)
		}
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, cf1884D)
}