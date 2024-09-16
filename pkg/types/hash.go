package types

import (
	"encoding/hex"
	"math/big"
)

// HashLength is the expected length of the hash
const HashLength = 32

// Hash represents the 32 byte Keccak256 hash of arbitrary data.
type Hash [HashLength]byte

// Bytes returns the hash as a byte slice.
func (h Hash) Bytes() []byte { return h[:] }

// Big converts the hash to a big integer.
func (h Hash) Big() *big.Int {
	return new(big.Int).SetBytes(h[:])
}

// Hex returns the hex string representation of the hash.
func (h Hash) Hex() string { return "0x" + hex.EncodeToString(h[:]) }

// TerminalString returns a shortened version of the hash for logging purposes.
func (h Hash) TerminalString() string {
	return hex.EncodeToString(h[:3]) + "…" + hex.EncodeToString(h[29:])
}

// Equal checks if two hashes are equal.
func (h Hash) Equal(b Hash) bool {
	return h == b
}
