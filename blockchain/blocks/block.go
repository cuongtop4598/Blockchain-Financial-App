package blocks

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block is...
type Block struct {
	Index     int    // is the position of the data record in the blockchain
	Timestamp string // is automaticall determined and is the time the data is written
	BPM       int    // beats per minute
	Hash      string // hash is a SHA256 identifier representing this data record
	PrevHash  string // is the SHA256 identifier of the previous record in the chain
	Validator string
}

// CalculateHash is a simple SHA256 hashing function
func CalculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
func CalculateBlockHash(block Block) string {
	record := string(rune(block.Index)) + block.Timestamp + string(rune(block.BPM)) + block.PrevHash
	return CalculateHash(record)
}

// GenerateBlock creates a new block using previous block's hash
func GenerateBlock(oldBlock Block, BPM int, address string) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateBlockHash(newBlock)
	newBlock.Validator = address
	return newBlock, nil
}

// IsBlockValid makes sure block is valid by checking index
// and comparing the hash of the previous block
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
