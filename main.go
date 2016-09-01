package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//var bracketStack = 0
var braceStack = 0

func main() {
	err := setUp()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer cleanUp()
	mainloop()
}

func mainloop() {
	reader := bufio.NewReader(os.Stdin)
	// init current buffer
	curBuffer := increBuffer
	for {
		fmt.Print(">>> ")
		str, _ := reader.ReadString('\n')
		if strings.Trim(str, " \n") == "quit" {
			break
		}
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
