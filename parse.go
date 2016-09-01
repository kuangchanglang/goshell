package main

import "regexp"

// statement type
const (
	stImport       = iota // import "fmt"
	stAssign              // a := b
	stLeftBracket         // (
	stRightBracket        // )
	stLeftBrace           // {
	stRightBrace          // }
	stFunc                // func
	stEmptyStr            // empty string
	stDefault
)

// regexp for each statement type
const (
	reImport       = "^\\s*import\\s+"
	reAssign       = "^[^\"']+=.+|^[^\"']+:=.+|^\\s*var\\s*.+"
	reLeftBracket  = ".*(\\s*$"
	reRightBracket = "^\\s*)\\s*$"
	reLeftBrace    = ".*{\\s*$"
	reRightBrace   = "^\\s*}\\s*$"
	reFunc         = "^\\s*func\\s+"
	reEmptyStr     = "^\\s*$"
)

func statementType(s string) int {
	m, err := regexp.MatchString(reEmptyStr, s)
	if err == nil && m {
		return stEmptyStr
	}

	m, err = regexp.MatchString(reImport, s)
	if err == nil && m {
		return stImport
	}

	m, err = regexp.MatchString(reLeftBracket, s)
	if err == nil && m {
		return stLeftBracket
	}

	m, err = regexp.MatchString(reRightBracket, s)
	if err == nil && m {
		return stRightBracket
	}

	m, err = regexp.MatchString(reLeftBrace, s)
	if err == nil && m {
		m, err = regexp.MatchString(reFunc, s)
		if err == nil && m {
			return stFunc
		}
		return stLeftBrace
	}

	m, err = regexp.MatchString(reRightBrace, s)
	if err == nil && m {
		return stRightBrace
	}

	m, err = regexp.MatchString(reAssign, s)
	if err == nil && m {
		return stAssign
	}

	return stDefault
}
