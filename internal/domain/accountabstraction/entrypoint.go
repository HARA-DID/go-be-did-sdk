package accountabstractiondomain

import EntryPointSDK "github.com/meQlause/account-abstraction-sdk/pkg/entrypoint"

type ExecuteOperationInput struct {
	PrivKey string
	Input   EntryPointSDK.HandleOpsParams
}
