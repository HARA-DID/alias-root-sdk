package aliasstorage

import (
	"math/big"

	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

// SetFactoryContractParams for setFactoryContract function
type SetFactoryContractParams struct {
	FactoryContract utils.Address
}

func (p SetFactoryContractParams) ToArgs() []any {
	return []any{p.FactoryContract}
}

// AliasData represents the alias data structure (from getAliasData return)
type AliasData struct {
	Expired   *big.Int
	IsRevoked bool
}
