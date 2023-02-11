package main

import (
	"fmt"
	"log"
	"strings"
)

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s chain %d %s \n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
}

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	blockChain := NewBlockchain()
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 25.5)
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(3, previousHash)
	blockChain.Print()

	blockChain.AddTransaction("D", "C", 4.5)
	blockChain.AddTransaction("X", "Y", 3.5)

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()

}
