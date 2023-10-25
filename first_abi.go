package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

const (
	MantleUrl = "https://explorer.testnet.mantle.xyz/"
	myAddr    = "0xF01Db8e37dE8C2F0c509721d4C63a5b5aa4D4021"
)

func main() {
	cli, err := ethclient.Dial(MantleUrl)
	if err != nil {
		fmt.Printf("ethclient.Dial error: %s\n", err.Error())
	}

	chainid, err := cli.ChainID(context.Background())
	if err != nil {
		fmt.Printf("get cli.ChainID error: %s\n", err.Error())
	}
	fmt.Printf("get chainid %s", chainid)

	addr := common.HexToAddress(myAddr)
	blkNum := big.NewInt(24618854) // contract create
	// get account's nonce at blkNum
	nonce, err := cli.NonceAt(context.Background(), addr, blkNum)
	if err != nil {
		fmt.Printf("get nonce err: %s", err.Error())
	}
	fmt.Printf("get addr %s  nonce %d at blknum %d", myAddr, nonce, blkNum)
}
