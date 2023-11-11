package main

// RPC admin_peers
type GetResultPeer []ResultPeer

type PeerHeader struct {
	Jsonrpc string       `json:"jsonrpc"`
	ID      int64        `json:"id"`
	Result  []ResultPeer `json:"result"`
}

type ResultPeer struct {
	Caps      []string  `json:"caps"`
	Enode     string    `json:"enode"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Network   Network   `json:"network"`
	Protocols Protocols `json:"protocols"`
}

type Network struct {
	Inbound       bool   `json:"inbound"`
	LocalAddress  string `json:"localAddress"`
	RemoteAddress string `json:"remoteAddress"`
	Static        bool   `json:"static"`
	Trusted       bool   `json:"trusted"`
}

type Protocols struct {
	Eth  Eth `json:"eth"`
	Snap Eth `json:"snap"`
}

type Eth struct {
	Version int64 `json:"version"`
}

// RPC net_peerCount
type PeerCount struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}
