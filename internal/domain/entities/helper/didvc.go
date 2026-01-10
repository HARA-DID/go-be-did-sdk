package helperdomain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/meQlause/did-verifiable-credentials-sdk/pkg/vcfactory"
)

type EncodeIssueCredentialParamInput struct {
	Address common.Address
	VCParam vcfactory.IssueCredentialParams
}

type EncodeBurnCredentialParamInput struct {
	Address common.Address
	VCParam vcfactory.BurnCredentialParams
}

type EncodeUpdateMetadataParamInput struct {
	Address common.Address
	VCParam vcfactory.UpdateMetadataParams
}

type EncodeRevokeCredentialParamInput struct {
	Address common.Address
	VCParam vcfactory.RevokeCredentialParams
}

type EncodeClaimCredentialParamInput struct {
	Address common.Address
	VCParam vcfactory.ClaimCredentialParams
}
