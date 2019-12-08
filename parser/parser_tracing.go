package parser

import (
	"fmt"
	"strings"
)

var Enable = false
var traceLevel int

const traceIdentPlaceholder string = "  "

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	if Enable {
		fmt.Printf("%s%s\n", identLevel(), fs)
	}
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
