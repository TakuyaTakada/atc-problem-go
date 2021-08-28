package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func run(n, q int, a, k []int) []int {
	ans := make([]int, q)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = a[i] - (i + 1)
	}
	for i := 0; i < q; i++ {
		if k[i] > c[n-1] {
			ans[i] = k[i] + n
		} else {
			ok, ng := n-1, -1
			for abs(ok-ng) > 1 {
				mid := (ok + ng) / 2
				if c[mid] >= k[i] {
					ok = mid
				} else {
					ng = mid
				}
			}
			ans[i] = a[ok] - 1 - (c[ok] - k[i])
		}
	}
	return ans
}

func main() {
	io := NewIO()
	defer io.Output()

	n, q := io.Int(), io.Int()
	a := io.IntN(n)
	k := io.IntN(q)

	ans := run(n, q, a, k)
	for _, v := range ans {
		io.Write(v)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
