# n-gram

Get [n-grams](https://en.wikipedia.org/wiki/N-gram) in Golang.

## Installation

`$ go get -u github.com/fernandoporazzi/n-gram`

## Usage

```go
package main

import (
	"fmt"

	ngram "github.com/fernandoporazzi/n-gram"
)

func main() {
	fmt.Println(ngram.NGram(2, "fernando"))  // [fe er rn na an nd do]
	fmt.Println(ngram.NGram(3, "amsterdam")) // [ams mst ste ter erd rda dam]
}

```
