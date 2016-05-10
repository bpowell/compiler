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

func (s *SrcFile) rewind() error {
	_, err := s.File.Seek(1, 1)
	return err
}

func (s *SrcFile) nextToken() (string, error) {
	var token []byte
	for {
		ch, err := s.nextByte()
		if err == io.EOF {
			return string(token), err
		}

		switch ch {
		case '\n',
			' ':
			return string(token), nil
		}

		token = append(token, ch)
	}

	return string(token), nil
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
		token, err := src1.nextToken()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(token)
		src1.printStat()
	}
}
