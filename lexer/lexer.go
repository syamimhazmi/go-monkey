package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current character)
	readPosition int  // current reading position in input (after current character)
	ch           byte // current char under exemination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readCharacter()
	return l
}

// for the sake of simplicity, readCharacter method only support ASCII characters
// readCharacter method init character byte based on input readPosition
// and increase readPosition into 1(which mean check on the next character)
func (l *Lexer) readCharacter() {
	if l.readPosition >= len(l.input) {
		// set current char into 0 to indicate that we are at the end of file
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	// change current position in input with current reading position input
	l.position = l.readPosition
	// we need to add one to change our current reading position for the next char
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekCharacter() == '=' {
			tok = makeTwoCharacterToken(l, token.EQ)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	case '(':
		tok = newToken(token.LPAREN, l.ch)

	case ')':
		tok = newToken(token.RPAREN, l.ch)

	case ',':
		tok = newToken(token.COMMA, l.ch)

	case '+':
		tok = newToken(token.PLUS, l.ch)

	case '-':
		tok = newToken(token.MINUS, l.ch)

	case '*':
		tok = newToken(token.ASTERISK, l.ch)

	case '/':
		tok = newToken(token.SLASH, l.ch)

	case '!':
		if l.peekCharacter() == '=' {
			tok = makeTwoCharacterToken(l, token.NOT_EQ)
		} else {
			tok = newToken(token.BANG, l.ch)
		}

	case '<':
		tok = newToken(token.LT, l.ch)

	case '>':
		tok = newToken(token.GT, l.ch)

	case '{':
		tok = newToken(token.LBRACE, l.ch)

	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifer()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readCharacter()
	return tok
}

func (l *Lexer) readIdentifer() string {
	position := l.position
	for isLetter(l.ch) {
		l.readCharacter()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readCharacter()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekCharacter() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readCharacter()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func makeTwoCharacterToken(lexer *Lexer, tokenType token.TokenType) token.Token {
	ch := lexer.ch
	lexer.readCharacter()
	literal := string(ch) + string(lexer.ch)
	return token.Token{Type: tokenType, Literal: literal}
}
