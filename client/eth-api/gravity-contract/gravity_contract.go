// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LogicCallArgs is an auto generated low-level Go binding around an user-defined struct.
type LogicCallArgs struct {
	TransferAmounts        []*big.Int
	TransferTokenContracts []common.Address
	FeeAmounts             []*big.Int
	FeeTokenContracts      []common.Address
	LogicContractAddress   common.Address
	Payload                []byte
	TimeOut                *big.Int
	InvalidationId         [32]byte
	InvalidationNonce      *big.Int
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// ValsetArgs is an auto generated low-level Go binding around an user-defined struct.
type ValsetArgs struct {
	Validators   []common.Address
	Powers       []*big.Int
	ValsetNonce  *big.Int
	RewardAmount *big.Int
	RewardToken  common.Address
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_gravityId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchTimedOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectCheckpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cumulativePower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"powerThreshold\",\"type\":\"uint256\"}],\"name\":\"InsufficientPower\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidBatchNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLogicCallFees\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidLogicCallNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLogicCallTransfers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSendToCosmos\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidValsetNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LogicCallTimedOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MalformedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MalformedCurrentValidatorSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MalformedNewValidatorSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_cosmosDenom\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"ERC20DeployedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_invalidationId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_invalidationNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"LogicCallEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_destination\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"SendToCosmosEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"TransactionBatchExecutedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_newValsetNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"ValsetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cosmosDenom\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"deployERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Address\",\"type\":\"address\"}],\"name\":\"lastBatchNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_invalidation_id\",\"type\":\"bytes32\"}],\"name\":\"lastLogicCallNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_destination\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendToCosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_gravityId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"state_invalidationMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"state_lastBatchNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastEventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_currentValset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_destinations\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_batchTimeout\",\"type\":\"uint256\"}],\"name\":\"submitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_currentValset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"transferAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"transferTokenContracts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"feeTokenContracts\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"logicContractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timeOut\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"invalidationId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"invalidationNonce\",\"type\":\"uint256\"}],\"internalType\":\"structLogicCallArgs\",\"name\":\"_args\",\"type\":\"tuple\"}],\"name\":\"submitLogicCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_currentValset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_theHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"}],\"name\":\"testCheckValidatorSignatures\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_valsetArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_gravityId\",\"type\":\"bytes32\"}],\"name\":\"testMakeCheckpoint\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_newValset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_currentValset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"updateValset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Contract *ContractCaller) LastBatchNonce(opts *bind.CallOpts, _erc20Address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastBatchNonce", _erc20Address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Contract *ContractSession) LastBatchNonce(_erc20Address common.Address) (*big.Int, error) {
	return _Contract.Contract.LastBatchNonce(&_Contract.CallOpts, _erc20Address)
}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Contract *ContractCallerSession) LastBatchNonce(_erc20Address common.Address) (*big.Int, error) {
	return _Contract.Contract.LastBatchNonce(&_Contract.CallOpts, _erc20Address)
}

