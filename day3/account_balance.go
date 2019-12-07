package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	headNumber, _ := client.HeaderByNumber(context.Background(), nil)
	fmt.Println("最新区块:", headNumber.Number)
	//查询该账号最新余额 blockNumber 传入nil
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("区块最新余额为:", balance)

	//指定区块 9064113 当前余额
	blockNumber := big.NewInt(9064113)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("在区块高度: 5532993 时的余额为: ", balanceAt)

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("以太坊余额:", ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("待处理账户余额:", pendingBalance)
}
