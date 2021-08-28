package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func run(a, b, c int) string {
	var ans string
	if c % 2 == 0 {
		aabs := abs(a)
		babs := abs(b)
		if aabs == babs {
			ans = "="
		} else if aabs < babs {
			ans = "<"
		} else {
			ans = ">"
		}
	} else {
		if a == b {
			ans = "="
		} else if a < b {
			ans = "<"
		} else {
			ans = ">"
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	io := NewIO()
	defer io.Output()

	a, b, c := io.Int(), io.Int(), io.Int()

	io.Write(run(a, b, c))
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
