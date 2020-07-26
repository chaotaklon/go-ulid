# ULID

A go lang ULID library.


### Usage

`
$ go get github.com/chaotaklon/go-ulid
`

`
package main

import (
	"fmt"
	"time"
	
	"github.com/chaotaklon/go-ulid"  // the ulid module
)

func main() {
	ulid.Init()
	
	for i := 0; i < 10; i++ {
		fmt.Println(ulid.New())
		time.Sleep(time.Duration(500) * time.Microsecond)
	}
}
`


### License

Apache License 2.0


### What is ULID

https://github.com/ulid/spec

