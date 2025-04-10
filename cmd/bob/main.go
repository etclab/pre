package main

import (
	"github.com/etclab/pre"
)

func main() {

	// setup
	// request pp from proxy
	bob := pre.KeyGen(pp)
	// proxy responds with request for public key, which we send

	// wait for messages from proxy containing ct2...
	m2 := pre.Decrypt2(pp, ct2, bob.SK)
	// process message, make response (make it uppercase for example)
	// send response to proxy

}
