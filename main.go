package main

import (
	"fmt"
	"io"
	"os"
)

type SrcFile struct {
	Name   string
	File   *os.File
	Lineno int
	Charno int
}

func (s *SrcFile) nextByte() (byte, error) {
	ch := make([]byte, 1)

	_, err := s.File.Read(ch)
	if err == io.EOF {
		return ch[0], err
	} else if err != nil {
		panic(err)
	}

	s.Charno++
	if ch[0] == '\n' {
		s.Lineno++
		s.Charno = 1
	}

	return ch[0], nil
}

func (s *SrcFile) printStat() {
	fmt.Printf("Line number: %d\tChar number: %d\n", s.Lineno, s.Charno)
}

func main() {
	file, err := os.Open("test.c")
	if err != nil {
		panic(err)
	}

	src1 := SrcFile{File: file, Lineno: 1}
	for {
		ch, err := src1.nextByte()
		if err == io.EOF {
			fmt.Println("Done!")
			break
		}

		fmt.Printf("%d %s\n", ch, string(ch))
		src1.printStat()
	}
}
