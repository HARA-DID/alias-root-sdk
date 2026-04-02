package aliasstorage

import (
	"context"
	"fmt"

	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/contract"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/wallet"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
	"math/big"
)

type TxParams interface {
	ToArgs() []any
}

type AliasStorage struct {
	blockchain  *blockchain.Blockchain
	ContractABI utils.ABI
	Contract    *contract.Contract
	Address     utils.Address
}

func NewAliasStorage(
	contractAddress utils.Address,
	contractABI utils.ABI,
	bc *blockchain.Blockchain,
	contract *contract.Contract,
) *AliasStorage {
	return &AliasStorage{
		blockchain:  bc,
		ContractABI: contractABI,
		Contract:    contract,
		Address:     contractAddress,
	}
}

func NewAliasStorageWithHNS(
	ctx context.Context,
	hnsURI string,
	bc *blockchain.Blockchain,
) (*AliasStorage, error) {
	contract, err := bc.ContractWithHNS(ctx, hnsURI)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve contract with HNS: %w", err)
	}

	return &AliasStorage{
		blockchain:  bc,
		Contract:    contract,
		ContractABI: contract.ABI,
		Address:     contract.Address,
	}, nil
}

func (as *AliasStorage) buildAndSendTx(
	ctx context.Context,
	wallet *wallet.Wallet,
	methodName string,
	params TxParams,
	multipleRPCCalls bool,
) ([]string, error) {
	method, ok := as.ContractABI.Methods[methodName]
	if !ok {
		return nil, fmt.Errorf("method %s not found in ABI", methodName)
	}

	inputs, err := method.Inputs.Pack(params.ToArgs()...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s arguments: %w", methodName, err)
	}

	calldata := append(method.ID, inputs...)

	sender, err := wallet.GetAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet address: %w", err)
	}

	nonce, err := as.blockchain.Network.PendingNonce(ctx, sender)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending nonce: %w", err)
	}

	txParams := utils.TransactionParams{
		Nonce:    nonce,
		To:       as.Address,
		Value:    big.NewInt(0),
		GasLimit: 30000000,
		GasPrice: big.NewInt(0),
		Data:     calldata,
	}

	tx := as.blockchain.BuildTx(txParams)

	hashes, err := as.blockchain.CallContractWrite(ctx, wallet, tx, multipleRPCCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return hashes, nil
}
