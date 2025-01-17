package warden

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	wardencommon "github.com/warden-protocol/wardenprotocol/precompiles/common"
	types "github.com/warden-protocol/wardenprotocol/warden/x/warden/types/v1beta3"
)

func newAllKeysRequest(method *abi.Method, args []interface{}) (*types.QueryAllKeysRequest, error) {
	if len(args) != 2 {
		return nil, wardencommon.WrongArgsNumber{Expected: 2, Got: len(args)}
	}

	var input allKeysInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to allKeysInput struct: %w", err)
	}

	pagination := mapEthPageRequest(input.PageRequest)
	return &types.QueryAllKeysRequest{
		Pagination:      &pagination,
		DeriveAddresses: input.DeriveAddresses,
	}, nil
}

type allKeysInput struct {
	PageRequest     TypesPageRequest
	DeriveAddresses []types.AddressType
}

func (o *KeyResponse) FromResponse(res *types.QueryKeyResponse) (*KeyResponse, error) {
	if res == nil {
		return nil, errors.New("received nil QueryKeyResponse")
	}

	key := Key{
		Id:                res.Key.Id,
		SpaceId:           res.Key.SpaceId,
		KeychainId:        res.Key.KeychainId,
		KeyType:           int32(res.Key.Type),
		PublicKey:         res.Key.PublicKey,
		ApproveTemplateId: res.Key.ApproveTemplateId,
		RejectTemplateId:  res.Key.RejectTemplateId,
	}

	addresses := make([]AddressesResponse, len(res.Addresses))

	for j, a := range res.Addresses {
		ethAddress, err := wardencommon.AddressFromBech32Str(a.Address)
		if err != nil {
			return nil, err
		}

		addresses[j] = AddressesResponse{
			AddressValue: ethAddress,
			AddressType:  int32(a.Type),
		}
	}

	return &KeyResponse{
		Key:       key,
		Addresses: addresses,
	}, nil
}

func (o *KeyResponse) Pack(args abi.Arguments) ([]byte, error) {
	return args.Pack(o)
}

type keysOutput struct {
	KeysResponse []KeyResponse
	PageResponse TypesPageResponse
}

func (o *keysOutput) FromResponse(res *types.QueryKeysResponse) (*keysOutput, error) {
	if res == nil || res.Keys == nil {
		return nil, errors.New("received nil QueryKeyResponse")
	}

	o.KeysResponse = make([]KeyResponse, len(res.Keys))
	for i, k := range res.Keys {
		keyResponse, err := new(KeyResponse).FromResponse(&k)
		if err != nil {
			return nil, err
		}

		o.KeysResponse[i] = *keyResponse
	}

	if res.Pagination != nil {
		o.PageResponse.Total = res.Pagination.Total
		o.PageResponse.NextKey = res.Pagination.NextKey
	}

	return o, nil
}

func (o *keysOutput) Pack(args abi.Arguments) ([]byte, error) {
	return args.Pack(o.KeysResponse, o.PageResponse)
}

func newKeyByIdRequest(method *abi.Method, args []interface{}) (*types.QueryKeyByIdRequest, error) {
	if len(args) != 2 {
		return nil, wardencommon.WrongArgsNumber{Expected: 2, Got: len(args)}
	}

	var input keyByIdInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to keyByIdInput struct: %w", err)
	}

	return &types.QueryKeyByIdRequest{
		Id:              input.Id,
		DeriveAddresses: input.DeriveAddresses,
	}, nil
}

type keyByIdInput struct {
	Id              uint64
	DeriveAddresses []types.AddressType
}

func newKeysBySpaceIdRequest(method *abi.Method, args []interface{}) (*types.QueryKeysBySpaceIdRequest, error) {
	if len(args) != 3 {
		return nil, wardencommon.WrongArgsNumber{Expected: 3, Got: len(args)}
	}

	var input keysBySpaceIdInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to keysBySpaceIdInput struct: %w", err)
	}

	return &types.QueryKeysBySpaceIdRequest{
		Pagination:      &input.PageRequest,
		SpaceId:         input.SpaceId,
		DeriveAddresses: input.DeriveAddresses,
	}, nil
}

type keysBySpaceIdInput struct {
	PageRequest     query.PageRequest
	SpaceId         uint64
	DeriveAddresses []types.AddressType
}

func newKeyRequestByIdRequest(method *abi.Method, args []interface{}) (*types.QueryKeyRequestByIdRequest, error) {
	if len(args) != 1 {
		return nil, wardencommon.WrongArgsNumber{Expected: 1, Got: len(args)}
	}

	var input struct{ Id uint64 }
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("failed to unpack arguments into key request id input: %w", err)
	}

	return &types.QueryKeyRequestByIdRequest{
		Id: input.Id,
	}, nil
}

