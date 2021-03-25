package node

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/my/repo/github.com/cuongtop4598/blockchain/blocks"
)

// Blockchain is a series of validated Blocks
var Blockchain []blocks.Block
var TempBlocks []blocks.Block

// candidateBlocks handles incoming blocks for validation
var CandidateBlocks = make(chan blocks.Block)

// announcements broadcasts winning validator to all nodes
var Announcements = make(chan string)

var mutex = &sync.Mutex{}

// validators keeps track of open validators and balances
var validators = make(map[string]int)

// CalculateHash is a simple SHA256 hashing function
func CalculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	go func() {
		for {
			msg := <-Announcements
			io.WriteString(conn, msg)
		}
	}()
	// validator address
	var address string

	// allow user to allocate number of tokens to stake
	// the greater the number of tokens, the greater chance to forging a new block
	io.WriteString(conn, "Enter token balance: ") // write string to conn
	scanBalance := bufio.NewScanner(conn)         // return a new scanner read from conn
	for scanBalance.Scan() {
		_, err := strconv.Atoi(scanBalance.Text())
		if err != nil {
			log.Printf("%v not a number: %v", scanBalance.Text(), err)
			return
		}
		t := time.Now()
		address = CalculateHash(t.String())
		fmt.Println(validators)
		break
	}
	io.WriteString(conn, "\nEnter a new BPM:")

	scanBPM := bufio.NewScanner(conn)

	go func() {
		for {
			// take in BPM from stdin and add it to blockchain after conducting necessary validation
			for scanBPM.Scan() {
				bpm, err := strconv.Atoi(scanBPM.Text())
				//if malicious part tries to mutate the chain with a bad input, delete them as
				// a validator and they lose their staked tokens
				if err != nil {
					log.Printf("%v not a bumber: %v", scanBPM.Text(), err)
					delete(validators, address)
					conn.Close()
				}
				mutex.Lock()
				oldLastIndex := Blockchain[len(Blockchain)-1]
				mutex.Unlock()
				// create newBlock for consideration to be forged
				newBlock, err := blocks.GenerateBlock(oldLastIndex, bpm, address)
				if err != nil {
					log.Println(err)
					continue
				}
				if blocks.IsBlockValid(newBlock, oldLastIndex) {
					CandidateBlocks <- newBlock
				}
				io.WriteString(conn, "\nEnter a new BPM:")

			}
		}
	}()
	// simulate receiving broadcast
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
