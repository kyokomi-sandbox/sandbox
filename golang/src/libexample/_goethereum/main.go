package main

import (
	"fmt"
	"log"

	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var key = os.Getenv("ETH_KEY")
var password = os.Getenv("ETH_KEY_PASSWORD")

func main() {
	log.SetFlags(log.Lshortfile)

	infuraAccessToken := os.Getenv("INFURA_ACCESS_TOKEN")
	ethApiEndpoint := os.Getenv("ETH_API_ENDPOINT")
	contractAddress := os.Getenv("ETH_CONTRACT_ADDRESS")
	toAddress := "0x5fd15A13Ea65EB827F378D9ABDBf89dcc6448Fb5" // TODO: APIの引数にする kyokomidev
	tokenValue := int64(1)                                    // TODO: 設定値:環境変数にする 単位1e18

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(ethApiEndpoint + infuraAccessToken)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	token, err := NewToken(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token name:", name)

	auth, err := bind.NewTransactor(strings.NewReader(key), password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	tx, err := token.Transfer(auth, common.HexToAddress(toAddress), big.NewInt(tokenValue*1e18))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
}
