package aliasfactory

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/HARA-DID/hara-core-blockchain-lib/pkg/blockchain"
	"github.com/HARA-DID/hara-core-blockchain-lib/utils"
)

func (af *AliasFactory) call(ctx context.Context, method string, args ...any) ([]byte, error) {
	data, err := af.ContractABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("abi pack error for %s: %w", method, err)
	}
	fmt.Println(af.Address)
	raw := "0x" + utils.Bytes2Hex(data)
	return af.blockchain.Network.Call(ctx, af.Address, raw)
}

func unwrapDoubleEncoding(out []byte) ([]byte, error) {
	if len(out) > 2 && out[0] == 0x22 && out[len(out)-1] == 0x22 {
		asciiStr := string(out)
		innerHex := strings.Trim(asciiStr, "\"")
		innerBytes, err := hex.DecodeString(strings.TrimPrefix(innerHex, "0x"))
		if err != nil {
			return nil, fmt.Errorf("failed to decode inner hex: %w", err)
		}
		return innerBytes, nil
	}
	return out, nil
}

func (af *AliasFactory) GetRegistrationPeriod(
	ctx context.Context,
	period uint8,
) (*big.Int, error) {
	out, err := af.call(ctx, "getRegistrationPeriod", period)
	if err != nil {
		return nil, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return nil, err
	}

	values, err := af.ContractABI.Methods["getRegistrationPeriod"].Outputs.Unpack(out)
	if err != nil {
		return nil, fmt.Errorf("decode getRegistrationPeriod: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("unexpected getRegistrationPeriod result length: %d", len(values))
	}
	resp, ok := values[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("unexpected getRegistrationPeriod type %T", values[0])
	}
	return resp, nil
}

func (af *AliasFactory) Namehash(
	ctx context.Context,
	name string,
) (utils.Hash, error) {
	out, err := af.call(ctx, "namehash", name)
	if err != nil {
		return utils.Hash{}, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return utils.Hash{}, err
	}

	values, err := af.ContractABI.Methods["namehash"].Outputs.Unpack(out)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("decode namehash: %w", err)
	}
	if len(values) != 1 {
		return utils.Hash{}, fmt.Errorf("unexpected namehash result length: %d", len(values))
	}

	var hash utils.Hash
	switch v := values[0].(type) {
	case [32]byte:
		hash = utils.Hash(v)
	case utils.Hash:
		hash = v
	default:
		return utils.Hash{}, fmt.Errorf("unexpected resolveFromString type %T", values[0])
	}

	return hash, nil
}

func (af *AliasFactory) Resolve(
	ctx context.Context,
	node utils.Hash,
) (utils.Hash, error) {
	out, err := af.call(ctx, "resolve", node)
	if err != nil {
		return utils.Hash{}, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return utils.Hash{}, err
	}

	values, err := af.ContractABI.Methods["resolve"].Outputs.Unpack(out)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("decode resolve: %w", err)
	}
	if len(values) != 1 {
		return utils.Hash{}, fmt.Errorf("unexpected resolve result length: %d", len(values))
	}

	var did utils.Hash
	switch v := values[0].(type) {
	case [32]byte:
		did = utils.Hash(v)
	case utils.Hash:
		did = v
	default:
		return utils.Hash{}, fmt.Errorf("unexpected resolveFromString type %T", values[0])
	}

	var zeroHash utils.Hash
	if did == zeroHash {
		return utils.Hash{}, fmt.Errorf("no DID registered to this name")
	}

	return did, nil
}

func (af *AliasFactory) ResolveFromString(
	ctx context.Context,
	name string,
) (utils.Hash, error) {
	out, err := af.call(ctx, "resolveFromString", name)
	if err != nil {
		return utils.Hash{}, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return utils.Hash{}, err
	}

	values, err := af.ContractABI.Methods["resolveFromString"].Outputs.Unpack(out)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("decode resolveFromString: %w", err)
	}
	if len(values) != 1 {
		return utils.Hash{}, fmt.Errorf("unexpected resolveFromString result length: %d", len(values))
	}

	var did utils.Hash
	switch v := values[0].(type) {
	case [32]byte:
		did = utils.Hash(v)
	case utils.Hash:
		did = v
	default:
		return utils.Hash{}, fmt.Errorf("unexpected resolveFromString type %T", values[0])
	}

	var zeroHash utils.Hash
	if did == zeroHash {
		return utils.Hash{}, fmt.Errorf("no DID registered to this name")
	}

	return did, nil
}

func (af *AliasFactory) GetAliasStatus(
	ctx context.Context,
	node utils.Hash,
) (*big.Int, bool, bool, error) {
	out, err := af.call(ctx, "getAliasStatus", node)
	if err != nil {
		return nil, false, false, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return nil, false, false, err
	}

	values, err := af.ContractABI.Methods["getAliasStatus"].Outputs.Unpack(out)
	if err != nil {
		return nil, false, false, fmt.Errorf("decode getAliasStatus: %w", err)
	}
	if len(values) != 3 {
		return nil, false, false, fmt.Errorf("unexpected getAliasStatus result length: %d", len(values))
	}

	expired, ok := values[0].(*big.Int)
	if !ok {
		return nil, false, false, fmt.Errorf("unexpected expired type %T", values[0])
	}
	isRevoked, ok := values[1].(bool)
	if !ok {
		return nil, false, false, fmt.Errorf("unexpected isRevoked type %T", values[1])
	}
	isValid, ok := values[2].(bool)
	if !ok {
		return nil, false, false, fmt.Errorf("unexpected isValid type %T", values[2])
	}

	return expired, isRevoked, isValid, nil
}

