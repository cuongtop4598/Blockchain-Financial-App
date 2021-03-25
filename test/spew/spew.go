package main

import "github.com/my/repo/github.com/davecgh/go-spew/spew"

func main() {
	var hello string
	hello = "hello word"
	spew.Dump(hello)
	//spew.Fdump(hello)
}
