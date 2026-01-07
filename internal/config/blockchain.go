package config

import (
	"math/big"

	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/network"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

var (
	bcInstance  *blockchain.Blockchain
	netInstance *network.Network
)

func InitBlockchain() {
	onceBlockchain.Do(func() {
		netInstance = network.NewNetwork(
			config.RPC.Endpoints,
			"2.0",
			1,
			utils.DefaultLogConfig(),
		)

		bcInstance = blockchain.NewBlockchain(
			netInstance,
			big.NewInt(1212),
		)
	})
}

func Blockchain() *blockchain.Blockchain {
	if bcInstance == nil {
		panic("bootstrap: Init() must be called before Blockchain()")
	}
	return bcInstance
}

func Network() *network.Network {
	if netInstance == nil {
		panic("bootstrap: Init() must be called before Network()")
	}
	return netInstance
}

func Get() (*blockchain.Blockchain, *network.Network) {
	if bcInstance == nil || netInstance == nil {
		panic("bootstrap: Init() must be called before Get()")
	}
	return bcInstance, netInstance
}
