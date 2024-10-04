package main

import (
	"blockchain_assignment/assignment01bca"
)

func main() {
	chain := assignment01bca.Blockchain{}

	chain.NewBlock("bob to alice", 1, "")
	chain.NewBlock("alice to carol", 2, chain.Blocks[len(chain.Blocks)-1].Hash)
	chain.NewBlock("carol to dave", 3, chain.Blocks[len(chain.Blocks)-1].Hash)

	chain.ListBlocks()

	chain.ChangeBlock(1, "bob to dave")

	chain.VerifyChain()
}