// LastLogicCallNonce is a free data retrieval call binding the contract method 0xc9d194d5.
//
// Solidity: function lastLogicCallNonce(bytes32 _invalidation_id) view returns(uint256)
func (_Contract *ContractCaller) LastLogicCallNonce(opts *bind.CallOpts, _invalidation_id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastLogicCallNonce", _invalidation_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLogicCallNonce is a free data retrieval call binding the contract method 0xc9d194d5.
//
// Solidity: function lastLogicCallNonce(bytes32 _invalidation_id) view returns(uint256)
func (_Contract *ContractSession) LastLogicCallNonce(_invalidation_id [32]byte) (*big.Int, error) {
	return _Contract.Contract.LastLogicCallNonce(&_Contract.CallOpts, _invalidation_id)
}

// LastLogicCallNonce is a free data retrieval call binding the contract method 0xc9d194d5.
//
// Solidity: function lastLogicCallNonce(bytes32 _invalidation_id) view returns(uint256)
func (_Contract *ContractCallerSession) LastLogicCallNonce(_invalidation_id [32]byte) (*big.Int, error) {
	return _Contract.Contract.LastLogicCallNonce(&_Contract.CallOpts, _invalidation_id)
}

// StateGravityId is a free data retrieval call binding the contract method 0xbdda81d4.
//
// Solidity: function state_gravityId() view returns(bytes32)
func (_Contract *ContractCaller) StateGravityId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "state_gravityId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateGravityId is a free data retrieval call binding the contract method 0xbdda81d4.
//
// Solidity: function state_gravityId() view returns(bytes32)
func (_Contract *ContractSession) StateGravityId() ([32]byte, error) {
	return _Contract.Contract.StateGravityId(&_Contract.CallOpts)
}

// StateGravityId is a free data retrieval call binding the contract method 0xbdda81d4.
//
// Solidity: function state_gravityId() view returns(bytes32)
func (_Contract *ContractCallerSession) StateGravityId() ([32]byte, error) {
	return _Contract.Contract.StateGravityId(&_Contract.CallOpts)
}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_Contract *ContractCaller) StateInvalidationMapping(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "state_invalidationMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_Contract *ContractSession) StateInvalidationMapping(arg0 [32]byte) (*big.Int, error) {
	return _Contract.Contract.StateInvalidationMapping(&_Contract.CallOpts, arg0)
}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_Contract *ContractCallerSession) StateInvalidationMapping(arg0 [32]byte) (*big.Int, error) {
	return _Contract.Contract.StateInvalidationMapping(&_Contract.CallOpts, arg0)
}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Contract *ContractCaller) StateLastBatchNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "state_lastBatchNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Contract *ContractSession) StateLastBatchNonces(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.StateLastBatchNonces(&_Contract.CallOpts, arg0)
}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Contract *ContractCallerSession) StateLastBatchNonces(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.StateLastBatchNonces(&_Contract.CallOpts, arg0)
}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Contract *ContractCaller) StateLastEventNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "state_lastEventNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Contract *ContractSession) StateLastEventNonce() (*big.Int, error) {
	return _Contract.Contract.StateLastEventNonce(&_Contract.CallOpts)
}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Contract *ContractCallerSession) StateLastEventNonce() (*big.Int, error) {
	return _Contract.Contract.StateLastEventNonce(&_Contract.CallOpts)
}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Contract *ContractCaller) StateLastValsetCheckpoint(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "state_lastValsetCheckpoint")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Contract *ContractSession) StateLastValsetCheckpoint() ([32]byte, error) {
	return _Contract.Contract.StateLastValsetCheckpoint(&_Contract.CallOpts)
}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Contract *ContractCallerSession) StateLastValsetCheckpoint() ([32]byte, error) {
	return _Contract.Contract.StateLastValsetCheckpoint(&_Contract.CallOpts)
}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Contract *ContractCaller) StateLastValsetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "state_lastValsetNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Contract *ContractSession) StateLastValsetNonce() (*big.Int, error) {
	return _Contract.Contract.StateLastValsetNonce(&_Contract.CallOpts)
}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Contract *ContractCallerSession) StateLastValsetNonce() (*big.Int, error) {
	return _Contract.Contract.StateLastValsetNonce(&_Contract.CallOpts)
}

// TestCheckValidatorSignatures is a free data retrieval call binding the contract method 0x00901153.
//
// Solidity: function testCheckValidatorSignatures((address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs, bytes32 _theHash, uint256 _powerThreshold) pure returns()
func (_Contract *ContractCaller) TestCheckValidatorSignatures(opts *bind.CallOpts, _currentValset ValsetArgs, _sigs []Signature, _theHash [32]byte, _powerThreshold *big.Int) error {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "testCheckValidatorSignatures", _currentValset, _sigs, _theHash, _powerThreshold)

	if err != nil {
		return err
	}

	return err

}

// TestCheckValidatorSignatures is a free data retrieval call binding the contract method 0x00901153.
//
// Solidity: function testCheckValidatorSignatures((address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs, bytes32 _theHash, uint256 _powerThreshold) pure returns()
func (_Contract *ContractSession) TestCheckValidatorSignatures(_currentValset ValsetArgs, _sigs []Signature, _theHash [32]byte, _powerThreshold *big.Int) error {
	return _Contract.Contract.TestCheckValidatorSignatures(&_Contract.CallOpts, _currentValset, _sigs, _theHash, _powerThreshold)
}

