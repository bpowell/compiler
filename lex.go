package main

import (
	"fmt"
	"unicode/utf8"
)

const (
	TOK_NUMBER int = iota
	TOK_DECIMAL
)

const eof = -1

type Token struct {
	Tok int
	Val string
}

type stateFn func(*lexer) stateFn

type lexer struct {
	name   string  // name of file
	input  string  // file data
	pos    int     // current position in input
	start  int     // start of current token
	width  int     // width of last rune from input
	state  stateFn // next state to go to
	tokens []Token // list of all tokens
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.rewind()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) rewind() {
	l.pos -= l.width
}

// emit a token to the lexer
func (l *lexer) emit(t int) {
	l.tokens = append(l.tokens, Token{t, l.input[l.start:l.pos]})
}

func (l *lexer) run() {
	for l.state = startState; l.state != nil; {
		l.state = l.state(l)
	}
}

func startState(l *lexer) stateFn {
	fmt.Println("startState")
	return nil
}
