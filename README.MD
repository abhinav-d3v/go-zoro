# what is go-zoro
go-zoro is a Go package designed to simplify the process of dynamically creating structs for Ethereum Virtual Machine (EVM) smart contract events and decoding their values. It is particularly useful for developers working with blockchain data, allowing them to efficiently handle and interpret events emitted by smart contracts on the Ethereum network.
## Features
* Dynamic Struct Generation: go-zoro enables the automatic creation of Go structs based on the ABI (Application Binary Interface) of smart contract events. This eliminates the need for manual struct definition, saving time and reducing errors.

* Event Value Decoding: The package provides functionality to decode the values of smart contract events, making it easy to extract and work with the data emitted by smart contracts.

* Ethereum Compatibility: go-zoro is compatible with Ethereum and other EVM-compatible blockchains, providing a seamless experience for developers building applications on these networks.
## Data type conversion

| Solidity  | Golang         |
|-----------|----------------|
| uint256   | *bigInt        |
| address   | common.address |

## Quick start
