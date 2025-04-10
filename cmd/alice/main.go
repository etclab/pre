package main

import (
	"github.com/etclab/pre"
)

func main() {

	// setup
	alice := pre.KeyGen(pp)

	// proxy requested a re-encryption key
	// alice recieves bob's pk and pp both from proxy
	rkAB := pre.ReEncryptionKeyGen(pp, alice.SK, bob.PK)

	// normal decryption, no PRE
	m1 := pre.Decrypt1(pp, ct1, alice.SK)
	// process message, make response (make it uppercase for example)
	// send response to proxy
}
