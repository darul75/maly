package main

import (
	"bufio"
	"fmt"
	/*"io"*/
	"log"
	/*"math"*/
	"os"
	/*"sort"*/
	"strconv"
	"strings"
)

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

func (info *Info) ReadStdIn() {
	/*fmt.Scan(&info.Num)
	info.LinesString = make([]LineString, info.Num)
	for i := 0; i < info.Num; i++ {
		fmt.Scan(&info.LinesString[i])
	}*/
}

func (info *Info) ReadInputFile(input string) {
	// test case
	in, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}
	inputScanner := bufio.NewScanner(in)
	inputScanner.Split(bufio.ScanLines)
	info.LinesString = make([]LineString, 101)
	info.LinesInt = make([]LineInt, 101)

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

	//mask, _ := strconv.ParseInt("11111", 2, 8)

	var max int64 = 0
	team := 0

	for i := 0; i < len(lines)-1; i++ {
		fmt.Println(lines[i])
		st1, _ := strconv.ParseInt(string(lines[i]), 2, 64)
		fmt.Println(st1)
		for j := i + 1; j < len(lines); j++ {
			st2, _ := strconv.ParseInt(string(lines[j]), 2, 64)
			val := st1 | st2
			//fmt.Println(strconv.FormatInt(val, 2))
			if val > max {
				max = val
				team = 1
			} else if val == max {
				team += 1
			}
		}
	}

	maxBinaryStr := strconv.FormatInt(max, 2)
	fmt.Println(strings.Count(maxBinaryStr, "1"))
	fmt.Println(team)

	// compute, find best unfairness

	//fmt.Println(nums)

	/*for _, v := range info.LinesString {
		fmt.Println(v)
	}*/

}
