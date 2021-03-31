// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package homeWork

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

// ApinameABI is the input ABI used to generate the binding from.
const ApinameABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Login\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"Register\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"Usedpasswd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"setPassword\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"SetPassword\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getIP\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// GetIP is a free data retrieval call binding the contract method 0x3aed55e2.
//
// Solidity: function getIP(uint8 id) view returns(address)
func (_Apiname *ApinameCaller) GetIP(opts *bind.CallOpts, id uint8) (common.Address, error) {
	var out []interface{}
	err := _Apiname.contract.Call(opts, &out, "getIP", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIP is a free data retrieval call binding the contract method 0x3aed55e2.
//
// Solidity: function getIP(uint8 id) view returns(address)
func (_Apiname *ApinameSession) GetIP(id uint8) (common.Address, error) {
	return _Apiname.Contract.GetIP(&_Apiname.CallOpts, id)
}

// GetIP is a free data retrieval call binding the contract method 0x3aed55e2.
//
// Solidity: function getIP(uint8 id) view returns(address)
func (_Apiname *ApinameCallerSession) GetIP(id uint8) (common.Address, error) {
	return _Apiname.Contract.GetIP(&_Apiname.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Apiname *ApinameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Apiname.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Apiname *ApinameSession) Owner() (common.Address, error) {
	return _Apiname.Contract.Owner(&_Apiname.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Apiname *ApinameCallerSession) Owner() (common.Address, error) {
	return _Apiname.Contract.Owner(&_Apiname.CallOpts)
}

// Login is a paid mutator transaction binding the contract method 0x9d1cba1a.
//
// Solidity: function login(uint8 id, string passwd) returns(bool)
func (_Apiname *ApinameTransactor) Login(opts *bind.TransactOpts, id uint8, passwd string) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "login", id, passwd)
}

// Login is a paid mutator transaction binding the contract method 0x9d1cba1a.
//
// Solidity: function login(uint8 id, string passwd) returns(bool)
func (_Apiname *ApinameSession) Login(id uint8, passwd string) (*types.Transaction, error) {
	return _Apiname.Contract.Login(&_Apiname.TransactOpts, id, passwd)
}

// Login is a paid mutator transaction binding the contract method 0x9d1cba1a.
//
// Solidity: function login(uint8 id, string passwd) returns(bool)
func (_Apiname *ApinameTransactorSession) Login(id uint8, passwd string) (*types.Transaction, error) {
	return _Apiname.Contract.Login(&_Apiname.TransactOpts, id, passwd)
}

// Register is a paid mutator transaction binding the contract method 0xd3eb9541.
//
// Solidity: function register(uint8 id, string passwd) returns(bool)
func (_Apiname *ApinameTransactor) Register(opts *bind.TransactOpts, id uint8, passwd string) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "register", id, passwd)
}

// Register is a paid mutator transaction binding the contract method 0xd3eb9541.
//
// Solidity: function register(uint8 id, string passwd) returns(bool)
func (_Apiname *ApinameSession) Register(id uint8, passwd string) (*types.Transaction, error) {
	return _Apiname.Contract.Register(&_Apiname.TransactOpts, id, passwd)
}

// Register is a paid mutator transaction binding the contract method 0xd3eb9541.
//
// Solidity: function register(uint8 id, string passwd) returns(bool)
func (_Apiname *ApinameTransactorSession) Register(id uint8, passwd string) (*types.Transaction, error) {
	return _Apiname.Contract.Register(&_Apiname.TransactOpts, id, passwd)
}

// SetPassword is a paid mutator transaction binding the contract method 0xd9efca05.
//
// Solidity: function setPassword(uint8 id, string Usedpasswd, string passwd) returns(bool)
func (_Apiname *ApinameTransactor) SetPassword(opts *bind.TransactOpts, id uint8, Usedpasswd string, passwd string) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "setPassword", id, Usedpasswd, passwd)
}

// SetPassword is a paid mutator transaction binding the contract method 0xd9efca05.
//
// Solidity: function setPassword(uint8 id, string Usedpasswd, string passwd) returns(bool)
func (_Apiname *ApinameSession) SetPassword(id uint8, Usedpasswd string, passwd string) (*types.Transaction, error) {
	return _Apiname.Contract.SetPassword(&_Apiname.TransactOpts, id, Usedpasswd, passwd)
}

// SetPassword is a paid mutator transaction binding the contract method 0xd9efca05.
//
// Solidity: function setPassword(uint8 id, string Usedpasswd, string passwd) returns(bool)
func (_Apiname *ApinameTransactorSession) SetPassword(id uint8, Usedpasswd string, passwd string) (*types.Transaction, error) {
	return _Apiname.Contract.SetPassword(&_Apiname.TransactOpts, id, Usedpasswd, passwd)
}

