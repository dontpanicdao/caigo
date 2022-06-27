package rpc

import (
	"context"
	"testing"
)

// TestGetTransactionByBlockHashAndIndex tests transaction by blockHash and txIndex
func TestGetTransactionByBlockHashAndIndex(t *testing.T) {
	testConfig := beforeEach(t)
	defer testConfig.client.Close()

	type testSetType struct {
		BlockHash                  string
		TxIndex                    int
		ExpectedTxHash             string
		ExpectedContractAddress    string
		ExpectedEntrypointSelector string
	}
	testSet := map[string][]testSetType{
		"mock": {
			{
				BlockHash:                  "0xdeadbeef",
				TxIndex:                    7,
				ExpectedTxHash:             "0xdeadbeef",
				ExpectedContractAddress:    "0xdeadbeef",
				ExpectedEntrypointSelector: "0xdeadbeef",
			},
		},
		"testnet": {
			{
				BlockHash:                  "0x115aa451e374dbfdeb6f8d4c70133a39c6bb7b2948a4a3f0c9d5dda30f94044",
				TxIndex:                    3,
				ExpectedTxHash:             "0x179124db5707ea54a44c7e9cc2e654a0160a0f6fa9a3ef8f3f062a659224da1",
				ExpectedContractAddress:    "0x6bc8601525b88448ecc22de7ce94caa4cd7b3d594ad5c8360f21c7d6bfa5085",
				ExpectedEntrypointSelector: "0x15d40a3d6ca2ac30f4031e42be28da9b056fef9bb7357ac5e85627ee876e5ad",
			},
		},
		"mainnet": {
			{
				BlockHash:                  "0x6f8e6413281c43bfcb9f96e315a08c57c619c9da4b10e2cb7d33369f3fb75a0",
				TxIndex:                    1,
				ExpectedTxHash:             "0x2b8ff0898ab240a45082fa2d2e0118bfc2a30959a2d898bf3668d8c453963cb",
				ExpectedContractAddress:    "0x443df1c1f5b55878b44c623557dc0bf7bee18c01813fe09fc6e47811a467976",
				ExpectedEntrypointSelector: "0x12ead94ae9d3f9d2bdb6b847cf255f1f398193a1f88884a0ae8e18f24a037b6",
			},
		},
	}[testEnv]

	for _, test := range testSet {
		tx, err := testConfig.client.GetTransactionByBlockHashAndIndex(context.Background(), test.BlockHash, test.TxIndex)
		if err != nil {
			t.Fatal(err)
		}
		if tx == nil || tx.TransactionHash != test.ExpectedTxHash {
			t.Fatal("transaction should exist and match the tx hash")
		}
		if tx.ContractAddress != test.ExpectedContractAddress {
			t.Fatalf("expecting contract %s, got %s", test.ExpectedContractAddress, tx.ContractAddress)
		}
		if tx.EntryPointSelector != test.ExpectedEntrypointSelector {
			t.Fatalf("expecting entrypoint %s, got %s", test.ExpectedEntrypointSelector, tx.EntryPointSelector)
		}
	}
}

// TestGetTransactionByBlockNumberAndIndex tests transaction by blockHash and txIndex
func TestGetTransactionByBlockNumberAndIndex(t *testing.T) {
	testConfig := beforeEach(t)
	defer testConfig.client.Close()

	type testSetType struct {
		BlockNumberOrTag           interface{}
		TxIndex                    int
		ExpectedTxHash             string
		ExpectedContractAddress    string
		ExpectedEntrypointSelector string
	}
	testSet := map[string][]testSetType{
		"mock": {
			{
				BlockNumberOrTag:           7,
				TxIndex:                    7,
				ExpectedTxHash:             "0xdeadbeef",
				ExpectedContractAddress:    "0xdeadbeef",
				ExpectedEntrypointSelector: "0xdeadbeef",
			},
		},
		"testnet": {
			{
				BlockNumberOrTag:           242060,
				TxIndex:                    3,
				ExpectedTxHash:             "0x179124db5707ea54a44c7e9cc2e654a0160a0f6fa9a3ef8f3f062a659224da1",
				ExpectedContractAddress:    "0x6bc8601525b88448ecc22de7ce94caa4cd7b3d594ad5c8360f21c7d6bfa5085",
				ExpectedEntrypointSelector: "0x15d40a3d6ca2ac30f4031e42be28da9b056fef9bb7357ac5e85627ee876e5ad",
			},
		},
		"mainnet": {
			{
				BlockNumberOrTag:           1500,
				TxIndex:                    1,
				ExpectedTxHash:             "0x2b8ff0898ab240a45082fa2d2e0118bfc2a30959a2d898bf3668d8c453963cb",
				ExpectedContractAddress:    "0x443df1c1f5b55878b44c623557dc0bf7bee18c01813fe09fc6e47811a467976",
				ExpectedEntrypointSelector: "0x12ead94ae9d3f9d2bdb6b847cf255f1f398193a1f88884a0ae8e18f24a037b6",
			},
		},
	}[testEnv]

	for _, test := range testSet {
		tx, err := testConfig.client.GetTransactionByBlockNumberAndIndex(context.Background(), test.BlockNumberOrTag, test.TxIndex)
		if err != nil {
			t.Fatal(err)
		}
		if tx == nil || tx.TransactionHash != test.ExpectedTxHash {
			t.Fatal("transaction should exist and match the tx hash")
		}
		if tx.ContractAddress != test.ExpectedContractAddress {
			t.Fatalf("expecting contract %s, got %s", test.ExpectedContractAddress, tx.ContractAddress)
		}
		if tx.EntryPointSelector != test.ExpectedEntrypointSelector {
			t.Fatalf("expecting entrypoint %s, got %s", test.ExpectedEntrypointSelector, tx.EntryPointSelector)
		}
	}
}

