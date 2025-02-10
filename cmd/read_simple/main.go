package main

import (
	"fmt"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/mogo/fmt/fmtutil"
)

func main() {
	chars := gameofthrones.Characters()
	fmtutil.MustPrintJSON(chars)
	fmtutil.MustPrintJSON(chars[0])
	fmt.Println("DONE")
}
