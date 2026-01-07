package repository

import (
	"context"
	"math/big"

	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type DIDAliasRepository interface {
	Resolve(ctx context.Context, input didaliasdomain.ResolveInput) (utils.Hash, error)
	ResolveFromString(ctx context.Context, input didaliasdomain.ResolveFromStringInput) (utils.Hash, error)
	GetAliasStatus(ctx context.Context, input didaliasdomain.GetAliasStatusInput) (*didaliasdomain.AliasStatus, error)
	GetAliasStatusFromString(ctx context.Context, input didaliasdomain.GetAliasStatusFromStringInput) (*didaliasdomain.AliasStatus, error)
	GetOwner(ctx context.Context, input didaliasdomain.GetOwnerInput) (string, error)
	GetOwnerFromString(ctx context.Context, input didaliasdomain.GetOwnerFromStringInput) (string, error)
	GetDID(ctx context.Context, input didaliasdomain.GetDIDInput) (utils.Hash, error)
	GetDIDFromString(ctx context.Context, input didaliasdomain.GetDIDFromStringInput) (utils.Hash, error)
	Namehash(ctx context.Context, input didaliasdomain.NamehashInput) (utils.Hash, error)
	GetRegistrationPeriod(ctx context.Context, input didaliasdomain.GetRegistrationPeriodInput) (*big.Int, error)
}
