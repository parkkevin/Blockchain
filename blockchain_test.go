// Kevin Park
// 301322108

package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"encoding/hex"
	"bytes"
	"fmt"
)

func TestInitial(t *testing.T) {
	b := Initial(7)
	assert.Equal(t, "0000000000000000000000000000000000000000000000000000000000000000", hex.EncodeToString(b.PrevHash), "They should be the same.")
	assert.Equal(t, uint64(0), b.Generation, "They should be the same.")
}

func TestNext(t *testing.T) {
	b0 := Initial(7)
	b0.Mine(1)
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	assert.Equal(t, b0.Hash, b1.PrevHash, "They should be the same.")
}

func TestDifficultyGenerationPrevHash(t *testing.T) {
	b0 := Initial(7)
	b0.Mine(1)
	b := Blockchain{}
	b.Add(b0)
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	b.Add(b1)
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	b.Add(b2)

	for i := 0; i < len(b.Chain); i++ {
		if i > 0 {
			assert.Equal(t, b.Chain[0].Difficulty, b.Chain[i].Difficulty, "They should be the same.")
			assert.Equal(t, b.Chain[i-1].Generation+1, b.Chain[i].Generation, "They should be the same.")
			assert.Equal(t, b.Chain[i-1].Hash, b.Chain[i].PrevHash, "They should be the same.")
		}
	}
}

func TestNullBytesHash(t *testing.T) {
	b0 := Initial(7)
	b0.Mine(1)
	b := Blockchain{}
	b.Add(b0)
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	b.Add(b1)
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	b.Add(b2)

	for i := 0; i < len(b.Chain); i++ {
		if !b.Chain[i].ValidHash() {
			t.Errorf("%s has invalid null bytes\n", hex.EncodeToString(b.Chain[i].Hash))
		}
		if !bytes.Equal(b.Chain[i].Hash, b.Chain[i].CalcHash()) {
			t.Errorf("%s has invalid hash\n", hex.EncodeToString(b.Chain[i].Hash))
			fmt.Printf("PrevHash: %s\n", hex.EncodeToString(b.Chain[i].PrevHash))
			fmt.Printf("Generation: %d\n", b.Chain[i].Generation)
			fmt.Printf("Difficulty: %d\n", b.Chain[i].Difficulty)
			fmt.Printf("Data: %s\n", b.Chain[i].Data)
			fmt.Printf("Proof: %d\n", b.Chain[i].Proof)
		}
	}
}
