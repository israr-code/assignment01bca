package main

import (
	"blockchain_assignment/assignment01bca"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create a new blockchain
	chain := assignment01bca.Blockchain{}

	// Ask the user how many blocks they want to create initially
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many blocks do you want to create? ")
	blockCountStr, _ := reader.ReadString('\n')
	blockCountStr = strings.TrimSpace(blockCountStr) // Remove newline or spaces
	blockCount, err := strconv.Atoi(blockCountStr)

	if err != nil || blockCount <= 0 {
		fmt.Println("Invalid input. Please enter a valid number of blocks.")
		return
	}

	// Add the specified number of blocks initially
	for i := 0; i < blockCount; i++ {
		fmt.Printf("Enter transaction for Block #%d (e.g., 'bob to alice'): ", i+1)
		transaction, _ := reader.ReadString('\n')
		transaction = strings.TrimSpace(transaction) // Trim spaces and newline

		previousHash := ""
		if len(chain.Blocks) > 0 {
			previousHash = chain.Blocks[len(chain.Blocks)-1].Hash
		}
		chain.NewBlock(transaction, i+1, previousHash)
	}

	// List all blocks in the blockchain
	fmt.Println("\nBlockchain Details:")
	chain.ListBlocks()

	// Function to handle tampering prompt
	handleTampering(&chain)

	// Option to append new blocks at the end before finishing
	for {
		fmt.Print("\nDo you want to append a new block? (yes/no): ")
		appendAnswer, _ := reader.ReadString('\n')
		appendAnswer = strings.TrimSpace(appendAnswer)

		if appendAnswer == "no" {
			fmt.Println("Exiting program.")
			break
		} else if appendAnswer == "yes" {
			// Prompt for new block transaction
			fmt.Print("Enter transaction for the new block: ")
			newTransaction, _ := reader.ReadString('\n')
			newTransaction = strings.TrimSpace(newTransaction)

			// Append the new block
			chain.AppendBlock(newTransaction)
			fmt.Println("\nUpdated Blockchain Details:")
			chain.ListBlocks()

			// After appending, ask if they want to tamper with any block
			handleTampering(&chain)
		} else {
			fmt.Println("Invalid input. Please type 'yes' or 'no'.")
		}
	}
}

// Function to handle the tampering process
func handleTampering(chain *assignment01bca.Blockchain) {
	reader := bufio.NewReader(os.Stdin)

	// Ask the user if they want to tamper with a block
	fmt.Print("\nDo you want to tamper with any block? (yes/no): ")
	tamperAnswer, _ := reader.ReadString('\n')
	tamperAnswer = strings.TrimSpace(tamperAnswer)

	if tamperAnswer == "yes" {
		fmt.Print("Which block number do you want to tamper with? (1 to ", len(chain.Blocks), "): ")
		blockToTamperStr, _ := reader.ReadString('\n')
		blockToTamperStr = strings.TrimSpace(blockToTamperStr)
		blockToTamper, err := strconv.Atoi(blockToTamperStr)

		if err == nil && blockToTamper > 0 && blockToTamper <= len(chain.Blocks) {
			fmt.Print("Enter new transaction for Block #", blockToTamper, ": ")
			newTransaction, _ := reader.ReadString('\n')
			newTransaction = strings.TrimSpace(newTransaction)
			chain.ChangeBlock(blockToTamper-1, newTransaction)

			// Display confirmation of the tampering
			fmt.Printf("\nBlock %d transaction changed to: %s\n", blockToTamper, newTransaction)
		} else {
			fmt.Println("Invalid block number")
		}
	}

	// Verify the blockchain after tampering
	fmt.Println("\nBlockchain Verification:")
	chain.VerifyChain()
}
