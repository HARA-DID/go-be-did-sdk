package backendutils

const (
	TypeGeneralExecute         uint8 = 0
	TypeCreateDID              uint8 = 1
	TypeUpdateDID              uint8 = 2
	TypeDeactivateDID          uint8 = 3
	TypeReactivateDID          uint8 = 4
	TypeTransferDID            uint8 = 5
	TypeStoreData              uint8 = 6
	TypeDeleteData             uint8 = 7
	TypeAddKey                 uint8 = 8
	TypeRemoveKey              uint8 = 9
	TypeAddClaim               uint8 = 10
	TypeRemoveClaim            uint8 = 11
	TypeSetDIDRootStorage      uint8 = 12
	TypeRegisterTLD            uint8 = 13
	TypeRegisterDomain         uint8 = 14
	TypeRegisterSubdomain      uint8 = 15
	TypeSetDID                 uint8 = 16
	TypeExtendRegistration     uint8 = 17
	TypeRevokeAlias            uint8 = 18
	TypeUnrevokeAlias          uint8 = 19
	TypeTransferAliasOwnership uint8 = 20
	TypeIssueCredential        uint8 = 21
	TypeBurnCredential         uint8 = 22
	TypeUpdateMetadata         uint8 = 23
	TypeRevokeCredential       uint8 = 24
	TypeClaimCredential        uint8 = 25
)

const (
	ServiceDIDRoot  uint8 = 0
	ServiceDIDAlias uint8 = 1
	ServiceDIDVC    uint8 = 2
)