package BLC

import (
	"time"
)

// step1:创建block结构体
type Block struct {
	// 区块高度
	Height int64
	// 上一个区块的hash
	PrevBlockHash []byte
	// 数据
	Data []byte
	// 时间戳
	TimeStamp int64
	// 当前区块哈希
	Hash  []byte
	Nonce int64
}

// step2:创建一个区块
func NewBlock(data string, prevBlockHash []byte, height int64) *Block {
	//创建区块
	block := &Block{height, provBlockHash, []byte(data), time.Now().Unix(), nil, 0}
	//step5：设置block的hash和nonce
	//设置哈希
	//block.SetHash()
	//调用工作量证明的方法，并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

//step4:创建创世区块：
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, make([]byte, 32, 32), 0)
}
