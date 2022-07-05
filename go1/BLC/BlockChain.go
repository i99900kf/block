package BLC

//step5 创建区块链
type BlockChain struct {
	//区块链的所有区块
	Blocks []*Block
}

//step6 创建一个区块链
func CreateBlockChainWithGenesisBlock(data string) *BlockChain {
	//创建一个区块
	genesisBlock := CreateGenesisBlock(data)
	//创建一个区块链
	blockChain := BlockChain{[]*Block{genesisBlock}}
	return &blockChain
}

// step7:添加区块到区块链中
func (bc *BlockChain) AddBlockToBlockChain(data string, height int64, prevHash []byte) *Block {
	//创建一个区块
	block := NewBlock(data, prevHash, height)
	//添加到区块链中
	bc.Blocks = append(bc.Blocks, block)
	return block
}
