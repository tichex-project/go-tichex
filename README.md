# Tichex Blockchain

This repository contains the source code of the Tichex Blockchain

**WARNING**: The repository is in development.

**Note**: Requires [Go 1.12.4+](https://golang.org/dl/)

# Install Tichex Blockchain

There are several ways you can install Tichex Blockchain Testnet node on your machine.

## Using Binary
1. **Download Tichex**
Get [latest binary](https://github.com/tichex-project/go-tichex/releases) build suitable for your architecture and unpack it to desired folder.

2. **Run Tichex**
	```bash
	./tichexd start
	```
## From Source
1. **Install Go** by following the [official docs](https://golang.org/doc/install). Remember to set your `$GOPATH`, `$GOBIN`, and `$PATH` environment variables, for example:
	```bash
	mkdir -p $HOME/go/bin
	echo  "export GOPATH=$HOME/go" >> ~/.bash_profile
	echo  "export GOBIN=\$GOPATH/bin" >> ~/.bash_profile
	echo  "export PATH=\$PATH:\$GOBIN" >> ~/.bash_profile
	echo  "export GO111MODULE=on" >> ~/.bash_profile
	source ~/.bash_profile
	```
2. **Clone Tichex source code to your machine**
	```bash
	mkdir -p $GOPATH/src/github.com/tichex-project
	cd $GOPATH/src/github.com/tichex-project
	git clone https://github.com/tichex-project/go-tichex.git
	cd go-tickex
	```
  3. **Compile**
		```bash
		# Install the app into your $GOBIN
		make install
		# Now you should be able to run the following commands:
		tichexd help
		tichexcli help
		```
		The latest `go-tichex version` is now installed.
3. **Run Minter**
	```bash
	tichexd start
	```

## Install on Digital Ocean
1. **Clone repository**
    ```bash
	git clone https://github.com/tichex-project/go-tichex.git
    chmod +x go-tichex/scripts/install/install_ubuntu.sh
	```
2. **Run the script**
    ```bash
    go-tichex/scripts/install/install_ubuntu.sh
    source ~/.profile
	```
3. Now you should be able to run the following commands:
	```bash
	tichexd help
	tichexcli help
	```
    The latest `go-tichex version` is now installed.

## Running the test network and using the commands

To initialize configuration and a `genesis.json` file for your application and an account for the transactions, start by running:

>  _*NOTE*_: In the below commands addresses are are pulled using terminal utilities. You can also just input the raw strings saved from creating keys, shown below. The commands require [`jq`](https://stedolan.github.io/jq/download/) to be installed on your machine.

>  _*NOTE*_: If you have run the tutorial before, you can start from scratch with a `tichexd unsafe-reset-all` or by deleting both of the home folders `rm -rf ~/.tichex*`

>  _*NOTE*_: If you have the Cosmos app for ledger and you want to use it, when you create the key with `tichexcli keys add jack` just add `--ledger` at the end. That's all you need. When you sign, `jack` will be recognized as a Ledger key and will require a device.

```bash
# Initialize configuration files and genesis file
tichexd init --chain-id tichex-test-network-1

# Copy the `Address` output here and save it for later use
# [optional] add "--ledger" at the end to use a Ledger Nano S
tichexcli keys add jack

# Copy the `Address` output here and save it for later use
tichexcli keys add alice

# Add both accounts, with coins to the genesis file
tichexd add-genesis-account $(tichexcli keys show jack -a) 1000theur,1000thx
tichexd add-genesis-account $(tichexcli keys show alice -a) 1000theur,1000thx

# Configure your CLI to eliminate need for chain-id flag
tichexcli config chain-id tichex-test-network-1
tichexcli config output json
tichexcli config indent true
tichexcli config trust-node true
```

You can now start `tichexd` by calling `tichexd start`. You will see logs begin streaming that represent blocks being produced, this will take a couple of seconds.

Open another terminal to run commands against the network you have just created:

```bash
# First check the accounts to ensure they have funds
tichexcli query account $(tichexcli keys show jack -a)
tichexcli query account $(tichexcli keys show alice -a)
```

# Transactions
You can now start the first transaction

```bash
tichexcli tx send --from=$(tichexcli keys show jack -a)  $(tichexcli keys show alice -a) 10theur
```

# Query
Query an account

```bash
tichexcli query account $(tichexcli keys show jack -a)
```