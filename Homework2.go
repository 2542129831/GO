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

// PersonDIDPerson is an auto generated low-level Go binding around an user-defined struct.
type PersonDIDPerson struct {
	Id    uint8
	Age   uint8
	Name  string
	Story string
}

// ApinameABI is the input ABI used to generate the binding from.
const ApinameABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AddPerson\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"story\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AddStory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ip\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"giveAccess\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PersonInfo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"story\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"PersonName\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"story\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"Right\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"story\",\"type\":\"string\"}],\"name\":\"addPerson\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfPersons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"}],\"name\":\"getPersonAge\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"Yourid\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"ip\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"time\",\"type\":\"uint8\"}],\"name\":\"giveRight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPersonExsist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"seeId\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"Yourid\",\"type\":\"uint8\"}],\"name\":\"myself\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"story\",\"type\":\"string\"}],\"internalType\":\"structPersonDID.Person\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"}],\"name\":\"setAdminer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"}],\"name\":\"setPersonAgeMem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"}],\"name\":\"setPersonAgeSto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Apiname is an auto generated Go binding around an Ethereum contract.
type Apiname struct {
	ApinameCaller     // Read-only binding to the contract
	ApinameTransactor // Write-only binding to the contract
	ApinameFilterer   // Log filterer for contract events
}

// ApinameCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApinameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApinameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApinameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApinameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApinameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApinameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApinameSession struct {
	Contract     *Apiname          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApinameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApinameCallerSession struct {
	Contract *ApinameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ApinameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApinameTransactorSession struct {
	Contract     *ApinameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ApinameRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApinameRaw struct {
	Contract *Apiname // Generic contract binding to access the raw methods on
}

// ApinameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApinameCallerRaw struct {
	Contract *ApinameCaller // Generic read-only contract binding to access the raw methods on
}

// ApinameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApinameTransactorRaw struct {
	Contract *ApinameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApiname creates a new instance of Apiname, bound to a specific deployed contract.
func NewApiname(address common.Address, backend bind.ContractBackend) (*Apiname, error) {
	contract, err := bindApiname(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Apiname{ApinameCaller: ApinameCaller{contract: contract}, ApinameTransactor: ApinameTransactor{contract: contract}, ApinameFilterer: ApinameFilterer{contract: contract}}, nil
}

// NewApinameCaller creates a new read-only instance of Apiname, bound to a specific deployed contract.
func NewApinameCaller(address common.Address, caller bind.ContractCaller) (*ApinameCaller, error) {
	contract, err := bindApiname(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApinameCaller{contract: contract}, nil
}

// NewApinameTransactor creates a new write-only instance of Apiname, bound to a specific deployed contract.
func NewApinameTransactor(address common.Address, transactor bind.ContractTransactor) (*ApinameTransactor, error) {
	contract, err := bindApiname(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApinameTransactor{contract: contract}, nil
}

// NewApinameFilterer creates a new log filterer instance of Apiname, bound to a specific deployed contract.
func NewApinameFilterer(address common.Address, filterer bind.ContractFilterer) (*ApinameFilterer, error) {
	contract, err := bindApiname(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApinameFilterer{contract: contract}, nil
}

// bindApiname binds a generic wrapper to an already deployed contract.
func bindApiname(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApinameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apiname *ApinameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Apiname.Contract.ApinameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apiname *ApinameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apiname.Contract.ApinameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apiname *ApinameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apiname.Contract.ApinameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apiname *ApinameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Apiname.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apiname *ApinameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apiname.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apiname *ApinameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apiname.Contract.contract.Transact(opts, method, params...)
}

// PersonInfo is a free data retrieval call binding the contract method 0x44c7c373.
//
// Solidity: function PersonInfo(address ) view returns(uint8 id, uint8 age, string name, string story)
func (_Apiname *ApinameCaller) PersonInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Id    uint8
	Age   uint8
	Name  string
	Story string
}, error) {
	var out []interface{}
	err := _Apiname.contract.Call(opts, &out, "PersonInfo", arg0)

	outstruct := new(struct {
		Id    uint8
		Age   uint8
		Name  string
		Story string
	})

	outstruct.Id = out[0].(uint8)
	outstruct.Age = out[1].(uint8)
	outstruct.Name = out[2].(string)
	outstruct.Story = out[3].(string)

	return *outstruct, err

}

// PersonInfo is a free data retrieval call binding the contract method 0x44c7c373.
//
// Solidity: function PersonInfo(address ) view returns(uint8 id, uint8 age, string name, string story)
func (_Apiname *ApinameSession) PersonInfo(arg0 common.Address) (struct {
	Id    uint8
	Age   uint8
	Name  string
	Story string
}, error) {
	return _Apiname.Contract.PersonInfo(&_Apiname.CallOpts, arg0)
}

// PersonInfo is a free data retrieval call binding the contract method 0x44c7c373.
//
// Solidity: function PersonInfo(address ) view returns(uint8 id, uint8 age, string name, string story)
func (_Apiname *ApinameCallerSession) PersonInfo(arg0 common.Address) (struct {
	Id    uint8
	Age   uint8
	Name  string
	Story string
}, error) {
	return _Apiname.Contract.PersonInfo(&_Apiname.CallOpts, arg0)
}

// PersonName is a free data retrieval call binding the contract method 0xa1758722.
//
// Solidity: function PersonName(uint8 ) view returns(uint8 id, uint8 age, string name, string story)
func (_Apiname *ApinameCaller) PersonName(opts *bind.CallOpts, arg0 uint8) (struct {
	Id    uint8
	Age   uint8
	Name  string
	Story string
}, error) {
	var out []interface{}
	err := _Apiname.contract.Call(opts, &out, "PersonName", arg0)

	outstruct := new(struct {
		Id    uint8
		Age   uint8
		Name  string
		Story string
	})

	outstruct.Id = out[0].(uint8)
	outstruct.Age = out[1].(uint8)
	outstruct.Name = out[2].(string)
	outstruct.Story = out[3].(string)

	return *outstruct, err

}

// PersonName is a free data retrieval call binding the contract method 0xa1758722.
//
// Solidity: function PersonName(uint8 ) view returns(uint8 id, uint8 age, string name, string story)
func (_Apiname *ApinameSession) PersonName(arg0 uint8) (struct {
	Id    uint8
	Age   uint8
	Name  string
	Story string
}, error) {
	return _Apiname.Contract.PersonName(&_Apiname.CallOpts, arg0)
}

// PersonName is a free data retrieval call binding the contract method 0xa1758722.
//
// Solidity: function PersonName(uint8 ) view returns(uint8 id, uint8 age, string name, string story)
func (_Apiname *ApinameCallerSession) PersonName(arg0 uint8) (struct {
	Id    uint8
	Age   uint8
	Name  string
	Story string
}, error) {
	return _Apiname.Contract.PersonName(&_Apiname.CallOpts, arg0)
}

// Right is a free data retrieval call binding the contract method 0x8f35152b.
//
// Solidity: function Right(uint8 ) view returns(bool)
func (_Apiname *ApinameCaller) Right(opts *bind.CallOpts, arg0 uint8) (bool, error) {
	var out []interface{}
	err := _Apiname.contract.Call(opts, &out, "Right", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Right is a free data retrieval call binding the contract method 0x8f35152b.
//
// Solidity: function Right(uint8 ) view returns(bool)
func (_Apiname *ApinameSession) Right(arg0 uint8) (bool, error) {
	return _Apiname.Contract.Right(&_Apiname.CallOpts, arg0)
}

// Right is a free data retrieval call binding the contract method 0x8f35152b.
//
// Solidity: function Right(uint8 ) view returns(bool)
func (_Apiname *ApinameCallerSession) Right(arg0 uint8) (bool, error) {
	return _Apiname.Contract.Right(&_Apiname.CallOpts, arg0)
}

// GetNumberOfPersons is a free data retrieval call binding the contract method 0xcaac5da7.
//
// Solidity: function getNumberOfPersons() view returns(uint256)
func (_Apiname *ApinameCaller) GetNumberOfPersons(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Apiname.contract.Call(opts, &out, "getNumberOfPersons")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfPersons is a free data retrieval call binding the contract method 0xcaac5da7.
//
// Solidity: function getNumberOfPersons() view returns(uint256)
func (_Apiname *ApinameSession) GetNumberOfPersons() (*big.Int, error) {
	return _Apiname.Contract.GetNumberOfPersons(&_Apiname.CallOpts)
}

// GetNumberOfPersons is a free data retrieval call binding the contract method 0xcaac5da7.
//
// Solidity: function getNumberOfPersons() view returns(uint256)
func (_Apiname *ApinameCallerSession) GetNumberOfPersons() (*big.Int, error) {
	return _Apiname.Contract.GetNumberOfPersons(&_Apiname.CallOpts)
}

// GetPersonAge is a free data retrieval call binding the contract method 0x0cfce82e.
//
// Solidity: function getPersonAge(address ip) view returns(uint8)
func (_Apiname *ApinameCaller) GetPersonAge(opts *bind.CallOpts, ip common.Address) (uint8, error) {
	var out []interface{}
	err := _Apiname.contract.Call(opts, &out, "getPersonAge", ip)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPersonAge is a free data retrieval call binding the contract method 0x0cfce82e.
//
// Solidity: function getPersonAge(address ip) view returns(uint8)
func (_Apiname *ApinameSession) GetPersonAge(ip common.Address) (uint8, error) {
	return _Apiname.Contract.GetPersonAge(&_Apiname.CallOpts, ip)
}

// GetPersonAge is a free data retrieval call binding the contract method 0x0cfce82e.
//
// Solidity: function getPersonAge(address ip) view returns(uint8)
func (_Apiname *ApinameCallerSession) GetPersonAge(ip common.Address) (uint8, error) {
	return _Apiname.Contract.GetPersonAge(&_Apiname.CallOpts, ip)
}

// IsPersonExsist is a free data retrieval call binding the contract method 0x974b6cfb.
//
// Solidity: function isPersonExsist(address ) view returns(bool)
func (_Apiname *ApinameCaller) IsPersonExsist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Apiname.contract.Call(opts, &out, "isPersonExsist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPersonExsist is a free data retrieval call binding the contract method 0x974b6cfb.
//
// Solidity: function isPersonExsist(address ) view returns(bool)
func (_Apiname *ApinameSession) IsPersonExsist(arg0 common.Address) (bool, error) {
	return _Apiname.Contract.IsPersonExsist(&_Apiname.CallOpts, arg0)
}

// IsPersonExsist is a free data retrieval call binding the contract method 0x974b6cfb.
//
// Solidity: function isPersonExsist(address ) view returns(bool)
func (_Apiname *ApinameCallerSession) IsPersonExsist(arg0 common.Address) (bool, error) {
	return _Apiname.Contract.IsPersonExsist(&_Apiname.CallOpts, arg0)
}

// AddPerson is a paid mutator transaction binding the contract method 0x36ccecaf.
//
// Solidity: function addPerson(uint8 id, uint8 age, string name, string story) returns(bool)
func (_Apiname *ApinameTransactor) AddPerson(opts *bind.TransactOpts, id uint8, age uint8, name string, story string) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "addPerson", id, age, name, story)
}

// AddPerson is a paid mutator transaction binding the contract method 0x36ccecaf.
//
// Solidity: function addPerson(uint8 id, uint8 age, string name, string story) returns(bool)
func (_Apiname *ApinameSession) AddPerson(id uint8, age uint8, name string, story string) (*types.Transaction, error) {
	return _Apiname.Contract.AddPerson(&_Apiname.TransactOpts, id, age, name, story)
}

// AddPerson is a paid mutator transaction binding the contract method 0x36ccecaf.
//
// Solidity: function addPerson(uint8 id, uint8 age, string name, string story) returns(bool)
func (_Apiname *ApinameTransactorSession) AddPerson(id uint8, age uint8, name string, story string) (*types.Transaction, error) {
	return _Apiname.Contract.AddPerson(&_Apiname.TransactOpts, id, age, name, story)
}

// GiveRight is a paid mutator transaction binding the contract method 0xc0e839a6.
//
// Solidity: function giveRight(uint8 Yourid, uint8 ip, uint8 time) returns(bool)
func (_Apiname *ApinameTransactor) GiveRight(opts *bind.TransactOpts, Yourid uint8, ip uint8, time uint8) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "giveRight", Yourid, ip, time)
}

// GiveRight is a paid mutator transaction binding the contract method 0xc0e839a6.
//
// Solidity: function giveRight(uint8 Yourid, uint8 ip, uint8 time) returns(bool)
func (_Apiname *ApinameSession) GiveRight(Yourid uint8, ip uint8, time uint8) (*types.Transaction, error) {
	return _Apiname.Contract.GiveRight(&_Apiname.TransactOpts, Yourid, ip, time)
}

// GiveRight is a paid mutator transaction binding the contract method 0xc0e839a6.
//
// Solidity: function giveRight(uint8 Yourid, uint8 ip, uint8 time) returns(bool)
func (_Apiname *ApinameTransactorSession) GiveRight(Yourid uint8, ip uint8, time uint8) (*types.Transaction, error) {
	return _Apiname.Contract.GiveRight(&_Apiname.TransactOpts, Yourid, ip, time)
}

// Myself is a paid mutator transaction binding the contract method 0x55de3eae.
//
// Solidity: function myself(uint8 seeId, uint8 Yourid) returns((uint8,uint8,string,string))
func (_Apiname *ApinameTransactor) Myself(opts *bind.TransactOpts, seeId uint8, Yourid uint8) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "myself", seeId, Yourid)
}

// Myself is a paid mutator transaction binding the contract method 0x55de3eae.
//
// Solidity: function myself(uint8 seeId, uint8 Yourid) returns((uint8,uint8,string,string))
func (_Apiname *ApinameSession) Myself(seeId uint8, Yourid uint8) (*types.Transaction, error) {
	return _Apiname.Contract.Myself(&_Apiname.TransactOpts, seeId, Yourid)
}

// Myself is a paid mutator transaction binding the contract method 0x55de3eae.
//
// Solidity: function myself(uint8 seeId, uint8 Yourid) returns((uint8,uint8,string,string))
func (_Apiname *ApinameTransactorSession) Myself(seeId uint8, Yourid uint8) (*types.Transaction, error) {
	return _Apiname.Contract.Myself(&_Apiname.TransactOpts, seeId, Yourid)
}

// SetAdminer is a paid mutator transaction binding the contract method 0x187b45bf.
//
// Solidity: function setAdminer(address ip) returns(bool)
func (_Apiname *ApinameTransactor) SetAdminer(opts *bind.TransactOpts, ip common.Address) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "setAdminer", ip)
}

