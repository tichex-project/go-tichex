#!/bin/bash

PASSWORD="12345678"

rm -rf ~/.gaia*

# Initialize the genesis.json file that will help you to bootstrap the network
gaiad init MyValidator --chain-id=tichex-devnet-1

gaiacli config chain-id tichex-devnet-1
gaiacli config output json
gaiacli config indent true
gaiacli config trust-node true

# Change default bond token genesis.json
sed -i 's/stake/utcx/g' ~/.gaiad/config/genesis.json

# Create a key to hold your validator account
echo ${PASSWORD} | gaiacli keys add validator

# Add that key into the genesis.app_state.accounts array in the genesis file
# NOTE: this command lets you set the number of coins. Make sure this account has some coins
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
gaiad add-genesis-account validator 100000000000utcx

# Generate the transaction that creates your validator
echo ${PASSWORD} | gaiad gentx --name validator --amount=90000000000utcx

# Add the generated bonding transaction to the genesis file
gaiad collect-gentxs
gaiad validate-genesis

# Now its safe to start `gaiad`
gaiad start