package liquidstake

import (
	"fmt"

	"github.com/cosmos/ibc-go/v8/modules/core/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	sdk "github.com/cosmos/cosmos-sdk/types"
	cmn "github.com/cosmos/evm/precompiles/common"
	keeper "github.com/cosmos/evm/x/liquidstake/keeper"
)

const (
	LiquidStakeMethod   = "liquidStake"
	StakeToLPMethod     = "stakeToLP"
	LiquidUnstakeMethod = "liquidUnstake"

	UpdateParams                = "updateParams"
	UpdateWhitelistedValidators = "updateWhitelistedValidators"
	SetModulePaused             = "setModulePaused"
)

func (p Precompile) LiquidStake(
	ctx sdk.Context,
	origin common.Address,
	contract *vm.Contract,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	bondDenom, err := p.liquidStakeKeeper.BondDenom(ctx)
	if err != nil {
		return nil, err
	}

	msg, delegatorHexAddr, err := NewMsgLiquidStake(args, bondDenom)
	if err != nil {
		return nil, err
	}

	msgSender := contract.Caller()
	if msgSender != delegatorHexAddr {
		return nil, fmt.Errorf(cmn.ErrRequesterIsNotMsgSender, msgSender.String(), delegatorHexAddr.String())
	}

	// Execute the transaction using the message server
	msgSrv := keeper.NewMsgServerImpl(p.liquidStakeKeeper)
	if _, err = msgSrv.LiquidStake(ctx, msg); err != nil {
		return nil, err
	}

	// Emit event after successful transaction
	if err := p.EmitLiquidStakeEvent(ctx, stateDB, msg, delegatorHexAddr); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p Precompile) StakeToLP(
	ctx sdk.Context,
	origin common.Address,
	contract *vm.Contract,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	bondDenom, err := p.liquidStakeKeeper.BondDenom(ctx)
	if err != nil {
		return nil, err
	}

	msg, delegatorHexAddr, err := NewMsgStakeToLP(args, bondDenom)
	if err != nil {
		return nil, err
	}

	msgSender := contract.Caller()
	if msgSender != delegatorHexAddr {
		return nil, fmt.Errorf(cmn.ErrRequesterIsNotMsgSender, msgSender.String(), delegatorHexAddr.String())
	}

	msgSrv := keeper.NewMsgServerImpl(p.liquidStakeKeeper)

	// Execute the transaction using the message server
	if _, err = msgSrv.StakeToLP(ctx, msg); err != nil {
		return nil, err
	}

	// Emit event after successful transaction
	if err := p.EmitStakeToLPEvent(ctx, stateDB, msg, delegatorHexAddr); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p Precompile) LiquidUnstake(
	ctx sdk.Context,
	origin common.Address,
	contract *vm.Contract,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	bondDenom := p.liquidStakeKeeper.LiquidBondDenom(ctx)

	msg, delegatorHexAddr, err := NewMsgLiquidUnstake(args, bondDenom)
	if err != nil {
		return nil, err
	}

	msgSender := contract.Caller()
	if msgSender != delegatorHexAddr {
		return nil, fmt.Errorf(cmn.ErrRequesterIsNotMsgSender, msgSender.String(), delegatorHexAddr.String())
	}

	msgSrv := keeper.NewMsgServerImpl(p.liquidStakeKeeper)

	// Execute the transaction using the message server
	responce, err := msgSrv.LiquidUnstake(ctx, msg)
	if err != nil {
		return nil, err
	}

	// Emit event after successful transaction
	if err := p.EmitLiquidUnstakeEvent(ctx, stateDB, msg, delegatorHexAddr); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(responce.CompletionTime.Unix())
}

