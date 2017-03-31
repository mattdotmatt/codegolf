package main

import . "fmt"

func main() {
	var c, j int
	Scanf("%d", &c)
	for i := 1; c > 0; i, j = i+j, i {
		Print(i)
		c--
	}
}
