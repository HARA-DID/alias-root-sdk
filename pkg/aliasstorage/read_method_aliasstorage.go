package aliasstorage

import (
	"context"
	"fmt"

	"github.com/HARA-DID/hara-core-blockchain-lib/pkg/blockchain"
	"github.com/HARA-DID/hara-core-blockchain-lib/utils"
)

func (as *AliasStorage) GetNodeOwner(
	ctx context.Context,
	node [32]byte,
) (utils.Address, error) {
	result, err := as.blockchain.CallContract(
		ctx,
		as.Contract,
		"getNodeOwner",
		[]any{node},
	)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to call contract: %w", err)
	}

	method := as.ContractABI.Methods["getNodeOwner"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].(utils.Address), nil
}

func (as *AliasStorage) GetTLDOwner(
	ctx context.Context,
	tldNode [32]byte,
) (utils.Address, error) {
	result, err := as.blockchain.CallContract(
		ctx,
		as.Contract,
		"getTLDOwner",
		[]any{tldNode},
	)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to call contract: %w", err)
	}

	method := as.ContractABI.Methods["getTLDOwner"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].(utils.Address), nil
}

func (as *AliasStorage) GetNodeDID(
	ctx context.Context,
	tldNode [32]byte,
	node [32]byte,
) ([32]byte, error) {
	result, err := as.blockchain.CallContract(
		ctx,
		as.Contract,
		"getNodeDID",
		[]any{tldNode, node},
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to call contract: %w", err)
	}

	method := as.ContractABI.Methods["getNodeDID"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].([32]byte), nil
}

func (as *AliasStorage) GetAliasData(
	ctx context.Context,
	node [32]byte,
) (*AliasData, error) {
	result, err := as.blockchain.CallContract(
		ctx,
		as.Contract,
		"getAliasData",
		[]any{node},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var aliasData AliasData
	err = as.ContractABI.UnpackIntoInterface(&aliasData, "getAliasData", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	return &aliasData, nil
}

func (as *AliasStorage) IsAliasValid(
	ctx context.Context,
	node [32]byte,
) (bool, error) {
	result, err := as.blockchain.CallContract(
		ctx,
		as.Contract,
		"isAliasValid",
		[]any{node},
	)
	if err != nil {
		return false, fmt.Errorf("failed to call contract: %w", err)
	}

	method := as.ContractABI.Methods["isAliasValid"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return false, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].(bool), nil
}

func (as *AliasStorage) GetDIDRootStorage(ctx context.Context) (utils.Address, error) {
	result, err := as.blockchain.CallContract(
		ctx,
		as.Contract,
		"getDIDRootStorage",
		[]any{},
	)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to call contract: %w", err)
	}

	method := as.ContractABI.Methods["getDIDRootStorage"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].(utils.Address), nil
}

func (as *AliasStorage) GetDIDOrgStorage(ctx context.Context) (utils.Address, error) {
	result, err := as.blockchain.CallContract(
		ctx,
		as.Contract,
		"getDIDOrgStorage",
		[]any{},
	)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to call contract: %w", err)
	}

	method := as.ContractABI.Methods["getDIDOrgStorage"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].(utils.Address), nil
}

func (as *AliasStorage) GetAddress() utils.Address {
	return as.Address
}

func (as *AliasStorage) GetBlockchain() *blockchain.Blockchain {
	return as.blockchain
}
