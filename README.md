# programming-language-detection


This is a simple library to guess which programming language is a project.

Simple usage:

```go
package main

import (
	"os"
	"github.com/loft-sh/programming-language-detection/pkg/detector"
)

func main() {
	lang := detector.GetLanguage("/path/to/my/project", -1)
	println(lang)
}
```

it will return None if it didn't match anything, else it will return one of:

- JavaScript
- TypeScript
- Python
- C
- Cpp
- DotNet
- Go
- PHP
- Java
- Rust
- Ruby
