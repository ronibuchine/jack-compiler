package syntaxAnalyzer

import (
	"bytes"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	var b bytes.Buffer
	BuildXML(Parse(Tokenize("test2.jack")), &b)
	fmt.Print(b.String())
}

func TestTokenize(t *testing.T) {
	var b bytes.Buffer
	TokenToXML(Tokenize("test2.jack"), &b)
	fmt.Print(b.String())
}