func (p Precompile) UpdateParams(
	ctx sdk.Context,
	origin common.Address,
	contract *vm.Contract,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	bondDenom := p.liquidStakeKeeper.LiquidBondDenom(ctx)

	AdminAccAddr, err := sdk.AccAddressFromBech32(p.liquidStakeKeeper.GetParams(ctx).WhitelistAdminAddress)
	var AdminBytes []byte
	if err == nil {
		AdminBytes = AdminAccAddr.Bytes()
	} else {
		AdminValAddr, err := sdk.ValAddressFromBech32(p.liquidStakeKeeper.GetParams(ctx).WhitelistAdminAddress)
		if err != nil {
			return nil, err
		}
		AdminBytes = AdminValAddr.Bytes()
	}

	adminAddr := common.BytesToAddress(AdminBytes)

	if adminAddr != contract.Caller() {
		return nil, errors.ErrUnauthorized
	}

	msg, err := NewMsgUpdateParams(args, bondDenom, adminAddr)
	if err != nil {
		return nil, err
	}

	msgSrv := keeper.NewMsgServerImpl(p.liquidStakeKeeper)

	// Execute the transaction using the message server
	if _, err = msgSrv.UpdateParams(ctx, msg); err != nil {
		return nil, err
	}

	// Emit event after successful transaction
	if err := p.EmitUpdateParamsEvent(ctx, stateDB, msg); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p Precompile) UpdateWhitelistedValidators(
	ctx sdk.Context,
	origin common.Address,
	contract *vm.Contract,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	bondDenom := p.liquidStakeKeeper.LiquidBondDenom(ctx)

	AdminAccAddr, err := sdk.AccAddressFromBech32(p.liquidStakeKeeper.GetParams(ctx).WhitelistAdminAddress)
	var AdminBytes []byte
	if err == nil {
		AdminBytes = AdminAccAddr.Bytes()
	} else {
		AdminValAddr, err := sdk.ValAddressFromBech32(p.liquidStakeKeeper.GetParams(ctx).WhitelistAdminAddress)
		if err != nil {
			return nil, err
		}
		AdminBytes = AdminValAddr.Bytes()
	}

	adminAddr := common.BytesToAddress(AdminBytes)

	if adminAddr != contract.Caller() {
		return nil, errors.ErrUnauthorized
	}

	msg, err := NewMsgUpdateWhitelistedValidators(args, bondDenom, adminAddr)
	if err != nil {
		return nil, err
	}

	msgSrv := keeper.NewMsgServerImpl(p.liquidStakeKeeper)

	// Execute the transaction using the message server
	if _, err = msgSrv.UpdateWhitelistedValidators(ctx, msg); err != nil {
		return nil, err
	}

	// Emit event after successful transaction
	if err := p.EmitUpdateWhitelistedValidatorEvent(ctx, stateDB, msg); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p Precompile) SetModulePaused(
	ctx sdk.Context,
	origin common.Address,
	contract *vm.Contract,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	bondDenom := p.liquidStakeKeeper.LiquidBondDenom(ctx)

	AdminAccAddr, err := sdk.AccAddressFromBech32(p.liquidStakeKeeper.GetParams(ctx).WhitelistAdminAddress)
	var AdminBytes []byte
	if err == nil {
		AdminBytes = AdminAccAddr.Bytes()
	} else {
		AdminValAddr, err := sdk.ValAddressFromBech32(p.liquidStakeKeeper.GetParams(ctx).WhitelistAdminAddress)
		if err != nil {
			return nil, err
		}
		AdminBytes = AdminValAddr.Bytes()
	}

	adminAddr := common.BytesToAddress(AdminBytes)

	if adminAddr != contract.Caller() {
		return nil, errors.ErrUnauthorized
	}

	msg, err := NewMsgSetModulePaused(args, bondDenom, adminAddr)
	if err != nil {
		return nil, err
	}

	msgSrv := keeper.NewMsgServerImpl(p.liquidStakeKeeper)

	// Execute the transaction using the message server
	if _, err = msgSrv.SetModulePaused(ctx, msg); err != nil {
		return nil, err
	}

	// Emit event after successful transaction
	if err := p.EmitSetModulePausedEvent(ctx, stateDB, msg); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
