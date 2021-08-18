package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func run(s string) string {
	var ans string
	for i := len(s); i > 0; i-- {
		if s[i-1:i] == "6" {
			ans += "9"
		} else if s[i-1:i] == "9" {
			ans += "6"
		} else {
			ans += s[i-1:i]
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