func (kr *KeyRequest) FromResponse(res *types.QueryKeyRequestByIdResponse) (*KeyRequest, error) {
	if res == nil || res.KeyRequest == nil {
		return nil, errors.New("received nil response or key request")
	}
	return kr.mapKeyRequest(*res.KeyRequest)
}

func (kr *KeyRequest) Pack(args abi.Arguments) ([]byte, error) {
	return args.Pack(kr)
}

func newKeyRequestsRequest(method *abi.Method, args []interface{}) (*types.QueryKeyRequestsRequest, error) {
	if len(args) != 4 {
		return nil, wardencommon.WrongArgsNumber{Expected: 4, Got: len(args)}
	}

	var input keyRequestsInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to keyRequestsInput struct: %w", err)
	}

	if _, ok := types.KeyRequestStatus_name[input.Status]; !ok {
		return nil, fmt.Errorf("invalid Status value: %d", input.Status)
	}

	return &types.QueryKeyRequestsRequest{
		Pagination: &input.PageRequest,
		KeychainId: input.KeychainId,
		Status:     types.KeyRequestStatus(input.Status),
		SpaceId:    input.SpaceId,
	}, nil
}

type keyRequestsInput struct {
	PageRequest query.PageRequest
	KeychainId  uint64
	Status      int32
	SpaceId     uint64
}

type keyRequestsOutput struct {
	Pagination  TypesPageResponse
	KeyRequests []KeyRequest
}

func (o *keyRequestsOutput) FromResponse(res *types.QueryKeyRequestsResponse) (*keyRequestsOutput, error) {
	if res == nil || res.KeyRequests == nil {
		return nil, errors.New("received nil QueryKeyRequestsResponse")
	}

	o.KeyRequests = make([]KeyRequest, len(res.KeyRequests))
	for i, k := range res.KeyRequests {
		keyRequestResponse, err := new(KeyRequest).mapKeyRequest(*k)

		if err != nil {
			return nil, err
		}

		o.KeyRequests[i] = *keyRequestResponse
	}

	if res.Pagination != nil {
		o.Pagination = mapSdkPageResponse(*res.Pagination)
	}

	return o, nil
}

func (o *keyRequestsOutput) Pack(args abi.Arguments) ([]byte, error) {
	return args.Pack(o.KeyRequests, o.Pagination)
}

func newKeychainRequest(method *abi.Method, args []interface{}) (*types.QueryKeychainByIdRequest, error) {
	if len(args) != 1 {
		return nil, wardencommon.WrongArgsNumber{Expected: 1, Got: len(args)}
	}

	var input struct{ Id uint64 }
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("failed to unpack arguments into keychain id input: %w", err)
	}

	return &types.QueryKeychainByIdRequest{
		Id: input.Id,
	}, nil
}

func (k *Keychain) FromResponse(res *types.QueryKeychainByIdResponse) (*Keychain, error) {
	if res == nil || res.Keychain == nil {
		return nil, errors.New("received nil response or keychain")
	}
	return k.mapKeychain(*res.Keychain)
}

func (k *Keychain) Pack(args abi.Arguments) ([]byte, error) {
	return args.Pack(k)
}

func newKeychainsRequest(method *abi.Method, args []interface{}) (*types.QueryKeychainsRequest, error) {
	if len(args) != 1 {
		return nil, wardencommon.WrongArgsNumber{Expected: 1, Got: len(args)}
	}

	var input keychainsRequestsInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to keychainsRequestsInput struct: %w", err)
	}

	return &types.QueryKeychainsRequest{
		Pagination: &input.PageRequest,
	}, nil
}

type keychainsRequestsInput struct {
	PageRequest query.PageRequest
}

type keychainsOutput struct {
	Pagination query.PageResponse
	Keychains  []Keychain
}

func (o *keychainsOutput) FromResponse(res *types.QueryKeychainsResponse) (*keychainsOutput, error) {
	if res == nil || res.Keychains == nil {
		return nil, errors.New("received nil QueryKeychainsResponse")
	}

	o.Keychains = make([]Keychain, len(res.Keychains))

	for i, k := range res.Keychains {
		keychain, err := new(Keychain).mapKeychain(k)
		if err != nil {
			return nil, err
		}
		o.Keychains[i] = *keychain
	}

	if res.Pagination != nil {
		o.Pagination = *res.Pagination
	}

	return o, nil
}

func (o *keychainsOutput) Pack(args abi.Arguments) ([]byte, error) {
	return args.Pack(o.Keychains, o.Pagination)
}

