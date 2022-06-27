## Below a list of features to implement

The list below provides a number of things to implement on Caigo:

1. Add support for starknet_estimateFee
2. Add support for starknet_addInvokeTransaction
3. Integrate test in Github Actions
4. Add test and support for devnet on RPC
5. Rely on reflect.DeepCompare to "snapshot" output and track changes
6. Find whitelisted contract to test Deploy and Declare on mainnet
7. Implement and tests the API with the Gateway
8. Create tests to detect when Pathfinder implements new methods and implement nightly tests
9. Create docs to better document the project
10.  provide abigen to create a higher level API from the ABI file (already created)
11. Provide a release note and plan the version to come 

In addition, I would plan 15-min community call every 2 weeks to help with the roadmap and get some feedback.

In addition, I have a few questions that I would like to address:

- How are sequencers implemented and how do they integrate with nodes?
- What is the business model for pathfinder/juno?
- What are good ideas to work on starknet? Would it make sense to develop [checkpoint](https://checkpoint.fyi) as a service?
