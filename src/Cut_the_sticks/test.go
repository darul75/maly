package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	numSticks   int
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
	info.numSticks = num1

	info.LinesString = make([]LineString, 1)
	info.LinesInt = make([]LineInt, 1)

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

	line := info.LinesString[0]

	infos := strings.Split(string(line), " ")

	sticks := make([]int, len(infos))

	for i := 0; i < len(infos); i++ {
		val1, _ := strconv.Atoi(infos[i])
		sticks[i] = val1
	}

	fmt.Println(len(sticks))

	cut := 0
	done := false

	for !done {
		// ascending
		sort.Ints(sticks)
		cut = sticks[0]
		slice := 0

		for i, v := range sticks {
			if cut == v {
				slice++
				continue
			}
			sticks[i] = sticks[i] - cut
		}

		sticks = sticks[slice:]

		if len(sticks) == 0 {
			done = true
		} else {
			fmt.Println(len(sticks))
		}

	}

}
