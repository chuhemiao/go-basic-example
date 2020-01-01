package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

/*
* 安装 go get github.com/davecgh/go-spew/spew
* 安装 go get github.com/joho/godotenv
* 如果出现 dial tcp 216.58.200.49:443: i/o timeout
* 请访问https://www.idiot6.com/2019/07/23/go-gin-golang-x/ 设置GO MODULE
* macos运行 go run main.go 然后打开新窗口输入sudo nc localhost 9000 输入token数量和BPM数量，即可等待验证Pos
*
*
 */

//Block 是每个区块的内容
//Blockchain 是我们的官方区块链，它只是一串经过验证的区块集合。每个区块中的 PrevHash 与前面块的 Hash 相比较，以确保我们的链是正确的。 tempBlocks 是临时存储单元，在区块被选出来并添加到 BlockChain 之前，临时存储在这里
//candidateBlocks 是 Block 的通道，任何一个节点在提出一个新块时都将它发送到这个通道
//announcements 也是一个通道，我们的主Go TCP服务器将向所有节点广播最新的区块链
//mutex是一个标准变量，允许我们控制读/写和防止数据竞争
//validators 是节点的存储map，同时也会保存每个节点持有的token数(持币数)

// 定义结构体
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
	Validator string
}

var Blockchain []Block // 创建区块链
var tempBlocks []Block // tempBlocks是临时存储单元，在区块被选出来并添加到BlockChain之前，临时存储在这里

var candidateBlocks = make(chan Block)

var announcements = make(chan string)

var mutex = &sync.Mutex{}

var validators = make(map[string]int)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// create genesis block
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, calculateBlockHash(genesisBlock), "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	httpPort := os.Getenv("PORT")

	// start TCP and serve TCP server
	server, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTTP Server Listening on port :", httpPort)
	defer server.Close()

	go func() {
		for candidate := range candidateBlocks {
			mutex.Lock()
			tempBlocks = append(tempBlocks, candidate)
			mutex.Unlock()
		}
	}()

	go func() {
		for {
			pickWinner()
		}
	}()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

// pos 主要逻辑
// 每隔30秒，我们选出一个胜利者，这样对于每个验证者来说，都有时间提议新的区块，参与到竞争中来。
// 接着创建一个lotteryPool，它会持有所有验证者的地址，这些验证者都有机会成为一个胜利者。
// 然后，对于提议块的暂存区域，我们会通过if len(temp) > 0来判断是否已经有了被提议的区块。
func pickWinner() {
	time.Sleep(30 * time.Second)
	mutex.Lock()
	temp := tempBlocks
	mutex.Unlock()

	lotteryPool := []string{}
	if len(temp) > 0 {

		// slightly modified traditional proof of stake algorithm
		// from all validators who submitted a block, weight them by the number of staked tokens
		// in traditional proof of stake, validators can participate without submitting a block to be forged
	OUTER:
		for _, block := range temp {
			// 检查暂存区域是否和 lotteryPool 中存在同样的验证者，如果存在，则跳过
			for _, node := range lotteryPool {
				if block.Validator == node {
					continue OUTER
				}
			}

			// lock list of validators to prevent data race
			mutex.Lock()
			setValidators := validators
			mutex.Unlock()
			//  获取验证者的tokens
			k, ok := setValidators[block.Validator]
			if ok {
				for i := 0; i < k; i++ {
					// 向 lotteryPool 追加 k 条数据，k 代表的是当前验证者的tokens
					lotteryPool = append(lotteryPool, block.Validator)
				}
			}
		}

		// 通过随机获得获胜节点的地址
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]

		// 把获胜者的区块添加到整条区块链上，然后通知所有节点关于胜利者的消息
		for _, block := range temp {
			if block.Validator == lotteryWinner {
				mutex.Lock()
				Blockchain = append(Blockchain, block)
				mutex.Unlock()
				for _ = range validators {
					announcements <- "\nwinning validator: " + lotteryWinner + "\n"
				}
				break
			}
		}
	}

	mutex.Lock()
	// 清空tempBlocks，以便下次提议的进行
	tempBlocks = []Block{}
	mutex.Unlock()
}

/*
验证者连接到的TCP服务，需要提供一些函数达到以下目标：

	输入令牌的余额（之前提到过，不做钱包等逻辑)
	接收区块链的最新广播
	接收验证者赢得区块的广播信息
	将自身节点添加到全局的验证者列表中（validators)
	输入Block的BPM数据- BPM是每个验证者的人体脉搏值
	提议创建一个新的区块
*/

func handleConn(conn net.Conn) {
	defer conn.Close()

	go func() {
		for {
			msg := <-announcements
			io.WriteString(conn, msg)
		}
	}()
	// validator address
	var address string

	// tokens数量由用户从控制台输入
	io.WriteString(conn, "Enter token balance:")
	scanBalance := bufio.NewScanner(conn)
	// 获取用户输入的balance值，并打印出来
	for scanBalance.Scan() {
		balance, err := strconv.Atoi(scanBalance.Text())
		if err != nil {
			log.Printf("%v not a number: %v", scanBalance.Text(), err)
			return
		}
		t := time.Now()
		address = calculateHash(t.String())
		validators[address] = balance
		fmt.Println(validators)
		break // 只循环一次
	}

	// 循环输入BPM
	io.WriteString(conn, "\nEnter a new BPM:")

	scanBPM := bufio.NewScanner(conn)

	go func() {
		for {
			// take in BPM from stdin and add it to blockchain after conducting necessary validation
			for scanBPM.Scan() {
				bpm, err := strconv.Atoi(scanBPM.Text())
				// if malicious party tries to mutate the chain with a bad input, delete them as a validator and they lose their staked tokens
				if err != nil {
					log.Printf("%v not a number: %v", scanBPM.Text(), err)
					delete(validators, address)
					conn.Close()
				}

				mutex.Lock()
				oldLastIndex := Blockchain[len(Blockchain)-1]
				mutex.Unlock()

				// 生成新的区块
				newBlock, err := generateBlock(oldLastIndex, bpm, address)
				if err != nil {
					log.Println(err)
					continue
				}
				if isBlockValid(newBlock, oldLastIndex) {
					// main func 中 for candidate := range candidateBlocks ,将newBlock 追加到 tempBlocks中
					candidateBlocks <- newBlock
				}
				io.WriteString(conn, "\nEnter a new BPM:")
			}
		}
	}()

	// 循环打印出最新的区块链，这样每个验证者都能获知最新的状态
	for {
		time.Sleep(time.Minute)
		mutex.Lock()
		output, err := json.Marshal(Blockchain)
		mutex.Unlock()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, string(output)+"\n")
	}

}

/*
通过检查 Index 来确保它们按预期递增。
检查以确保我们 PrevHash 的确与 Hash 前一个区块相同。
最后，通过在当前块上 calculateBlockHash 再次运行该函数来检查当前块的散列。
*/
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// calculateHash 函数会接受一个 string ，并且返回一个SHA256 hash

func calculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// calculateBlockHash 是对一个 block 进行 hash，将一个 block 的所有字段连接到一起后，再调用 calculateHash 将字符串转为 SHA256 hash
func calculateBlockHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	return calculateHash(record)
}

/*
generateBlock	创建新块
newBlock.PrevHash 存储的是上一个区块的 Hash
newBlock.Hash 是通过 calculateBlockHash(newBlock) 生成的 Hash 。
newBlock.Validator 存储的是获取记账权的节点地址
*/
func generateBlock(oldBlock Block, BPM int, address string) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateBlockHash(newBlock)
	newBlock.Validator = address

	return newBlock, nil
}
