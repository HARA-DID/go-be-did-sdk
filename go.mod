module github.com/meQlause/go-be-did

go 1.25.5

replace github.com/meQlause/account-abstraction-sdk => ./internal/infrastructure/lib/core-account-abstraction-sdk/account-abstraction-sdk

replace github.com/meQlause/hara-core-blockchain-lib => ./internal/infrastructure/lib/core-general-sdk

replace github.com/meQlause/did-root-sdk => ./internal/infrastructure/lib/did-root-sdk/did-root-sdk

replace github.com/meQlause/alias-root-sdk => ./internal/infrastructure/lib/did-alias-sdk/did-alias-sdk

require (
	github.com/ethereum/go-ethereum v1.16.7
	github.com/go-playground/validator/v10 v10.30.1
	github.com/gofiber/fiber/v2 v2.52.10
	github.com/gofiber/swagger v1.1.1
	github.com/joho/godotenv v1.5.1
	github.com/meQlause/account-abstraction-sdk v0.0.0-00010101000000-000000000000
	github.com/meQlause/alias-root-sdk v0.0.0-00010101000000-000000000000
	github.com/meQlause/did-root-sdk v0.0.0-00010101000000-000000000000
	github.com/meQlause/hara-core-blockchain-lib v1.9.0
	github.com/swaggo/swag v1.16.6
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/ProjectZKM/Ziren/crates/go-runtime/zkvm_runtime v0.0.0-20260104020744-7268a54d0358 // indirect
	github.com/andybalholm/brotli v1.2.0 // indirect
	github.com/bits-and-blooms/bitset v1.24.4 // indirect
	github.com/clipperhouse/uax29/v2 v2.2.0 // indirect
	github.com/consensys/gnark-crypto v0.19.2 // indirect
	github.com/crate-crypto/go-eth-kzg v1.4.0 // indirect
	github.com/crate-crypto/go-ipa v0.0.0-20240724233137-53bbb0ceb27a // indirect
	github.com/deckarep/golang-set/v2 v2.8.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0 // indirect
	github.com/ethereum/c-kzg-4844/v2 v2.1.5 // indirect
	github.com/ethereum/go-verkle v0.2.2 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.12 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.22.4 // indirect
	github.com/go-openapi/jsonreference v0.21.4 // indirect
	github.com/go-openapi/spec v0.22.3 // indirect
	github.com/go-openapi/swag/conv v0.25.4 // indirect
	github.com/go-openapi/swag/jsonname v0.25.4 // indirect
	github.com/go-openapi/swag/jsonutils v0.25.4 // indirect
	github.com/go-openapi/swag/loading v0.25.4 // indirect
	github.com/go-openapi/swag/stringutils v0.25.4 // indirect
	github.com/go-openapi/swag/typeutils v0.25.4 // indirect
	github.com/go-openapi/swag/yamlutils v0.25.4 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/klauspost/compress v1.18.2 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.19 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/supranational/blst v0.3.16 // indirect
	github.com/swaggo/files/v2 v2.0.2 // indirect
	github.com/tklauser/go-sysconf v0.3.16 // indirect
	github.com/tklauser/numcpus v0.11.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.69.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/mod v0.31.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	golang.org/x/tools v0.40.0 // indirect
)
