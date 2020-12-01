package main

import (
	"fmt"
	"math/big"
	"encoding/hex"
	"github.com/btcsuite/btcd/blockchain"
)

func main() {
	factor := int64(128)

	oldForkBlockBits := uint32(0x1b0ffff0)
    	bn := blockchain.CompactToBig(oldForkBlockBits)
	bn = bn.Mul(bn, big.NewInt(factor))
	newForkBlockBits := blockchain.BigToCompact(bn)


	oldPowLimit := "00000fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	b, _ := hex.DecodeString(oldPowLimit)
        bn = big.NewInt(0).SetBytes(b)
	bn = bn.Mul(bn, big.NewInt(factor))
	newPowLimit := hex.EncodeToString(bn.Bytes())

	fmt.Printf("Factor used: %d\n", factor)
	fmt.Printf("Previous fork block bits: %x\n", oldForkBlockBits)
        fmt.Printf("New fork block bits: %x\n", newForkBlockBits)
        fmt.Printf("Previous powLimit: %s\n", oldPowLimit)
        fmt.Printf("New powLimit: %s\n", newPowLimit)

}
