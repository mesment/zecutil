package zecutil

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	// "github.com/codahale/blake2"
	"github.com/minio/blake2b-simd"
)

/*
// blake2bHash zcash hash func
func blake2bHash(data, key []byte) (h chainhash.Hash, err error) {
	bHash := blake2.New(&blake2.Config{
		Size:     32,
		Personal: key,
	})

	if _, err = bHash.Write(data); err != nil {
		return h, err
	}

	err = (&h).SetBytes(bHash.Sum(nil))
	return h, err
}

 */




// blake2bHash zcash hash func
func blake2bHash(data, key []byte) (h chainhash.Hash, err error) {
	bl, _ := blake2b.New(&blake2b.Config{
		Size:   32,
		Person: key,
	})

	if _, err = bl.Write(data); err != nil {
		return h, err
	}
	err = (&h).SetBytes(bl.Sum(nil))
	return h, nil
}