package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"

	proofofwork "github.com/orenvadi/blockchain-learn/internal/blockchain/proof-of-work"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b Block) GetTimestamp() int64 {
	return b.Timestamp
}

func (b Block) GetData() []byte {
	return b.Data
}

func (b Block) GetPrevBlockHash() []byte {
	return b.PrevBlockHash
}

func (b Block) GetHash() []byte {
	return b.Hash
}

func (b Block) GetNonce() int {
	return b.Nonce
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func New(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := proofofwork.New(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return New("Genesis Block", []byte{})
}