func (af *AliasFactory) GetAliasStatusFromString(
	ctx context.Context,
	name string,
) (*big.Int, bool, bool, error) {
	out, err := af.call(ctx, "getAliasStatusFromString", name)
	if err != nil {
		return nil, false, false, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return nil, false, false, err
	}

	values, err := af.ContractABI.Methods["getAliasStatusFromString"].Outputs.Unpack(out)
	if err != nil {
		return nil, false, false, fmt.Errorf("decode getAliasStatusFromString: %w", err)
	}
	if len(values) != 3 {
		return nil, false, false, fmt.Errorf("unexpected getAliasStatusFromString result length: %d", len(values))
	}

	expired, ok := values[0].(*big.Int)
	if !ok {
		return nil, false, false, fmt.Errorf("unexpected expired type %T", values[0])
	}
	isRevoked, ok := values[1].(bool)
	if !ok {
		return nil, false, false, fmt.Errorf("unexpected isRevoked type %T", values[1])
	}
	isValid, ok := values[2].(bool)
	if !ok {
		return nil, false, false, fmt.Errorf("unexpected isValid type %T", values[2])
	}

	return expired, isRevoked, isValid, nil
}

func (af *AliasFactory) GetOwner(
	ctx context.Context,
	node utils.Hash,
) (string, error) {
	out, err := af.call(ctx, "getOwner", node)
	if err != nil {
		return "", err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return "", err
	}

	values, err := af.ContractABI.Methods["getOwner"].Outputs.Unpack(out)
	if err != nil {
		return "", fmt.Errorf("decode getOwner: %w", err)
	}
	if len(values) != 1 {
		return "", fmt.Errorf("unexpected getOwner result length: %d", len(values))
	}
	address, ok := values[0].(utils.Address)
	if !ok {
		return "", fmt.Errorf("unexpected getOwner type %T", values[0])
	}
	return address.Hex(), nil
}

func (af *AliasFactory) GetOwnerFromString(
	ctx context.Context,
	name string,
) (string, error) {
	out, err := af.call(ctx, "getOwnerFromString", name)
	if err != nil {
		return "", err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return "", err
	}

	values, err := af.ContractABI.Methods["getOwnerFromString"].Outputs.Unpack(out)
	if err != nil {
		return "", fmt.Errorf("decode getOwnerFromString: %w", err)
	}
	if len(values) != 1 {
		return "", fmt.Errorf("unexpected getOwnerFromString result length: %d", len(values))
	}
	address, ok := values[0].(utils.Address)
	if !ok {
		return "", fmt.Errorf("unexpected getOwnerFromString type %T", values[0])
	}
	return address.Hex(), nil
}

func (af *AliasFactory) GetDID(
	ctx context.Context,
	node utils.Hash,
) (utils.Hash, error) {
	out, err := af.call(ctx, "getDID", node)
	if err != nil {
		return utils.Hash{}, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return utils.Hash{}, err
	}

	values, err := af.ContractABI.Methods["getDID"].Outputs.Unpack(out)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("decode getDID: %w", err)
	}
	if len(values) != 1 {
		return utils.Hash{}, fmt.Errorf("unexpected getDID result length: %d", len(values))
	}

	var did utils.Hash
	switch v := values[0].(type) {
	case [32]byte:
		did = utils.Hash(v)
	case utils.Hash:
		did = v
	default:
		return utils.Hash{}, fmt.Errorf("unexpected getDID type %T", values[0])
	}

	var zeroHash utils.Hash
	if did == zeroHash {
		return utils.Hash{}, fmt.Errorf("no DID registered to this name")
	}

	return did, nil
}

func (af *AliasFactory) GetDIDFromString(
	ctx context.Context,
	name string,
) (utils.Hash, error) {
	out, err := af.call(ctx, "getDIDFromString", name)
	if err != nil {
		return utils.Hash{}, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return utils.Hash{}, err
	}

	values, err := af.ContractABI.Methods["getDIDFromString"].Outputs.Unpack(out)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("decode getDIDFromString: %w", err)
	}
	if len(values) != 1 {
		return utils.Hash{}, fmt.Errorf("unexpected getDIDFromString result length: %d", len(values))
	}

	var did utils.Hash
	switch v := values[0].(type) {
	case [32]byte:
		did = utils.Hash(v)
	case utils.Hash:
		did = v
	default:
		return utils.Hash{}, fmt.Errorf("unexpected getDIDFromString type %T", values[0])
	}

	var zeroHash utils.Hash
	if did == zeroHash {
		return utils.Hash{}, fmt.Errorf("no DID registered to this name")
	}

	return did, nil
}

func (af *AliasFactory) GetTLDOwner(
	ctx context.Context,
	tld string,
) (string, error) {
	out, err := af.call(ctx, "getTLDOwner", tld)
	if err != nil {
		return "", err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return "", err
	}

	values, err := af.ContractABI.Methods["getTLDOwner"].Outputs.Unpack(out)
	if err != nil {
		return "", fmt.Errorf("decode getTLDOwner: %w", err)
	}
	if len(values) != 1 {
		return "", fmt.Errorf("unexpected getTLDOwner result length: %d", len(values))
	}
	address, ok := values[0].(utils.Address)
	if !ok {
		return "", fmt.Errorf("unexpected getTLDOwner type %T", values[0])
	}
	return address.Hex(), nil
}

func (af *AliasFactory) GetAddress() utils.Address {
	return af.Address
}

func (af *AliasFactory) GetBlockchain() *blockchain.Blockchain {
	return af.blockchain
}
