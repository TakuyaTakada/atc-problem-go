package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func run(ms Mountains) string {
	var ans string
	sort.Sort(ms)
	ans = ms[len(ms)-2].s
	return ans
}

type Mount struct {
	s string
	t int
}

type Mountains []*Mount
func (h Mountains) Len() int           { return len(h) }
func (h Mountains) Less(i, j int) bool { return h[i].t < h[j].t }
func (h Mountains) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func main() {
	io := NewIO()
	defer io.Output()

	n := io.Int()
	mt := make(Mountains, n)
	for i := 0; i < n; i++ {
		m := &Mount{
			s: io.Read(),
			t: io.Int(),
		}
		mt[i] = m
	}

	io.Write(run(mt))
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
