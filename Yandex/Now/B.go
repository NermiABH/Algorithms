package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, x, t, M int
	fmt.Fscan(in, &n, &x, &t)
	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		var V int
		fmt.Fscan(in, &V)
		V = int(math.Abs(float64(x - V)))
		a[i] = [2]int{V, i + 1}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	var sum int
	var ans string
	for _, v := range a {
		if sum += v[0]; sum > t {
			break
		}
		M++
		ans += strconv.Itoa(v[1]) + " "
	}
	fmt.Printf("%v\n%s", M, ans)
}
