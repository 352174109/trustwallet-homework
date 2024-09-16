# Service Command Line Interface (CLI) Usage Guide

This document provides instructions on how to use the CLI commands for interacting with the `Service`. The `Service` provides functionalities to monitor blockchain blocks and transactions, as well as subscribe to specific addresses for transaction updates.

## Prerequisites

Before using this tool, ensure that you have the following dependencies:

- Go installed on your system.
- Properly set up blockchain parser and its associated context within the `Service`.

## Building from source
```shell
$ mkdir -p $GOPATH/src/github.com/352174109
$ cd $GOPATH/src/github.com/352174109
$ git clone https://github.com/352174109/trustwallet-homework.git
$ cd trustwallet-homework
$ make build
$ bin/trustwallet-homework 
```

## Available Commands

### 1. `getCurrentBlock`

This command retrieves the current block number from the blockchain parser.

**Usage:**

```bash
> getCurrentBlock
```
Example Output
```plaintext
Current Block: 15045234
```

### 2. `subscribe <address>`
This command subscribes to a specific blockchain address to monitor its activity.

**Usage:**

```bash
> subscribe <address>
```
* `<address>`: The blockchain address you want to subscribe to.

Example:
```shell
> subscribe 0x123456789abcdef
```
Example Output
```plaintext
Subscribed to address: 0x123456789abcdef
```
If the subscription fails, you will see the following output:
```shell
Failed to subscribe to address: 0x123456789abcdef
```

### 3. `getTransactions <address>`
This command retrieves all transactions related to a specific blockchain address.

**Usage:**

```bash
> getTransactions <address>
```

*  `<address>`: The blockchain address you want to retrieve transactions.

Example:
```shell
> getTransactions 0x123456789abcdef
```
Example Output
```yaml
Transactions:
- {TransactionID: abc123, Amount: 100, From: 0x123456789abcdef, To: 0x987654321abcdef, ...}
- {TransactionID: def456, Amount: 50, From: 0x123456789abcdef, To: 0xabcdef123456789, ...}
```
If no transactions are found, the output will be:
```css
No transactions found for address: 0x123456789abcdef
```

### 4. `help`
This command prints a list of available commands along with their usage.

**Usage:**

```bash
> help
```

*  `<address>`: The blockchain address you want to retrieve transactions.

Example Output
```
Usage:
  getCurrentBlock            - Subscribed the latest block number
  subscribe <address>        - Subscribe to monitor a specific address
  getTransactions <address>  - Subscribed transactions related to a specific address
  help                       - Show available commands and usage

```

## Error Handling
* If no command is provided, the system will print: No command provided.
* For unknown commands, the system will display: Unknown command: '<command>' and suggest using the help command.

## Mocked Cases [Already load into the service]
### 1. `getCurrentBlock`
```shell
getCurrentBlock 
```
Output
```shell
2024/09/17 01:26:14 INFO: Received command: getCurrentBlock
2024/09/17 01:26:14 INFO: Current Block: 20764652
```

### 2. `subscribe`
```shell
subscribe 0x00000000009e50a7ddb7a7b0e2ee6604fd120e49
``` 
Output
```
2024/09/17 01:29:28 INFO: Received command: subscribe 0x00000000009e50a7ddb7a7b0e2ee6604fd120e49
2024/09/17 01:29:28 INFO: Address [0x00000000009e50a7ddb7a7b0e2ee6604fd120e49] subscribed successful
2024/09/17 01:29:28 INFO: Subscribed to address: 0x00000000009e50a7ddb7a7b0e2ee6604fd120e49
```


### 3. `getTransactions`
```shell
getTransactions 0x00000000009e50a7ddb7a7b0e2ee6604fd120e49
``` 
Output
```
2024/09/17 01:29:55 INFO: Received command: getTransactions 0x00000000009e50a7ddb7a7b0e2ee6604fd120e49
2024/09/17 01:29:55 INFO: TransactionByAddr addr [0x00000000009e50a7ddb7a7b0e2ee6604fd120e49]
2024/09/17 01:29:55 INFO: Transactions:
2024/09/17 01:29:55 INFO: - {"chainId":"0x1","blockNumber":"0x13cd296","hash":"0x1b0d6db9bee0b358beda0da81de82cec9ebba5b4488f460943b979fe2315f3c3","nonce":"0x2c08b","from":"0xe75ed6f453c602bd696ce27af11565edc9b46b0d","to":"0x00000000009e50a7ddb7a7b0e2ee6604fd120e49","value":"0xf5232269","gas":"0x46b47","gasPrice":"0x1a5c9d8f9","input":"0x960d1f9afe7a4e6c6aa2f928b71a512b2e6644d7a7e5593d148b89b41a0889322bba387c825180ebfb62bd8e6969ebe5b5e52d02aa1efb3c159d81db1c006d"}
2024/09/17 01:29:55 INFO: - {"chainId":"0x1","blockNumber":"0x13cd7f3","hash":"0x49e72b9cb343d22a1b1e1b7376933e51c6ba69170386649b84b367eb06221316","nonce":"0x2c108","from":"0xe75ed6f453c602bd696ce27af11565edc9b46b0d","to":"0x00000000009e50a7ddb7a7b0e2ee6604fd120e49","value":"0x17b6e08","gas":"0xb226f","gasPrice":"0x1673fe572","input":"0xf30d1ff5e875b9f457f2dd8112bd68999eb72befb17b033c430e9ab61b3cbd3106a0a076bedae847652f42ef07fd58589e001fc02aaa39b223fe8d0a0e5c4f27ead9083c756cc205017b000000000000017e3689e18d1d13b986fd52697f16be888bfad2c5bf12cd67ce834b0002000000000000000006342a9e577618b81865f86c4f39225572f4032756330c58fb71f0393c00c82cfc00000000000000577618b80000002406f603058769e699f6fb2bd8df8d49a39447b706a1b054e8ec623d0791399219006d2a9e27b96f1918cea525a05aa552ca870a8a7e1b13fb2e69caf93701cc2cfc0000000000000027b96f19000000243fa300440000000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2a9059cbb0000000000000000000000008a326ab6ba2f19db9a17b13d473c974b04ff7b7f00000000000000000000000000000000000000000000000000243cd8900000003fa3004400000000000000000000000000008a326ab6ba2f19db9a17b13d473c974b04ff7b7fbd6015b4000000000000000000000000cea525a05aa552ca870a8a7e1b13fb2e69caf9370000000000000000000000000000000000000000000000000000000000000000006d21e6477132d9287b4b3e3097eac69c5d1042983bef7022d658bd362cde0000000000000194a77200000000243cbd3106a0a076bedae847652f42ef07fd58589e001fc02aaa39b223fe8d0a0e5c4f27ead9083c756cc22a017b0000000000000196fbbf56844d1eb986fd52697f16be888bfad2c5bf12cd67ce834b0002000000000000000006342d8024006d"}
2024/09/17 01:29:55 INFO: - {"chainId":"0x1","blockNumber":"0x13cd7ff","hash":"0xe10c64f840dadc59ce8c8a2098b11b3a41fb7ec72021b96b8f099bdfc4e0c196","nonce":"0x1fc54","from":"0xfc9928f6590d853752824b0b403a6ae36785e535","to":"0x00000000009e50a7ddb7a7b0e2ee6604fd120e49","value":"0xac2658ed","gas":"0x3f615","gasPrice":"0x660fc79587d","input":"0xff1c961a5cbe206379252ed32b85cf8d1f53195c6daac758010f3268d66f784b49c2f3acf80e549cde65c81a0a1e12be2f8c2b3c2d8024006d"}
```
