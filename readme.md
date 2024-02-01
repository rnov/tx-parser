# tx-Parser

Made with :blue_heart: by rnov.

tx-Parser parses ethereum transactions every block for a given address and stores them in memory, it exposes a REST API to query the data.

### Quick Start

To build and start the service locally:
```sh
make all
```

To stop the service:
```sh
make stop
```

To remove bin directory:
```sh
make clean
```

To start making request to the service:

Get current block number
```sh
make get-block
```

Subscribe address to be monitored.
```sh
make subscribe address="0xdac17f958d2ee523a2206206994597c13d831ec7"
```

Get transactions for a subscribed address (empty if no transactions or address not subscribed)
```sh
make get-txs address="0xdac17f958d2ee523a2206206994597c13d831ec7" 
```

Alternatively you can use the following script to make requests (add jq for pretty printing):
```sh
# get current block number
./tools/shell/requests.sh block | jq

# subscribe address to be monitored.
./tools/shell/requests.sh subscribe 0xdac17f958d2ee523a2206206994597c13d831ec7 | jq

# get transactions for a subscribed address (empty if no transactions or address not subscribed)
./tools/shell/requests.sh transactions 0xdac17f958d2ee523a2206206994597c13d831ec7 | jq
```

### Description

The project initialises a http server that exposes the following endpoints:

- `GET /block`: Returns the current block number.
- `PUT /subscribe`: Subscribes an address to be monitored.
- `GET /transactions`: Returns the transactions for a subscribed address.

Prior to initialize the REST server and expose the endpoints, the service starts in parallel a
monitoring process that checks for new blocks and parses the transactions for the subscribed addresses, updating the
storage accordingly.

When data is requested from the exposed API, the storage is queried and the data is returned (exception for getBlock, check
notes).

### Design

The project's design largely adheres to
the [golang-standards/project-layout](https://github.com/golang-standards/project-layout):

- `cmd`: Contains the main applications for the project.
- `config`: Holds the config file/s.
- `internal`: Houses all the logic intended for internal use. Notably:
    - `http/node`: Implements client that's used to make call to ethereum nodes(JSONRPC).
    - `storage`: Defines the storage interface and its implementations.
    - `parser`: Contains the logic for monitoring and parsing the transactions based on the subscribed addresses.
- `pkg`: Houses all the logic that could be imported by other projects. Notably:
    - `data`: Contains the data structures of the block and transaction response.
    - `config`: Shared logic and structures for the configuration of the project.

### Notes

- `/subscribe` is a PUT request because it's idempotent and can return `200` regardless of the address being already
    subscribed or not, that way we avoid complex logic of returning different status codes that would require a POST request.
    Not a fan of this approach, but it's a tradeoff for simplicity.
- Since the addresses are not always consistent lower/upper case depend on where we got them with the response from the node,
  they are stored in lower case, although there is better approaches to handle this (e.g: checksum), for the sake of simplicity.
- The address used in the example `0xdac17f958d2ee523a2206206994597c13d831ec7` belongs to USDT (Tether) and has plenty of txs each block
  another address that can be used is `0x28C6c06298d514Db089934071355E5743bf21d60` it belongs to binance however it has less txs per block,
  making easier to read the responses.


