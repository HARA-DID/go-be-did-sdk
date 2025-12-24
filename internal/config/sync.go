package config

import (
	"sync"
)

var (
	onceConfig     sync.Once
	onceBlockchain sync.Once
)
