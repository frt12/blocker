package types

import ("testing"

	"github.com/frt12/blocker/crypto"
	"github.com/frt12/blocker/util"
	"github.com/frt12/blocker/proto")

func TestingNewTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	toPrivKey := crypto.GeneratePrivateKey()
	toAddress := toPrivKey.Public().Address().Bytes()
	fromAddress := fromPrivKey.Public().Address().Bytes()
	
	
	input := &proto.TxInput{
		PrevTxHash: util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey: fromPrivKey.Public().Bytes(),
		
	}
	output1 := &proto.TxOutput{
		amount: 5,
		address: toAddress,
		}
	
	output2 := &proto.TxOutput{
		amount: 95,
		address: fromAddress,
	}
	
	tx := &proto.Transaction{
		version: 1,
		inputs: []*proto.TxInput(input),
		output: []*proto.TxOutput(output1, output2),
		 
	}
	
	sig := SignTransaction(fromPrivKey, tx)
	input.Signature = sig.Bites()
	
	fmt.Printf("%v/n", tx)
}