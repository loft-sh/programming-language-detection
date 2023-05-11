package main

import (
	"os"

	"github.com/loft-sh/programming-language-detection/pkg/detector"
)

func main() {
	lang := detector.GetLanguage(os.Args[1], -1)
	println(lang)
}
