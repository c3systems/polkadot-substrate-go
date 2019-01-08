package client

import (
	"github.com/c3systems/c3-go/rpc"
	"github.com/c3systems/go-substrate/p2p"
)

// TODO ...

// Config ...
type Config struct {
	//Chain ChainName,
	//DB DbConfig
	//Dev DevConfig
	P2P   *p2p.Config
	RPC   *rpc.Config
	Roles []string
	//Telemetry TelemetryConfig
	//Wasm WasmConfig
}
