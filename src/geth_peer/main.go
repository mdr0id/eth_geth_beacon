package main

import (
	"context"
	"fmt"

	"github.com/ybbus/jsonrpc/v3"
)

func main() {
	for {
		getPeerInfo()
		getPeerCount()
	}
}

func getPeerInfo() {
	rpcClient := jsonrpc.NewClient("http://geth:8552")

	// sanity check wire
	// resp, err := rpcClient.Call(context.Background(), "admin_peers")
	// if err != nil {
	// 	fmt.Println("Error: RPC admin_peers JSON decode error.")
	// 	return
	// }
	// fmt.Println(resp.Result)

	var peerResults *GetResultPeer
	if err := rpcClient.CallFor(context.Background(), &peerResults, "admin_peers"); err != nil {
		fmt.Println("Error: RPC admin_peers JSON decode error.", err)
	} else {
		for _, rp := range *peerResults {
			fmt.Println("ID:", rp.ID, "RemoteIP:", rp.Network.RemoteAddress, "Name:", rp.Name, "Proto:", rp.Protocols)
		}

	}
	//if we want to poll some duration on long calls etc
	//time.Sleep(time.Duration(2) * time.Second)
}

func getPeerCount() {
	rpcClient := jsonrpc.NewClient("http://geth:8552")

	body, err := rpcClient.Call(context.Background(), "net_peerCount")
	if err != nil {
		fmt.Println("Error: RPC net_peerCount JSON decode error")
		return
	}
	fmt.Printf("Peers: %s \n", body.Result)

	//if we want to poll some duration on long calls etc
	//time.Sleep(time.Duration(2) * time.Second)
}
