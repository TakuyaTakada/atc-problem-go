package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func run(n int, s, t []int) []int {
	for i := 0; i < n*2; i++ {
		t[(i+1)%n] = min(t[(i+1)%n], t[i%n] + s[i%n])
	}
	return t
}

func main() {
	io := NewIO()
	defer io.Output()

	n := io.Int()
	s := io.IntN(n)
	t := io.IntN(n)

	result := run(n, s, t)
	for _, r := range result {
		io.Write(r)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
