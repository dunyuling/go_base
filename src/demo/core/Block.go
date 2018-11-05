package core

import (
    "crypto/sha256"
    "encoding/hex"
    "time"
)


type Block struct {
    //区块头
    Index int64   //区块编号
    Timestamp int64 //区块时间戳 
    PrevBlockHash string //上一个区块hash
    Hash string //当前区块hash

    //区块体
    Data string //区块数据
}

func calculateHash(b Block) string {
    blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
    hashInBytes := sha256.Sum256([]byte(blockData)) 
    return hex.EncodeToString(hashInBytes[:])
}

func GenerateNewBlock(preBlock Block, data string) Block {
    newBlock := Block{}
    newBlock.Index = preBlock.Index + 1
    newBlock.PrevBlockHash = preBlock.Hash
    newBlock.Timestamp = time.Now().Unix()
    newBlock.Data = data
    newBlock.Hash = calculateHash(newBlock)
    return newBlock
}

func GenerateGenesisBlock() Block {
    preBlock := Block{} 
    preBlock.Index = -1
    preBlock.Hash = ""
    return GenerateNewBlock(preBlock,"Genesis Block")
}
