package repl

import (
	"bufio"
	"fmt"
	"github.com/Zhalkhas/monkeyLang/lexer"
	"github.com/Zhalkhas/monkeyLang/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanned := scanner.Scan()
	if !scanned {
		return
	}
	l := lexer.New(scanner.Text())
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
