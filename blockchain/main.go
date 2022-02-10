package main

import (
	"fmt"
	// "rsc.io/quote"
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	/* Block representation */
	Hash		[]byte // hash of this block
	Data		[]byte // data in this block
	PrevHash	[]byte // hash of prev block
}

func (b *Block) DeriveHash() { // https://stackoverflow.com/questions/34031801/function-declaration-syntax-things-in-parenthesis-before-function-name
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]

}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func GenesisBlock() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenesisBlock()}}
}

func main() {
	fmt.Println("Initiating GoBlockChain")
	chain := InitBlockChain()

	fmt.Println("Adding Blocks...")
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")
	fmt.Printf("Total number of blocks in GoBlockChain: %d\n",len(chain.blocks))

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}

}