package rpc

// func TestIntegrationNodePendingTransactions(t *testing.T) {
// 	godotenv.Load()
// 	if os.Getenv("INTEGRATION") != "true" {
// 		t.Skip("Skipping integration test")
// 	}
// 	node, err := NewNode("node")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	_, err = node.PendingTransactions(context.Background())
// 	// TODO: this code is not yet enable in pathfinder v0.2.2-alpha.
// 	if err == nil || !strings.Contains(err.Error(), "Method not found") {
// 		t.Fatal(err)
// 	}
// }

// //go:embed test/openzeppelin.json
// var openzeppelin []byte

// func TestIntegrationNodeAddDeclareTransaction(t *testing.T) {
// 	godotenv.Load()
// 	if os.Getenv("INTEGRATION") != "true" {
// 		t.Skip("Skipping integration test")
// 	}
// 	node, err := NewNode("node")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	classFile := ClassFile{}
// 	err = json.Unmarshal(openzeppelin, &classFile)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	declare, err := node.AddDeclareTransaction(context.Background(), classFile.Program, classFile.EntryPointsByType)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("Class: %s\n", declare.ClassHash)
// 	fmt.Printf("Tx: %s\n", declare.TxHash)
// }

// func TestIntegrationNodeaddInvokeTransaction(t *testing.T) {
// 	godotenv.Load()
// 	if os.Getenv("INTEGRATION") != "true" {
// 		t.Skip("Skipping integration test")
// 	}
// 	node, err := NewNode("node")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	contractAddress := "0x6c49d194c895308b3b4267dd41326afb1a360cfc5d7670c499c13dc3b0fd8ed"
// 	methodName := "constructor"
// 	callData := []string{"0x5ff5eff3bed10c5109c25ab3618323d74a436e7e0b66a512ca6dbab27f08a6"}
// 	invoke, err := node.AddInvokeTransaction(context.Background(), contractAddress, methodName, callData)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("%v\n", invoke)
// }