// TestTransactionByHash tests transaction by hash
func TestTransactionByHash(t *testing.T) {
	testConfig := beforeEach(t)
	defer testConfig.client.Close()

	type testSetType struct {
		TxHash                     string
		ExpectedContractAddress    string
		ExpectedEntrypointSelector string
	}
	testSet := map[string][]testSetType{
		"mock": {
			{
				TxHash:                     "0xdeadbeef",
				ExpectedContractAddress:    "0xdeadbeef",
				ExpectedEntrypointSelector: "0xdeadbeef",
			},
		},
		"testnet": {
			{
				TxHash:                     "0x705547f8f2f8fdfb10ed533d909f76482bb293c5a32648d476774516a0bebd0",
				ExpectedContractAddress:    "0x315e364b162653e5c7b23efd34f8da27ba9c069b68e3042b7d76ce1df890313",
				ExpectedEntrypointSelector: "0x15d40a3d6ca2ac30f4031e42be28da9b056fef9bb7357ac5e85627ee876e5ad",
			},
		},
		"mainnet": {
			{
				TxHash:                     "0x5f904b9185d4ed442846ac7e26bc4c60249a2a7f0bb85376c0bc7459665bae6",
				ExpectedContractAddress:    "0x3b4be7def2fc08589348966255e101824928659ebb724855223ff3a8c831efa",
				ExpectedEntrypointSelector: "0x2913ee03e5e3308c41e308bd391ea4faac9b9cb5062c76a6b3ab4f65397e106",
			},
		},
	}[testEnv]

	for _, test := range testSet {
		tx, err := testConfig.client.TransactionByHash(context.Background(), test.TxHash)
		if err != nil {
			t.Fatal(err)
		}
		if tx == nil || tx.TransactionHash != test.TxHash {
			t.Fatal("transaction should exist and match the tx hash")
		}
		if tx.ContractAddress != test.ExpectedContractAddress {
			t.Fatalf("expecting contract %s, got %s", test.ExpectedContractAddress, tx.ContractAddress)
		}
		if tx.EntryPointSelector != test.ExpectedEntrypointSelector {
			t.Fatalf("expecting entrypoint %s, got %s", test.ExpectedEntrypointSelector, tx.EntryPointSelector)
		}
	}
}

// TestTransactionReceipt tests transaction receipt
func TestTransactionReceipt(t *testing.T) {

	testConfig := beforeEach(t)
	defer testConfig.client.Close()

	type testSetType struct {
		TxHash         string
		ExpectedStatus string
	}
	testSet := map[string][]testSetType{
		"mock": {
			{
				TxHash:         "0xdeadbeef",
				ExpectedStatus: "ACCEPTED_ON_L1",
			},
		},
		"testnet": {
			{
				TxHash:         "0x705547f8f2f8fdfb10ed533d909f76482bb293c5a32648d476774516a0bebd0",
				ExpectedStatus: "ACCEPTED_ON_L1",
			},
		},
		"mainnet": {
			{
				TxHash:         "0x5f904b9185d4ed442846ac7e26bc4c60249a2a7f0bb85376c0bc7459665bae6",
				ExpectedStatus: "ACCEPTED_ON_L1",
			},
		},
	}[testEnv]

	for _, test := range testSet {
		txReceipt, err := testConfig.client.TransactionReceipt(context.Background(), test.TxHash)
		if err != nil {
			t.Fatal(err)
		}
		if txReceipt == nil || txReceipt.TransactionHash != test.TxHash {
			t.Fatal("transaction should exist and match the tx hash")
		}
		if txReceipt.Status != test.ExpectedStatus {
			t.Fatalf("expecting status %s, got %s", test.ExpectedStatus, txReceipt.Status)
		}
	}
}
