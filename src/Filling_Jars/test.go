package main

import (
	"bufio"
	"fmt"
	/*"io"*/
	"log"
	"os"
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

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])

	}

}
