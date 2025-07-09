package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}