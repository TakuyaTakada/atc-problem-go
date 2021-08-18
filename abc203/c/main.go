package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func run(n, k int, ab IntHeap) int {
	var ans int
	sort.Sort(ab)
	for i := 0; i < n; i++ {
		if ab[i].a > k {
			break
		}
		k += ab[i].b
	}
	ans = k
	return ans
}

func main() {
	io := NewIO()
	defer io.Output()

	n, k := io.Int(), io.Int()

	ab := make(IntHeap, n)
	for i := 0; i < n; i++ {
		ab[i] = Friend{
			a: io.Int(),
			b: io.Int(),
		}
	}

	io.Write(run(n, k, ab))
}

type Friend struct {
	a int
	b int
}

type IntHeap []Friend
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].a < h[j].a }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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
