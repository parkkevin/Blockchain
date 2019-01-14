package blockchain

import (
	"work_queue"
)

type miningWorker struct {
	block Block
	width uint64
	rangeEnd uint64
}

type MiningResult struct {
	Proof uint64
	Found bool
}

func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	w := end - start
	q := work_queue.Create(uint(workers), uint(chunks))
	mr := MiningResult{Found: false}
	mw := miningWorker{block: blk, width: w, rangeEnd: end}
	j := start
	for i := uint64(0); i < chunks; i++ {
		for ; j<=end && j<j+w; j+=w {
			mw.block.Proof = j
			q.Enqueue(mw)
		}
		if j > end { break }
	}
	for r := range q.Results {
		if r != nil {
			mr = r.(MiningResult)
			q.Shutdown()
			break
		}
	}
	return mr
}

func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}

func (mw miningWorker) Run() interface{} {
	mr := MiningResult{Found: false}
	if tempB, ok := GetProof(&mw.block, mw.width, mw.rangeEnd); ok {
		mr.Proof = tempB.Proof
		mr.Found = true
		mw.block.Proof = tempB.Proof
		mw.block.Hash = tempB.Hash
	}
	return mr
}

func GetProof(block *Block, width uint64, end uint64) (Block, bool) {
	b := block
	var foundFlag bool
	b.Hash = b.CalcHash()
	for i := b.Proof; i<end && i<i+width; i++ {
		if b.ValidHash() {
			foundFlag = true
			block.Hash = b.Hash
		} else {
			b.Proof++
			b.Hash = b.CalcHash()
			foundFlag = false
		}
	}
	return *b, foundFlag
}
