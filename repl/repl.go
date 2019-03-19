package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexedText := lexer.New(line)

		for tok := lexedText.NextToken(); tok.Type != token.EOF; tok = lexedText.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
