package config

import (
	"math/big"
	"sync"

	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/network"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

var (
	once sync.Once

	bcInstance  *blockchain.Blockchain
	netInstance *network.Network
)

func InitBlockchain() {
	once.Do(func() {

		netInstance = network.NewNetwork(
			[]string{
				"http://20.198.228.24:5625",
				"http://13.214.26.197:5625",
				"http://70.153.16.221:5628",
				"http://70.153.192.125:5625",
				"http://70.153.16.216:5625",
			},
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
