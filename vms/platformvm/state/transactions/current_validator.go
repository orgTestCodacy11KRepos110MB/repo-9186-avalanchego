// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package transactions

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/transactions/signed"
	"github.com/ava-labs/avalanchego/vms/platformvm/transactions/unsigned"
)

var _ currentValidator = &currentValidatorImpl{}

type currentValidator interface {
	validator

	AddValidatorTx() (*unsigned.AddValidatorTx, ids.ID)

	// Weight of delegations to this validator. Doesn't include the stake
	// provided by this validator.
	DelegatorWeight() uint64

	PotentialReward() uint64
}

type currentValidatorImpl struct {
	// delegators are sorted in order of removal.
	validatorImpl

	addValidator    signed.ValidatorAndID
	delegatorWeight uint64
	potentialReward uint64
}

func (v *currentValidatorImpl) AddValidatorTx() (*unsigned.AddValidatorTx, ids.ID) {
	return v.addValidator.UnsignedAddValidatorTx, v.addValidator.TxID
}

func (v *currentValidatorImpl) DelegatorWeight() uint64 {
	return v.delegatorWeight
}

func (v *currentValidatorImpl) PotentialReward() uint64 {
	return v.potentialReward
}