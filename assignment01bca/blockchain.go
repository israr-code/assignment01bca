package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Block structure representing a block in the blockchain
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
	OriginalHash string // Stores the original hash of the block
}

// Blockchain structure representing the whole blockchain
type Blockchain struct {
	Blocks []Block
}

// NewBlock creates and adds a new block to the blockchain
func (bc *Blockchain) NewBlock(transaction string, nonce int, previousHash string) *Block {
	hash := CalculateHash(transaction, nonce, previousHash)
	block := Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Hash:         hash,
		OriginalHash: hash, // Store the original hash at the time of block creation
	}
	bc.Blocks = append(bc.Blocks, block)
	return &block
}

// ListBlocks prints all blocks in the blockchain in a clearer format
func (bc *Blockchain) ListBlocks() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block #%d\n", i+1)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.Hash)
		fmt.Println(strings.Repeat("-", 50)) // Separator for clarity
	}
}

// ChangeBlock modifies the transaction of a given block and recalculates the hash
func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
	if index >= 0 && index < len(bc.Blocks) {
		bc.Blocks[index].Transaction = newTransaction
		bc.Blocks[index].Hash = CalculateHash(newTransaction, bc.Blocks[index].Nonce, bc.Blocks[index].PreviousHash)
	} else {
		fmt.Println("Invalid block index")
	}
}

// VerifyChain checks if the blockchain is valid by verifying both the current block's hash
// and the chain integrity between the previous hash and the previous block's hash.
func (bc *Blockchain) VerifyChain() {
	chainBroken := false // Once we detect tampering, the rest of the chain should be considered broken

	for i := range bc.Blocks {
		currentBlock := bc.Blocks[i]

		// First, check if the current block's stored hash matches its original hash
		if currentBlock.Hash != currentBlock.OriginalHash {
			fmt.Printf("Block %d is tampered! Stored hash does not match original hash.\n", i+1)
			chainBroken = true // Mark the chain as broken
		}

		// If the chain is broken, mark all subsequent blocks as tampered
		if chainBroken {
			if i > 0 {
				fmt.Printf("Block %d is tampered because of a broken chain!\n", i+1)
			}
			continue
		}

		// Verify the chain integrity by checking the current block's previous hash matches the previous block's hash
		if i > 0 {
			previousBlock := bc.Blocks[i-1]
			if currentBlock.PreviousHash != previousBlock.Hash {
				fmt.Printf("Block %d is tampered due to mismatch with Block %d's hash!\n", i+1, i)
				chainBroken = true // Chain is now broken
			}
		}

		// If no tampering is detected, mark the block as valid
		if !chainBroken {
			fmt.Printf("Block %d is valid.\n", i+1)
		}
	}
}
// New function to append a block dynamically
func (bc *Blockchain) AppendBlock(transaction string) {
	// Get the previous hash of the last block in the chain
	previousHash := ""
	if len(bc.Blocks) > 0 {
		previousHash = bc.Blocks[len(bc.Blocks)-1].Hash
	}

	// Append the new block
	nonce := len(bc.Blocks) + 1 // Assuming nonce is just the block index + 1
	bc.NewBlock(transaction, nonce, previousHash)
	fmt.Printf("New block appended with transaction: '%s'\n", transaction)
}


// CalculateHash generates the hash for a block using SHA256
func CalculateHash(transaction string, nonce int, previousHash string) string {
	hashInput := fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}
