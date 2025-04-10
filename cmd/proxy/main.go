package main

import (
	"github.com/etclab/pre"
)

func main() {
	// setup, sends pp to replicas when they boot
	// NOTE: does it make sense for this to be here or in sender?
	pp := pre.NewPublicParams()

	// get alice's public key when she boots

	// get rkAB from bob when he boots

	// recieve requests from sender (containing ct1)...
	// re-encrypt message to bob (if alice can't handle it herself)
	ct2 := pre.ReEncrypt(pp, rkAB, ct1)
	// send ct2 to bob

	// if alice can handle it herself, just send it along to alice

	// recieve response from alice or bob
	// send response to sender (source)
}
