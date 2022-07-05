package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
	Hash []byte
}

// step2:创建一个区块
func NewBlock(data string, prevBlockHash []byte, height int64) *Block {
	// 创建区块
	block := &Block{}
	// 设置区块高度
	block.Height = height
	// 设置上一个区块哈希
	block.PrevBlockHash = prevBlockHash
	// 设置数据
	block.Data = []byte(data)
	// 设置时间戳
	block.TimeStamp = time.Now().Unix()
	// 设置当前区块哈希
	block.Hash = block.SetHash()
	return block
}

// step3 :设置区块哈希
func (block *Block) SetHash() []byte {
	//1.将高度转为字节数组
	heightBytes := IntToHex(block.Height)
	//fmt.Println(heightBytes)
	//2.时间戳转为字节数组
	//timeBytes:=IntToHex(block.TimeStamp)
	//转为二进制的字符串
	//fmt.Println(block.TimeStamp)
	//fmt.Printf("%x,%b\n",block.TimeStamp,block.TimeStamp)
	timeString := strconv.FormatInt(block.TimeStamp, 2)
	//fmt.Println("timeString:",timeString)
	timeBytes := []byte(timeString)
	//fmt.Println("timeStamp:",timeBytes)
	//3.拼接所有的属性
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeBytes}, []byte{})
	//4.生成哈希值
	hash := sha256.Sum256(blockBytes) //数组长度32位
	block.Hash = hash[:]
	return block.Hash
}

//step4:创建创世区块：
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, make([]byte, 32, 32), 0)
}
