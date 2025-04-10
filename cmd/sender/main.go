package main

import (
	bls "github.com/cloudflare/circl/ecc/bls12381"
	"github.com/etclab/pre"
)

func randomGt() *bls.Gt {
	a := pre.RandomScalar()
	b := pre.RandomScalar()

	g1 := bls.G1Generator()
	g2 := bls.G2Generator()

	g1.ScalarMult(a, g1)
	g2.ScalarMult(b, g2)

	z := bls.Pair(g1, g2)
	return z
}

func main() {
	m := randomGt()

	// request (alice's) public key from proxy

	// encrypt message to alice
	ct1 := pre.Encrypt(pp, m, alice.PK)

	// send ciphertext to proxy
	// wait for response from proxy
	// print response
}
