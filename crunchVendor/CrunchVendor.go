// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crunchVendor

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

// CrunchVendorMetaData contains all meta data concerning the CrunchVendor contract.
var CrunchVendorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spaceAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"invitee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"Commission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"invitee\",\"type\":\"address\"}],\"name\":\"Invite\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TopUpSuccess\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"boxOffice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contentID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"inviter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inviter_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161134d38038061134d83398181016040528101906100319190610253565b825f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100a2575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161009991906102b2565b60405180910390fd5b6100b18161010160201b60201c565b508160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806002819055505050506102cb565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101ef826101c6565b9050919050565b6101ff816101e5565b8114610209575f80fd5b50565b5f8151905061021a816101f6565b92915050565b5f819050919050565b61023281610220565b811461023c575f80fd5b50565b5f8151905061024d81610229565b92915050565b5f805f6060848603121561026a576102696101c2565b5b5f6102778682870161020c565b93505060206102888682870161020c565b92505060406102998682870161023f565b9150509250925092565b6102ac816101e5565b82525050565b5f6020820190506102c55f8301846102a3565b92915050565b611075806102d85f395ff3fe60806040526004361061009b575f3560e01c80638da5cb5b116100635780638da5cb5b1461016f57806391b7f5ed1461019957806398d5fdca146101c1578063d6f7ddf9146101eb578063ee8f0b7a14610207578063f2fde38b146102435761009b565b80630103c92b1461009f57806302d05d3f146100db5780630641def214610105578063312150f81461012f578063715018a614610159575b5f80fd5b3480156100aa575f80fd5b506100c560048036038101906100c09190610b80565b61026b565b6040516100d29190610bc3565b60405180910390f35b3480156100e6575f80fd5b506100ef6102b1565b6040516100fc9190610beb565b60405180910390f35b348015610110575f80fd5b506101196102bf565b6040516101269190610bc3565b60405180910390f35b34801561013a575f80fd5b506101436102c8565b6040516101509190610bc3565b60405180910390f35b348015610164575f80fd5b5061016d6102d1565b005b34801561017a575f80fd5b506101836102e4565b6040516101909190610beb565b60405180910390f35b3480156101a4575f80fd5b506101bf60048036038101906101ba9190610c2e565b61030b565b005b3480156101cc575f80fd5b506101d561031d565b6040516101e29190610bc3565b60405180910390f35b61020560048036038101906102009190610c59565b610326565b005b348015610212575f80fd5b5061022d60048036038101906102289190610b80565b6108dd565b60405161023a9190610beb565b60405180910390f35b34801561024e575f80fd5b5061026960048036038101906102649190610b80565b610942565b005b5f60065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f6102ba6102e4565b905090565b5f600554905090565b5f600254905090565b6102d96109c6565b6102e25f610a4d565b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6103136109c6565b8060038190555050565b5f600354905090565b806003546103349190610cc4565b3414610375576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036c90610d5f565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141580156103dd57503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b801561047057505f73ffffffffffffffffffffffffffffffffffffffff1660045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b1561054b578160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fa7f319d32ea1111130a0941caa632eb49779c9ca37820ae07e85f9f0ebbb51c860405160405180910390a35b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633e4eb36c6040518163ffffffff1660e01b81526004015f60405180830381865afa1580156105b5573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906105dd9190610ee1565b90505f3490505f5b8251811015610792575f8111156106565760045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1694505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff160315610792575f60648483815181106106a05761069f610f28565b5b6020026020010151866106b39190610cc4565b6106bd9190610f82565b90508573ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610702573d5f803e3d5ffd5b503373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f05456de91d83e21ad7c41a09ae7cb41836049c49e6ddaf07bdfc40c2231885d2836001866107609190610fb2565b60405161076e929190610fe5565b60405180910390a38083610782919061100c565b92505080806001019150506105e5565b508260065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546107df9190610fb2565b925050819055508260055f8282546107f79190610fb2565b9250508190555060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ac56b64b826002546040518363ffffffff1660e01b815260040161085b9190610bc3565b5f604051808303818588803b158015610872575f80fd5b505af1158015610884573d5f803e3d5ffd5b50505050503373ffffffffffffffffffffffffffffffffffffffff167f41fb22dbb233b52cfc8961f58617ef5917854e057128d2834e76b11bc6f31eab846040516108cf9190610bc3565b60405180910390a250505050565b5f60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b61094a6109c6565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109ba575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016109b19190610beb565b60405180910390fd5b6109c381610a4d565b50565b6109ce610b0e565b73ffffffffffffffffffffffffffffffffffffffff166109ec6102e4565b73ffffffffffffffffffffffffffffffffffffffff1614610a4b57610a0f610b0e565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610a429190610beb565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b4f82610b26565b9050919050565b610b5f81610b45565b8114610b69575f80fd5b50565b5f81359050610b7a81610b56565b92915050565b5f60208284031215610b9557610b94610b1e565b5b5f610ba284828501610b6c565b91505092915050565b5f819050919050565b610bbd81610bab565b82525050565b5f602082019050610bd65f830184610bb4565b92915050565b610be581610b45565b82525050565b5f602082019050610bfe5f830184610bdc565b92915050565b610c0d81610bab565b8114610c17575f80fd5b50565b5f81359050610c2881610c04565b92915050565b5f60208284031215610c4357610c42610b1e565b5b5f610c5084828501610c1a565b91505092915050565b5f8060408385031215610c6f57610c6e610b1e565b5b5f610c7c85828601610b6c565b9250506020610c8d85828601610c1a565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610cce82610bab565b9150610cd983610bab565b9250828202610ce781610bab565b91508282048414831517610cfe57610cfd610c97565b5b5092915050565b5f82825260208201905092915050565b7f7072696365206e6f74206d6174636800000000000000000000000000000000005f82015250565b5f610d49600f83610d05565b9150610d5482610d15565b602082019050919050565b5f6020820190508181035f830152610d7681610d3d565b9050919050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610dc782610d81565b810181811067ffffffffffffffff82111715610de657610de5610d91565b5b80604052505050565b5f610df8610b15565b9050610e048282610dbe565b919050565b5f67ffffffffffffffff821115610e2357610e22610d91565b5b602082029050602081019050919050565b5f80fd5b5f81519050610e4681610c04565b92915050565b5f610e5e610e5984610e09565b610def565b90508083825260208201905060208402830185811115610e8157610e80610e34565b5b835b81811015610eaa5780610e968882610e38565b845260208401935050602081019050610e83565b5050509392505050565b5f82601f830112610ec857610ec7610d7d565b5b8151610ed8848260208601610e4c565b91505092915050565b5f60208284031215610ef657610ef5610b1e565b5b5f82015167ffffffffffffffff811115610f1357610f12610b22565b5b610f1f84828501610eb4565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610f8c82610bab565b9150610f9783610bab565b925082610fa757610fa6610f55565b5b828204905092915050565b5f610fbc82610bab565b9150610fc783610bab565b9250828201905080821115610fdf57610fde610c97565b5b92915050565b5f604082019050610ff85f830185610bb4565b6110056020830184610bb4565b9392505050565b5f61101682610bab565b915061102183610bab565b925082820390508181111561103957611038610c97565b5b9291505056fea264697066735822122001d4930cb8c96477c8f2abf2b6f8ab279c86690f7a99bd8ff486565f6d75458864736f6c634300081a0033",
}

