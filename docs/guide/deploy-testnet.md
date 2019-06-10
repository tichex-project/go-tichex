# Deploy your own testnet

This document describes 3 ways to setup a network of `tichexd` nodes

1. Single-node, local, manual testnet

## Single-node, local, manual testnet

This guide helps you create a single validator node that runs a network locally for testing and other development related uses.

### Requirements

- [Install Tichex](../../README.md)

### Create genesis file and start the network

```bash
# You can run all of these commands from your home directory
cd $HOME

# Initialize the genesis.json file that will help you to bootstrap the network
tichexd init --chain-id=testing testing

# Create a key to hold your validator account
tichexcli keys add validator

# Add that key into the genesis.app_state.accounts array in the genesis file
# NOTE: this command lets you set the number of coins. Make sure this account has some coins
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
tichexd add-genesis-account $(tichexcli keys show validator -a) 100000000000thx

# Generate the transaction that creates your validator
tichexd gentx --name validator --amount 1000000000uthx

# Add the generated bonding transaction to the genesis file
tichexd collect-gentxs

# Now its safe to start `tichexd`
tichexd start
```

This setup puts all the data for `tichexd` in `~/.tichexd`. You can examine the genesis file you created at `~/.tichexd/config/genesis.json`. With this configuration `tichexcli` is also ready to use and has an account with tokens (both staking and custom).