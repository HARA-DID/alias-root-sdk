package aliasfactory

import (
	"context"

	"github.com/meQlause/hara-core-blockchain-lib/pkg/wallet"
)

func (af *AliasFactory) SetDIDRootStorage(
	ctx context.Context,
	wallet *wallet.Wallet,
	params SetDIDRootStorageParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"setDIDRootStorage",
		params,
		multipleRPCCalls,
	)
}

func (af *AliasFactory) SetDIDOrgStorage(
	ctx context.Context,
	wallet *wallet.Wallet,
	params SetDIDOrgStorageParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"setDIDOrgStorage",
		params,
		multipleRPCCalls,
	)
}

func (af *AliasFactory) RegisterTLD(
	ctx context.Context,
	wallet *wallet.Wallet,
	params RegisterTLDParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"registerTLD",
		params,
		multipleRPCCalls,
	)
}

func (af *AliasFactory) RegisterDomain(
	ctx context.Context,
	wallet *wallet.Wallet,
	params RegisterDomainParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"registerDomain",
		params,
		multipleRPCCalls,
	)
}

func (af *AliasFactory) SetDID(
	ctx context.Context,
	wallet *wallet.Wallet,
	params SetDIDParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"setDID",
		params,
		multipleRPCCalls,
	)
}

func (af *AliasFactory) SetDIDOrg(
	ctx context.Context,
	wallet *wallet.Wallet,
	params SetDIDOrgParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"setDIDOrg",
		params,
		multipleRPCCalls,
	)
}

func (af *AliasFactory) ExtendRegistration(
	ctx context.Context,
	wallet *wallet.Wallet,
	params ExtendRegistrationParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"extendRegistration",
		params,
		multipleRPCCalls,
	)
}

func (af *AliasFactory) RevokeAlias(
	ctx context.Context,
	wallet *wallet.Wallet,
	params NodeOnlyParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"revokeAlias",
		params,
		multipleRPCCalls,
	)
}

func (af *AliasFactory) UnrevokeAlias(
	ctx context.Context,
	wallet *wallet.Wallet,
	params NodeOnlyParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"unrevokeAlias",
		params,
		multipleRPCCalls,
	)
}

func (af *AliasFactory) RegisterSubdomain(
	ctx context.Context,
	wallet *wallet.Wallet,
	params RegisterSubdomainParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"registerSubdomain",
		params,
		multipleRPCCalls,
	)
}

func (af *AliasFactory) TransferAliasOwnership(
	ctx context.Context,
	wallet *wallet.Wallet,
	params TransferAliasOwnershipParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"transferAliasOwnership",
		params,
		multipleRPCCalls,
	)
}

func (af *AliasFactory) TransferTLD(
	ctx context.Context,
	wallet *wallet.Wallet,
	params TransferTLDParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return af.buildAndSendTx(
		ctx,
		wallet,
		"transferTLD",
		params,
		multipleRPCCalls,
	)
}
