package rpc

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/dontpanicdao/caigo/types"
)

// AddDeployTransactionOutput provides the output for AddDeployTransaction.
type AddDeployTransactionOutput struct {
	TransactionHash string `json:"transaction_hash"`
	ContractAddress string `json:"contract_address"`
}

func (sc *Client) Invoke(context.Context, types.Transaction) (*types.AddTxResponse, error) {
	panic("'starknet_addInvokeTransaction' not implemented")
}

func (sc *Client) Declare(context.Context, types.Transaction) (*types.AddTxResponse, error) {
	panic("'starknet_addDeclareTransaction' not implemented")
}

// AddDeployTransaction allows to declare a class and instantiate the
// associated contract in one command. This function will be deprecated and
// replaced by AddDeclareTransaction to declare a class, followed by
// AddInvokeTransaction to instantiate the contract. For now, it remains the only
// way to deploy an account without being charged for it.
func (sc *Client) AddDeployTransaction(ctx context.Context, contractAddressSalt string, constructorCallData []string, contractDefinition types.ContractClass) (*AddDeployTransactionOutput, error) {
	program, ok := contractDefinition.Program.(string)
	if !ok {
		data, err := json.Marshal(contractDefinition.Program)
		if err != nil {
			return nil, err
		}
		program, err = encodeProgram(data)
		if err != nil {
			return nil, err
		}
	}
	contractDefinition.Program = program

	var result AddDeployTransactionOutput
	err := sc.do(ctx, "starknet_addDeployTransaction", &result, contractAddressSalt, constructorCallData, contractDefinition)

	return &result, err
}

func encodeProgram(content []byte) (string, error) {
	buf := bytes.NewBuffer(nil)
	gzipContent := gzip.NewWriter(buf)
	_, err := gzipContent.Write(content)
	if err != nil {
		return "", err
	}
	gzipContent.Close()
	program := base64.StdEncoding.EncodeToString(buf.Bytes())
	return program, nil
}
