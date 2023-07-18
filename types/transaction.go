package types

import (
	"crypto/sha256"

	"github.com/frt12/blocker/crypto"
	"github.com/frt12/blocker/proto"
	pb "github.com/golang/protobuf/proto"
)

func SignTransaction(pk *crypto.PryvateKey, tx *proto.Transaction) *crypto.Signature {
	pk.Sign(HashTransaction(tx))
}

func HashTransaction(block *proto.Block) []byte {
	b, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)
	return hash[:]
}
