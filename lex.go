package main

import (
	"fmt"
	"unicode/utf8"
)

const (
	TOK_NUMBER int = iota
	TOK_DECIMAL
	TOK_PLUS
	TOK_MINUS
	TOK_MULTIPLY
	TOK_DIVIDE
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
	if l.input[l.start] == ' ' {
		l.start += 1
	}

	l.tokens = append(l.tokens, Token{t, l.input[l.start:l.pos]})
	l.start = l.pos
}

func (l *lexer) run() {
	for l.state = startState; l.state != nil; {
		l.state = l.state(l)
	}
}

func isNumeric(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isAlphaNumeric(ch rune) bool {
	return isAlpha(ch) || isNumeric(ch)
}

func isOperator(ch rune) bool {
	switch ch {
	case '+', '-', '*', '/':
		return true
	}

	return false
}

func (l *lexer) eat() {
	for {
		if !isAlphaNumeric(l.peek()) {
			break
		}

		l.next()
	}
}

func errorState(l *lexer) stateFn {
	l.eat()
	fmt.Printf("Invalid token %s at position %d\n", l.input[l.start:l.pos], l.start)
	l.start = l.pos

	return startState
}

func decimalState(l *lexer) stateFn {
	defer l.next()

	switch ch := l.peek(); {
	case ch == eof:
		l.emit(TOK_DECIMAL)
		return nil
	case isNumeric(ch):
		return decimalState
	case isSpace(ch):
		l.emit(TOK_DECIMAL)
		return startState
	case isOperator(ch):
		l.emit(TOK_DECIMAL)
		l.rewind()
		return startState
	default:
		return errorState
	}

	return nil
}

func numericState(l *lexer) stateFn {
	defer l.next()

	switch ch := l.peek(); {
	case ch == eof:
		l.emit(TOK_NUMBER)
		return nil
	case ch == '.':
		return decimalState
	case isNumeric(ch):
		return numericState
	case isSpace(ch):
		l.emit(TOK_NUMBER)
		return startState
	case isOperator(ch):
		l.emit(TOK_NUMBER)
		l.rewind()
		return startState
	default:
		return errorState
	}

	return nil
}

func plusState(l *lexer) stateFn {
	defer l.next()

	switch ch := l.peek(); {
	case ch == eof:
		l.emit(TOK_PLUS)
		return nil
	case isNumeric(ch) || isSpace(ch):
		l.emit(TOK_PLUS)
		return startState
	default:
		return errorState
	}

	return nil
}

func minusState(l *lexer) stateFn {
	defer l.next()

	switch ch := l.peek(); {
	case ch == eof:
		l.emit(TOK_MINUS)
		return nil
	case isNumeric(ch) || isSpace(ch):
		l.emit(TOK_MINUS)
		return startState
	default:
		return errorState
	}

	return nil
}

func multiplyState(l *lexer) stateFn {
	defer l.next()

	switch ch := l.peek(); {
	case ch == eof:
		l.emit(TOK_MULTIPLY)
		return nil
	case isNumeric(ch) || isSpace(ch):
		l.emit(TOK_MULTIPLY)
		return startState
	default:
		return errorState
	}

	return nil
}

func divideState(l *lexer) stateFn {
	defer l.next()

	switch ch := l.peek(); {
	case ch == eof:
		l.emit(TOK_DIVIDE)
		return nil
	case isNumeric(ch) || isSpace(ch):
		l.emit(TOK_DIVIDE)
		return startState
	default:
		return errorState
	}

	return nil
}

func startState(l *lexer) stateFn {
	switch ch := l.next(); {
	case ch == eof:
		return nil
	case isNumeric(ch):
		return numericState
	case isSpace(ch):
		return startState
	case ch == '+':
		return plusState
	case ch == '-':
		return minusState
	case ch == '*':
		return multiplyState
	case ch == '/':
		return divideState
	default:
		return errorState
	}

	return nil
}
