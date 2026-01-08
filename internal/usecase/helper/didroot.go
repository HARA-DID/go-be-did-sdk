package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
)

func (huc *HelperUseCase) EncodeAddClaimParam(input helperdo.EncodeAddClaimParamInput) (string, error) {
	return huc.repo.EncodeAddClaimParam(input)
}

func (huc *HelperUseCase) EncodeAddKeyParam(input helperdo.EncodeAddKeyParamInput) (string, error) {
	return huc.repo.EncodeAddKeyParam(input)
}

func (huc *HelperUseCase) EncodeCreateDIDParam(input helperdo.EncodeCreateDIDParamInput) (string, error) {
	return huc.repo.EncodeCreateDIDParam(input)
}

func (huc *HelperUseCase) EncodeDeactiveDIDParam(input helperdo.EncodeDeactiveDIDParamInput) (string, error) {
	return huc.repo.EncodeDeactiveDIDParam(input)
}

func (huc *HelperUseCase) EncodeDeleteDataParam(input helperdo.EncodeDeleteDataParamInput) (string, error) {
	return huc.repo.EncodeDeleteDataParam(input)
}

func (huc *HelperUseCase) EncodeReactiveDIDParam(input helperdo.EncodeReactiveDIDParamInput) (string, error) {
	return huc.repo.EncodeReactiveDIDParam(input)
}

func (huc *HelperUseCase) EncodeRemoveClaimParam(input helperdo.EncodeRemoveClaimParamInput) (string, error) {
	return huc.repo.EncodeRemoveClaimParam(input)
}

func (huc *HelperUseCase) EncodeRemoveKeyParam(input helperdo.EncodeRemoveKeyParamInput) (string, error) {
	return huc.repo.EncodeRemoveKeyParam(input)
}

func (huc *HelperUseCase) EncodeStoreDataParam(input helperdo.EncodeStoreDataParamInput) (string, error) {
	return huc.repo.EncodeStoreDataParam(input)
}

func (huc *HelperUseCase) EncodeUpdateDIDParam(input helperdo.EncodeUpdateDIDParamInput) (string, error) {
	return huc.repo.EncodeUpdateDIDParam(input)
}

func (huc *HelperUseCase) EncodeTransferDIDOwnerParam(input helperdo.EncodeTransferDIDOwnerParamInput) (string, error) {
	return huc.repo.EncodeTransferDIDOwnerParam(input)
}
