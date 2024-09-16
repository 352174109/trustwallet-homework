package types

import (
	"bytes"
	"encoding/hex"
	"math/big"
)

// AddressLength is the expected length of the address
const AddressLength = 20

// Address represents the 20 byte address of an Ethereum account
type Address [AddressLength]byte

// Bytes returns the address as a byte slice
func (a *Address) Bytes() []byte { return a[:] }

// Hex returns an EIP55-compliant hex string representation of the address
func (a *Address) Hex() string { return ("0x") + hex.EncodeToString(a[:]) }

// Big converts the Ethereum address to a big integer.
func (a *Address) Big() *big.Int {
	return new(big.Int).SetBytes(a[:])
}

// Equal returns true if the two addresses are the same
func (a *Address) Equal(b Address) bool {
	return bytes.Equal(a[:], b[:])
}

func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

func (a *Address) String() string {
	return a.Hex()
}

// BuildAddress converts a hex string to an Address
func BuildAddress(s string) Address {
	var a Address
	a.SetBytes([]byte(s))
	return a
}
