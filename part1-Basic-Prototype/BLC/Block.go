package BLC

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	//区块高度
	Height int64
	//上一个区块的HASH
	PrevBlockHash []byte
	//交易数据
	Data []byte
	//时间戳
	Timestamp int64
	//当前区块的HASH
	Hash []byte
}

func(block *Block) SetHash(){
	//将高度转成[]byte
	heightBytes := IntToHex(block.Height)
	//将时间转成[]byte
	timeString := strconv.FormatInt(block.Timestamp,2)
	timeBytes := []byte(timeString)
	//拼接[]byte
	blockBytes := bytes.Join([][]byte{heightBytes,block.PrevBlockHash,block.Data,timeBytes,block.Hash},[]byte{})
	//生成hash
	hash := sha256.Sum256(blockBytes)

	block.Hash = hash[:]
}

//1. 创建新的区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	//创建区块
	block := &Block{height,[]byte(data),prevBlockHash,time.Now().Unix(),nil}

	//设置hash
	block.SetHash()

	return block
}

//2. 生成创世区块
func CreateGenesisBlock(data string) *Block{
	return NewBlock(data,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}
