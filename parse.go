package main

import "regexp"

// statement type
const (
	ST_TYPE_IMPORT        = iota // import "fmt"
	ST_TYPE_ASSIGN               // a := b
	ST_TYPE_LEFT_BRACKET         // (
	ST_TYPE_RIGHT_BRACKET        // )
	ST_TYPE_LEFT_BRACE           // {
	ST_TYPE_RIGHT_BRACE          // }
	ST_TYPE_FUNC                 // func
	ST_TYPE_EMPTY_STR            // empty string
	ST_TYPE_DEFAULT
)

// regexp for each statement type
const (
	RE_IMPORT        = "^\\s*import\\s+"
	RE_ASSIGN        = "^[^\"']+=.+|^[^\"']+:=.+|^\\s*var\\s*.+"
	RE_LEFT_BRACKET  = ".*(\\s*$"
	RE_RIGHT_BRACKET = "^\\s*)\\s*$"
	RE_LEFT_BRACE    = ".*{\\s*$"
	RE_RIGHT_BRACE   = "^\\s*}\\s*$"
	RE_FUNC          = "^\\s*func\\s+"
	RE_EMPTY_STR     = "^\\s*$"
)

func statementType(s string) int {
	m, err := regexp.MatchString(RE_EMPTY_STR, s)
	if err == nil && m {
		return ST_TYPE_EMPTY_STR
	}

	m, err = regexp.MatchString(RE_IMPORT, s)
	if err == nil && m {
		return ST_TYPE_IMPORT
	}

	m, err = regexp.MatchString(RE_LEFT_BRACKET, s)
	if err == nil && m {
		return ST_TYPE_LEFT_BRACKET
	}

	m, err = regexp.MatchString(RE_RIGHT_BRACKET, s)
	if err == nil && m {
		return ST_TYPE_RIGHT_BRACKET
	}

	m, err = regexp.MatchString(RE_LEFT_BRACE, s)
	if err == nil && m {
		m, err = regexp.MatchString(RE_FUNC, s)
		if err == nil && m {
			return ST_TYPE_FUNC
		}
		return ST_TYPE_LEFT_BRACE
	}

	m, err = regexp.MatchString(RE_RIGHT_BRACE, s)
	if err == nil && m {
		return ST_TYPE_RIGHT_BRACE
	}

	m, err = regexp.MatchString(RE_ASSIGN, s)
	if err == nil && m {
		return ST_TYPE_ASSIGN
	}

	return ST_TYPE_DEFAULT
}
