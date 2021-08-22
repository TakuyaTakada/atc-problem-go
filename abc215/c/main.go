package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func run(s string, k int) string {
	var ans string
	chars := Chars(strings.Split(s, ""))
	sort.Strings(chars)
	for i := 1; i < k; i++ {
		nextPermutation(chars)
	}
	ans = strings.Join(chars, "")
	return ans
}

type Chars []string
func (h Chars) Len() int           { return len(h) }
func (h Chars) Less(i, j int) bool { return h[i] < h[j] }
func (h Chars) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func nextPermutation(s sort.Interface) bool {
	n := s.Len()
	if n < 2 {
		return false
	}
	i := n - 2
	for ; !s.Less(i, i+1); i-- {
		if i == 0 {
			return false
		}
	}
	j := n - 1
	for !s.Less(i, j) {
		j--
	}
	s.Swap(i, j)
	for k, l := i+1, n-1; k < l; {
		s.Swap(k, l)
		k++
		l--
	}
	return true
}

func main() {
	io := NewIO()
	defer io.Output()

	s, k := io.Read(), io.Int()

	io.Write(run(s, k))
}

type IO struct {
	sc  *bufio.Scanner
	out *bufio.Writer
}

func NewIO() *IO {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	bufSize := math.MaxInt64
	sc.Buffer([]byte{}, bufSize)

	out := bufio.NewWriter(os.Stdout)

	return &IO{
		sc:  sc,
		out: out,
	}
}

func (io *IO) Read() string {
	ok := io.sc.Scan()
	if !ok {
		panic(io.sc.Err())
	}
	return io.sc.Text()
}

func (io *IO) Write(val ...interface{}) {
	for _, v := range val {
		if _, err := fmt.Fprintln(io.out, v); err != nil {
			panic(err)
		}
	}
}

func (io *IO) Output() {
	err := io.out.Flush()
	if err != nil {
		panic(err)
	}
}

func (io *IO) Int() int {
	i, err := strconv.Atoi(io.Read())
	if err != nil {
		panic(err)
	}
	return i
}

func (io *IO) IntN(n int) []int {
	intN := make([]int, n)
	for i := 0; i < n; i++ {
		intN[i] = io.Int()
	}
	return intN
}
