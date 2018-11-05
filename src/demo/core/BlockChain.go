package core

import (
	"log"
	"fmt"
)

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	genesisBlock := GenerateGenesisBlock()
	blockChain := BlockChain{}
	blockChain.AppendBlock(&genesisBlock)
	return &blockChain
}

func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(&newBlock)
}

func (bc *BlockChain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}

}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index:%d\n", block.Index)
		fmt.Printf("PrevHash:%s\n", block.PrevBlockHash)
		fmt.Printf("Hash:%s\n", block.Hash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Timestamp:%d\n", block.Timestamp)
		fmt.Println()
	}
}