// ApinameLoginIterator is returned from FilterLogin and is used to iterate over the raw logs and unpacked data for Login events raised by the Apiname contract.
type ApinameLoginIterator struct {
	Event *ApinameLogin // Event containing the contract specifics and raw log

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
func (it *ApinameLoginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApinameLogin)
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
		it.Event = new(ApinameLogin)
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
func (it *ApinameLoginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApinameLoginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApinameLogin represents a Login event raised by the Apiname contract.
type ApinameLogin struct {
	Id   uint8
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogin is a free log retrieval operation binding the contract event 0x554145b27dd70d8d66a4693ccd8b2b4793b89e6fd4e0c1596c71f5ff4844b042.
//
// Solidity: event Login(uint8 id, uint256 time)
func (_Apiname *ApinameFilterer) FilterLogin(opts *bind.FilterOpts) (*ApinameLoginIterator, error) {

	logs, sub, err := _Apiname.contract.FilterLogs(opts, "Login")
	if err != nil {
		return nil, err
	}
	return &ApinameLoginIterator{contract: _Apiname.contract, event: "Login", logs: logs, sub: sub}, nil
}

// WatchLogin is a free log subscription operation binding the contract event 0x554145b27dd70d8d66a4693ccd8b2b4793b89e6fd4e0c1596c71f5ff4844b042.
//
// Solidity: event Login(uint8 id, uint256 time)
func (_Apiname *ApinameFilterer) WatchLogin(opts *bind.WatchOpts, sink chan<- *ApinameLogin) (event.Subscription, error) {

	logs, sub, err := _Apiname.contract.WatchLogs(opts, "Login")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApinameLogin)
				if err := _Apiname.contract.UnpackLog(event, "Login", log); err != nil {
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

// ParseLogin is a log parse operation binding the contract event 0x554145b27dd70d8d66a4693ccd8b2b4793b89e6fd4e0c1596c71f5ff4844b042.
//
// Solidity: event Login(uint8 id, uint256 time)
func (_Apiname *ApinameFilterer) ParseLogin(log types.Log) (*ApinameLogin, error) {
	event := new(ApinameLogin)
	if err := _Apiname.contract.UnpackLog(event, "Login", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ApinameRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the Apiname contract.
type ApinameRegisterIterator struct {
	Event *ApinameRegister // Event containing the contract specifics and raw log

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
func (it *ApinameRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApinameRegister)
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
		it.Event = new(ApinameRegister)
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
func (it *ApinameRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApinameRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApinameRegister represents a Register event raised by the Apiname contract.
type ApinameRegister struct {
	Id     uint8
	Passwd string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0x657a2ffaaa160de47a5bd694dddcc6127ec9d1b12ff968fa358181753432c92e.
//
// Solidity: event Register(uint8 id, string passwd)
func (_Apiname *ApinameFilterer) FilterRegister(opts *bind.FilterOpts) (*ApinameRegisterIterator, error) {

	logs, sub, err := _Apiname.contract.FilterLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return &ApinameRegisterIterator{contract: _Apiname.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0x657a2ffaaa160de47a5bd694dddcc6127ec9d1b12ff968fa358181753432c92e.
//
// Solidity: event Register(uint8 id, string passwd)
func (_Apiname *ApinameFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *ApinameRegister) (event.Subscription, error) {

	logs, sub, err := _Apiname.contract.WatchLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApinameRegister)
				if err := _Apiname.contract.UnpackLog(event, "Register", log); err != nil {
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

// ParseRegister is a log parse operation binding the contract event 0x657a2ffaaa160de47a5bd694dddcc6127ec9d1b12ff968fa358181753432c92e.
//
// Solidity: event Register(uint8 id, string passwd)
func (_Apiname *ApinameFilterer) ParseRegister(log types.Log) (*ApinameRegister, error) {
	event := new(ApinameRegister)
	if err := _Apiname.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ApinameSetPasswordIterator is returned from FilterSetPassword and is used to iterate over the raw logs and unpacked data for SetPassword events raised by the Apiname contract.
type ApinameSetPasswordIterator struct {
	Event *ApinameSetPassword // Event containing the contract specifics and raw log

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
func (it *ApinameSetPasswordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApinameSetPassword)
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
		it.Event = new(ApinameSetPassword)
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
func (it *ApinameSetPasswordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApinameSetPasswordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApinameSetPassword represents a SetPassword event raised by the Apiname contract.
type ApinameSetPassword struct {
	Passwd string
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPassword is a free log retrieval operation binding the contract event 0x719222a086573305529dd5e415194e75cdf7edada37716f67f86fc1d42ca41b9.
//
// Solidity: event SetPassword(string passwd, uint256 time)
func (_Apiname *ApinameFilterer) FilterSetPassword(opts *bind.FilterOpts) (*ApinameSetPasswordIterator, error) {

	logs, sub, err := _Apiname.contract.FilterLogs(opts, "SetPassword")
	if err != nil {
		return nil, err
	}
	return &ApinameSetPasswordIterator{contract: _Apiname.contract, event: "SetPassword", logs: logs, sub: sub}, nil
}

// WatchSetPassword is a free log subscription operation binding the contract event 0x719222a086573305529dd5e415194e75cdf7edada37716f67f86fc1d42ca41b9.
//
// Solidity: event SetPassword(string passwd, uint256 time)
func (_Apiname *ApinameFilterer) WatchSetPassword(opts *bind.WatchOpts, sink chan<- *ApinameSetPassword) (event.Subscription, error) {

	logs, sub, err := _Apiname.contract.WatchLogs(opts, "SetPassword")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApinameSetPassword)
				if err := _Apiname.contract.UnpackLog(event, "SetPassword", log); err != nil {
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

// ParseSetPassword is a log parse operation binding the contract event 0x719222a086573305529dd5e415194e75cdf7edada37716f67f86fc1d42ca41b9.
//
// Solidity: event SetPassword(string passwd, uint256 time)
func (_Apiname *ApinameFilterer) ParseSetPassword(log types.Log) (*ApinameSetPassword, error) {
	event := new(ApinameSetPassword)
	if err := _Apiname.contract.UnpackLog(event, "SetPassword", log); err != nil {
		return nil, err
	}
	return event, nil
}
