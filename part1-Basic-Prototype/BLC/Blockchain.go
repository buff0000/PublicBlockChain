package BLC

type Blockchain struct {
	//存储有序的区块
	Blocks []*Block
}

//1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain{
	genesisBlock := CreateGenesisBlock("Genesis Data......")
	return &Blockchain{[]*Block{genesisBlock}}
}

//2. 增加区块到区块链里面
func (blc *Blockchain)AddBlockToBlockChain(data string, height int64, preHash []byte){
	// 创建新区块
	newBlock := NewBlock(data,height, preHash)
	//往链上添加区块
	blc.Blocks = append(blc.Blocks, newBlock)
}