// CrunchVendorABI is the input ABI used to generate the binding from.
// Deprecated: Use CrunchVendorMetaData.ABI instead.
var CrunchVendorABI = CrunchVendorMetaData.ABI

// CrunchVendorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrunchVendorMetaData.Bin instead.
var CrunchVendorBin = CrunchVendorMetaData.Bin

// DeployCrunchVendor deploys a new Ethereum contract, binding an instance of CrunchVendor to it.
func DeployCrunchVendor(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address, spaceAddress common.Address, tokenID *big.Int) (common.Address, *types.Transaction, *CrunchVendor, error) {
	parsed, err := CrunchVendorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrunchVendorBin), backend, initialOwner, spaceAddress, tokenID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrunchVendor{CrunchVendorCaller: CrunchVendorCaller{contract: contract}, CrunchVendorTransactor: CrunchVendorTransactor{contract: contract}, CrunchVendorFilterer: CrunchVendorFilterer{contract: contract}}, nil
}

// CrunchVendor is an auto generated Go binding around an Ethereum contract.
type CrunchVendor struct {
	CrunchVendorCaller     // Read-only binding to the contract
	CrunchVendorTransactor // Write-only binding to the contract
	CrunchVendorFilterer   // Log filterer for contract events
}

// CrunchVendorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrunchVendorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrunchVendorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrunchVendorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrunchVendorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrunchVendorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrunchVendorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrunchVendorSession struct {
	Contract     *CrunchVendor     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrunchVendorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrunchVendorCallerSession struct {
	Contract *CrunchVendorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CrunchVendorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrunchVendorTransactorSession struct {
	Contract     *CrunchVendorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CrunchVendorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrunchVendorRaw struct {
	Contract *CrunchVendor // Generic contract binding to access the raw methods on
}

// CrunchVendorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrunchVendorCallerRaw struct {
	Contract *CrunchVendorCaller // Generic read-only contract binding to access the raw methods on
}

// CrunchVendorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrunchVendorTransactorRaw struct {
	Contract *CrunchVendorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrunchVendor creates a new instance of CrunchVendor, bound to a specific deployed contract.
func NewCrunchVendor(address common.Address, backend bind.ContractBackend) (*CrunchVendor, error) {
	contract, err := bindCrunchVendor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrunchVendor{CrunchVendorCaller: CrunchVendorCaller{contract: contract}, CrunchVendorTransactor: CrunchVendorTransactor{contract: contract}, CrunchVendorFilterer: CrunchVendorFilterer{contract: contract}}, nil
}

// NewCrunchVendorCaller creates a new read-only instance of CrunchVendor, bound to a specific deployed contract.
func NewCrunchVendorCaller(address common.Address, caller bind.ContractCaller) (*CrunchVendorCaller, error) {
	contract, err := bindCrunchVendor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrunchVendorCaller{contract: contract}, nil
}

// NewCrunchVendorTransactor creates a new write-only instance of CrunchVendor, bound to a specific deployed contract.
func NewCrunchVendorTransactor(address common.Address, transactor bind.ContractTransactor) (*CrunchVendorTransactor, error) {
	contract, err := bindCrunchVendor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrunchVendorTransactor{contract: contract}, nil
}

// NewCrunchVendorFilterer creates a new log filterer instance of CrunchVendor, bound to a specific deployed contract.
func NewCrunchVendorFilterer(address common.Address, filterer bind.ContractFilterer) (*CrunchVendorFilterer, error) {
	contract, err := bindCrunchVendor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrunchVendorFilterer{contract: contract}, nil
}

// bindCrunchVendor binds a generic wrapper to an already deployed contract.
func bindCrunchVendor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrunchVendorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrunchVendor *CrunchVendorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrunchVendor.Contract.CrunchVendorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrunchVendor *CrunchVendorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrunchVendor.Contract.CrunchVendorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrunchVendor *CrunchVendorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrunchVendor.Contract.CrunchVendorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrunchVendor *CrunchVendorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrunchVendor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrunchVendor *CrunchVendorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrunchVendor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrunchVendor *CrunchVendorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrunchVendor.Contract.contract.Transact(opts, method, params...)
}

