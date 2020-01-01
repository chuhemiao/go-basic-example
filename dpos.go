package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var blockChain []Block

//存放代理人的地址
var delegate = []string{"aaa", "bbb", "ccc", "ddd"}

// 声明区块的结构体

type Block struct {
	BMP int

	PrefHash string

	HashCode string

	TimeStamp string

	Index int

	//区块验证者
	Validator string
}

func main() {
	var firstBlock Block

	blockChain = append(blockChain, firstBlock)

	//通过变量n按顺序让代理人作为矿工
	var n = 0

	var temp = 0

	for {

		randeDelegate()
		//每隔5秒产生一个新区快
		time.Sleep(5 * time.Second)

		var nextBlock = GenerateNextBlock(temp, blockChain[temp], delegate[n])

		blockChain = append(blockChain, nextBlock)

		n++

		temp++

		n = n % len(delegate)

		fmt.Println(blockChain)
	}
}

//随机调换
func randeDelegate() {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(delegate))
	index1 := rand.Intn(len(delegate))

	if index1 == index {
	Label:
		index1 = rand.Intn(len(delegate))
		if index1 == index {
			goto Label
		}
	}

	tempDelegate := delegate[index]
	delegate[index] = delegate[index1]
	delegate[index1] = tempDelegate
}

// 声明产生新区快的方法
func GenerateNextBlock(bmp int, oldBlock Block, validator string) Block {

	var newBlock Block
	newBlock.PrefHash = oldBlock.HashCode
	newBlock.TimeStamp = time.Now().String()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Validator = validator
	newBlock.BMP = bmp
	newBlock.HashCode = SetHash(newBlock)
	return newBlock
}

// 产生区块哈希的方法
func SetHash(b Block) string {

	hashCode := []byte(b.Validator + strconv.Itoa(b.Index) + b.TimeStamp + b.PrefHash + strconv.Itoa(b.BMP))

	sha := sha256.New()

	sha.Write(hashCode)

	hash := sha.Sum(nil)

	return hex.EncodeToString(hash)
}