func newSignRequestByIdRequest(args []interface{}) (*types.QuerySignRequestByIdRequest, error) {
	if len(args) != 1 {
		return nil, wardencommon.WrongArgsNumber{Expected: 1, Got: len(args)}
	}

	id, ok := args[0].(uint64)
	if !ok {
		return nil, fmt.Errorf("expected uint64 for id, got %T", args[0])
	}

	return &types.QuerySignRequestByIdRequest{
		Id: id,
	}, nil
}

func (o *SignRequest) FromResponse(res *types.QuerySignRequestByIdResponse) (*SignRequest, error) {
	if res == nil || res.SignRequest == nil {
		return nil, errors.New("received nil QuerySignRequestByIdResponse")
	}
	return o.mapSignRequest(res.SignRequest)
}

func (o *SignRequest) mapSignRequest(signRequest *types.SignRequest) (*SignRequest, error) {
	ethCreator, err := wardencommon.AddressFromBech32Str(signRequest.Creator)
	if err != nil {
		return nil, err
	}

	o.Creator = ethCreator

	o.Id = signRequest.Id
	o.KeyId = signRequest.KeyId
	o.DataForSigning = signRequest.DataForSigning
	o.Status = int32(signRequest.Status)

	result := signRequest.Result
	if signRequest.Status == types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED {
		if signedData, ok := result.(*types.SignRequest_SignedData); ok {
			o.Result = signedData.SignedData
		} else {
			return nil, errors.New("unexpected result type for fulfilled sign request")
		}
	} else if signRequest.Status == types.SignRequestStatus_SIGN_REQUEST_STATUS_REJECTED {
		if rejectReason, ok := result.(*types.SignRequest_RejectReason); ok {
			o.Result = []byte(rejectReason.RejectReason)
		} else {
			return nil, errors.New("unexpected result type for rejected sign request")
		}
	}

	o.EncryptionKey = signRequest.EncryptionKey
	o.DeductedKeychainFees = mapSdkCoins(signRequest.DeductedKeychainFees)

	return o, nil
}

func (o *SignRequest) Pack(args abi.Arguments) ([]byte, error) {
	return args.Pack(o)
}

func newSignRequestsRequest(method *abi.Method, args []interface{}) (*types.QuerySignRequestsRequest, error) {
	if len(args) != 3 {
		return nil, wardencommon.WrongArgsNumber{Expected: 3, Got: len(args)}
	}

	var input signRequestsInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to signRequestsInput struct: %w", err)
	}

	if _, ok := types.SignRequestStatus_name[input.Status]; !ok {
		return nil, fmt.Errorf("invalid Status value: %d", input.Status)
	}

	return &types.QuerySignRequestsRequest{
		Pagination: &input.PageRequest,
		KeychainId: input.KeychainId,
		Status:     types.SignRequestStatus(input.Status),
	}, nil
}

type signRequestsInput struct {
	PageRequest query.PageRequest
	KeychainId  uint64
	Status      int32
}

type signRequestsOutput struct {
	SignRequests []SignRequest
	Pagination   query.PageResponse
}

func (o *signRequestsOutput) FromResponse(res *types.QuerySignRequestsResponse) (*signRequestsOutput, error) {
	if res == nil || res.SignRequests == nil {
		return nil, errors.New("received nil QuerySignRequestsResponse")
	}

	o.SignRequests = make([]SignRequest, len(res.SignRequests))
	for i, k := range res.SignRequests {
		signRequest, err := new(SignRequest).mapSignRequest(k)
		if err != nil {
			return nil, err
		}
		o.SignRequests[i] = *signRequest
	}

	if res.Pagination != nil {
		o.Pagination = *res.Pagination
	}

	return o, nil
}

func (o *signRequestsOutput) Pack(args abi.Arguments) ([]byte, error) {
	return args.Pack(o.SignRequests, o.Pagination)
}

func newSpaceByIdRequest(args []interface{}) (*types.QuerySpaceByIdRequest, error) {
	if len(args) != 1 {
		return nil, wardencommon.WrongArgsNumber{Expected: 1, Got: len(args)}
	}

	id, ok := args[0].(uint64)
	if !ok {
		return nil, fmt.Errorf("expected uint64 for id, got %T", args[0])
	}

	return &types.QuerySpaceByIdRequest{
		Id: id,
	}, nil
}

func (o *Space) FromResponse(res *types.QuerySpaceByIdResponse) (*Space, error) {
	if res == nil || res.Space == nil {
		return nil, errors.New("received nil QuerySpaceByIdResponse")
	}
	return o.mapSpace(res.Space)
}

