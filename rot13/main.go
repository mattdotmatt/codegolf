package main

import (
	b "bufio"
	. "fmt"
	. "os"
)

func main() {
	x := b.NewReader(Stdin)
	c, _ := x.ReadString('\n')

	for _, r := range c {
		if r > '`' && r < '{' {
			r = 'a' + (r-'a'+13)%26
		}
		Print(string(r))
	}
}