// SetAdminer is a paid mutator transaction binding the contract method 0x187b45bf.
//
// Solidity: function setAdminer(address ip) returns(bool)
func (_Apiname *ApinameSession) SetAdminer(ip common.Address) (*types.Transaction, error) {
	return _Apiname.Contract.SetAdminer(&_Apiname.TransactOpts, ip)
}

// SetAdminer is a paid mutator transaction binding the contract method 0x187b45bf.
//
// Solidity: function setAdminer(address ip) returns(bool)
func (_Apiname *ApinameTransactorSession) SetAdminer(ip common.Address) (*types.Transaction, error) {
	return _Apiname.Contract.SetAdminer(&_Apiname.TransactOpts, ip)
}

// SetPersonAgeMem is a paid mutator transaction binding the contract method 0x25b9a2c8.
//
// Solidity: function setPersonAgeMem(address ip, uint8 age) returns()
func (_Apiname *ApinameTransactor) SetPersonAgeMem(opts *bind.TransactOpts, ip common.Address, age uint8) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "setPersonAgeMem", ip, age)
}

// SetPersonAgeMem is a paid mutator transaction binding the contract method 0x25b9a2c8.
//
// Solidity: function setPersonAgeMem(address ip, uint8 age) returns()
func (_Apiname *ApinameSession) SetPersonAgeMem(ip common.Address, age uint8) (*types.Transaction, error) {
	return _Apiname.Contract.SetPersonAgeMem(&_Apiname.TransactOpts, ip, age)
}

