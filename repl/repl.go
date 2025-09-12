package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/moges7624/MonkeyInterpreter/evaluator"
	"github.com/moges7624/MonkeyInterpreter/lexer"
	"github.com/moges7624/MonkeyInterpreter/parser"
	"github.com/moges7624/MonkeyInterpreter/token"
)

const PROMPT = ">> "
const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

// returns the tokens of the input program
func Start(in io.Reader, out io.Writer) {
  scanner := bufio.NewScanner(in)

  for {
    fmt.Fprintf(out, PROMPT)
    scanned := scanner.Scan()
    if !scanned {
      return
    }

    line := scanner.Text()
    l := lexer.New(line)

    for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
      fmt.Fprintf(out, "%+v\n", tok)
    }
  }
}

// returns the ast of the input program
func StartParse(in io.Reader, out io.Writer) {
  io.WriteString(out, MONKEY_FACE)
  scanner := bufio.NewScanner(in)

  for  {
    fmt.Fprintf(out, PROMPT)
    scanned := scanner.Scan()
    if !scanned {
      return
    }

    line := scanner.Text()
    l := lexer.New(line)
    p := parser.New(l)

    program := p.ParseProgram()
    if len(p.Errors()) != 0 {
      printParserErrors(out, p.Errors())
      continue
    }

    io.WriteString(out, program.String())
    io.WriteString(out, "\n")
  }
}


// returns the evaluated response of the input program
func StartEvaluate(in io.Reader, out io.Writer) {
  io.WriteString(out, MONKEY_FACE)
  scanner := bufio.NewScanner(in)

  for  {
    fmt.Fprintf(out, PROMPT)
    scanned := scanner.Scan()
    if !scanned {
      return
    }

    line := scanner.Text()
    l := lexer.New(line)
    p := parser.New(l)

    program := p.ParseProgram()
    if len(p.Errors()) != 0 {
      printParserErrors(out, p.Errors())
      continue
    }

    evaluated := evaluator.Eval(program)
    if evaluated != nil {
      io.WriteString(out, evaluated.Inspect())
      io.WriteString(out, "\n")
    }
  }
}

func printParserErrors(out io.Writer, errors []string) {
  io.WriteString(out, "Woops! We ran into some monkey business here!\n")
  io.WriteString(out, " parser errors:\n")
  for _, msg := range errors {
    io.WriteString(out, "\t"+msg+"\n")
  }
}
