package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var A = "123"
var a = "123"

func PrintIndexValue(rangeValue []int) {
	for index, value := range rangeValue {
		fmt.Println(index, value)
	}
}

func ReadFromBuffer() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func GenerateGIF() {
	//var palette = []color.Color{color.White, color.Black}
	const (
		whiteIndex = 0 // first color in palette
		blackIndex = 1 // next color in palette
	)

}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.01 // first
		size    = 100
		nframes = 54
		delay   = 8
	)

}
