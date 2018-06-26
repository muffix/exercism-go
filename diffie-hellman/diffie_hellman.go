// Package diffiehellman implements the Diffie-Hellman-Merkle key exchange
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PublicKey returns a new public key
func PublicKey(private, p *big.Int, g int64) *big.Int {
	public := new(big.Int)

	public.Exp(big.NewInt(g), private, p)

	return public
}

// PrivateKey returns a new private key
func PrivateKey(p *big.Int) *big.Int {
	two := big.NewInt(2)

	max := new(big.Int)
	max.Sub(p, two)

	random, _ := rand.Int(rand.Reader, max)

	key := new(big.Int)
	key.Add(random, two)

	return key
}

// NewPair returns a new public/private key pair
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey returns the secret key to be shared between Alice and Bob
func SecretKey(private1, public2, p *big.Int) *big.Int {
	secret := new(big.Int)
	return secret.Exp(public2, private1, p)
}
