package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {

	// https://mainnet.infura.io  主网地址
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("连接以太坊主网成功")
	_ = client
}
