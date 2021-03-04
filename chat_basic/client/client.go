package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type message struct {
	Author string `json:"author"`
	Msg    string `json:"msg"`
}

func main() {
	connection, err := net.Dial("tcp", "localhost:8005")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Enter your name >>")
	namReader := bufio.NewReader(os.Stdin)
	nameInput, _ := namReader.ReadString('\n')
	nameInput = nameInput[:len(nameInput)-1]
	fmt.Println("---------------MESSAGE---------------")
	go onMessage(connection)
	for {
		fmt.Print("--send--\n")
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		msg = msg[:len(msg)-1]
		if err != nil {
			break
		}
		msgWrite := fmt.Sprintf("\n%s: %s\n", nameInput, msg)
		connection.Write([]byte(msgWrite))
		savedMessage(nameInput, msg)
	}
	connection.Close()
}

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('$')
		fmt.Print(msg)
	}
}

func savedMessage(name string, msg string) {
	var ctx = context.Background() // empty context => it's never canceld,has no values and has no deadline
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	jsonMsg, err := json.Marshal(message{Author: name, Msg: msg})
	if err != nil {
		log.Fatal(err)
	}
	key := generateKey(name)
	err = client.Set(ctx, key, jsonMsg, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func getMessage(id int) {
}

func generateKey(name string) string {
	currentTime := time.Now()
	key := "message:" + name + "-" + currentTime.Format("2006.01.02 15:04:05 PM")
	return key
}