// SetPersonAgeMem is a paid mutator transaction binding the contract method 0x25b9a2c8.
//
// Solidity: function setPersonAgeMem(address ip, uint8 age) returns()
func (_Apiname *ApinameTransactorSession) SetPersonAgeMem(ip common.Address, age uint8) (*types.Transaction, error) {
	return _Apiname.Contract.SetPersonAgeMem(&_Apiname.TransactOpts, ip, age)
}

// SetPersonAgeSto is a paid mutator transaction binding the contract method 0xe8159e9e.
//
// Solidity: function setPersonAgeSto(address ip, uint8 age) returns()
func (_Apiname *ApinameTransactor) SetPersonAgeSto(opts *bind.TransactOpts, ip common.Address, age uint8) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "setPersonAgeSto", ip, age)
}

// SetPersonAgeSto is a paid mutator transaction binding the contract method 0xe8159e9e.
//
// Solidity: function setPersonAgeSto(address ip, uint8 age) returns()
func (_Apiname *ApinameSession) SetPersonAgeSto(ip common.Address, age uint8) (*types.Transaction, error) {
	return _Apiname.Contract.SetPersonAgeSto(&_Apiname.TransactOpts, ip, age)
}

// SetPersonAgeSto is a paid mutator transaction binding the contract method 0xe8159e9e.
//
// Solidity: function setPersonAgeSto(address ip, uint8 age) returns()
func (_Apiname *ApinameTransactorSession) SetPersonAgeSto(ip common.Address, age uint8) (*types.Transaction, error) {
	return _Apiname.Contract.SetPersonAgeSto(&_Apiname.TransactOpts, ip, age)
}

