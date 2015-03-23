package main

import (
	"bufio"
	"fmt"
	/*"io"*/
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	/*"sort"*/ /*"strconv"*/ /*"strings"*/)

// INTERFACE

type Hack interface {
	Play()
}

type LineString string
type LineInt int

type Info struct {
	LinesString []LineString
	LinesInt    []LineInt
}

func main() {

	// program arguments
	args := os.Args[1:]
	// stdin
	info := &Info{}

	if len(args) == 2 {
		info.ReadInputFile(args[0])
	}

	info.Play()
}

func (info *Info) ReadInputFile(input string) {
	// test case
	in, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}
	inputScanner := bufio.NewScanner(in)
	inputScanner.Split(bufio.ScanLines)
	info.LinesString = make([]LineString, 4)
	info.LinesInt = make([]LineInt, 4)

	for i := 0; inputScanner.Scan(); i++ {
		s := inputScanner.Text()
		/*num, err := strconv.Atoi(s)
		if err == nil {
			info.LinesInt[i] = LineInt(num)
		} else {*/
		info.LinesString[i] = LineString(s)

	}
}

func (info *Info) Play() {
	doStuff(info)
}

// EXERCICE STUFF

func doStuff(info *Info) {

	lines := info.LinesString[1:]

	/*n := 5*/
	//m := 3

	type Operation struct {
		firstIdx  int
		secondIdx int
		candies   int
	}

	ops := make([]Operation, 3)
	jars := make([]int, 5)

	for i := 0; i < len(lines); i++ {
		nums := strings.Split(string(lines[i]), " ")
		val1, _ := strconv.Atoi(nums[0])
		val2, _ := strconv.Atoi(nums[1])
		val3, _ := strconv.Atoi(nums[2])
		op := Operation{val1, val2, val3}
		ops[i] = op
	}

	fmt.Println("... %v", ops)

	for _, op := range ops {
		for i := op.firstIdx - 1; i < op.secondIdx; i++ {
			jars[i] += op.candies
		}
	}

	// sum
	var sum int = 0
	for _, v := range jars {
		sum += v
	}

	fmt.Println("... %v", math.Floor(float64(sum/5)))

}
