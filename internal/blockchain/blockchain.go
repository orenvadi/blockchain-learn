package blockchain

import (
	"iter"

	"github.com/orenvadi/blockchain-learn/internal/blockchain/block"
)

type BlockChain struct {
	blocks []*block.Block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := block.New(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func New() *BlockChain {
	return &BlockChain{[]*block.Block{block.NewGenesisBlock()}}
}

// To range over blocks without having ability to change any block
func (bc BlockChain) All() iter.Seq2[int, *block.Block] {
	return func(yield func(int, *block.Block) bool) {
		for i, v := range bc.blocks {
			if !yield(i, v) {
				return
			}
		}
	}
}
