// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transfersRestrictedBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transferRestrictionLiftedNoLaterThan\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintingRestrictedBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintMaxPercent\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegatedPowerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"TransferAllowlistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transfersRestrictedBefore\",\"type\":\"uint256\"}],\"name\":\"TransfersRestrictedBeforeUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATE_BY_TYPE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_VERSION\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_MAX_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_MIN_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_RESTRICTION_LIFTED_NO_LATER_THAN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_mintingRestrictedBefore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_propositionPowerDelegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_propositionPowerSnapshots\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_propositionPowerSnapshotsCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenTransferAllowlist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_totalSupplySnapshots\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalSupplySnapshotsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_transfersRestrictedBefore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_votingDelegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_votingSnapshots\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_votingSnapshotsCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addressesToAdd\",\"type\":\"address[]\"}],\"name\":\"addToTokenTransferAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"delegateByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateByTypeBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getDelegateeByType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addressesToRemove\",\"type\":\"address[]\"}],\"name\":\"removeFromTokenTransferAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transfersRestrictedBefore\",\"type\":\"uint256\"}],\"name\":\"updateTransfersRestrictedBefore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MainBin is the compiled bytecode used for deploying new contracts.
var MainBin = "0x60e06040523480156200001157600080fd5b506040516200622738038062006227833981810160405260a08110156200003757600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291905050506040518060400160405280600481526020017f64596458000000000000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f44594458000000000000000000000000000000000000000000000000000000008152508160039080519060200190620000f492919062001134565b5080600490805190602001906200010d92919062001134565b506012600560006101000a81548160ff021916908360ff160217905550505060006200013e6200046c60201b60201c565b905080600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35060004690507f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6040518060400160405280600481526020017f6459645800000000000000000000000000000000000000000000000000000000815250805190602001206040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250805190602001208330604051602001808681526020018581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff16815260200195505050505050604051602081830303815290604052805190602001206080818152505042851162000338576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180620062026025913960400191505060405180910390fd5b8385111562000393576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180620061e06022913960400191505060405180910390fd5b428311620003ed576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180620061bd6023913960400191505060405180910390fd5b846011819055508360c08181525050826010819055508160a0818152505062000429866b033b2e3c9fd0803ce80000006200047460201b60201c565b7fd7b9c9321b627ff004969698b4616502d6b7305a588e685489e91c46fca509a9856040518082815260200191505060405180910390a1505050505050620011ea565b600033905090565b6200048b8282620005a060201b6200329f1760201c565b6000600e54905060004390506000620004a96200077e60201b60201c565b90506040518060400160405280836fffffffffffffffffffffffffffffffff168152602001826fffffffffffffffffffffffffffffffff16815250600d600085815260200190815260200160002060008201518160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060208201518160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550905050620005936001846200078860201b620034661790919060201c565b600e819055505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000644576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b62000658600083836200081160201b60201c565b62000674816002546200078860201b620034661790919060201c565b600281905550620006d2816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546200078860201b620034661790919060201c565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b6000600254905090565b60008082840190508381101562000807576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600062000826846009620008a060201b60201c565b905060006200083d846009620008a060201b60201c565b90506200085482828560006200094f60201b60201c565b60006200086986600c620008a060201b60201c565b905060006200088086600c620008a060201b60201c565b90506200089782828760016200094f60201b60201c565b50505050505050565b6000808260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562000944578391505062000949565b809150505b92915050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156200098a5762000d1a565b6000806200099e8362000d2060201b60201c565b5091509150600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161462000b5d576000808260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000811462000ab1578360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600060018303815260200190815260200160002060000160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16915062000ac5565b62000ac28862000d6c60201b60201c565b91505b600062000ae1878462000db460201b620034ee1790919060201c565b905062000af785858b8462000e0660201b60201c565b8873ffffffffffffffffffffffffffffffffffffffff167fa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f82886040518083815260200182600181111562000b4857fe5b81526020019250505060405180910390a25050505b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161462000d17576000808260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000811462000c6b578360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600060018303815260200190815260200160002060000160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16915062000c7f565b62000c7c8762000d6c60201b60201c565b91505b600062000c9b87846200078860201b620034661790919060201c565b905062000cb185858a8462000e0660201b60201c565b8773ffffffffffffffffffffffffffffffffffffffff167fa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f82886040518083815260200182600181111562000d0257fe5b81526020019250505060405180910390a25050505b50505b50505050565b600080600080600181111562000d3257fe5b84600181111562000d3f57fe5b141562000d585760076008600992509250925062000d65565b600a600b600c9250925092505b9193909250565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600062000dfe83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506200107060201b60201c565b905092915050565b600043905060008460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000821415801562000ef95750826fffffffffffffffffffffffffffffffff1681600060018503815260200190815260200160002060000160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16145b1562000f54578381600060018503815260200190815260200160002060000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555062001067565b6040518060400160405280846fffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff1681525081600084815260200190815260200160002060008201518160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060208201518160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550905050600182018660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b50505050505050565b600083831115829062001121576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015620010e5578082015181840152602081019050620010c8565b50505050905090810190601f168015620011135780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826200116c5760008555620011b8565b82601f106200118757805160ff1916838001178555620011b8565b82800160010185558215620011b8579182015b82811115620011b75782518255916020019190600101906200119a565b5b509050620011c79190620011cb565b5090565b5b80821115620011e6576000816000905550600101620011cc565b5090565b60805160a05160c051614f886200123560003980612378528061280352806137395250806115aa528061180b525080611389528061251e52806129865280612fce5250614f886000f3fe608060405234801561001057600080fd5b50600436106102a05760003560e01c8063715018a611610167578063c2ffbb91116100ce578063dd62ed3e11610087578063dd62ed3e14610f99578063e1b11da414611011578063eccec5a81461102f578063f0eb7f8e146110b2578063f2fde38b14611120578063f713d8a814611164576102a0565b8063c2ffbb9114610d8b578063c3cda52014610dfa578063cdfeb0c114610e73578063d17deb9f14610e91578063d505accf14610eaf578063dc937e1c14610f48576102a0565b806395d89b411161012057806395d89b4114610b8f578063a457c2d714610c12578063a9059cbb14610c76578063aa9fbe0214610cda578063ad36dafd14610cf8578063b2f4201d14610d26576102a0565b8063715018a6146109af5780637bb73c97146109b95780637ecebe0014610a115780638aaee23414610a695780638d48f4f114610ae25780638da5cb5b14610b5b576102a0565b806340c10f191161020b5780635c19a95c116101c45780635c19a95c146107cf5780635f3b687614610813578063657c7a85146108a057806369e3b0d0146108be5780636f50458d146108dc57806370a0823114610957576102a0565b806340c10f19146105b657806341cbf54a1461060457806348032ec114610622578063537f215c1461067c578063549aa8a3146106ea5780635b3cc0cf14610742576102a0565b80632b4e1a5b1161025d5780632b4e1a5b146104b95780632ff2e9dc146104d757806330adf81f146104f5578063313ce567146105135780633644e515146105345780633950935114610552576102a0565b806306fdde03146102a5578063095ea7b31461032857806313929bbe1461038c57806318160ddd146103f95780631ff06fdf1461041757806323b872dd14610435575b600080fd5b6102ad6111ea565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102ed5780820151818401526020810190506102d2565b50505050905090810190601f16801561031a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103746004803603604081101561033e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061128c565b60405180821515815260200191505060405180910390f35b6103b8600480360360208110156103a257600080fd5b81019080803590602001909291905050506112aa565b60405180836fffffffffffffffffffffffffffffffff168152602001826fffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b610401611306565b6040518082815260200191505060405180910390f35b61041f611310565b6040518082815260200191505060405180910390f35b6104a16004803603606081101561044b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611316565b60405180821515815260200191505060405180910390f35b6104c1611336565b6040518082815260200191505060405180910390f35b6104df61133c565b6040518082815260200191505060405180910390f35b6104fd61134c565b6040518082815260200191505060405180910390f35b61051b611370565b604051808260ff16815260200191505060405180910390f35b61053c611387565b6040518082815260200191505060405180910390f35b61059e6004803603604081101561056857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506113ab565b60405180821515815260200191505060405180910390f35b610602600480360360408110156105cc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061145e565b005b61060c61168d565b6040518082815260200191505060405180910390f35b6106646004803603602081101561063857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116b1565b60405180821515815260200191505060405180910390f35b6106be6004803603602081101561069257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116d1565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61072c6004803603602081101561070057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611704565b6040518082815260200191505060405180910390f35b61078e6004803603604081101561075857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061171c565b60405180836fffffffffffffffffffffffffffffffff168152602001826fffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b610811600480360360208110156107e557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611785565b005b61085f6004803603604081101561082957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506117a0565b60405180836fffffffffffffffffffffffffffffffff168152602001826fffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b6108a8611809565b6040518082815260200191505060405180910390f35b6108c661182d565b6040518082815260200191505060405180910390f35b61092b600480360360408110156108f257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff169060200190929190505050611835565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6109996004803603602081101561096d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611858565b6040518082815260200191505060405180910390f35b6109b76118a0565b005b6109fb600480360360208110156109cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a2b565b6040518082815260200191505060405180910390f35b610a5360048036036020811015610a2757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a43565b6040518082815260200191505060405180910390f35b610ae060048036036020811015610a7f57600080fd5b8101908080359060200190640100000000811115610a9c57600080fd5b820183602082011115610aae57600080fd5b80359060200191846020830284011164010000000083111715610ad057600080fd5b9091929391929390505050611a8c565b005b610b5960048036036020811015610af857600080fd5b8101908080359060200190640100000000811115610b1557600080fd5b820183602082011115610b2757600080fd5b80359060200191846020830284011164010000000083111715610b4957600080fd5b9091929391929390505050611d41565b005b610b63611ff5565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610b9761201f565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610bd7578082015181840152602081019050610bbc565b50505050905090810190601f168015610c045780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610c5e60048036036040811015610c2857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506120c1565b60405180821515815260200191505060405180910390f35b610cc260048036036040811015610c8c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061218e565b60405180821515815260200191505060405180910390f35b610ce26121b3565b6040518082815260200191505060405180910390f35b610d2460048036036020811015610d0e57600080fd5b81019080803590602001909291905050506121d7565b005b610d7560048036036040811015610d3c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff16906020019092919050505061244e565b6040518082815260200191505060405180910390f35b610de460048036036060811015610da157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803560ff169060200190929190505050612477565b6040518082815260200191505060405180910390f35b610e71600480360360c0811015610e1057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803560ff16906020019092919080359060200190929190803590602001909291905050506124a1565b005b610e7b6127fb565b6040518082815260200191505060405180910390f35b610e99612801565b6040518082815260200191505060405180910390f35b610f46600480360360e0811015610ec557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803560ff1690602001909291908035906020019092919080359060200190929190505050612825565b005b610f9760048036036040811015610f5e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff169060200190929190505050612c08565b005b610ffb60048036036040811015610faf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612c17565b6040518082815260200191505060405180910390f35b611019612c9e565b6040518082815260200191505060405180910390f35b611037612cc2565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561107757808201518184015260208101905061105c565b50505050905090810190601f1680156110a45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6110f4600480360360208110156110c857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612cfb565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6111626004803603602081101561113657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612d2e565b005b6111e8600480360360e081101561117a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff1690602001909291908035906020019092919080359060200190929190803560ff1690602001909291908035906020019092919080359060200190929190505050612f3e565b005b606060038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112825780601f1061125757610100808354040283529160200191611282565b820191906000526020600020905b81548152906001019060200180831161126557829003601f168201915b5050505050905090565b60006112a0611299613538565b8484613540565b6001905092915050565b600d6020528060005260406000206000915090508060000160009054906101000a90046fffffffffffffffffffffffffffffffff16908060000160109054906101000a90046fffffffffffffffffffffffffffffffff16905082565b6000600254905090565b60115481565b60006113228484613737565b61132d848484613884565b90509392505050565b600e5481565b6b033b2e3c9fd0803ce800000081565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b6000600560009054906101000a900460ff16905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006114546113b8613538565b8461144f85600160006113c9613538565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461346690919063ffffffff16565b613540565b6001905092915050565b611466613538565b73ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611528576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6010544210156115a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f4d494e545f544f4f5f4541524c5900000000000000000000000000000000000081525060200191505060405180910390fd5b6115ed60646115df7f00000000000000000000000000000000000000000000000000000000000000006115d1611306565b61395d90919063ffffffff16565b6139e390919063ffffffff16565b811115611662576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f4d41585f4d494e545f455843454544454400000000000000000000000000000081525060200191505060405180910390fd5b6116796301e133804261346690919063ffffffff16565b6010819055506116898282613a2d565b5050565b7f9a9a49b990ba9bb39f8048c490a40ab25c18f55d208d5fbcf958261a9b48716d81565b600f6020528060005260406000206000915054906101000a900460ff1681565b60096020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600b6020528060005260406000206000915090505481565b6007602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a90046fffffffffffffffffffffffffffffffff16908060000160109054906101000a90046fffffffffffffffffffffffffffffffff16905082565b61179133826000613b3d565b61179d33826001613b3d565b50565b600a602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a90046fffffffffffffffffffffffffffffffff16908060000160109054906101000a90046fffffffffffffffffffffffffffffffff16905082565b7f000000000000000000000000000000000000000000000000000000000000000081565b6301e1338081565b60008061184183613d0b565b9250505061184f8482613d53565b91505092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6118a8613538565b73ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461196a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60086020528060005260406000206000915090505481565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b611a94613538565b73ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611b56576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b60005b82829050811015611d3c57600f6000848484818110611b7457fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611c2f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180614e986024913960400191505060405180910390fd5b6001600f6000858585818110611c4157fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f80e7b8bab7ab72e47957c0c472ede1b89bb3e7d2ba30cd37c2d6b000b49a960a838383818110611cdc57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff166001604051808373ffffffffffffffffffffffffffffffffffffffff16815260200182151581526020019250505060405180910390a18080600101915050611b59565b505050565b611d49613538565b73ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611e0b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b60005b82829050811015611ff057600f6000848484818110611e2957fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611ee3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180614ee0602c913960400191505060405180910390fd5b6000600f6000858585818110611ef557fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f80e7b8bab7ab72e47957c0c472ede1b89bb3e7d2ba30cd37c2d6b000b49a960a838383818110611f9057fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff166000604051808373ffffffffffffffffffffffffffffffffffffffff16815260200182151581526020019250505060405180910390a18080600101915050611e0e565b505050565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156120b75780601f1061208c576101008083540402835291602001916120b7565b820191906000526020600020905b81548152906001019060200180831161209a57829003601f168201915b5050505050905090565b60006121846120ce613538565b8461217f85604051806060016040528060258152602001614f2e60259139600160006120f8613538565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054613e009092919063ffffffff16565b613540565b6001905092915050565b60006121a161219b613538565b84613737565b6121ab8383613ec0565b905092915050565b7f10d8d059343739efce7dad10d09f0806da52b252b3e6a7951920d2d6ec4102e581565b6121df613538565b73ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146122a1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6000601154905080421061231d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f5452414e534645525f5245535452494354494f4e5f454e44454400000000000081525060200191505060405180910390fd5b81811115612376576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180614f0c6022913960400191505060405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000082111561240c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f41465445525f4d41585f5452414e534645525f5245535452494354494f4e000081525060200191505060405180910390fd5b816011819055507fd7b9c9321b627ff004969698b4616502d6b7305a588e685489e91c46fca509a9826040518082815260200191505060405180910390a15050565b600080600061245c84613d0b565b509150915061246d82828743613ede565b9250505092915050565b600080600061248584613d0b565b509150915061249682828888613ede565b925050509392505050565b60007f9a9a49b990ba9bb39f8048c490a40ab25c18f55d208d5fbcf958261a9b48716d878787604051602001808581526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200194505050505060405160208183030381529060405280519060200120905060007f00000000000000000000000000000000000000000000000000000000000000008260405160200180807f190100000000000000000000000000000000000000000000000000000000000081525060020183815260200182815260200192505050604051602081830303815290604052805190602001209050600060018287878760405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156125f1573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156126a0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f494e56414c49445f5349474e415455524500000000000000000000000000000081525060200191505060405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050558814612762576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f494e56414c49445f4e4f4e43450000000000000000000000000000000000000081525060200191505060405180910390fd5b864211156127d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f494e56414c49445f45585049524154494f4e000000000000000000000000000081525060200191505060405180910390fd5b6127e4818a6000613b3d565b6127f0818a6001613b3d565b505050505050505050565b60105481565b7f000000000000000000000000000000000000000000000000000000000000000081565b600073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614156128c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f494e56414c49445f4f574e45520000000000000000000000000000000000000081525060200191505060405180910390fd5b8342111561293e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f494e56414c49445f45585049524154494f4e000000000000000000000000000081525060200191505060405180910390fd5b6000600660008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060007f00000000000000000000000000000000000000000000000000000000000000007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98a8a8a868b604051602001808781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200182815260200196505050505050506040516020818303038152906040528051906020012060405160200180807f19010000000000000000000000000000000000000000000000000000000000008152506002018381526020018281526020019250505060405160208183030381529060405280519060200120905060018186868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015612af1573d6000803e3d6000fd5b5050506020604051035173ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1614612b9b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f494e56414c49445f5349474e415455524500000000000000000000000000000081525060200191505060405180910390fd5b612baf60018361346690919063ffffffff16565b600660008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550612bfd898989613540565b505050505050505050565b612c13338383613b3d565b5050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81565b6040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525081565b600c6020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b612d36613538565b73ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612df8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612e7e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180614dbc6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60007f10d8d059343739efce7dad10d09f0806da52b252b3e6a7951920d2d6ec4102e588886001811115612f6e57fe5b8888604051602001808681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018381526020018281526020019550505050505060405160208183030381529060405280519060200120905060007f00000000000000000000000000000000000000000000000000000000000000008260405160200180807f190100000000000000000000000000000000000000000000000000000000000081525060020183815260200182815260200192505050604051602081830303815290604052805190602001209050600060018287878760405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156130a1573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415613150576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f494e56414c49445f5349474e415455524500000000000000000000000000000081525060200191505060405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050558814613212576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f494e56414c49445f4e4f4e43450000000000000000000000000000000000000081525060200191505060405180910390fd5b86421115613288576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f494e56414c49445f45585049524154494f4e000000000000000000000000000081525060200191505060405180910390fd5b613293818b8b613b3d565b50505050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415613342576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b61334e6000838361437c565b6133638160025461346690919063ffffffff16565b6002819055506133ba816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461346690919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b6000808284019050838110156134e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600061353083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250613e00565b905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156135c6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180614ebc6024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561364c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180614de26022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b7f000000000000000000000000000000000000000000000000000000000000000042108015613767575060115442105b1561388057600f60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff168061380d5750600f60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b61387f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4e4f4e5f414c4c4f574c4953545f5452414e53464552535f44495341424c454481525060200191505060405180910390fd5b5b5050565b60006138918484846143db565b6139528461389d613538565b61394d85604051806060016040528060288152602001614e4b60289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000613903613538565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054613e009092919063ffffffff16565b613540565b600190509392505050565b60008083141561397057600090506139dd565b600082840290508284828161398157fe5b04146139d8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180614e2a6021913960400191505060405180910390fd5b809150505b92915050565b6000613a2583836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061469c565b905092915050565b613a37828261329f565b6000600e54905060004390506000613a4d611306565b90506040518060400160405280836fffffffffffffffffffffffffffffffff168152602001826fffffffffffffffffffffffffffffffff16815250600d600085815260200190815260200160002060008201518160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060208201518160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550905050613b3060018461346690919063ffffffff16565b600e819055505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415613be0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f494e56414c49445f44454c45474154454500000000000000000000000000000081525060200191505060405180910390fd5b6000613beb82613d0b565b925050506000613bfa85611858565b90506000613c088684613d53565b9050848360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550613c9381868487614762565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd8660405180826001811115613cf357fe5b815260200191505060405180910390a3505050505050565b6000806000806001811115613d1c57fe5b846001811115613d2857fe5b1415613d3f57600760086009925092509250613d4c565b600a600b600c9250925092505b9193909250565b6000808260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415613df55783915050613dfa565b809150505b92915050565b6000838311158290613ead576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015613e72578082015181840152602081019050613e57565b50505050905090810190601f168015613e9f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b6000613ed4613ecd613538565b84846143db565b6001905092915050565b600043821115613f56576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f494e56414c49445f424c4f434b5f4e554d42455200000000000000000000000081525060200191505060405180910390fd5b60008460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000811415613fb357613fab84611858565b915050614374565b828660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600060018403815260200190815260200160002060000160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16116140c9578560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600060018303815260200190815260200160002060000160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16915050614374565b828660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080815260200190815260200160002060000160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16111561415b576000915050614374565b6000806001830390505b818111156142ed57600060028383038161417b57fe5b0482039050614188614d5a565b8960008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000206040518060400160405290816000820160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020016000820160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff168152505090508681600001516fffffffffffffffffffffffffffffffff1614156142b95780602001516fffffffffffffffffffffffffffffffff1695505050505050614374565b8681600001516fffffffffffffffffffffffffffffffff1610156142df578193506142e6565b6001820392505b5050614165565b8760008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060000160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1693505050505b949350505050565b6000614389846009613d53565b90506000614398846009613d53565b90506143a78282856000614762565b60006143b486600c613d53565b905060006143c386600c613d53565b90506143d28282876001614762565b50505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415614461576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180614e736025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156144e7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180614d996023913960400191505060405180910390fd5b6144f283838361437c565b61455d81604051806060016040528060268152602001614e04602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054613e009092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506145f0816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461346690919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b60008083118290614748576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561470d5780820151818401526020810190506146f2565b50505050905090810190601f16801561473a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161475457fe5b049050809150509392505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561479b57614aed565b6000806147a783613d0b565b5091509150600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161461494b576000808260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600081146148b7578360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600060018303815260200190815260200160002060000160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1691506148c3565b6148c088611858565b91505b60006148d887846134ee90919063ffffffff16565b90506148e685858b84614af3565b8873ffffffffffffffffffffffffffffffffffffffff167fa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f82886040518083815260200182600181111561493657fe5b81526020019250505060405180910390a25050505b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614614aea576000808260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008114614a56578360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600060018303815260200190815260200160002060000160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff169150614a62565b614a5f87611858565b91505b6000614a77878461346690919063ffffffff16565b9050614a8585858a84614af3565b8773ffffffffffffffffffffffffffffffffffffffff167fa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f828860405180838152602001826001811115614ad557fe5b81526020019250505060405180910390a25050505b50505b50505050565b600043905060008460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008214158015614be55750826fffffffffffffffffffffffffffffffff1681600060018503815260200190815260200160002060000160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16145b15614c3e578381600060018503815260200190815260200160002060000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550614d51565b6040518060400160405280846fffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff1681525081600084815260200190815260200160002060008201518160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060208201518160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550905050600182018660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b50505050505050565b604051806040016040528060006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff168152509056fe45524332303a207472616e7366657220746f20746865207a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7745524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f2061646472657373414444524553535f4558495354535f494e5f5452414e534645525f414c4c4f574c49535445524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373414444524553535f444f45535f4e4f545f45584953545f494e5f5452414e534645525f414c4c4f574c4953544e45575f5452414e534645525f5245535452494354494f4e5f544f4f5f4541524c5945524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220e6e2b5cc7dbdafea1051496375448d0d398424d6dd50eaf98e87aea2d5bdd37264736f6c634300070500334d494e54494e475f524553545249435445445f4245464f52455f544f4f5f4541524c594d41585f5452414e534645525f5245535452494354494f4e5f544f4f5f4541524c595452414e53464552535f524553545249435445445f4245464f52455f544f4f5f4541524c59"

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend, distributor common.Address, transfersRestrictedBefore *big.Int, transferRestrictionLiftedNoLaterThan *big.Int, mintingRestrictedBefore *big.Int, mintMaxPercent *big.Int) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MainBin), backend, distributor, transfersRestrictedBefore, transferRestrictionLiftedNoLaterThan, mintingRestrictedBefore, mintMaxPercent)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_Main *MainCaller) DELEGATEBYTYPETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "DELEGATE_BY_TYPE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_Main *MainSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _Main.Contract.DELEGATEBYTYPETYPEHASH(&_Main.CallOpts)
}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_Main *MainCallerSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _Main.Contract.DELEGATEBYTYPETYPEHASH(&_Main.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_Main *MainCaller) DELEGATETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "DELEGATE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_Main *MainSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _Main.Contract.DELEGATETYPEHASH(&_Main.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_Main *MainCallerSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _Main.Contract.DELEGATETYPEHASH(&_Main.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Main *MainCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Main *MainSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Main.Contract.DOMAINSEPARATOR(&_Main.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Main *MainCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Main.Contract.DOMAINSEPARATOR(&_Main.CallOpts)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(bytes32)
func (_Main *MainCaller) EIP712DOMAIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "EIP712_DOMAIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(bytes32)
func (_Main *MainSession) EIP712DOMAIN() ([32]byte, error) {
	return _Main.Contract.EIP712DOMAIN(&_Main.CallOpts)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(bytes32)
func (_Main *MainCallerSession) EIP712DOMAIN() ([32]byte, error) {
	return _Main.Contract.EIP712DOMAIN(&_Main.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(bytes)
func (_Main *MainCaller) EIP712VERSION(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "EIP712_VERSION")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(bytes)
func (_Main *MainSession) EIP712VERSION() ([]byte, error) {
	return _Main.Contract.EIP712VERSION(&_Main.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(bytes)
func (_Main *MainCallerSession) EIP712VERSION() ([]byte, error) {
	return _Main.Contract.EIP712VERSION(&_Main.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Main *MainCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "INITIAL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Main *MainSession) INITIALSUPPLY() (*big.Int, error) {
	return _Main.Contract.INITIALSUPPLY(&_Main.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Main *MainCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _Main.Contract.INITIALSUPPLY(&_Main.CallOpts)
}

// MINTMAXPERCENT is a free data retrieval call binding the contract method 0x657c7a85.
//
// Solidity: function MINT_MAX_PERCENT() view returns(uint256)
func (_Main *MainCaller) MINTMAXPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "MINT_MAX_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTMAXPERCENT is a free data retrieval call binding the contract method 0x657c7a85.
//
// Solidity: function MINT_MAX_PERCENT() view returns(uint256)
func (_Main *MainSession) MINTMAXPERCENT() (*big.Int, error) {
	return _Main.Contract.MINTMAXPERCENT(&_Main.CallOpts)
}

// MINTMAXPERCENT is a free data retrieval call binding the contract method 0x657c7a85.
//
// Solidity: function MINT_MAX_PERCENT() view returns(uint256)
func (_Main *MainCallerSession) MINTMAXPERCENT() (*big.Int, error) {
	return _Main.Contract.MINTMAXPERCENT(&_Main.CallOpts)
}

// MINTMININTERVAL is a free data retrieval call binding the contract method 0x69e3b0d0.
//
// Solidity: function MINT_MIN_INTERVAL() view returns(uint256)
func (_Main *MainCaller) MINTMININTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "MINT_MIN_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTMININTERVAL is a free data retrieval call binding the contract method 0x69e3b0d0.
//
// Solidity: function MINT_MIN_INTERVAL() view returns(uint256)
func (_Main *MainSession) MINTMININTERVAL() (*big.Int, error) {
	return _Main.Contract.MINTMININTERVAL(&_Main.CallOpts)
}

// MINTMININTERVAL is a free data retrieval call binding the contract method 0x69e3b0d0.
//
// Solidity: function MINT_MIN_INTERVAL() view returns(uint256)
func (_Main *MainCallerSession) MINTMININTERVAL() (*big.Int, error) {
	return _Main.Contract.MINTMININTERVAL(&_Main.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Main *MainCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Main *MainSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Main.Contract.PERMITTYPEHASH(&_Main.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Main *MainCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Main.Contract.PERMITTYPEHASH(&_Main.CallOpts)
}

// TRANSFERRESTRICTIONLIFTEDNOLATERTHAN is a free data retrieval call binding the contract method 0xd17deb9f.
//
// Solidity: function TRANSFER_RESTRICTION_LIFTED_NO_LATER_THAN() view returns(uint256)
func (_Main *MainCaller) TRANSFERRESTRICTIONLIFTEDNOLATERTHAN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "TRANSFER_RESTRICTION_LIFTED_NO_LATER_THAN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TRANSFERRESTRICTIONLIFTEDNOLATERTHAN is a free data retrieval call binding the contract method 0xd17deb9f.
//
// Solidity: function TRANSFER_RESTRICTION_LIFTED_NO_LATER_THAN() view returns(uint256)
func (_Main *MainSession) TRANSFERRESTRICTIONLIFTEDNOLATERTHAN() (*big.Int, error) {
	return _Main.Contract.TRANSFERRESTRICTIONLIFTEDNOLATERTHAN(&_Main.CallOpts)
}

// TRANSFERRESTRICTIONLIFTEDNOLATERTHAN is a free data retrieval call binding the contract method 0xd17deb9f.
//
// Solidity: function TRANSFER_RESTRICTION_LIFTED_NO_LATER_THAN() view returns(uint256)
func (_Main *MainCallerSession) TRANSFERRESTRICTIONLIFTEDNOLATERTHAN() (*big.Int, error) {
	return _Main.Contract.TRANSFERRESTRICTIONLIFTEDNOLATERTHAN(&_Main.CallOpts)
}

// MintingRestrictedBefore is a free data retrieval call binding the contract method 0xcdfeb0c1.
//
// Solidity: function _mintingRestrictedBefore() view returns(uint256)
func (_Main *MainCaller) MintingRestrictedBefore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "_mintingRestrictedBefore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintingRestrictedBefore is a free data retrieval call binding the contract method 0xcdfeb0c1.
//
// Solidity: function _mintingRestrictedBefore() view returns(uint256)
func (_Main *MainSession) MintingRestrictedBefore() (*big.Int, error) {
	return _Main.Contract.MintingRestrictedBefore(&_Main.CallOpts)
}

// MintingRestrictedBefore is a free data retrieval call binding the contract method 0xcdfeb0c1.
//
// Solidity: function _mintingRestrictedBefore() view returns(uint256)
func (_Main *MainCallerSession) MintingRestrictedBefore() (*big.Int, error) {
	return _Main.Contract.MintingRestrictedBefore(&_Main.CallOpts)
}

// PropositionPowerDelegates is a free data retrieval call binding the contract method 0xf0eb7f8e.
//
// Solidity: function _propositionPowerDelegates(address ) view returns(address)
func (_Main *MainCaller) PropositionPowerDelegates(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "_propositionPowerDelegates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PropositionPowerDelegates is a free data retrieval call binding the contract method 0xf0eb7f8e.
//
// Solidity: function _propositionPowerDelegates(address ) view returns(address)
func (_Main *MainSession) PropositionPowerDelegates(arg0 common.Address) (common.Address, error) {
	return _Main.Contract.PropositionPowerDelegates(&_Main.CallOpts, arg0)
}

// PropositionPowerDelegates is a free data retrieval call binding the contract method 0xf0eb7f8e.
//
// Solidity: function _propositionPowerDelegates(address ) view returns(address)
func (_Main *MainCallerSession) PropositionPowerDelegates(arg0 common.Address) (common.Address, error) {
	return _Main.Contract.PropositionPowerDelegates(&_Main.CallOpts, arg0)
}

// PropositionPowerSnapshots is a free data retrieval call binding the contract method 0x5f3b6876.
//
// Solidity: function _propositionPowerSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Main *MainCaller) PropositionPowerSnapshots(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "_propositionPowerSnapshots", arg0, arg1)

	outstruct := new(struct {
		BlockNumber *big.Int
		Value       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PropositionPowerSnapshots is a free data retrieval call binding the contract method 0x5f3b6876.
//
// Solidity: function _propositionPowerSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Main *MainSession) PropositionPowerSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Main.Contract.PropositionPowerSnapshots(&_Main.CallOpts, arg0, arg1)
}

// PropositionPowerSnapshots is a free data retrieval call binding the contract method 0x5f3b6876.
//
// Solidity: function _propositionPowerSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Main *MainCallerSession) PropositionPowerSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Main.Contract.PropositionPowerSnapshots(&_Main.CallOpts, arg0, arg1)
}

// PropositionPowerSnapshotsCounts is a free data retrieval call binding the contract method 0x549aa8a3.
//
// Solidity: function _propositionPowerSnapshotsCounts(address ) view returns(uint256)
func (_Main *MainCaller) PropositionPowerSnapshotsCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "_propositionPowerSnapshotsCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PropositionPowerSnapshotsCounts is a free data retrieval call binding the contract method 0x549aa8a3.
//
// Solidity: function _propositionPowerSnapshotsCounts(address ) view returns(uint256)
func (_Main *MainSession) PropositionPowerSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.PropositionPowerSnapshotsCounts(&_Main.CallOpts, arg0)
}

// PropositionPowerSnapshotsCounts is a free data retrieval call binding the contract method 0x549aa8a3.
//
// Solidity: function _propositionPowerSnapshotsCounts(address ) view returns(uint256)
func (_Main *MainCallerSession) PropositionPowerSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.PropositionPowerSnapshotsCounts(&_Main.CallOpts, arg0)
}

// TokenTransferAllowlist is a free data retrieval call binding the contract method 0x48032ec1.
//
// Solidity: function _tokenTransferAllowlist(address ) view returns(bool)
func (_Main *MainCaller) TokenTransferAllowlist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "_tokenTransferAllowlist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenTransferAllowlist is a free data retrieval call binding the contract method 0x48032ec1.
//
// Solidity: function _tokenTransferAllowlist(address ) view returns(bool)
func (_Main *MainSession) TokenTransferAllowlist(arg0 common.Address) (bool, error) {
	return _Main.Contract.TokenTransferAllowlist(&_Main.CallOpts, arg0)
}

// TokenTransferAllowlist is a free data retrieval call binding the contract method 0x48032ec1.
//
// Solidity: function _tokenTransferAllowlist(address ) view returns(bool)
func (_Main *MainCallerSession) TokenTransferAllowlist(arg0 common.Address) (bool, error) {
	return _Main.Contract.TokenTransferAllowlist(&_Main.CallOpts, arg0)
}

// TotalSupplySnapshots is a free data retrieval call binding the contract method 0x13929bbe.
//
// Solidity: function _totalSupplySnapshots(uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Main *MainCaller) TotalSupplySnapshots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "_totalSupplySnapshots", arg0)

	outstruct := new(struct {
		BlockNumber *big.Int
		Value       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TotalSupplySnapshots is a free data retrieval call binding the contract method 0x13929bbe.
//
// Solidity: function _totalSupplySnapshots(uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Main *MainSession) TotalSupplySnapshots(arg0 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Main.Contract.TotalSupplySnapshots(&_Main.CallOpts, arg0)
}

// TotalSupplySnapshots is a free data retrieval call binding the contract method 0x13929bbe.
//
// Solidity: function _totalSupplySnapshots(uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Main *MainCallerSession) TotalSupplySnapshots(arg0 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Main.Contract.TotalSupplySnapshots(&_Main.CallOpts, arg0)
}

// TotalSupplySnapshotsCount is a free data retrieval call binding the contract method 0x2b4e1a5b.
//
// Solidity: function _totalSupplySnapshotsCount() view returns(uint256)
func (_Main *MainCaller) TotalSupplySnapshotsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "_totalSupplySnapshotsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplySnapshotsCount is a free data retrieval call binding the contract method 0x2b4e1a5b.
//
// Solidity: function _totalSupplySnapshotsCount() view returns(uint256)
func (_Main *MainSession) TotalSupplySnapshotsCount() (*big.Int, error) {
	return _Main.Contract.TotalSupplySnapshotsCount(&_Main.CallOpts)
}

// TotalSupplySnapshotsCount is a free data retrieval call binding the contract method 0x2b4e1a5b.
//
// Solidity: function _totalSupplySnapshotsCount() view returns(uint256)
func (_Main *MainCallerSession) TotalSupplySnapshotsCount() (*big.Int, error) {
	return _Main.Contract.TotalSupplySnapshotsCount(&_Main.CallOpts)
}

// TransfersRestrictedBefore is a free data retrieval call binding the contract method 0x1ff06fdf.
//
// Solidity: function _transfersRestrictedBefore() view returns(uint256)
func (_Main *MainCaller) TransfersRestrictedBefore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "_transfersRestrictedBefore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransfersRestrictedBefore is a free data retrieval call binding the contract method 0x1ff06fdf.
//
// Solidity: function _transfersRestrictedBefore() view returns(uint256)
func (_Main *MainSession) TransfersRestrictedBefore() (*big.Int, error) {
	return _Main.Contract.TransfersRestrictedBefore(&_Main.CallOpts)
}

// TransfersRestrictedBefore is a free data retrieval call binding the contract method 0x1ff06fdf.
//
// Solidity: function _transfersRestrictedBefore() view returns(uint256)
func (_Main *MainCallerSession) TransfersRestrictedBefore() (*big.Int, error) {
	return _Main.Contract.TransfersRestrictedBefore(&_Main.CallOpts)
}

// VotingDelegates is a free data retrieval call binding the contract method 0x537f215c.
//
// Solidity: function _votingDelegates(address ) view returns(address)
func (_Main *MainCaller) VotingDelegates(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "_votingDelegates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingDelegates is a free data retrieval call binding the contract method 0x537f215c.
//
// Solidity: function _votingDelegates(address ) view returns(address)
func (_Main *MainSession) VotingDelegates(arg0 common.Address) (common.Address, error) {
	return _Main.Contract.VotingDelegates(&_Main.CallOpts, arg0)
}

// VotingDelegates is a free data retrieval call binding the contract method 0x537f215c.
//
// Solidity: function _votingDelegates(address ) view returns(address)
func (_Main *MainCallerSession) VotingDelegates(arg0 common.Address) (common.Address, error) {
	return _Main.Contract.VotingDelegates(&_Main.CallOpts, arg0)
}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Main *MainCaller) VotingSnapshots(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "_votingSnapshots", arg0, arg1)

	outstruct := new(struct {
		BlockNumber *big.Int
		Value       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Main *MainSession) VotingSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Main.Contract.VotingSnapshots(&_Main.CallOpts, arg0, arg1)
}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Main *MainCallerSession) VotingSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Main.Contract.VotingSnapshots(&_Main.CallOpts, arg0, arg1)
}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_Main *MainCaller) VotingSnapshotsCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "_votingSnapshotsCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_Main *MainSession) VotingSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.VotingSnapshotsCounts(&_Main.CallOpts, arg0)
}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_Main *MainCallerSession) VotingSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.VotingSnapshotsCounts(&_Main.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Main *MainCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Main *MainSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Main.Contract.Allowance(&_Main.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Main *MainCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Main.Contract.Allowance(&_Main.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Main *MainCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Main *MainSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Main.Contract.BalanceOf(&_Main.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Main *MainCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Main.Contract.BalanceOf(&_Main.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Main *MainCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Main *MainSession) Decimals() (uint8, error) {
	return _Main.Contract.Decimals(&_Main.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Main *MainCallerSession) Decimals() (uint8, error) {
	return _Main.Contract.Decimals(&_Main.CallOpts)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_Main *MainCaller) GetDelegateeByType(opts *bind.CallOpts, delegator common.Address, delegationType uint8) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getDelegateeByType", delegator, delegationType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_Main *MainSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _Main.Contract.GetDelegateeByType(&_Main.CallOpts, delegator, delegationType)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_Main *MainCallerSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _Main.Contract.GetDelegateeByType(&_Main.CallOpts, delegator, delegationType)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_Main *MainCaller) GetPowerAtBlock(opts *bind.CallOpts, user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getPowerAtBlock", user, blockNumber, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_Main *MainSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _Main.Contract.GetPowerAtBlock(&_Main.CallOpts, user, blockNumber, delegationType)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_Main *MainCallerSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _Main.Contract.GetPowerAtBlock(&_Main.CallOpts, user, blockNumber, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_Main *MainCaller) GetPowerCurrent(opts *bind.CallOpts, user common.Address, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getPowerCurrent", user, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_Main *MainSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _Main.Contract.GetPowerCurrent(&_Main.CallOpts, user, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_Main *MainCallerSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _Main.Contract.GetPowerCurrent(&_Main.CallOpts, user, delegationType)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainSession) Name() (string, error) {
	return _Main.Contract.Name(&_Main.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainCallerSession) Name() (string, error) {
	return _Main.Contract.Name(&_Main.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Main *MainCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Main *MainSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Main.Contract.Nonces(&_Main.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Main *MainCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Main.Contract.Nonces(&_Main.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCallerSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainSession) Symbol() (string, error) {
	return _Main.Contract.Symbol(&_Main.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainCallerSession) Symbol() (string, error) {
	return _Main.Contract.Symbol(&_Main.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainSession) TotalSupply() (*big.Int, error) {
	return _Main.Contract.TotalSupply(&_Main.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainCallerSession) TotalSupply() (*big.Int, error) {
	return _Main.Contract.TotalSupply(&_Main.CallOpts)
}

// AddToTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8aaee234.
//
// Solidity: function addToTokenTransferAllowlist(address[] addressesToAdd) returns()
func (_Main *MainTransactor) AddToTokenTransferAllowlist(opts *bind.TransactOpts, addressesToAdd []common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addToTokenTransferAllowlist", addressesToAdd)
}

// AddToTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8aaee234.
//
// Solidity: function addToTokenTransferAllowlist(address[] addressesToAdd) returns()
func (_Main *MainSession) AddToTokenTransferAllowlist(addressesToAdd []common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddToTokenTransferAllowlist(&_Main.TransactOpts, addressesToAdd)
}

// AddToTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8aaee234.
//
// Solidity: function addToTokenTransferAllowlist(address[] addressesToAdd) returns()
func (_Main *MainTransactorSession) AddToTokenTransferAllowlist(addressesToAdd []common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddToTokenTransferAllowlist(&_Main.TransactOpts, addressesToAdd)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Main *MainTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Main *MainSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Approve(&_Main.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Main *MainTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Approve(&_Main.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Main *MainTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Main *MainSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Main.Contract.DecreaseAllowance(&_Main.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Main *MainTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Main.Contract.DecreaseAllowance(&_Main.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Main *MainTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Main *MainSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Main.Contract.Delegate(&_Main.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Main *MainTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Main.Contract.Delegate(&_Main.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.Contract.DelegateBySig(&_Main.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.Contract.DelegateBySig(&_Main.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_Main *MainTransactor) DelegateByType(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "delegateByType", delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_Main *MainSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _Main.Contract.DelegateByType(&_Main.TransactOpts, delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_Main *MainTransactorSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _Main.Contract.DelegateByType(&_Main.TransactOpts, delegatee, delegationType)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainTransactor) DelegateByTypeBySig(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "delegateByTypeBySig", delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.Contract.DelegateByTypeBySig(&_Main.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainTransactorSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.Contract.DelegateByTypeBySig(&_Main.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Main *MainTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Main *MainSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Main.Contract.IncreaseAllowance(&_Main.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Main *MainTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Main.Contract.IncreaseAllowance(&_Main.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_Main *MainTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "mint", recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_Main *MainSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Mint(&_Main.TransactOpts, recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_Main *MainTransactorSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Mint(&_Main.TransactOpts, recipient, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.Contract.Permit(&_Main.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Main *MainTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Main.Contract.Permit(&_Main.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RemoveFromTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8d48f4f1.
//
// Solidity: function removeFromTokenTransferAllowlist(address[] addressesToRemove) returns()
func (_Main *MainTransactor) RemoveFromTokenTransferAllowlist(opts *bind.TransactOpts, addressesToRemove []common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeFromTokenTransferAllowlist", addressesToRemove)
}

// RemoveFromTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8d48f4f1.
//
// Solidity: function removeFromTokenTransferAllowlist(address[] addressesToRemove) returns()
func (_Main *MainSession) RemoveFromTokenTransferAllowlist(addressesToRemove []common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveFromTokenTransferAllowlist(&_Main.TransactOpts, addressesToRemove)
}

// RemoveFromTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8d48f4f1.
//
// Solidity: function removeFromTokenTransferAllowlist(address[] addressesToRemove) returns()
func (_Main *MainTransactorSession) RemoveFromTokenTransferAllowlist(addressesToRemove []common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveFromTokenTransferAllowlist(&_Main.TransactOpts, addressesToRemove)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainSession) RenounceOwnership() (*types.Transaction, error) {
	return _Main.Contract.RenounceOwnership(&_Main.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Main.Contract.RenounceOwnership(&_Main.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Main *MainTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Main *MainSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Transfer(&_Main.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Main *MainTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Transfer(&_Main.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Main *MainTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Main *MainSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom(&_Main.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Main *MainTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom(&_Main.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// UpdateTransfersRestrictedBefore is a paid mutator transaction binding the contract method 0xad36dafd.
//
// Solidity: function updateTransfersRestrictedBefore(uint256 transfersRestrictedBefore) returns()
func (_Main *MainTransactor) UpdateTransfersRestrictedBefore(opts *bind.TransactOpts, transfersRestrictedBefore *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "updateTransfersRestrictedBefore", transfersRestrictedBefore)
}

// UpdateTransfersRestrictedBefore is a paid mutator transaction binding the contract method 0xad36dafd.
//
// Solidity: function updateTransfersRestrictedBefore(uint256 transfersRestrictedBefore) returns()
func (_Main *MainSession) UpdateTransfersRestrictedBefore(transfersRestrictedBefore *big.Int) (*types.Transaction, error) {
	return _Main.Contract.UpdateTransfersRestrictedBefore(&_Main.TransactOpts, transfersRestrictedBefore)
}

// UpdateTransfersRestrictedBefore is a paid mutator transaction binding the contract method 0xad36dafd.
//
// Solidity: function updateTransfersRestrictedBefore(uint256 transfersRestrictedBefore) returns()
func (_Main *MainTransactorSession) UpdateTransfersRestrictedBefore(transfersRestrictedBefore *big.Int) (*types.Transaction, error) {
	return _Main.Contract.UpdateTransfersRestrictedBefore(&_Main.TransactOpts, transfersRestrictedBefore)
}

// MainApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Main contract.
type MainApprovalIterator struct {
	Event *MainApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainApproval represents a Approval event raised by the Main contract.
type MainApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Main *MainFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MainApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MainApprovalIterator{contract: _Main.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Main *MainFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MainApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainApproval)
				if err := _Main.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Main *MainFilterer) ParseApproval(log types.Log) (*MainApproval, error) {
	event := new(MainApproval)
	if err := _Main.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the Main contract.
type MainDelegateChangedIterator struct {
	Event *MainDelegateChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainDelegateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainDelegateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainDelegateChanged represents a DelegateChanged event raised by the Main contract.
type MainDelegateChanged struct {
	Delegator      common.Address
	Delegatee      common.Address
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_Main *MainFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, delegatee []common.Address) (*MainDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &MainDelegateChangedIterator{contract: _Main.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_Main *MainFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *MainDelegateChanged, delegator []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainDelegateChanged)
				if err := _Main.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateChanged is a log parse operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_Main *MainFilterer) ParseDelegateChanged(log types.Log) (*MainDelegateChanged, error) {
	event := new(MainDelegateChanged)
	if err := _Main.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainDelegatedPowerChangedIterator is returned from FilterDelegatedPowerChanged and is used to iterate over the raw logs and unpacked data for DelegatedPowerChanged events raised by the Main contract.
type MainDelegatedPowerChangedIterator struct {
	Event *MainDelegatedPowerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainDelegatedPowerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainDelegatedPowerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainDelegatedPowerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainDelegatedPowerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainDelegatedPowerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainDelegatedPowerChanged represents a DelegatedPowerChanged event raised by the Main contract.
type MainDelegatedPowerChanged struct {
	User           common.Address
	Amount         *big.Int
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegatedPowerChanged is a free log retrieval operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_Main *MainFilterer) FilterDelegatedPowerChanged(opts *bind.FilterOpts, user []common.Address) (*MainDelegatedPowerChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return &MainDelegatedPowerChangedIterator{contract: _Main.contract, event: "DelegatedPowerChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatedPowerChanged is a free log subscription operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_Main *MainFilterer) WatchDelegatedPowerChanged(opts *bind.WatchOpts, sink chan<- *MainDelegatedPowerChanged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainDelegatedPowerChanged)
				if err := _Main.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatedPowerChanged is a log parse operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_Main *MainFilterer) ParseDelegatedPowerChanged(log types.Log) (*MainDelegatedPowerChanged, error) {
	event := new(MainDelegatedPowerChanged)
	if err := _Main.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Main contract.
type MainOwnershipTransferredIterator struct {
	Event *MainOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainOwnershipTransferred represents a OwnershipTransferred event raised by the Main contract.
type MainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MainOwnershipTransferredIterator{contract: _Main.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainOwnershipTransferred)
				if err := _Main.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) ParseOwnershipTransferred(log types.Log) (*MainOwnershipTransferred, error) {
	event := new(MainOwnershipTransferred)
	if err := _Main.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Main contract.
type MainTransferIterator struct {
	Event *MainTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransfer represents a Transfer event raised by the Main contract.
type MainTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Main *MainFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MainTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MainTransferIterator{contract: _Main.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Main *MainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MainTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransfer)
				if err := _Main.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Main *MainFilterer) ParseTransfer(log types.Log) (*MainTransfer, error) {
	event := new(MainTransfer)
	if err := _Main.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransferAllowlistUpdatedIterator is returned from FilterTransferAllowlistUpdated and is used to iterate over the raw logs and unpacked data for TransferAllowlistUpdated events raised by the Main contract.
type MainTransferAllowlistUpdatedIterator struct {
	Event *MainTransferAllowlistUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainTransferAllowlistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransferAllowlistUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainTransferAllowlistUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainTransferAllowlistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransferAllowlistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransferAllowlistUpdated represents a TransferAllowlistUpdated event raised by the Main contract.
type MainTransferAllowlistUpdated struct {
	Account   common.Address
	IsAllowed bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferAllowlistUpdated is a free log retrieval operation binding the contract event 0x80e7b8bab7ab72e47957c0c472ede1b89bb3e7d2ba30cd37c2d6b000b49a960a.
//
// Solidity: event TransferAllowlistUpdated(address account, bool isAllowed)
func (_Main *MainFilterer) FilterTransferAllowlistUpdated(opts *bind.FilterOpts) (*MainTransferAllowlistUpdatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "TransferAllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return &MainTransferAllowlistUpdatedIterator{contract: _Main.contract, event: "TransferAllowlistUpdated", logs: logs, sub: sub}, nil
}

// WatchTransferAllowlistUpdated is a free log subscription operation binding the contract event 0x80e7b8bab7ab72e47957c0c472ede1b89bb3e7d2ba30cd37c2d6b000b49a960a.
//
// Solidity: event TransferAllowlistUpdated(address account, bool isAllowed)
func (_Main *MainFilterer) WatchTransferAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *MainTransferAllowlistUpdated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "TransferAllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransferAllowlistUpdated)
				if err := _Main.contract.UnpackLog(event, "TransferAllowlistUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferAllowlistUpdated is a log parse operation binding the contract event 0x80e7b8bab7ab72e47957c0c472ede1b89bb3e7d2ba30cd37c2d6b000b49a960a.
//
// Solidity: event TransferAllowlistUpdated(address account, bool isAllowed)
func (_Main *MainFilterer) ParseTransferAllowlistUpdated(log types.Log) (*MainTransferAllowlistUpdated, error) {
	event := new(MainTransferAllowlistUpdated)
	if err := _Main.contract.UnpackLog(event, "TransferAllowlistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransfersRestrictedBeforeUpdatedIterator is returned from FilterTransfersRestrictedBeforeUpdated and is used to iterate over the raw logs and unpacked data for TransfersRestrictedBeforeUpdated events raised by the Main contract.
type MainTransfersRestrictedBeforeUpdatedIterator struct {
	Event *MainTransfersRestrictedBeforeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainTransfersRestrictedBeforeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransfersRestrictedBeforeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainTransfersRestrictedBeforeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainTransfersRestrictedBeforeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransfersRestrictedBeforeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransfersRestrictedBeforeUpdated represents a TransfersRestrictedBeforeUpdated event raised by the Main contract.
type MainTransfersRestrictedBeforeUpdated struct {
	TransfersRestrictedBefore *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterTransfersRestrictedBeforeUpdated is a free log retrieval operation binding the contract event 0xd7b9c9321b627ff004969698b4616502d6b7305a588e685489e91c46fca509a9.
//
// Solidity: event TransfersRestrictedBeforeUpdated(uint256 transfersRestrictedBefore)
func (_Main *MainFilterer) FilterTransfersRestrictedBeforeUpdated(opts *bind.FilterOpts) (*MainTransfersRestrictedBeforeUpdatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "TransfersRestrictedBeforeUpdated")
	if err != nil {
		return nil, err
	}
	return &MainTransfersRestrictedBeforeUpdatedIterator{contract: _Main.contract, event: "TransfersRestrictedBeforeUpdated", logs: logs, sub: sub}, nil
}

// WatchTransfersRestrictedBeforeUpdated is a free log subscription operation binding the contract event 0xd7b9c9321b627ff004969698b4616502d6b7305a588e685489e91c46fca509a9.
//
// Solidity: event TransfersRestrictedBeforeUpdated(uint256 transfersRestrictedBefore)
func (_Main *MainFilterer) WatchTransfersRestrictedBeforeUpdated(opts *bind.WatchOpts, sink chan<- *MainTransfersRestrictedBeforeUpdated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "TransfersRestrictedBeforeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransfersRestrictedBeforeUpdated)
				if err := _Main.contract.UnpackLog(event, "TransfersRestrictedBeforeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfersRestrictedBeforeUpdated is a log parse operation binding the contract event 0xd7b9c9321b627ff004969698b4616502d6b7305a588e685489e91c46fca509a9.
//
// Solidity: event TransfersRestrictedBeforeUpdated(uint256 transfersRestrictedBefore)
func (_Main *MainFilterer) ParseTransfersRestrictedBeforeUpdated(log types.Log) (*MainTransfersRestrictedBeforeUpdated, error) {
	event := new(MainTransfersRestrictedBeforeUpdated)
	if err := _Main.contract.UnpackLog(event, "TransfersRestrictedBeforeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
