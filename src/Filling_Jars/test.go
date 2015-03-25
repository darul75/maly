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
	numJars     int
	numOps      int
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
	inputScanner.Scan()
	s := inputScanner.Text()
	infos := strings.Split(s, " ")
	num1, _ := strconv.Atoi(infos[0])
	num2, _ := strconv.Atoi(infos[1])
	info.numJars = num1
	info.numOps = num2

	info.LinesString = make([]LineString, num2)
	info.LinesInt = make([]LineInt, num2)

	for i := 0; inputScanner.Scan(); i++ {
		s = inputScanner.Text()
		info.LinesString[i] = LineString(s)
	}
}

func (info *Info) Play() {
	doStuff(info)
}

// EXERCICE STUFF

func doStuff(info *Info) {

	fmt.Println(len(info.LinesString))
	fmt.Println(info.numJars)
	fmt.Println(info.numOps)

	lines := info.LinesString[0:]

	sum := 0

	for i := 0; i < len(lines); i++ {
		nums := strings.Split(string(lines[i]), " ")
		val1, _ := strconv.Atoi(nums[0])
		val2, _ := strconv.Atoi(nums[1])
		val3, _ := strconv.Atoi(nums[2])
		sum += (val2 - val1 + 1) * (val3)
	}

	fmt.Println("%v", int64(math.Floor(float64(sum/info.numJars))))

}
