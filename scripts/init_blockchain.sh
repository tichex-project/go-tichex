#!/bin/bash

rm -rf ~/.tichex*

tichexd init --chain-id tichex-test-network-1

tichexcli keys add jack
tichexcli keys add alice

tichexd add-genesis-account $(tichexcli keys show jack -a) 1000theur,1000thx
tichexd add-genesis-account $(tichexcli keys show alice -a) 1000theur,1000thx

tichexcli config chain-id tichex-test-network-1
tichexcli config output json
tichexcli config indent true
tichexcli config trust-node true