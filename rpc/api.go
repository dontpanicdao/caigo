package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/types"
)

type Events struct {
	Events []Event `json:"events"`
}

type Event struct {
	*types.Event
	FromAddress     string `json:"from_address"`
	BlockHash       string `json:"block_hash"`
	BlockNumber     int    `json:"block_number"`
	TransactionHash string `json:"transaction_hash"`
}

type EventParams struct {
	FromBlock  uint64 `json:"fromBlock"`
	ToBlock    uint64 `json:"toBlock"`
	PageSize   uint64 `json:"page_size"`
	PageNumber uint64 `json:"page_number"`
}

// Call a starknet function without creating a StarkNet transaction.
func (sc *Client) Call(ctx context.Context, call types.FunctionCall, hash string) ([]string, error) {
	call.EntryPointSelector = caigo.BigToHex(caigo.GetSelectorFromName(call.EntryPointSelector))
	if len(call.Calldata) == 0 {
		call.Calldata = make([]string, 0)
	}

	var result []string
	if err := sc.do(ctx, "starknet_call", &result, call, hash); err != nil {
		return nil, err
	}

	return result, nil
}

// BlockNumber gets the most recent accepted block number.
func (sc *Client) BlockNumber(ctx context.Context) (*big.Int, error) {
	var blockNumber big.Int
	if err := sc.c.CallContext(ctx, &blockNumber, "starknet_blockNumber"); err != nil {
		return nil, err
	}

	return &blockNumber, nil
}

// BlockByHash gets block information given the block id.
func (sc *Client) BlockByHash(ctx context.Context, hash string, scope string) (*types.Block, error) {
	var block types.Block
	if err := sc.do(ctx, "starknet_getBlockByHash", &block, hash, scope); err != nil {
		return nil, err
	}

	return &block, nil
}

// BlockByNumber gets block information given the block number (its height).
func (sc *Client) BlockByNumber(ctx context.Context, number *big.Int, scope string) (*types.Block, error) {
	var block types.Block
	if err := sc.do(ctx, "starknet_getBlockByNumber", &block, toBlockNumArg(number), scope); err != nil {
		return nil, err
	}

	return &block, nil
}

// CodeAt returns the contract and class associated with the an address.
// Deprecated: you should use ClassAt and TransactionByHash to access the
// associated values.
func (sc *Client) CodeAt(ctx context.Context, address string) (*types.Code, error) {
	var contractRaw struct {
		Bytecode []string `json:"bytecode"`
		AbiRaw   string   `json:"abi"`
		Abi      types.ABI
	}
	if err := sc.do(ctx, "starknet_getCode", &contractRaw, address); err != nil {
		return nil, err
	}

	contract := types.Code{
		Bytecode: contractRaw.Bytecode,
	}
	if err := json.Unmarshal([]byte(contractRaw.AbiRaw), &contract.Abi); err != nil {
		return nil, err
	}

	return &contract, nil
}

// Class gets the contract class definition associated with the given hash.
func (sc *Client) Class(ctx context.Context, hash string) (*types.ContractClass, error) {
	var rawClass types.ContractClass
	if err := sc.do(ctx, "starknet_getClass", &rawClass, hash); err != nil {
		return nil, err
	}

	return &rawClass, nil
}

// ClassAt get the contract class definition at the given address.
func (sc *Client) ClassAt(ctx context.Context, address string) (*types.ContractClass, error) {
	var rawClass types.ContractClass
	if err := sc.do(ctx, "starknet_getClassAt", &rawClass, address); err != nil {
		return nil, err
	}

	return &rawClass, nil
}

// ClassHashAt gets the contract class hash for the contract deployed at the given address.
func (sc *Client) ClassHashAt(ctx context.Context, address string) (string, error) {
	result := new(string)
	if err := sc.do(ctx, "starknet_getClassHashAt", &result, address); err != nil {
		return "", err
	}

	return *result, nil
}

// GetStorageAt gets the value of the storage at the given address and key.
func (sc *Client) GetStorageAt(ctx context.Context, contractAddress, key, blockHashOrTag string) (string, error) {
	var value string
	hashKey := fmt.Sprintf("0x%s", caigo.GetSelectorFromName(key).Text(16))
	err := sc.do(ctx, "starknet_getStorageAt", &value, contractAddress, hashKey, blockHashOrTag)
	return value, err
}

// TransactionByHash gets the details and status of a submitted transaction.
func (sc *Client) TransactionByHash(ctx context.Context, hash string) (*types.Transaction, error) {
	var tx types.Transaction
	if err := sc.do(ctx, "starknet_getTransactionByHash", &tx, hash); err != nil {
		return nil, err
	} else if tx.TransactionHash == "" {
		return nil, ErrNotFound
	}

	return &tx, nil
}

// TransactionReceipt gets the transaction receipt by the transaction hash.
func (sc *Client) TransactionReceipt(ctx context.Context, hash string) (*types.TransactionReceipt, error) {
	var receipt types.TransactionReceipt
	err := sc.do(ctx, "starknet_getTransactionReceipt", &receipt, hash)
	if err != nil {
		return nil, err
	} else if receipt.TransactionHash == "" {
		return nil, ErrNotFound
	}

	return &receipt, nil
}

// Events returns all events matching the given filter
// TODO: check the query parameters as they include filter directives that have
// not been implemented. For more details, check the
// [specification](https://github.com/starkware-libs/starknet-specs/blob/master/api/starknet_api_openrpc.json)
func (sc *Client) Events(ctx context.Context, evParams EventParams) (*Events, error) {
	var result Events
	if err := sc.do(ctx, "starknet_getEvents", &result, evParams); err != nil {
		return nil, err
	}

	return &result, nil
}

// Estimate the fee for a given StarkNet transaction.
func (sc *Client) EstimateFee(context.Context, types.Transaction) (*types.FeeEstimate, error) {
	panic("not implemented")
}

// AccountNonce gets the latest nonce associated with the given address
func (sc *Client) AccountNonce(context.Context, string) (*big.Int, error) {
	panic("not implemented")
}

func toBlockNumArg(number *big.Int) interface{} {
	var numOrTag interface{}

	if number == nil {
		numOrTag = "latest"
	} else if number.Cmp(big.NewInt(-1)) == 0 {
		numOrTag = "pending"
	} else {
		numOrTag = number.Uint64()
	}

	return numOrTag
}
