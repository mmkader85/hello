package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	"github.com/mmkader85/hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!dlroW ,olleH"))
	fmt.Println(cmp.Diff("Hello, World!", "Hello, Go!"))
}
