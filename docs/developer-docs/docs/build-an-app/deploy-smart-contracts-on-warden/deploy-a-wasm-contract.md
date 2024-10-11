﻿---
sidebar_position: 2
---

# Deploy a WASM contract

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

## Overview

The [`x/wasm`](/learn/warden-protocol-modules/external-modules#xwasm) Warden module allows executing WebAssembly smart contracts developed with [CosmWasm](https://cosmwasm.com) and **Rust**.

This guide explains how to create and deploy a simple "Hello World" WASM contract on the Warden chain. Since it's intended for testing purposes, you'll be running a local chain.

## Prerequisites

Before you start, complete the following prerequisites:

- Install Rust by running the following:

  ```bash
  curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
  ```

- Set up the CosmWasm development environment:

  - [CosmWasm](https://book.cosmwasm.com/setting-up-env.html): The CosmWasm binary and its dependencies.

  - [cargo-generate](https://cargo-generate.github.io/cargo-generate/installation.html): A tool to help you get up and running quickly with a new Rust project by leveraging a pre-existing git repository as a template.

  - [wasm-opt](https://docs.rs/wasm-opt/latest/wasm_opt/index.html): A tool for optimizing the compiled WebAssembly (Wasm) code.
   
   To install these tools, run the following commands:

   ```bash
   rustup target add wasm32-unknown-unknown
   cargo install cargo-generate --features vendored-openssl
   brew install binaryen
   ```

- [Run a local chain](/operate-a-node/run-a-local-chain) and make sure you have `wardend` correctly installed.

  The next steps require your local account name, or key name. You can check the list of available keys by executing this command:

  ```bash
  wardend keys list
  ```
  :::tip
  If you used our `just` script to run the node with default settings, the local account name is `shulgin`.
  :::

## 1. Create a CosmWasm project

Create a new CosmWasm project by running the following:

```bash
cargo generate --git https://github.com/CosmWasm/cw-template.git --name hello-world
cd hello-world
```

## 2. Modify the contract code

1. Open `src/contract.rs` and replace its contents with this code:

   ```rust
   use cosmwasm_std::{
       entry_point, to_binary, Binary, Deps, DepsMut, Env, MessageInfo, Response, StdResult,
   };
   use cw2::set_contract_version;
   
   use crate::error::ContractError;
   use crate::msg::{ExecuteMsg, InstantiateMsg, QueryMsg};
   
   const CONTRACT_NAME: &str = "crates.io:hello-world";
   const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");
   
   #[entry_point]
   pub fn instantiate(
       deps: DepsMut,
       _env: Env,
       info: MessageInfo,
       _msg: InstantiateMsg,
   ) -> Result<Response, ContractError> {
       set_contract_version(deps.storage, CONTRACT_NAME, CONTRACT_VERSION)?;
       Ok(Response::new().add_attribute("method", "instantiate")
           .add_attribute("owner", info.sender))
   }
   
   #[entry_point]
   pub fn execute(
       _deps: DepsMut,
       _env: Env,
       info: MessageInfo,
       msg: ExecuteMsg,
   ) -> Result<Response, ContractError> {
       match msg {
           ExecuteMsg::SayHello {} => Ok(Response::new()
               .add_attribute("method", "say_hello")
               .add_attribute("sender", info.sender)),
       }
   }
   
   #[entry_point]
   pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
       match msg {
           QueryMsg::GetGreeting {} => to_binary(&"Hello, World!"),
       }
   }
   ```

2. Open `src/msg.rs` and replace its contents with this code:
   
   ```rust
   use cosmwasm_schema::{cw_serde, QueryResponses};
   
   #[cw_serde]
   pub struct InstantiateMsg {}
   
   #[cw_serde]
   pub enum ExecuteMsg {
       SayHello {},
   }
   
   #[cw_serde]
   #[derive(QueryResponses)]
   pub enum QueryMsg {
       #[returns(String)]
       GetGreeting {},
   }
   ```

## 3. Compile the contract

To compile the contract, run the following:

```bash
cargo wasm
```

The contract should be compiled without any errors.

## 4. Optimize the code

Now you need to optimize your compiled Wasm code:

```bash
wasm-opt -Os -o target/wasm32-unknown-unknown/release/hello_world.wasm /
target/wasm32-unknown-unknown/release/hello_world.wasm
```

## 5. Run the chain

If your local chain isn't running, start it:
   
```bash
wardend start
```

## 6. Store the contract on-chain

To store your contract on the Warden chain, run the command below. Specify your key name from [Prerequisites](#prerequisites) in the `--from` flag (typically `shulgin`).
   
<Tabs>
<TabItem value="default" label="Default node settings">
```bash
wardend tx wasm store target/wasm32-unknown-unknown/release/hello_world.wasm /
--from shulgin --gas auto --gas-adjustment 1.3 --gas-prices 100000000000award -y
```
</TabItem>
<TabItem value="custom" label="Custom node settings">
```bash
wardend tx wasm store target/wasm32-unknown-unknown/release/hello_world.wasm /
--from my-key-name --gas auto --gas-adjustment 1.3 --gas-prices 100000000000award -y
```
</TabItem>
</Tabs>

The transaction should be successful without any errors.

## 7. Get the code ID

Get the code ID that indentifies your Wasm code:

```bash
wardend query wasm list-code
```

Note down `code_id` from the output.

## 8. Instantiate the contract

You can instantiate the contract by using the command below.

Before you proceed, replace `1` with the actual code ID you retrieved in previous step and specify your key name in the `--from` flag. Also note that you can either define an admin or pass `--no-admin` to make it immutable, like in this example.

<Tabs>
<TabItem value="default" label="Default node settings">
```bash
wardend tx wasm instantiate 1 '{}' /
  --from shulgin --label "Hello World" /
  --gas auto --gas-adjustment 1.3 --gas-prices 100000000000award /
  --no-admin -y 
```
</TabItem>
<TabItem value="custom" label="Custom node settings">
```bash
wardend tx wasm instantiate 1 '{}' /
  --from my-key-name --label "Hello World" /
  --gas auto --gas-adjustment 1.3 --gas-prices 100000000000award /
  --no-admin -y 
```
</TabItem>
</Tabs>


## 9. Get the contract address

To get the contract address, run the following command. Replace `1` with the actual code ID:

```bash
wardend query wasm list-contract-by-code 1
```

Note down the contract address.

## 10. Execute the contract

Use the command below to exectute your contract. Replace `my-contract-address` with your contract address and specify your key name in the `--from` flag.

<Tabs>
<TabItem value="default" label="Default node settings">
```bash
wardend tx wasm execute my-contract-address '{"say_hello":{}}' /
--from shulgin --gas auto --gas-adjustment 1.3 --gas-prices 100000000000award -y
```
</TabItem>
<TabItem value="custom" label="Custom node settings">
```bash
wardend tx wasm execute my-contract-address '{"say_hello":{}}' /
--from my-key-name --gas auto --gas-adjustment 1.3 --gas-prices 100000000000award -y
```
</TabItem>
</Tabs>


## 11. Query the contract

You can query your contract with the following command. Replace `my-contract-address` with your contract address.

```bash
wardend query wasm contract-state smart my-contract-address '{"get_greeting":{}}'
```

In the output, you should see this: `data: Hello, World!`

If you encounter any issues, please reach out to us in [Discord](https://discord.com/invite/warden) or [Twitter](https://twitter.com/wardenprotocol).

Happy coding! 🚀