package blockchain

import (
	"encoding/hex"
	"bytes"
)

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	chain.Chain = append(chain.Chain, blk)
}

func (chain Blockchain) IsValid() bool {
	for i := 0;  i < len(chain.Chain); i++ {
		if i == 0 && ((hex.EncodeToString(chain.Chain[0].PrevHash) != "0000000000000000000000000000000000000000000000000000000000000000") || (chain.Chain[i].Generation != 0)) {
			return false
		} else if chain.Chain[0].Difficulty != chain.Chain[i].Difficulty {
			return false
		} else if i > 0 && (chain.Chain[i].Generation != (((chain.Chain[i-1].Generation)+1))) {
			return false
		} else if i > 0 && !bytes.Equal(chain.Chain[i].PrevHash, chain.Chain[i-1].Hash){
			return false
		} else if !bytes.Equal(chain.Chain[i].Hash, chain.Chain[i].CalcHash()) {
			return false
		} else if !chain.Chain[i].ValidHash() {
			return false
		}
	}
	return true
}
