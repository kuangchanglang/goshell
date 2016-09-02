package readline

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

var (
	isWindows = false
)

// WaitForResume need to call before current process got suspend.
// It will run a ticker until a long duration is occurs,
// which means this process is resumed.
func WaitForResume() chan struct{} {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(10 * time.Millisecond)
		t := time.Now()
		wg.Done()
		for {
			now := <-ticker.C
			if now.Sub(t) > 100*time.Millisecond {
				break
			}
			t = now
		}
		ticker.Stop()
		ch <- struct{}{}
	}()
	wg.Wait()
	return ch
}

func Restore(fd int, state *State) error {
	err := restoreTerm(fd, state)
	if err != nil {
		// errno 0 means everything is ok :)
		if err.Error() == "errno 0" {
			err = nil
		}
	}
	return nil
}

func IsPrintable(key rune) bool {
	isInSurrogateArea := key >= 0xd800 && key <= 0xdbff
	return key >= 32 && !isInSurrogateArea
}

// translate Esc[X
func escapeExKey(key *escapeKeyPair) rune {
	var r rune
	switch key.typ {
	case 'D':
		r = CharBackward
	case 'C':
		r = CharForward
	case 'A':
		r = CharPrev
	case 'B':
		r = CharNext
	case 'H':
		r = CharLineStart
	case 'F':
		r = CharLineEnd
	case '~':
		if key.attr == "3" {
			r = CharDelete
		}
	default:
	}
	return r
}

type escapeKeyPair struct {
	attr string
	typ  rune
}

func (e *escapeKeyPair) Get2() (int, int, bool) {
	sp := strings.Split(e.attr, ";")
	if len(sp) < 2 {
		return -1, -1, false
	}
	s1, err := strconv.Atoi(sp[0])
	if err != nil {
		return -1, -1, false
	}
	s2, err := strconv.Atoi(sp[1])
	if err != nil {
		return -1, -1, false
	}
	return s1, s2, true
}

func readEscKey(r rune, reader *bufio.Reader) *escapeKeyPair {
	p := escapeKeyPair{}
	buf := bytes.NewBuffer(nil)
	for {
		if r == ';' {
		} else if unicode.IsNumber(r) {
		} else {
			p.typ = r
			break
		}
		buf.WriteRune(r)
		r, _, _ = reader.ReadRune()
	}
	p.attr = buf.String()
	return &p
}

// translate EscX to Meta+X
func escapeKey(r rune, reader *bufio.Reader) rune {
	switch r {
	case 'b':
		r = MetaBackward
	case 'f':
		r = MetaForward
	case 'd':
		r = MetaDelete
	case CharTranspose:
		r = MetaTranspose
	case CharBackspace:
		r = MetaBackspace
	case 'O':
		d, _, _ := reader.ReadRune()
		switch d {
		case 'H':
			r = CharLineStart
		case 'F':
			r = CharLineEnd
		default:
			reader.UnreadRune()
		}
	case CharEsc:

	}
	return r
}

func SplitByLine(start, screenWidth int, rs []rune) []string {
	var ret []string
	buf := bytes.NewBuffer(nil)
	currentWidth := start
	for _, r := range rs {
		w := runes.Width(r)
		currentWidth += w
		buf.WriteRune(r)
		if currentWidth >= screenWidth {
			ret = append(ret, buf.String())
			buf.Reset()
			currentWidth = 0
		}
	}
	ret = append(ret, buf.String())
	return ret
}

// calculate how many lines for N character
func LineCount(screenWidth, w int) int {
	r := w / screenWidth
	if w%screenWidth != 0 {
		r++
	}
	return r
}

func IsWordBreak(i rune) bool {
	switch {
	case i >= 'a' && i <= 'z':
	case i >= 'A' && i <= 'Z':
	case i >= '0' && i <= '9':
	default:
		return true
	}
	return false
}

func GetInt(s []string, def int) int {
	if len(s) == 0 {
		return def
	}
	c, err := strconv.Atoi(s[0])
	if err != nil {
		return def
	}
	return c
}

type RawMode struct {
	state *State
}

func (r *RawMode) Enter() (err error) {
	r.state, err = MakeRaw(GetStdin())
	return err
}

func (r *RawMode) Exit() error {
	if r.state == nil {
		return nil
	}
	return Restore(GetStdin(), r.state)
}