package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func run(s, t int) int {
	var ans int
	for a := 0; a <= s; a++ {
		for b := 0; b <= s; b++ {
			for c := 0; c <= s; c++ {
				if a + b + c <= s && a * b * c <= t {
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	io := NewIO()
	defer io.Output()

	s, t := io.Int(), io.Int()

	io.Write(run(s, t))
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
