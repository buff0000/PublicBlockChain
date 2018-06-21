package main

import (
	"com.buff/PublicChain/part1-Basic-Prototype/BLC"
	"fmt"
)

//import "com.buff/PublicChain/part1-Basic-Prototype/BLC"

func main(){
	block := BLC.NewBlock("Genesis Block",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Println(block)

	genesisBlock := BLC.CreateGenesisBlock("Genesis Block")
	fmt.Println(genesisBlock)

	genesisBlockChain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(genesisBlockChain)
	fmt.Println(genesisBlockChain.Blocks)
	fmt.Println(genesisBlockChain.Blocks[0])

	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	blockchain.AddBlockToBlockChain("Send 100RMB To Buff", blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1,blockchain.Blocks[len(blockchain.Blocks) - 1].Hash)
	blockchain.AddBlockToBlockChain("Send 100RMB To Ruth", blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1,blockchain.Blocks[len(blockchain.Blocks) - 1].Hash)
	fmt.Println("---->",len(blockchain.Blocks))
	}
