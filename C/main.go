package main

// https://atcoder.jp/contests/dp/tasks/dp_c

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const INF = math.MaxInt32

func calc(h [][]int, dp [][]int, N int) int {
	for i := 1; i < N+1; i++ {
		dp[i][0] = h[i-1][0] + max(dp[i-1][1], dp[i-1][2])
		dp[i][1] = h[i-1][1] + max(dp[i-1][0], dp[i-1][2])
		dp[i][2] = h[i-1][2] + max(dp[i-1][0], dp[i-1][1])
	}
	return max(dp[N][0], max(dp[N][1], dp[N][2]))
}

func main() {
	// 標準入力の準備
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 1010000)
	sc.Split(bufio.ScanWords)
	N := nextInt(sc)
	h := make([][]int, N)
	for i := 0; i < N; i++ {
		h[i] = make([]int, 3)
		h[i][0] = nextInt(sc)
		h[i][1] = nextInt(sc)
		h[i][2] = nextInt(sc)
	}
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, 3)
	}
	ans := calc(h, dp, N)
	fmt.Println(ans)
	fmt.Println(dp)
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
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
