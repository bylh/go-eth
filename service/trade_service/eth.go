package trade_service

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接客户端
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/6d0358a8d5e7446bac5270194c5a9245")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")

	account := common.HexToAddress("0xEc5D1fbFB9e5fde5CA3d4C1B8F2dF2bB5bB43e66")
	balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        log.Fatal(err)
	}
	fmt.Println("共多少wei: ") // 25893180161173005034
	fmt.Println(balance) // 25893180161173005034

    // blockNumber := big.NewInt(5532993)
    // balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Println(balanceAt) // 25729324269165216042

    fbalance := new(big.Float)
    fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("共多少eth: ") // 25893180161173005034
    fmt.Println(ethValue) // 25.729324269165216041

	// pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	
    // fmt.Println(pendingBalance) // 25729324269165216042
}
