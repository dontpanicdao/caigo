package rpc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dontpanicdao/caigo/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// SyncResponse is the Starknet RPC type returned by the Syncing method.
type SyncResponse struct {
	StartingBlockHash string `json:"starting_block_hash"`
	StartingBlockNum  string `json:"starting_block_num"`
	CurrentBlockHash  string `json:"current_block_hash"`
	CurrentBlockNum   string `json:"current_block_num"`
	HighestBlockHash  string `json:"highest_block_hash"`
	HighestBlockNum   string `json:"highest_block_num"`
}

// ErrNotFound is returned by API methods if the requested item does not exist.
var ErrNotFound = errors.New("not found")

type callCloser interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	Close()
}

// Client provides the client type for caigo/rpc implementation.
type Client struct {
	c callCloser
}

var _ types.Provider = &Client{}

// Dial connects a client to the given URL. It creates a `go-ethereum/rpc` *Client and relies on context.Background().
func Dial(rawurl string) (*Client, error) {
	return DialContext(context.Background(), rawurl)
}

// DialContext connects a client to the given URL with an existing context. It creates a `go-ethereum/rpc` *Client.
func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

// NewClient creates a *Client from an existing `go-ethereum/rpc` *Client.
func NewClient(c *rpc.Client) *Client {
	return &Client{c: c}
}

// Close closes the underlying client.
func (sc *Client) Close() {
	sc.c.Close()
}

// ChainID retrieves the current chain ID for transaction replay protection.
func (sc *Client) ChainID(ctx context.Context) (string, error) {
	var result string
	err := sc.c.CallContext(ctx, &result, "starknet_chainId", []interface{}{}...)
	if err != nil {
		return "", err
	}
	return result, err
}

// Syncing checks the syncing status of the node.
func (sc *Client) Syncing(ctx context.Context) (*SyncResponse, error) {
	var result SyncResponse
	if err := sc.c.CallContext(ctx, &result, "starknet_syncing", []interface{}{}...); err != nil {
		return nil, err
	}
	return &result, nil
}

// ProtocolVersion returns the current starknet protocol version identifier, as supported by this sequencer.
func (sc *Client) ProtocolVersion(ctx context.Context) (string, error) {
	var protocol string
	err := sc.do(ctx, "starknet_protocolVersion", &protocol, []interface{}{}...)
	return protocol, err
}

func (sc *Client) do(ctx context.Context, method string, data interface{}, args ...interface{}) error {
	var raw json.RawMessage
	err := sc.c.CallContext(ctx, &raw, method, args...)
	if err != nil {
		return err
	}
	if len(raw) == 0 {
		return ErrNotFound
	}
	if err := json.Unmarshal(raw, &data); err != nil {
		return err
	}
	return nil
}
