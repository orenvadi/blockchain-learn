package proofofwork

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"

	// "github.com/orenvadi/blockchain-learn/internal/blockchain/block"
	"github.com/orenvadi/blockchain-learn/internal/lib/utils"
)

// TODO move to config
const targetBits = 16

var maxNonce = math.MaxInt64

type ProofOfWork struct {
	block  Block
	target *big.Int
}

type Block interface {
	SetHash()

	GetTimestamp() int64
	GetData() []byte
	GetPrevBlockHash() []byte
	GetHash() []byte
	GetNonce() int
}

func New(b Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.GetPrevBlockHash(),
			pow.block.GetData(),
			utils.IntToHex(pow.block.GetTimestamp()),
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.GetData())
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.GetNonce())
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