// BoxOffice is a free data retrieval call binding the contract method 0x0641def2.
//
// Solidity: function boxOffice() view returns(uint256)
func (_CrunchVendor *CrunchVendorCaller) BoxOffice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrunchVendor.contract.Call(opts, &out, "boxOffice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoxOffice is a free data retrieval call binding the contract method 0x0641def2.
//
// Solidity: function boxOffice() view returns(uint256)
func (_CrunchVendor *CrunchVendorSession) BoxOffice() (*big.Int, error) {
	return _CrunchVendor.Contract.BoxOffice(&_CrunchVendor.CallOpts)
}

// BoxOffice is a free data retrieval call binding the contract method 0x0641def2.
//
// Solidity: function boxOffice() view returns(uint256)
func (_CrunchVendor *CrunchVendorCallerSession) BoxOffice() (*big.Int, error) {
	return _CrunchVendor.Contract.BoxOffice(&_CrunchVendor.CallOpts)
}

// ContentID is a free data retrieval call binding the contract method 0x312150f8.
//
// Solidity: function contentID() view returns(uint256)
func (_CrunchVendor *CrunchVendorCaller) ContentID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrunchVendor.contract.Call(opts, &out, "contentID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContentID is a free data retrieval call binding the contract method 0x312150f8.
//
// Solidity: function contentID() view returns(uint256)
func (_CrunchVendor *CrunchVendorSession) ContentID() (*big.Int, error) {
	return _CrunchVendor.Contract.ContentID(&_CrunchVendor.CallOpts)
}

// ContentID is a free data retrieval call binding the contract method 0x312150f8.
//
// Solidity: function contentID() view returns(uint256)
func (_CrunchVendor *CrunchVendorCallerSession) ContentID() (*big.Int, error) {
	return _CrunchVendor.Contract.ContentID(&_CrunchVendor.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_CrunchVendor *CrunchVendorCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrunchVendor.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_CrunchVendor *CrunchVendorSession) Creator() (common.Address, error) {
	return _CrunchVendor.Contract.Creator(&_CrunchVendor.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_CrunchVendor *CrunchVendorCallerSession) Creator() (common.Address, error) {
	return _CrunchVendor.Contract.Creator(&_CrunchVendor.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_CrunchVendor *CrunchVendorCaller) GetPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrunchVendor.contract.Call(opts, &out, "getPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_CrunchVendor *CrunchVendorSession) GetPrice() (*big.Int, error) {
	return _CrunchVendor.Contract.GetPrice(&_CrunchVendor.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_CrunchVendor *CrunchVendorCallerSession) GetPrice() (*big.Int, error) {
	return _CrunchVendor.Contract.GetPrice(&_CrunchVendor.CallOpts)
}

// Inviter is a free data retrieval call binding the contract method 0xee8f0b7a.
//
// Solidity: function inviter(address user) view returns(address)
func (_CrunchVendor *CrunchVendorCaller) Inviter(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _CrunchVendor.contract.Call(opts, &out, "inviter", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inviter is a free data retrieval call binding the contract method 0xee8f0b7a.
//
// Solidity: function inviter(address user) view returns(address)
func (_CrunchVendor *CrunchVendorSession) Inviter(user common.Address) (common.Address, error) {
	return _CrunchVendor.Contract.Inviter(&_CrunchVendor.CallOpts, user)
}

// Inviter is a free data retrieval call binding the contract method 0xee8f0b7a.
//
// Solidity: function inviter(address user) view returns(address)
func (_CrunchVendor *CrunchVendorCallerSession) Inviter(user common.Address) (common.Address, error) {
	return _CrunchVendor.Contract.Inviter(&_CrunchVendor.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrunchVendor *CrunchVendorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrunchVendor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrunchVendor *CrunchVendorSession) Owner() (common.Address, error) {
	return _CrunchVendor.Contract.Owner(&_CrunchVendor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrunchVendor *CrunchVendorCallerSession) Owner() (common.Address, error) {
	return _CrunchVendor.Contract.Owner(&_CrunchVendor.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user_) view returns(uint256)
func (_CrunchVendor *CrunchVendorCaller) UserBalance(opts *bind.CallOpts, user_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrunchVendor.contract.Call(opts, &out, "userBalance", user_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user_) view returns(uint256)
func (_CrunchVendor *CrunchVendorSession) UserBalance(user_ common.Address) (*big.Int, error) {
	return _CrunchVendor.Contract.UserBalance(&_CrunchVendor.CallOpts, user_)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user_) view returns(uint256)
func (_CrunchVendor *CrunchVendorCallerSession) UserBalance(user_ common.Address) (*big.Int, error) {
	return _CrunchVendor.Contract.UserBalance(&_CrunchVendor.CallOpts, user_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrunchVendor *CrunchVendorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrunchVendor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrunchVendor *CrunchVendorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrunchVendor.Contract.RenounceOwnership(&_CrunchVendor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrunchVendor *CrunchVendorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrunchVendor.Contract.RenounceOwnership(&_CrunchVendor.TransactOpts)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 price_) returns()
func (_CrunchVendor *CrunchVendorTransactor) SetPrice(opts *bind.TransactOpts, price_ *big.Int) (*types.Transaction, error) {
	return _CrunchVendor.contract.Transact(opts, "setPrice", price_)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 price_) returns()
func (_CrunchVendor *CrunchVendorSession) SetPrice(price_ *big.Int) (*types.Transaction, error) {
	return _CrunchVendor.Contract.SetPrice(&_CrunchVendor.TransactOpts, price_)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 price_) returns()
func (_CrunchVendor *CrunchVendorTransactorSession) SetPrice(price_ *big.Int) (*types.Transaction, error) {
	return _CrunchVendor.Contract.SetPrice(&_CrunchVendor.TransactOpts, price_)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(address inviter_, uint256 amount) payable returns()
func (_CrunchVendor *CrunchVendorTransactor) TopUp(opts *bind.TransactOpts, inviter_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrunchVendor.contract.Transact(opts, "topUp", inviter_, amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(address inviter_, uint256 amount) payable returns()
func (_CrunchVendor *CrunchVendorSession) TopUp(inviter_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrunchVendor.Contract.TopUp(&_CrunchVendor.TransactOpts, inviter_, amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(address inviter_, uint256 amount) payable returns()
func (_CrunchVendor *CrunchVendorTransactorSession) TopUp(inviter_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrunchVendor.Contract.TopUp(&_CrunchVendor.TransactOpts, inviter_, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrunchVendor *CrunchVendorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CrunchVendor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrunchVendor *CrunchVendorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrunchVendor.Contract.TransferOwnership(&_CrunchVendor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrunchVendor *CrunchVendorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrunchVendor.Contract.TransferOwnership(&_CrunchVendor.TransactOpts, newOwner)
}

// CrunchVendorCommissionIterator is returned from FilterCommission and is used to iterate over the raw logs and unpacked data for Commission events raised by the CrunchVendor contract.
type CrunchVendorCommissionIterator struct {
	Event *CrunchVendorCommission // Event containing the contract specifics and raw log

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
func (it *CrunchVendorCommissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrunchVendorCommission)
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
		it.Event = new(CrunchVendorCommission)
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
func (it *CrunchVendorCommissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrunchVendorCommissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrunchVendorCommission represents a Commission event raised by the CrunchVendor contract.
type CrunchVendorCommission struct {
	Inviter common.Address
	Invitee common.Address
	Amount  *big.Int
	Level   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCommission is a free log retrieval operation binding the contract event 0x05456de91d83e21ad7c41a09ae7cb41836049c49e6ddaf07bdfc40c2231885d2.
//
// Solidity: event Commission(address indexed inviter, address indexed invitee, uint256 amount, uint256 level)
func (_CrunchVendor *CrunchVendorFilterer) FilterCommission(opts *bind.FilterOpts, inviter []common.Address, invitee []common.Address) (*CrunchVendorCommissionIterator, error) {

	var inviterRule []interface{}
	for _, inviterItem := range inviter {
		inviterRule = append(inviterRule, inviterItem)
	}
	var inviteeRule []interface{}
	for _, inviteeItem := range invitee {
		inviteeRule = append(inviteeRule, inviteeItem)
	}

	logs, sub, err := _CrunchVendor.contract.FilterLogs(opts, "Commission", inviterRule, inviteeRule)
	if err != nil {
		return nil, err
	}
	return &CrunchVendorCommissionIterator{contract: _CrunchVendor.contract, event: "Commission", logs: logs, sub: sub}, nil
}

// WatchCommission is a free log subscription operation binding the contract event 0x05456de91d83e21ad7c41a09ae7cb41836049c49e6ddaf07bdfc40c2231885d2.
//
// Solidity: event Commission(address indexed inviter, address indexed invitee, uint256 amount, uint256 level)
func (_CrunchVendor *CrunchVendorFilterer) WatchCommission(opts *bind.WatchOpts, sink chan<- *CrunchVendorCommission, inviter []common.Address, invitee []common.Address) (event.Subscription, error) {

	var inviterRule []interface{}
	for _, inviterItem := range inviter {
		inviterRule = append(inviterRule, inviterItem)
	}
	var inviteeRule []interface{}
	for _, inviteeItem := range invitee {
		inviteeRule = append(inviteeRule, inviteeItem)
	}

	logs, sub, err := _CrunchVendor.contract.WatchLogs(opts, "Commission", inviterRule, inviteeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrunchVendorCommission)
				if err := _CrunchVendor.contract.UnpackLog(event, "Commission", log); err != nil {
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

// ParseCommission is a log parse operation binding the contract event 0x05456de91d83e21ad7c41a09ae7cb41836049c49e6ddaf07bdfc40c2231885d2.
//
// Solidity: event Commission(address indexed inviter, address indexed invitee, uint256 amount, uint256 level)
func (_CrunchVendor *CrunchVendorFilterer) ParseCommission(log types.Log) (*CrunchVendorCommission, error) {
	event := new(CrunchVendorCommission)
	if err := _CrunchVendor.contract.UnpackLog(event, "Commission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrunchVendorInviteIterator is returned from FilterInvite and is used to iterate over the raw logs and unpacked data for Invite events raised by the CrunchVendor contract.
type CrunchVendorInviteIterator struct {
	Event *CrunchVendorInvite // Event containing the contract specifics and raw log

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
func (it *CrunchVendorInviteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrunchVendorInvite)
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
		it.Event = new(CrunchVendorInvite)
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
func (it *CrunchVendorInviteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrunchVendorInviteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrunchVendorInvite represents a Invite event raised by the CrunchVendor contract.
type CrunchVendorInvite struct {
	Inviter common.Address
	Invitee common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInvite is a free log retrieval operation binding the contract event 0xa7f319d32ea1111130a0941caa632eb49779c9ca37820ae07e85f9f0ebbb51c8.
//
// Solidity: event Invite(address indexed inviter, address indexed invitee)
func (_CrunchVendor *CrunchVendorFilterer) FilterInvite(opts *bind.FilterOpts, inviter []common.Address, invitee []common.Address) (*CrunchVendorInviteIterator, error) {

	var inviterRule []interface{}
	for _, inviterItem := range inviter {
		inviterRule = append(inviterRule, inviterItem)
	}
	var inviteeRule []interface{}
	for _, inviteeItem := range invitee {
		inviteeRule = append(inviteeRule, inviteeItem)
	}

	logs, sub, err := _CrunchVendor.contract.FilterLogs(opts, "Invite", inviterRule, inviteeRule)
	if err != nil {
		return nil, err
	}
	return &CrunchVendorInviteIterator{contract: _CrunchVendor.contract, event: "Invite", logs: logs, sub: sub}, nil
}

// WatchInvite is a free log subscription operation binding the contract event 0xa7f319d32ea1111130a0941caa632eb49779c9ca37820ae07e85f9f0ebbb51c8.
//
// Solidity: event Invite(address indexed inviter, address indexed invitee)
func (_CrunchVendor *CrunchVendorFilterer) WatchInvite(opts *bind.WatchOpts, sink chan<- *CrunchVendorInvite, inviter []common.Address, invitee []common.Address) (event.Subscription, error) {

	var inviterRule []interface{}
	for _, inviterItem := range inviter {
		inviterRule = append(inviterRule, inviterItem)
	}
	var inviteeRule []interface{}
	for _, inviteeItem := range invitee {
		inviteeRule = append(inviteeRule, inviteeItem)
	}

	logs, sub, err := _CrunchVendor.contract.WatchLogs(opts, "Invite", inviterRule, inviteeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrunchVendorInvite)
				if err := _CrunchVendor.contract.UnpackLog(event, "Invite", log); err != nil {
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

// ParseInvite is a log parse operation binding the contract event 0xa7f319d32ea1111130a0941caa632eb49779c9ca37820ae07e85f9f0ebbb51c8.
//
// Solidity: event Invite(address indexed inviter, address indexed invitee)
func (_CrunchVendor *CrunchVendorFilterer) ParseInvite(log types.Log) (*CrunchVendorInvite, error) {
	event := new(CrunchVendorInvite)
	if err := _CrunchVendor.contract.UnpackLog(event, "Invite", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrunchVendorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CrunchVendor contract.
type CrunchVendorOwnershipTransferredIterator struct {
	Event *CrunchVendorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrunchVendorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrunchVendorOwnershipTransferred)
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
		it.Event = new(CrunchVendorOwnershipTransferred)
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
func (it *CrunchVendorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrunchVendorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrunchVendorOwnershipTransferred represents a OwnershipTransferred event raised by the CrunchVendor contract.
type CrunchVendorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrunchVendor *CrunchVendorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrunchVendorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrunchVendor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrunchVendorOwnershipTransferredIterator{contract: _CrunchVendor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrunchVendor *CrunchVendorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrunchVendorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrunchVendor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrunchVendorOwnershipTransferred)
				if err := _CrunchVendor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CrunchVendor *CrunchVendorFilterer) ParseOwnershipTransferred(log types.Log) (*CrunchVendorOwnershipTransferred, error) {
	event := new(CrunchVendorOwnershipTransferred)
	if err := _CrunchVendor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrunchVendorTopUpSuccessIterator is returned from FilterTopUpSuccess and is used to iterate over the raw logs and unpacked data for TopUpSuccess events raised by the CrunchVendor contract.
type CrunchVendorTopUpSuccessIterator struct {
	Event *CrunchVendorTopUpSuccess // Event containing the contract specifics and raw log

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
func (it *CrunchVendorTopUpSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrunchVendorTopUpSuccess)
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
		it.Event = new(CrunchVendorTopUpSuccess)
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
func (it *CrunchVendorTopUpSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrunchVendorTopUpSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrunchVendorTopUpSuccess represents a TopUpSuccess event raised by the CrunchVendor contract.
type CrunchVendorTopUpSuccess struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTopUpSuccess is a free log retrieval operation binding the contract event 0x41fb22dbb233b52cfc8961f58617ef5917854e057128d2834e76b11bc6f31eab.
//
// Solidity: event TopUpSuccess(address indexed user, uint256 amount)
func (_CrunchVendor *CrunchVendorFilterer) FilterTopUpSuccess(opts *bind.FilterOpts, user []common.Address) (*CrunchVendorTopUpSuccessIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrunchVendor.contract.FilterLogs(opts, "TopUpSuccess", userRule)
	if err != nil {
		return nil, err
	}
	return &CrunchVendorTopUpSuccessIterator{contract: _CrunchVendor.contract, event: "TopUpSuccess", logs: logs, sub: sub}, nil
}

// WatchTopUpSuccess is a free log subscription operation binding the contract event 0x41fb22dbb233b52cfc8961f58617ef5917854e057128d2834e76b11bc6f31eab.
//
// Solidity: event TopUpSuccess(address indexed user, uint256 amount)
func (_CrunchVendor *CrunchVendorFilterer) WatchTopUpSuccess(opts *bind.WatchOpts, sink chan<- *CrunchVendorTopUpSuccess, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrunchVendor.contract.WatchLogs(opts, "TopUpSuccess", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrunchVendorTopUpSuccess)
				if err := _CrunchVendor.contract.UnpackLog(event, "TopUpSuccess", log); err != nil {
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

// ParseTopUpSuccess is a log parse operation binding the contract event 0x41fb22dbb233b52cfc8961f58617ef5917854e057128d2834e76b11bc6f31eab.
//
// Solidity: event TopUpSuccess(address indexed user, uint256 amount)
func (_CrunchVendor *CrunchVendorFilterer) ParseTopUpSuccess(log types.Log) (*CrunchVendorTopUpSuccess, error) {
	event := new(CrunchVendorTopUpSuccess)
	if err := _CrunchVendor.contract.UnpackLog(event, "TopUpSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
