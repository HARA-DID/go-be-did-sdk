package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
)

func (huc *HelperUseCase) EncodeIssueCredentialParam(input helperdo.EncodeIssueCredentialParamInput) (string, error) {
	return huc.repo.EncodeIssueCredentialParam(input)
}

func (huc *HelperUseCase) EncodeBurnCredentialParam(input helperdo.EncodeBurnCredentialParamInput) (string, error) {
	return huc.repo.EncodeBurnCredentialParam(input)
}

func (huc *HelperUseCase) EncodeUpdateMetadataParam(input helperdo.EncodeUpdateMetadataParamInput) (string, error) {
	return huc.repo.EncodeUpdateMetadataParam(input)
}

func (huc *HelperUseCase) EncodeRevokeCredentialParam(input helperdo.EncodeRevokeCredentialParamInput) (string, error) {
	return huc.repo.EncodeRevokeCredentialParam(input)
}

func (huc *HelperUseCase) EncodeClaimCredentialParam(input helperdo.EncodeClaimCredentialParamInput) (string, error) {
	return huc.repo.EncodeClaimCredentialParam(input)
}
