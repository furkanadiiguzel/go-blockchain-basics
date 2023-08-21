package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}
type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}
func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}
func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func main() {
	// create a new blockchain instance with a mining difficulty of 2
	blockchain := CreateBlockchain(2)

	blockchain.addBlock("Furkan", "Adıgüzel", 5)
	blockchain.addBlock("Nakruf", "Adıgüzel", 2)

	// check if the blockchain is valid; expecting true
	fmt.Println(blockchain.isValid())
	fmt.Println("Genesis Block: ", blockchain.genesisBlock)
	fmt.Println("Blockchain Length: ", len(blockchain.chain))
	fmt.Println("Blockchain Difficulty: ", blockchain.difficulty)
	fmt.Println("Blockchain Last Block: ", blockchain.chain[len(blockchain.chain)-1])
	fmt.Println("Blockchain Last Block Data: ", blockchain.chain[len(blockchain.chain)-1].data)
	fmt.Println("Blockchain Last Block Data From: ", blockchain.chain[len(blockchain.chain)-1].data["from"])
	fmt.Println("Blockchain Last Block Data To: ", blockchain.chain[len(blockchain.chain)-1].data["to"])
	fmt.Println("Blockchain Last Block Data Amount: ", blockchain.chain[len(blockchain.chain)-1].data["amount"])
	fmt.Println("Blockchain Last Block Hash: ", blockchain.chain[len(blockchain.chain)-1].hash)
	fmt.Println("Blockchain Last Block Previous Hash: ", blockchain.chain[len(blockchain.chain)-1].previousHash)
	fmt.Println("Blockchain Last Block Timestamp: ", blockchain.chain[len(blockchain.chain)-1].timestamp)
	fmt.Println("Blockchain Last Block Pow: ", blockchain.chain[len(blockchain.chain)-1].pow)

	if (blockchain.chain[len(blockchain.chain)-1].previousHash) == "00b470abdb0ed5a9e5d8305e61154b8037bfef774ab40fa9631768577cd9044e" {
		fmt.Println("True")
	}

}