// ApinameAddPersonIterator is returned from FilterAddPerson and is used to iterate over the raw logs and unpacked data for AddPerson events raised by the Apiname contract.
type ApinameAddPersonIterator struct {
	Event *ApinameAddPerson // Event containing the contract specifics and raw log

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
func (it *ApinameAddPersonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApinameAddPerson)
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
		it.Event = new(ApinameAddPerson)
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
func (it *ApinameAddPersonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApinameAddPersonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApinameAddPerson represents a AddPerson event raised by the Apiname contract.
type ApinameAddPerson struct {
	Id        uint8
	Age       uint8
	Name      string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddPerson is a free log retrieval operation binding the contract event 0x7b53ca1f4cd943128c4f1669446ee2cb8f999d03b07e755d0c3d24cdf39b9456.
//
// Solidity: event AddPerson(uint8 id, uint8 age, string name, uint256 timestamp)
func (_Apiname *ApinameFilterer) FilterAddPerson(opts *bind.FilterOpts) (*ApinameAddPersonIterator, error) {

	logs, sub, err := _Apiname.contract.FilterLogs(opts, "AddPerson")
	if err != nil {
		return nil, err
	}
	return &ApinameAddPersonIterator{contract: _Apiname.contract, event: "AddPerson", logs: logs, sub: sub}, nil
}

// WatchAddPerson is a free log subscription operation binding the contract event 0x7b53ca1f4cd943128c4f1669446ee2cb8f999d03b07e755d0c3d24cdf39b9456.
//
// Solidity: event AddPerson(uint8 id, uint8 age, string name, uint256 timestamp)
func (_Apiname *ApinameFilterer) WatchAddPerson(opts *bind.WatchOpts, sink chan<- *ApinameAddPerson) (event.Subscription, error) {

	logs, sub, err := _Apiname.contract.WatchLogs(opts, "AddPerson")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApinameAddPerson)
				if err := _Apiname.contract.UnpackLog(event, "AddPerson", log); err != nil {
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

// ParseAddPerson is a log parse operation binding the contract event 0x7b53ca1f4cd943128c4f1669446ee2cb8f999d03b07e755d0c3d24cdf39b9456.
//
// Solidity: event AddPerson(uint8 id, uint8 age, string name, uint256 timestamp)
func (_Apiname *ApinameFilterer) ParseAddPerson(log types.Log) (*ApinameAddPerson, error) {
	event := new(ApinameAddPerson)
	if err := _Apiname.contract.UnpackLog(event, "AddPerson", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ApinameAddStoryIterator is returned from FilterAddStory and is used to iterate over the raw logs and unpacked data for AddStory events raised by the Apiname contract.
type ApinameAddStoryIterator struct {
	Event *ApinameAddStory // Event containing the contract specifics and raw log

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
func (it *ApinameAddStoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApinameAddStory)
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
		it.Event = new(ApinameAddStory)
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
func (it *ApinameAddStoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApinameAddStoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApinameAddStory represents a AddStory event raised by the Apiname contract.
type ApinameAddStory struct {
	Id        uint8
	Story     string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddStory is a free log retrieval operation binding the contract event 0xb229e062178df6eb9852151c39ba2a3e2f850a19cf2c5f285ef97c998898e34b.
//
// Solidity: event AddStory(uint8 id, string story, uint256 timestamp)
func (_Apiname *ApinameFilterer) FilterAddStory(opts *bind.FilterOpts) (*ApinameAddStoryIterator, error) {

	logs, sub, err := _Apiname.contract.FilterLogs(opts, "AddStory")
	if err != nil {
		return nil, err
	}
	return &ApinameAddStoryIterator{contract: _Apiname.contract, event: "AddStory", logs: logs, sub: sub}, nil
}

// WatchAddStory is a free log subscription operation binding the contract event 0xb229e062178df6eb9852151c39ba2a3e2f850a19cf2c5f285ef97c998898e34b.
//
// Solidity: event AddStory(uint8 id, string story, uint256 timestamp)
func (_Apiname *ApinameFilterer) WatchAddStory(opts *bind.WatchOpts, sink chan<- *ApinameAddStory) (event.Subscription, error) {

	logs, sub, err := _Apiname.contract.WatchLogs(opts, "AddStory")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApinameAddStory)
				if err := _Apiname.contract.UnpackLog(event, "AddStory", log); err != nil {
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

// ParseAddStory is a log parse operation binding the contract event 0xb229e062178df6eb9852151c39ba2a3e2f850a19cf2c5f285ef97c998898e34b.
//
// Solidity: event AddStory(uint8 id, string story, uint256 timestamp)
func (_Apiname *ApinameFilterer) ParseAddStory(log types.Log) (*ApinameAddStory, error) {
	event := new(ApinameAddStory)
	if err := _Apiname.contract.UnpackLog(event, "AddStory", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ApinameGiveAccessIterator is returned from FilterGiveAccess and is used to iterate over the raw logs and unpacked data for GiveAccess events raised by the Apiname contract.
type ApinameGiveAccessIterator struct {
	Event *ApinameGiveAccess // Event containing the contract specifics and raw log

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
func (it *ApinameGiveAccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApinameGiveAccess)
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
		it.Event = new(ApinameGiveAccess)
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
func (it *ApinameGiveAccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApinameGiveAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApinameGiveAccess represents a GiveAccess event raised by the Apiname contract.
type ApinameGiveAccess struct {
	Id        uint8
	Ip        uint8
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGiveAccess is a free log retrieval operation binding the contract event 0xf13a37be7a293526045ec51b9f41b9681bf7d86406d550db74ab7ba6991ff663.
//
// Solidity: event giveAccess(uint8 id, uint8 ip, uint256 timestamp)
func (_Apiname *ApinameFilterer) FilterGiveAccess(opts *bind.FilterOpts) (*ApinameGiveAccessIterator, error) {

	logs, sub, err := _Apiname.contract.FilterLogs(opts, "giveAccess")
	if err != nil {
		return nil, err
	}
	return &ApinameGiveAccessIterator{contract: _Apiname.contract, event: "giveAccess", logs: logs, sub: sub}, nil
}

// WatchGiveAccess is a free log subscription operation binding the contract event 0xf13a37be7a293526045ec51b9f41b9681bf7d86406d550db74ab7ba6991ff663.
//
// Solidity: event giveAccess(uint8 id, uint8 ip, uint256 timestamp)
func (_Apiname *ApinameFilterer) WatchGiveAccess(opts *bind.WatchOpts, sink chan<- *ApinameGiveAccess) (event.Subscription, error) {

	logs, sub, err := _Apiname.contract.WatchLogs(opts, "giveAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApinameGiveAccess)
				if err := _Apiname.contract.UnpackLog(event, "giveAccess", log); err != nil {
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

// ParseGiveAccess is a log parse operation binding the contract event 0xf13a37be7a293526045ec51b9f41b9681bf7d86406d550db74ab7ba6991ff663.
//
// Solidity: event giveAccess(uint8 id, uint8 ip, uint256 timestamp)
func (_Apiname *ApinameFilterer) ParseGiveAccess(log types.Log) (*ApinameGiveAccess, error) {
	event := new(ApinameGiveAccess)
	if err := _Apiname.contract.UnpackLog(event, "giveAccess", log); err != nil {
		return nil, err
	}
	return event, nil
}
