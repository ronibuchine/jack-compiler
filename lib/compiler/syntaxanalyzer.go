package compiler

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// aliasing structs for xml objects
type Tokens struct {
	XMLName xml.Name `xml:"tokens"`
	Tokens  []Token  `xml:"tokens"`
}

type Token struct {
	XMLName xml.Name `xml:"token"`
	Type    string   `xml:"type,attr"`
	Token   string   `xml:"token"`
}

// Node struct for parsing and recursive descent
type Node struct {
	name     string
	children []Node
}

// globals for matching
var normalizedTokenStream [][]string
var current int = 0

// Takes an ordered token stream which is a map[string]string and parses it and returns the XML tree
func BuildXML(tokenStream map[string]string) *os.File {
	output, err := os.Create("output.xml")
	if err != nil {
		log.Fatal("Failed to create an ouput file.")
	}
	defer output.Close()
	// TODO: implement

	return output
}

// Reads in the token stream form the XML file, returns as a list of tokens to the global token stream
func ReadStream(tokenStream string) {
	input, err := os.Open(tokenStream)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	byteValue, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal("Error reading the bytes of the XML file.")
	}

	var tokens Tokens
	xml.Unmarshal(byteValue, &tokens)
	for _, token := range tokens.Tokens {
		fmt.Println("Type: " + token.Type)
		fmt.Println("XMLName: " + token.XMLName.Local)
	}
	getTokenStrings(tokens.Tokens)
}

// returns an array of tuples, first value in the tuple is the token, second value is the token type
func getTokenStrings(tokens []Token) {
	for i := 0; i < len(tokens); i++ {
		normalizedTokenStream[i][0] = tokens[i].Token
		normalizedTokenStream[i][1] = tokens[i].Type
	}
}

// global variable used for token parsing and matching
func match(token string) Node {
	if current >= len(normalizedTokenStream) {
		log.Fatal("end of token stream")
	}
	if token == normalizedTokenStream[current][0] {
		current++
		return Node{token, []Node{}}
	} else {
		return Node{"", []Node{}}
	}
}

func (parent Node) addChild(child Node) {
	parent.children = append(parent.children, child)
}

// functions for grammar
func letStatement() {

}

func whileStatement() {

}

func ifStatement() {

}

func doStatement() {

}

func statement() {

}

func statements() {

}

func expression() {

}

func term() {

}

func varName() {

}

func constant() Node {
	result := Node{"constant", []Node{}}
	value, _ := strconv.Atoi(normalizedTokenStream[current][0])
	if value >= 0 || value < 32000 {
		result.addChild(match(normalizedTokenStream[current][0]))
	} else {
		// TODO: error
	}
	return result
}

func op() Node {
	result := Node{"op", []Node{}}
	if normalizedTokenStream[current][0] == "+" {
		result.addChild(match("+"))
	} else if normalizedTokenStream[current][0] == "-" {
		result.addChild(match("-"))
	} else if normalizedTokenStream[current][0] == "=" {
		result.addChild(match("="))
	} else {
		// TODO:  error logging
	}
	return result
}

func relop() {
	if normalizedTokenStream[current][0] == "<" {

	} else if normalizedTokenStream[current][0] == ">" {

	} else {
		// TODO: error logging
	}
}

func boolOp() {
	if normalizedTokenStream[current][0] == "&" {

	} else if normalizedTokenStream[current][0] == "|" {

	} else {
		// TODO: error logging
	}
}

func unaryOp() {
	if normalizedTokenStream[current][0] == "-" {

	} else {
		// TODO: error logging
	}
}

func boolUnaryOp() {
	if normalizedTokenStream[current][0] == "~" {

	} else {
		// TODO: error logging
	}
}
