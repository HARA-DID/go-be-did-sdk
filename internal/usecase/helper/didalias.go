package helperusecase

import helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"

func (huc *HelperUseCase) EncodeSetDIDRootStorageParam(input helperdo.EncodeSetDIDRootStorageParamInput) (string, error) {
	return huc.repo.EncodeSetDIDRootStorageParam(input)
}

func (huc *HelperUseCase) EncodeRegisterDomainParam(input helperdo.EncodeRegisterDomainParamInput) (string, error) {
	return huc.repo.EncodeRegisterDomainParam(input)
}

func (huc *HelperUseCase) EncodeRegisterSubdomainParam(input helperdo.EncodeRegisterSubdomainParamInput) (string, error) {
	return huc.repo.EncodeRegisterSubdomainParam(input)
}

func (huc *HelperUseCase) EncodeSetDIDParam(input helperdo.EncodeSetDIDParamInput) (string, error) {
	return huc.repo.EncodeSetDIDParam(input)
}

func (huc *HelperUseCase) EncodeExtendRegistrationParam(input helperdo.EncodeExtendRegistrationParamInput) (string, error) {
	return huc.repo.EncodeExtendRegistrationParam(input)
}

func (huc *HelperUseCase) EncodeRevokeAliasParam(input helperdo.EncodeRevokeAliasParamInput) (string, error) {
	return huc.repo.EncodeRevokeAliasParam(input)
}

func (huc *HelperUseCase) EncodeUnrevokeAliasParam(input helperdo.EncodeUnrevokeAliasParamInput) (string, error) {
	return huc.repo.EncodeUnrevokeAliasParam(input)
}

func (huc *HelperUseCase) EncodeTransferAliasOwnershipParam(input helperdo.EncodeTransferAliasOwnershipParamInput) (string, error) {
	return huc.repo.EncodeTransferAliasOwnershipParam(input)
}
