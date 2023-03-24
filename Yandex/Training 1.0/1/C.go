package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var s string
	fmt.Fscan(in, &s)
	s = cutBack(s)
	for i := 0; i < 3; i++ {
		var l string
		fmt.Fscan(in, &l)
		if s == cutBack(l) {
			out.WriteString("YES\n")
		} else {
			out.WriteString("NO\n")
		}
	}
}

func cutBack(s string) string {
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	switch len(s) {
	case 11:
		s = s[1:]
	case 12:
		s = s[2:]
	case 7:
		s = "495" + s
	}
	return s
}
