package main

import "bytes"

func Countdown (out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}