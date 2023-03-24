package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, k int
	s    string
	p    []int
	d    []int
	S    string
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k, &s)
	p, d = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &d[i])
	}
	var k int
	var reap = make(map[uint8]int)
	for i := range s {
		sym := s[i] - 97
		m := reap[sym]
		if m == 0 {
			m++
			S += Sym(sym)
		} else {

		}

	}
}

func Sym(n uint8) string {
	return string(n + 97)
}
func repeat(sym int) int {

}
