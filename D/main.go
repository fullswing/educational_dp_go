package main

// https://atcoder.jp/contests/dp/tasks/dp_d

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const INF = math.MaxInt32

func calc(dp [][]int64, w []int, v []int64, N int, W int) int64 {
	for i := N; i >= 1; i-- {
		for j := 1; j <= W; j++ {
			if j < w[i] {
				dp[i][j] = dp[i+1][j]
			} else {
				// i番目の荷物を使うか使わないかの二択
				dp[i][j] = max(dp[i+1][j], dp[i+1][j-w[i]]+v[i])
			}
		}
	}
	return dp[1][W]
}

func main() {
	// 標準入力の準備
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 1010000)
	sc.Split(bufio.ScanWords)
	N := nextInt(sc)
	W := nextInt(sc)

	dp := make([][]int64, N+2)
	w := make([]int, N+1)
	v := make([]int64, N+1)
	dp[0] = make([]int64, W+1)
	dp[N+1] = make([]int64, W+1)
	for i := 1; i < N+1; i++ {
		w[i] = nextInt(sc)
		v[i] = int64(nextInt(sc))
		dp[i] = make([]int64, W+1)
	}
	ans := calc(dp, w, v, N, W)
	fmt.Println(ans)
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

// int型専用のContains関数
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