func (o *Space) mapSpace(space *types.Space) (*Space, error) {
	ethCreator, err := wardencommon.AddressFromBech32Str(space.Creator)
	if err != nil {
		return nil, err
	}

	ethOwners, err := wardencommon.AddressesFromBech32StrArray(space.Owners)
	if err != nil {
		return nil, err
	}

	o.Id = space.Id
	o.Creator = ethCreator
	o.Owners = ethOwners
	o.Nonce = space.Nonce
	o.ApproveAdminTemplateId = space.ApproveAdminTemplateId
	o.RejectAdminTemplateId = space.RejectAdminTemplateId
	o.ApproveSignTemplateId = space.ApproveSignTemplateId
	o.RejectSignTemplateId = space.RejectSignTemplateId

	return o, nil
}

func (o *Space) Pack(args abi.Arguments) ([]byte, error) {
	return args.Pack(o)
}

func newSpacesRequest(method *abi.Method, args []interface{}) (*types.QuerySpacesRequest, error) {
	if len(args) != 1 {
		return nil, wardencommon.WrongArgsNumber{Expected: 1, Got: len(args)}
	}

	var input spacesInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to spacesInput struct: %w", err)
	}

	return &types.QuerySpacesRequest{
		Pagination: &input.PageRequest,
	}, nil
}

type spacesInput struct {
	PageRequest query.PageRequest
}

type spacesOutput struct {
	Pagination query.PageResponse
	Spaces     []Space
}

func (o *spacesOutput) FromResponse(res *types.QuerySpacesResponse) (*spacesOutput, error) {
	if res == nil || res.Spaces == nil {
		return nil, errors.New("received nil QuerySpacesResponse")
	}
	o.Spaces = make([]Space, len(res.Spaces))
	for i, k := range res.Spaces {
		space, err := new(Space).mapSpace(&k)
		if err != nil {
			return nil, err
		}

		o.Spaces[i] = *space
	}

	if res.Pagination != nil {
		o.Pagination = *res.Pagination
	}

	return o, nil
}

func (o *spacesOutput) Pack(args abi.Arguments) ([]byte, error) {
	return args.Pack(o.Spaces, o.Pagination)
}

func newSpacesByOwnerRequest(method *abi.Method, args []interface{}) (*types.QuerySpacesByOwnerRequest, error) {
	if len(args) != 2 {
		return nil, wardencommon.WrongArgsNumber{Expected: 2, Got: len(args)}
	}

	var input spacesByOwnerInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to spacesInput struct: %w", err)
	}

	return &types.QuerySpacesByOwnerRequest{
		Pagination: &input.PageRequest,
		Owner:      wardencommon.Bech32StrFromAddress(input.Owner),
	}, nil
}

type spacesByOwnerInput struct {
	PageRequest query.PageRequest
	Owner       common.Address
}

func (kr *KeyRequest) mapKeyRequest(keyRequest types.KeyRequest) (*KeyRequest, error) {
	if keyRequest.Creator == "" {
		return nil, errors.New("keyRequest.Creator is empty")
	}

	ethCreator, err := wardencommon.AddressFromBech32Str(keyRequest.Creator)

	if err != nil {
		return nil, err
	}

	kr.Id = keyRequest.Id
	kr.Creator = ethCreator
	kr.SpaceId = keyRequest.SpaceId
	kr.KeychainId = keyRequest.KeychainId
	kr.KeyType = int32(keyRequest.KeyType)
	kr.Status = int32(keyRequest.Status)
	kr.RejectReason = keyRequest.RejectReason
	kr.ApproveTemplateId = keyRequest.ApproveTemplateId
	kr.RejectTemplateId = keyRequest.RejectTemplateId
	kr.DeductedKeychainFees = mapSdkCoins(keyRequest.DeductedKeychainFees)

	return kr, nil
}

func (k *Keychain) mapKeychain(keychain types.Keychain) (*Keychain, error) {
	var keybaseId string
	if keychain.KeybaseId != nil {
		keybaseId = keychain.KeybaseId.Value
	}

	ethCreator, err := wardencommon.AddressFromBech32Str(keychain.Creator)
	if err != nil {
		return nil, err
	}

	ethAdmins, err := wardencommon.AddressesFromBech32StrArray(keychain.Admins)
	if err != nil {
		return nil, err
	}

	ethWriters, err := wardencommon.AddressesFromBech32StrArray(keychain.Writers)
	if err != nil {
		return nil, err
	}

	k.Id = keychain.Id
	k.Creator = ethCreator
	k.Name = keychain.Name
	k.Admins = ethAdmins
	k.Writers = ethWriters
	k.Fees = mapSdkKeychainFees(keychain.Fees)
	k.Description = keychain.Description
	k.Url = keychain.Url
	k.KeybaseId = keybaseId

	return k, nil
}
