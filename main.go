package main

import "log"

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	blockChain := NewBlockchain()
	blockChain.AddTransaction("A", "B", 25.5)
	blockChain.Print()

	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(3, previousHash)
	blockChain.Print()

	blockChain.AddTransaction("D", "C", 4.5)
	blockChain.AddTransaction("X", "Y", 3.5)

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()

}
