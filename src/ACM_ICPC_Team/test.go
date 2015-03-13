package main

import (
	"bufio"
	"fmt"
	/*"io"*/
	"log"
	"math/big"
	"os"
	/*"sort"*/
	/*"strconv"*/ /*"strings"*/)

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

	var bigNum1 big.Int
	var bigNum2 big.Int
	var result big.Int

	var max int64 = 0
	team := 0

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			bits := int64(0)

			if _, ok := bigNum1.SetString(string(lines[i]), 2); !ok {
				panic("invalid num1")
			}
			if _, ok := bigNum2.SetString(string(lines[j]), 2); !ok {
				panic("invalid num2")
			}
			result.Or(&bigNum1, &bigNum2)

			for i := result.BitLen() - 1; i >= 0; i-- {
				bits += int64(result.Bit(i))
			}

			if bits > max {
				max = bits
				team = 1
			} else if bits == max {
				team += 1
			}

		}
	}

	fmt.Println(max)
	fmt.Println(team)
}
