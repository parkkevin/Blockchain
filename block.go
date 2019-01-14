package blockchain

import (
	"fmt"
	"encoding/hex"
	"crypto/sha256"
)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

func Initial(difficulty uint8) Block {
	hexZero, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	if err != nil {
		fmt.Println("Error while DecodeString()...")
	}
	b0 := Block{PrevHash: hexZero, Generation: 0, Difficulty: difficulty, Data: ""}
	b0.Hash = b0.CalcHash()
	return b0
}

func (prev_block Block) Next(data string) Block {
	b := Block{}
	b.PrevHash = prev_block.Hash
	b.Generation = prev_block.Generation + 1
	b.Difficulty = prev_block.Difficulty
	b.Data = data
	b.Hash = b.CalcHash()
	return b
}

func (blk Block) CalcHash() []byte {
	str := fmt.Sprintf("%s:%d:%d:%s:%d", hex.EncodeToString(blk.PrevHash), blk.Generation,
						blk.Difficulty, blk.Data, blk.Proof)
	strHashed := sha256.Sum256([]byte(str))
	return strHashed[:]
}

func (blk Block) ValidHash() bool {
	nBytes := blk.Difficulty / 8
	nBits := blk.Difficulty % 8
	var i int = 1
	for uint8(i) <= nBytes {
		if blk.Hash[len(blk.Hash)-i] == 0 {
			i++
		} else {
			return false
		}
	}
	if blk.Hash[len(blk.Hash)-i] % (1<<nBits) == 0 {
		return true
	} else {
		return false
	}
}

func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
