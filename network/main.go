package network

import (
	"github.com/openbankit/go-base/hash"
)

const (
	// PublicNetworkPassphrase is the pass phrase used for every transaction intended for the public stellar network
	PublicNetworkPassphrase = "Smart Money ; May 2016"
	// TestNetworkPassphrase is the pass phrase used for every transaction intended for the SDF-run test network
	TestNetworkPassphrase = "Dev Smart Money ; May 2016"
)

func ID(passphrase string) [32]byte {
	return hash.Hash([]byte(passphrase))
}
