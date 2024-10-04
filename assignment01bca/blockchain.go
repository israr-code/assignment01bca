package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Transaction string
	Nonce       int
	PreviousHash string
	Hash         string
}

type Blockchain struct {
	Blocks []Block
}

func (bc *Blockchain) NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := Block{
		Transaction: transaction,
		Nonce:       nonce,
		PreviousHash: previousHash,
		Hash:        CalculateHash(transaction, nonce, previousHash),
	}
	bc.Blocks = append(bc.Blocks, block)
	return &block
}

func (bc *Blockchain) ListBlocks() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block #%d\n", i+1)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n\n", block.Hash)
	}
}

func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
	if index >= 0 && index < len(bc.Blocks) {
		bc.Blocks[index].Transaction = newTransaction
		bc.Blocks[index].Hash = CalculateHash(newTransaction, bc.Blocks[index].Nonce, bc.Blocks[index].PreviousHash)
		fmt.Printf("Block %d transaction changed to: %s\n", index+1, newTransaction)
	} else {
		fmt.Println("Invalid block index")
	}
}

func (bc *Blockchain) VerifyChain() {
	for i := range bc.Blocks {
		if i > 0 {
			previousHash := bc.Blocks[i-1].Hash
			if bc.Blocks[i].PreviousHash != previousHash {
				fmt.Printf("Block %d is tampered!\n", i+1)
				return
			}
		}
		fmt.Printf("Block %d is valid.\n", i+1)
	}
}

func CalculateHash(transaction string, nonce int, previousHash string) string {
	hashInput := fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}