// TestCheckValidatorSignatures is a free data retrieval call binding the contract method 0x00901153.
//
// Solidity: function testCheckValidatorSignatures((address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs, bytes32 _theHash, uint256 _powerThreshold) pure returns()
func (_Contract *ContractCallerSession) TestCheckValidatorSignatures(_currentValset ValsetArgs, _sigs []Signature, _theHash [32]byte, _powerThreshold *big.Int) error {
	return _Contract.Contract.TestCheckValidatorSignatures(&_Contract.CallOpts, _currentValset, _sigs, _theHash, _powerThreshold)
}

// TestMakeCheckpoint is a free data retrieval call binding the contract method 0x01031525.
//
// Solidity: function testMakeCheckpoint((address[],uint256[],uint256,uint256,address) _valsetArgs, bytes32 _gravityId) pure returns()
func (_Contract *ContractCaller) TestMakeCheckpoint(opts *bind.CallOpts, _valsetArgs ValsetArgs, _gravityId [32]byte) error {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "testMakeCheckpoint", _valsetArgs, _gravityId)

	if err != nil {
		return err
	}

	return err

}

// TestMakeCheckpoint is a free data retrieval call binding the contract method 0x01031525.
//
// Solidity: function testMakeCheckpoint((address[],uint256[],uint256,uint256,address) _valsetArgs, bytes32 _gravityId) pure returns()
func (_Contract *ContractSession) TestMakeCheckpoint(_valsetArgs ValsetArgs, _gravityId [32]byte) error {
	return _Contract.Contract.TestMakeCheckpoint(&_Contract.CallOpts, _valsetArgs, _gravityId)
}

