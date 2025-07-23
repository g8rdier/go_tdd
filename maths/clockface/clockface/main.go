package main

import (
	"os"
	"time"

	"example.com/hello/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
