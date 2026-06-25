package main

import (
	"sync"

	"github.com/ss2552/3ds-friend/nex"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go nex.StartAuthenticationServer()
	go nex.StartSecureServer()

	wg.Wait()
}
