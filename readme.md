# tx-Parser

Made with :blue_heart: by rnov.

tx-Parser is an implementation of the
following [task](https://trustwallet.notion.site/Backend-Homework-Tx-Parser-abd431fca950427db75d73d90a0244a8).
The implementation and structure of the project follows all the requirements of the task.

### Quick Start

To build and start the service locally:

```sh
make all
```

To start making request to the service:

get current block number

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

Alternatively you can use the following script to make requests:

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
- `POST /subscribe`: Subscribes an address to be monitored.
- `GET /transactions`: Returns the transactions for a subscribed address.

Internally when started prior initialising and exposing the http server, the service initialises in parallel a
monitoring process that check for new blocks and parses the transactions for the subscribed addresses, updating the
storage accordingly.

When data is requested from the exposed API, the storage is queried and the data is returned (exception for block, check
notes).

### Design

The project's design largely adheres to
the [golang-standards/project-layout](https://github.com/golang-standards/project-layout):

- `cmd`: Contains the main applications for the project.
- `internal`: Houses all the logic intended for internal use. Notably:
    - `http/node`: Implements client that's used to make call to ethereum nodes(JSONRPC).
    - `storage`: Defines the storage interface and its implementations.
    - `parser`: Contains the logic for monitoring and parsing the transactions based on the subscribed addresses.
- `pkg`: Houses all the logic that could be imported by other projects. Notably:
    - `data`: Contains the data structures of the block and transaction response.
    - `config`: Shared logic and structures for the configuration of the project.

That being said, the project is structured in a way that it could be easily extended to support the other
functionalities
of the task as described: `Expose public interface for external usage either via code or command line or rest api that
will include supported list of operations defined in the Parser interface`. It allows to be extended adding support for
cli or code.

### Notes

- Followed all the points in `limitations` described in the task, the only external packages are for the mux and config loading.
- Followed the advice regarding time, simplicity and guidance (not production ready), therefore things like graceful shutdown,
  logging, metrics, etc. are not implemented or greatly simplified.
- Left some comments in the code to explain some decisions and possible improvements.

