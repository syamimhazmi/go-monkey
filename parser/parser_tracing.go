package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIdenfitiferPlaceholder string = "\t"

func identifierLevel() string {
	return strings.Repeat(traceIdenfitiferPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identifierLevel(), fs)
}

func incrementIdentifier() {
	traceLevel = traceLevel + 1
}

func decrementIdentifier() {
	traceLevel = traceLevel - 1
}

func trace(msg string) string {
	incrementIdentifier()

	tracePrint("BEGIN " + msg)

	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)

	decrementIdentifier()
}
