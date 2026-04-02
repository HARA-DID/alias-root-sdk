package aliasfactory

import (
	"math/big"

	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type RegistrationPeriod uint8

const (
	OneYear    RegistrationPeriod = 0
	TwoYears   RegistrationPeriod = 1
	ThreeYears RegistrationPeriod = 2
)

type AliasStatus struct {
	Expired   *big.Int
	IsRevoked bool
	IsValid   bool
}

type SetDIDRootStorageParams struct {
	DIDRootStorage string
}

func (p SetDIDRootStorageParams) ToArgs() []any {
	return []any{utils.HexToAddress(p.DIDRootStorage)}
}

type SetDIDOrgStorageParams struct {
	DIDOrgStorage string
}

func (p SetDIDOrgStorageParams) ToArgs() []any {
	return []any{utils.HexToAddress(p.DIDOrgStorage)}
}

type RegisterTLDParams struct {
	TLD   string
	Owner string
}

func (p RegisterTLDParams) ToArgs() []any {
	return []any{p.TLD, utils.HexToAddress(p.Owner)}
}

type RegisterDomainParams struct {
	Label  string
	TLD    string
	Period RegistrationPeriod
}

func (p RegisterDomainParams) ToArgs() []any {
	return []any{
		p.Label,
		p.TLD,
		uint8(p.Period),
	}
}

type SetDIDParams struct {
	Name string
	DID  [32]byte
}

func (p SetDIDParams) ToArgs() []any {
	return []any{p.Name, p.DID}
}

type SetDIDOrgParams struct {
	Name        string
	OrgDIDHash  [32]byte
	UserDIDHash [32]byte
}

func (p SetDIDOrgParams) ToArgs() []any {
	return []any{p.Name, p.OrgDIDHash, p.UserDIDHash}
}

type ExtendRegistrationParams struct {
	Node   [32]byte
	Period RegistrationPeriod
}

func (p ExtendRegistrationParams) ToArgs() []any {
	return []any{p.Node, uint8(p.Period)}
}

type NodeOnlyParams struct {
	Node [32]byte
}

func (p NodeOnlyParams) ToArgs() []any {
	return []any{p.Node}
}

type RegisterSubdomainParams struct {
	Label        string
	ParentDomain string
	Period       RegistrationPeriod
}

func (p RegisterSubdomainParams) ToArgs() []any {
	return []any{p.Label, p.ParentDomain, uint8(p.Period)}
}

type TransferAliasOwnershipParams struct {
	Node     [32]byte
	NewOwner string
}

func (p TransferAliasOwnershipParams) ToArgs() []any {
	return []any{p.Node, utils.HexToAddress(p.NewOwner)}
}

type TransferTLDParams struct {
	TLD      string
	NewOwner string
}

func (p TransferTLDParams) ToArgs() []any {
	return []any{p.TLD, utils.HexToAddress(p.NewOwner)}
}