// TestMakeCheckpoint is a free data retrieval call binding the contract method 0x01031525.
//
// Solidity: function testMakeCheckpoint((address[],uint256[],uint256,uint256,address) _valsetArgs, bytes32 _gravityId) pure returns()
func (_Contract *ContractCallerSession) TestMakeCheckpoint(_valsetArgs ValsetArgs, _gravityId [32]byte) error {
	return _Contract.Contract.TestMakeCheckpoint(&_Contract.CallOpts, _valsetArgs, _gravityId)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_Contract *ContractTransactor) DeployERC20(opts *bind.TransactOpts, _cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deployERC20", _cosmosDenom, _name, _symbol, _decimals)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_Contract *ContractSession) DeployERC20(_cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Contract.Contract.DeployERC20(&_Contract.TransactOpts, _cosmosDenom, _name, _symbol, _decimals)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_Contract *ContractTransactorSession) DeployERC20(_cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Contract.Contract.DeployERC20(&_Contract.TransactOpts, _cosmosDenom, _name, _symbol, _decimals)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x0f212357.
//
// Solidity: function sendToCosmos(address _tokenContract, string _destination, uint256 _amount) returns()
func (_Contract *ContractTransactor) SendToCosmos(opts *bind.TransactOpts, _tokenContract common.Address, _destination string, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sendToCosmos", _tokenContract, _destination, _amount)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x0f212357.
//
// Solidity: function sendToCosmos(address _tokenContract, string _destination, uint256 _amount) returns()
func (_Contract *ContractSession) SendToCosmos(_tokenContract common.Address, _destination string, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SendToCosmos(&_Contract.TransactOpts, _tokenContract, _destination, _amount)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x0f212357.
//
// Solidity: function sendToCosmos(address _tokenContract, string _destination, uint256 _amount) returns()
func (_Contract *ContractTransactorSession) SendToCosmos(_tokenContract common.Address, _destination string, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SendToCosmos(&_Contract.TransactOpts, _tokenContract, _destination, _amount)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x8690ff98.
//
// Solidity: function submitBatch((address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_Contract *ContractTransactor) SubmitBatch(opts *bind.TransactOpts, _currentValset ValsetArgs, _sigs []Signature, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitBatch", _currentValset, _sigs, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x8690ff98.
//
// Solidity: function submitBatch((address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_Contract *ContractSession) SubmitBatch(_currentValset ValsetArgs, _sigs []Signature, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SubmitBatch(&_Contract.TransactOpts, _currentValset, _sigs, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x8690ff98.
//
// Solidity: function submitBatch((address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_Contract *ContractTransactorSession) SubmitBatch(_currentValset ValsetArgs, _sigs []Signature, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SubmitBatch(&_Contract.TransactOpts, _currentValset, _sigs, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// SubmitLogicCall is a paid mutator transaction binding the contract method 0x6941db93.
//
// Solidity: function submitLogicCall((address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs, (uint256[],address[],uint256[],address[],address,bytes,uint256,bytes32,uint256) _args) returns()
func (_Contract *ContractTransactor) SubmitLogicCall(opts *bind.TransactOpts, _currentValset ValsetArgs, _sigs []Signature, _args LogicCallArgs) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitLogicCall", _currentValset, _sigs, _args)
}

// SubmitLogicCall is a paid mutator transaction binding the contract method 0x6941db93.
//
// Solidity: function submitLogicCall((address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs, (uint256[],address[],uint256[],address[],address,bytes,uint256,bytes32,uint256) _args) returns()
func (_Contract *ContractSession) SubmitLogicCall(_currentValset ValsetArgs, _sigs []Signature, _args LogicCallArgs) (*types.Transaction, error) {
	return _Contract.Contract.SubmitLogicCall(&_Contract.TransactOpts, _currentValset, _sigs, _args)
}

// SubmitLogicCall is a paid mutator transaction binding the contract method 0x6941db93.
//
// Solidity: function submitLogicCall((address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs, (uint256[],address[],uint256[],address[],address,bytes,uint256,bytes32,uint256) _args) returns()
func (_Contract *ContractTransactorSession) SubmitLogicCall(_currentValset ValsetArgs, _sigs []Signature, _args LogicCallArgs) (*types.Transaction, error) {
	return _Contract.Contract.SubmitLogicCall(&_Contract.TransactOpts, _currentValset, _sigs, _args)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xaca6b1c1.
//
// Solidity: function updateValset((address[],uint256[],uint256,uint256,address) _newValset, (address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Contract *ContractTransactor) UpdateValset(opts *bind.TransactOpts, _newValset ValsetArgs, _currentValset ValsetArgs, _sigs []Signature) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateValset", _newValset, _currentValset, _sigs)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xaca6b1c1.
//
// Solidity: function updateValset((address[],uint256[],uint256,uint256,address) _newValset, (address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Contract *ContractSession) UpdateValset(_newValset ValsetArgs, _currentValset ValsetArgs, _sigs []Signature) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValset(&_Contract.TransactOpts, _newValset, _currentValset, _sigs)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xaca6b1c1.
//
// Solidity: function updateValset((address[],uint256[],uint256,uint256,address) _newValset, (address[],uint256[],uint256,uint256,address) _currentValset, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Contract *ContractTransactorSession) UpdateValset(_newValset ValsetArgs, _currentValset ValsetArgs, _sigs []Signature) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValset(&_Contract.TransactOpts, _newValset, _currentValset, _sigs)
}

// ContractERC20DeployedEventIterator is returned from FilterERC20DeployedEvent and is used to iterate over the raw logs and unpacked data for ERC20DeployedEvent events raised by the Contract contract.
type ContractERC20DeployedEventIterator struct {
	Event *ContractERC20DeployedEvent // Event containing the contract specifics and raw log

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
func (it *ContractERC20DeployedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractERC20DeployedEvent)
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
		it.Event = new(ContractERC20DeployedEvent)
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
func (it *ContractERC20DeployedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractERC20DeployedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractERC20DeployedEvent represents a ERC20DeployedEvent event raised by the Contract contract.
type ContractERC20DeployedEvent struct {
	CosmosDenom   string
	TokenContract common.Address
	Name          string
	Symbol        string
	Decimals      uint8
	EventNonce    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterERC20DeployedEvent is a free log retrieval operation binding the contract event 0x82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7.
//
// Solidity: event ERC20DeployedEvent(string _cosmosDenom, address indexed _tokenContract, string _name, string _symbol, uint8 _decimals, uint256 _eventNonce)
func (_Contract *ContractFilterer) FilterERC20DeployedEvent(opts *bind.FilterOpts, _tokenContract []common.Address) (*ContractERC20DeployedEventIterator, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ERC20DeployedEvent", _tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &ContractERC20DeployedEventIterator{contract: _Contract.contract, event: "ERC20DeployedEvent", logs: logs, sub: sub}, nil
}

// WatchERC20DeployedEvent is a free log subscription operation binding the contract event 0x82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7.
//
// Solidity: event ERC20DeployedEvent(string _cosmosDenom, address indexed _tokenContract, string _name, string _symbol, uint8 _decimals, uint256 _eventNonce)
func (_Contract *ContractFilterer) WatchERC20DeployedEvent(opts *bind.WatchOpts, sink chan<- *ContractERC20DeployedEvent, _tokenContract []common.Address) (event.Subscription, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ERC20DeployedEvent", _tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractERC20DeployedEvent)
				if err := _Contract.contract.UnpackLog(event, "ERC20DeployedEvent", log); err != nil {
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

// ParseERC20DeployedEvent is a log parse operation binding the contract event 0x82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7.
//
// Solidity: event ERC20DeployedEvent(string _cosmosDenom, address indexed _tokenContract, string _name, string _symbol, uint8 _decimals, uint256 _eventNonce)
func (_Contract *ContractFilterer) ParseERC20DeployedEvent(log types.Log) (*ContractERC20DeployedEvent, error) {
	event := new(ContractERC20DeployedEvent)
	if err := _Contract.contract.UnpackLog(event, "ERC20DeployedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogicCallEventIterator is returned from FilterLogicCallEvent and is used to iterate over the raw logs and unpacked data for LogicCallEvent events raised by the Contract contract.
type ContractLogicCallEventIterator struct {
	Event *ContractLogicCallEvent // Event containing the contract specifics and raw log

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
func (it *ContractLogicCallEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogicCallEvent)
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
		it.Event = new(ContractLogicCallEvent)
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
func (it *ContractLogicCallEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogicCallEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogicCallEvent represents a LogicCallEvent event raised by the Contract contract.
type ContractLogicCallEvent struct {
	InvalidationId    [32]byte
	InvalidationNonce *big.Int
	ReturnData        []byte
	EventNonce        *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogicCallEvent is a free log retrieval operation binding the contract event 0x7c2bb24f8e1b3725cb613d7f11ef97d9745cc97a0e40f730621c052d684077a1.
//
// Solidity: event LogicCallEvent(bytes32 _invalidationId, uint256 _invalidationNonce, bytes _returnData, uint256 _eventNonce)
func (_Contract *ContractFilterer) FilterLogicCallEvent(opts *bind.FilterOpts) (*ContractLogicCallEventIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogicCallEvent")
	if err != nil {
		return nil, err
	}
	return &ContractLogicCallEventIterator{contract: _Contract.contract, event: "LogicCallEvent", logs: logs, sub: sub}, nil
}

// WatchLogicCallEvent is a free log subscription operation binding the contract event 0x7c2bb24f8e1b3725cb613d7f11ef97d9745cc97a0e40f730621c052d684077a1.
//
// Solidity: event LogicCallEvent(bytes32 _invalidationId, uint256 _invalidationNonce, bytes _returnData, uint256 _eventNonce)
func (_Contract *ContractFilterer) WatchLogicCallEvent(opts *bind.WatchOpts, sink chan<- *ContractLogicCallEvent) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogicCallEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogicCallEvent)
				if err := _Contract.contract.UnpackLog(event, "LogicCallEvent", log); err != nil {
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

// ParseLogicCallEvent is a log parse operation binding the contract event 0x7c2bb24f8e1b3725cb613d7f11ef97d9745cc97a0e40f730621c052d684077a1.
//
// Solidity: event LogicCallEvent(bytes32 _invalidationId, uint256 _invalidationNonce, bytes _returnData, uint256 _eventNonce)
func (_Contract *ContractFilterer) ParseLogicCallEvent(log types.Log) (*ContractLogicCallEvent, error) {
	event := new(ContractLogicCallEvent)
	if err := _Contract.contract.UnpackLog(event, "LogicCallEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSendToCosmosEventIterator is returned from FilterSendToCosmosEvent and is used to iterate over the raw logs and unpacked data for SendToCosmosEvent events raised by the Contract contract.
type ContractSendToCosmosEventIterator struct {
	Event *ContractSendToCosmosEvent // Event containing the contract specifics and raw log

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
func (it *ContractSendToCosmosEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSendToCosmosEvent)
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
		it.Event = new(ContractSendToCosmosEvent)
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
func (it *ContractSendToCosmosEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSendToCosmosEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSendToCosmosEvent represents a SendToCosmosEvent event raised by the Contract contract.
type ContractSendToCosmosEvent struct {
	TokenContract common.Address
	Sender        common.Address
	Destination   string
	Amount        *big.Int
	EventNonce    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSendToCosmosEvent is a free log retrieval operation binding the contract event 0x9e9794dbf94b0a0aa31a480f5b38550eda7f89115ac8fbf4953fa4dd219900c9.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, string _destination, uint256 _amount, uint256 _eventNonce)
func (_Contract *ContractFilterer) FilterSendToCosmosEvent(opts *bind.FilterOpts, _tokenContract []common.Address, _sender []common.Address) (*ContractSendToCosmosEventIterator, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SendToCosmosEvent", _tokenContractRule, _senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractSendToCosmosEventIterator{contract: _Contract.contract, event: "SendToCosmosEvent", logs: logs, sub: sub}, nil
}

// WatchSendToCosmosEvent is a free log subscription operation binding the contract event 0x9e9794dbf94b0a0aa31a480f5b38550eda7f89115ac8fbf4953fa4dd219900c9.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, string _destination, uint256 _amount, uint256 _eventNonce)
func (_Contract *ContractFilterer) WatchSendToCosmosEvent(opts *bind.WatchOpts, sink chan<- *ContractSendToCosmosEvent, _tokenContract []common.Address, _sender []common.Address) (event.Subscription, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SendToCosmosEvent", _tokenContractRule, _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSendToCosmosEvent)
				if err := _Contract.contract.UnpackLog(event, "SendToCosmosEvent", log); err != nil {
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

// ParseSendToCosmosEvent is a log parse operation binding the contract event 0x9e9794dbf94b0a0aa31a480f5b38550eda7f89115ac8fbf4953fa4dd219900c9.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, string _destination, uint256 _amount, uint256 _eventNonce)
func (_Contract *ContractFilterer) ParseSendToCosmosEvent(log types.Log) (*ContractSendToCosmosEvent, error) {
	event := new(ContractSendToCosmosEvent)
	if err := _Contract.contract.UnpackLog(event, "SendToCosmosEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransactionBatchExecutedEventIterator is returned from FilterTransactionBatchExecutedEvent and is used to iterate over the raw logs and unpacked data for TransactionBatchExecutedEvent events raised by the Contract contract.
type ContractTransactionBatchExecutedEventIterator struct {
	Event *ContractTransactionBatchExecutedEvent // Event containing the contract specifics and raw log

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
func (it *ContractTransactionBatchExecutedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransactionBatchExecutedEvent)
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
		it.Event = new(ContractTransactionBatchExecutedEvent)
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
func (it *ContractTransactionBatchExecutedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransactionBatchExecutedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransactionBatchExecutedEvent represents a TransactionBatchExecutedEvent event raised by the Contract contract.
type ContractTransactionBatchExecutedEvent struct {
	BatchNonce *big.Int
	Token      common.Address
	EventNonce *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionBatchExecutedEvent is a free log retrieval operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Contract *ContractFilterer) FilterTransactionBatchExecutedEvent(opts *bind.FilterOpts, _batchNonce []*big.Int, _token []common.Address) (*ContractTransactionBatchExecutedEventIterator, error) {

	var _batchNonceRule []interface{}
	for _, _batchNonceItem := range _batchNonce {
		_batchNonceRule = append(_batchNonceRule, _batchNonceItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TransactionBatchExecutedEvent", _batchNonceRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransactionBatchExecutedEventIterator{contract: _Contract.contract, event: "TransactionBatchExecutedEvent", logs: logs, sub: sub}, nil
}

// WatchTransactionBatchExecutedEvent is a free log subscription operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Contract *ContractFilterer) WatchTransactionBatchExecutedEvent(opts *bind.WatchOpts, sink chan<- *ContractTransactionBatchExecutedEvent, _batchNonce []*big.Int, _token []common.Address) (event.Subscription, error) {

	var _batchNonceRule []interface{}
	for _, _batchNonceItem := range _batchNonce {
		_batchNonceRule = append(_batchNonceRule, _batchNonceItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TransactionBatchExecutedEvent", _batchNonceRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransactionBatchExecutedEvent)
				if err := _Contract.contract.UnpackLog(event, "TransactionBatchExecutedEvent", log); err != nil {
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

// ParseTransactionBatchExecutedEvent is a log parse operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Contract *ContractFilterer) ParseTransactionBatchExecutedEvent(log types.Log) (*ContractTransactionBatchExecutedEvent, error) {
	event := new(ContractTransactionBatchExecutedEvent)
	if err := _Contract.contract.UnpackLog(event, "TransactionBatchExecutedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValsetUpdatedEventIterator is returned from FilterValsetUpdatedEvent and is used to iterate over the raw logs and unpacked data for ValsetUpdatedEvent events raised by the Contract contract.
type ContractValsetUpdatedEventIterator struct {
	Event *ContractValsetUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *ContractValsetUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValsetUpdatedEvent)
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
		it.Event = new(ContractValsetUpdatedEvent)
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
func (it *ContractValsetUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValsetUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValsetUpdatedEvent represents a ValsetUpdatedEvent event raised by the Contract contract.
type ContractValsetUpdatedEvent struct {
	NewValsetNonce *big.Int
	EventNonce     *big.Int
	RewardAmount   *big.Int
	RewardToken    common.Address
	Validators     []common.Address
	Powers         []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValsetUpdatedEvent is a free log retrieval operation binding the contract event 0x76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, uint256 _eventNonce, uint256 _rewardAmount, address _rewardToken, address[] _validators, uint256[] _powers)
func (_Contract *ContractFilterer) FilterValsetUpdatedEvent(opts *bind.FilterOpts, _newValsetNonce []*big.Int) (*ContractValsetUpdatedEventIterator, error) {

	var _newValsetNonceRule []interface{}
	for _, _newValsetNonceItem := range _newValsetNonce {
		_newValsetNonceRule = append(_newValsetNonceRule, _newValsetNonceItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValsetUpdatedEvent", _newValsetNonceRule)
	if err != nil {
		return nil, err
	}
	return &ContractValsetUpdatedEventIterator{contract: _Contract.contract, event: "ValsetUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchValsetUpdatedEvent is a free log subscription operation binding the contract event 0x76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, uint256 _eventNonce, uint256 _rewardAmount, address _rewardToken, address[] _validators, uint256[] _powers)
func (_Contract *ContractFilterer) WatchValsetUpdatedEvent(opts *bind.WatchOpts, sink chan<- *ContractValsetUpdatedEvent, _newValsetNonce []*big.Int) (event.Subscription, error) {

	var _newValsetNonceRule []interface{}
	for _, _newValsetNonceItem := range _newValsetNonce {
		_newValsetNonceRule = append(_newValsetNonceRule, _newValsetNonceItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValsetUpdatedEvent", _newValsetNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValsetUpdatedEvent)
				if err := _Contract.contract.UnpackLog(event, "ValsetUpdatedEvent", log); err != nil {
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

// ParseValsetUpdatedEvent is a log parse operation binding the contract event 0x76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, uint256 _eventNonce, uint256 _rewardAmount, address _rewardToken, address[] _validators, uint256[] _powers)
func (_Contract *ContractFilterer) ParseValsetUpdatedEvent(log types.Log) (*ContractValsetUpdatedEvent, error) {
	event := new(ContractValsetUpdatedEvent)
	if err := _Contract.contract.UnpackLog(event, "ValsetUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
