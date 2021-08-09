package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/binance-chain/go-sdk/client/rpc"
	ctypes "github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/types"
	"github.com/binance-chain/go-sdk/types/tx"
)

// https://github.com/binance-chain/node-binary/issues/53
func main() {

	client := rpc.NewRPCClient("testnet-dex.binance.org", ctypes.TestNetwork)
	_ = client
	codec := types.NewCodec()

	// good case from tbnbcli
	// txn_hex := []byte("c001f0625dee0a482a2c87fa0a200a14cee58554cc6439d3505eebb2a74db55c0e15fbfa12080a03424e4210904e12200a145ec295b63c86ae530f93d933aab169d71aac947212080a03424e4210904e12700a26eb5ae9872103b6d11a86e1df8da40eaed74bbb34706090c8dfec4b85ecd09b0b94d7bd7633011240e823766c7826b05005fcec7e197b36ad2b1a745bf361bd560fc133a00bb251f97da661de4e15e8bac2d813de3e1faab154e34c5363f70578096ed5fb1dd0d6a9188591022021")
	// txn_hex := []byte("c001f0625dee0a482a2c87fa0a200a14cee58554cc6439d3505eebb2a74db55c0e15fbfa12080a03424e4210904e12200a145ec295b63c86ae530f93d933aab169d71aac947212080a03424e4210904e12700a26eb5ae9872103b6d11a86e1df8da40eaed74bbb34706090c8dfec4b85ecd09b0b94d7bd7633011240e823766c7826b05005fcec7e197b36ad2b1a745bf361bd560fc133a00bb251f97da661de4e15e8bac2d813de3e1faab154e34c5363f70578096ed5fb1dd0d6a9188591022021")
	parsedTxs := make([]tx.StdTx, 1)

	txn := make([]byte, hex.DecodedLen(len(txn_hex)))
	n, err := hex.Decode(txn, txn_hex)
	if err != nil {
		fmt.Println("Error - decoding hex")
		return
	}

	err = codec.UnmarshalBinaryLengthPrefixed(txn[:n], &parsedTxs[0])
	// err = codec.UnmarshalBinaryBare(txn[:n], &parsedTxs[0])

	if err != nil {
		fmt.Println("Error - codec unmarshal", err)
		return
	}

	bz, err := json.Marshal(parsedTxs)

	if err != nil {
		fmt.Println("Error - json marshal")
		return
	}

	fmt.Println(string(bz))
}
