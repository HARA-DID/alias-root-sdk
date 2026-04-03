package aliasstorage

import (
	"context"

	"github.com/HARA-DID/hara-core-blockchain-lib/pkg/wallet"
)

func (as *AliasStorage) SetFactoryContract(
	ctx context.Context,
	wallet *wallet.Wallet,
	params SetFactoryContractParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return as.buildAndSendTx(
		ctx,
		wallet,
		"setFactoryContract",
		params,
		multipleRPCCalls,
	)
}
