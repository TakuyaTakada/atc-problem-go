package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func run(n int, a, b, c []int) int {
	var ans int
	count := make([]int, n, n)
	for j := 0; j < n; j++ {
		count[b[c[j]-1]-1]++
	}

	for i := 0; i < n; i++ {
		ans += count[a[i]-1]
	}
	return ans
}

func main() {
	io := NewIO()
	defer io.Output()

	n := io.Int()
	a, b, c := io.IntN(n), io.IntN(n), io.IntN(n)

	io.Write(run(n, a, b, c))
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
