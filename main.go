package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/chzyer/readline"
)

//var bracketStack = 0
var braceStack = 0

func main() {
	err := setUp()
	if err != nil {
		log.Fatal("setup error: %v", err)
		return
	}
	defer cleanUp()
	mainloop()
}

func mainloop() {
	// init current buffer
	curBuffer := increBuffer

	fmt.Println("Goshell v1.0 (2016-9-1 17:28:02)")
	fmt.Println(`Type "quit" to exit`)
	rl, err := readline.New(">>> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	for {
		//str, _ := reader.ReadString('\n')
		str, err := rl.Readline()
		if err != nil {
			log.Fatal("read line error: %v", err)
		}

		if strings.Trim(str, " \n") == "quit" {
			break
		}
		str += "\n"
		switch statementType(str) {
		case stImport: // import
			importBuffer.WriteString(str)
		/*
			case ST_TYPE_LEFT_BRACKET:
				curBuffer.WriteString(str)
				bracketStack += 1
			case ST_TYPE_RIGHT_BRACKET:
				curBuffer.WriteString(str)
				bracketStack -= 1
				if bracketStack <= 0 {
					bracketStack = 0
					run()
				}
		*/
		case stLeftBrace: // '{' found at the end
			curBuffer.WriteString(str)
			braceStack++
		case stRightBrace: // '}' found
			curBuffer.WriteString(str)
			braceStack--
			if braceStack <= 0 {
				braceStack = 0
				curBuffer = increBuffer
				execute()
			}
		case stAssign:
			curBuffer.WriteString(str)
		case stFunc:
			// begin func
			braceStack++
			curBuffer = funcBuffer
			curBuffer.WriteString(str)
		case stEmptyStr:
			// do nothing
		default:
			curBuffer.WriteString(str)
			if braceStack <= 0 {
				execute()
			}
		}
	}

}
