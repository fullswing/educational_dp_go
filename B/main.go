package main

// https://atcoder.jp/contests/dp/tasks/dp_b

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const INF = math.MaxInt32

func main() {
	// 標準入力の準備
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 1010000)
	sc.Split(bufio.ScanWords)
	N := nextInt(sc)
	K := nextInt(sc)
	hs := make([]int, N)
	for i := 0; i < N; i++ {
		hs[i] = nextInt(sc)
	}
	dp := make([]int, N)
	dp[0] = 0
	dp[1] = abs(hs[0], hs[1])
	for i := 2; i < N; i++ {
		cost := math.MaxInt32
		for k := 1; k <= K; k++ {
			if i-k < 0 {
				continue
			} else {
				curr := dp[i-k] + abs(hs[i-k], hs[i])
				cost = min(cost, curr)
			}
		}

		dp[i] = cost
	}
	fmt.Println(dp[N-1])
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

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
