package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func run(s string) int {
	var ans int
	obit := uint(0)
	xbit := uint(0)
	for i := 0; i < 10; i++ {
		if s[i:i+1] == "o" {
			obit |= 1<<i
		}
		if s[i:i+1] == "x" {
			xbit |= 1<<i
		}
	}

	for i := 0; i < 10000; i++ {
		bit := uint(0)
		n := i
		for j := 0; j < 4; j++ {
			bit |= 1<<(n%10)
			n /= 10
		}
		ok := true
		if bit & obit != obit || ^bit & xbit != xbit {
			ok = false
		}
		if ok {
			ans++
		}
	}

	return ans
}

func main() {
	io := NewIO()
	defer io.Output()

	s := io.Read()

	io.Write(run(s))
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
