# Changelog

## 1.0.0

### Bug Fixes

* (gaiad) [\#4113](https://github.com/cosmos/cosmos-sdk/issues/4113) Fix incorrect `$GOBIN` in `Install Go`
* (gaiacli) [\#3945](https://github.com/cosmos/cosmos-sdk/issues/3945) There's no check for chain-id in TxBuilder.SignStdTx
* (gaiacli) [\#4190](https://github.com/cosmos/cosmos-sdk/issues/4190) Fix redelegations-from by using the correct params and query endpoint.
* (gaiacli) [\#4219](https://github.com/cosmos/cosmos-sdk/issues/4219) Return an error when an empty mnemonic is provided during key recovery.
* (gaiacli) [\#4345](https://github.com/cosmos/cosmos-sdk/issues/4345) Improved Ledger Nano X detection

### Breaking Changes

* (gaiad) [\#3985](https://github.com/cosmos/cosmos-sdk/issues/3985) ValidatorPowerRank uses potential consensus power
* (gaiad) [\#4027](https://github.com/cosmos/cosmos-sdk/issues/4027) gaiad version command does not return the checksum of the go.sum file shipped along with the source release tarball.
  Go modules feature guarantees dependencies reproducibility and as long as binaries are built via the Makefile shipped with the sources, no dependendencies can break such guarantee.
* (gaiad) [\#4159](https://github.com/cosmos/cosmos-sdk/issues/4159) use module pattern and module manager for initialization
* (gaiad) [\#4272](https://github.com/cosmos/cosmos-sdk/issues/4272) Merge gaiareplay functionality into gaiad replay.
  Drop `gaiareplay` in favor of new `gaiad replay` command.
* (gaiacli) [\#3715](https://github.com/cosmos/cosmos-sdk/issues/3715) query distr rewards returns per-validator
  rewards along with rewards total amount.
* (gaiacli) [\#40](https://github.com/cosmos/cosmos-sdk/issues/40) rest-server's --cors option is now gone.
* (gaiacli) [\#4027](https://github.com/cosmos/cosmos-sdk/issues/4027) gaiacli version command dooes not return the checksum of the go.sum file anymore.
* (gaiacli) [\#4142](https://github.com/cosmos/cosmos-sdk/issues/4142) Turn gaiacli tx send's --from into a required argument.
  New shorter syntax: `gaiacli tx send FROM TO AMOUNT`
* (gaiacli) [\#4228](https://github.com/cosmos/cosmos-sdk/issues/4228) Merge gaiakeyutil functionality into gaiacli keys.
  Drop `gaiakeyutil` in favor of new `gaiacli keys parse` command. Syntax and semantic are preserved.
* (rest) [\#3715](https://github.com/cosmos/cosmos-sdk/issues/3715) Update /distribution/delegators/{delegatorAddr}/rewards GET endpoint
  as per new specs. For a given delegation, the endpoint now returns the
  comprehensive list of validator-reward tuples along with the grand total.
* (rest) [\#3942](https://github.com/cosmos/cosmos-sdk/issues/3942) Update pagination data in txs query.
* (rest) [\#4049](https://github.com/cosmos/cosmos-sdk/issues/4049) update tag MsgWithdrawValidatorCommission to match type
* (rest) The `/auth/accounts/{address}` now returns a `height` in the response. The
  account is now nested under `account`.

### Features

* (gaiad) Add `migrate` command to `gaiad` to provide the ability to migrate exported
  genesis state from one version to another.
* (gaiad) Update Gaia for community pool spend proposals per Cosmos Hub governance proposal [\#7](https://github.com/cosmos/cosmos-sdk/issues/7) "Activate the Community Pool"

### Improvements

* (gaiad) [\#4042](https://github.com/cosmos/cosmos-sdk/issues/4042) Update docs and scripts to include the correct `GO111MODULE=on` environment variable.
* (gaiad) [\#4066](https://github.com/cosmos/cosmos-sdk/issues/4066) Fix 'ExportGenesisFile() incorrectly overwrites genesis'
* (gaiad) [\#4064](https://github.com/cosmos/cosmos-sdk/issues/4064) Remove `dep` and `vendor` from `doc` and `version`.
* (gaiad) [\#4080](https://github.com/cosmos/cosmos-sdk/issues/4080) add missing invariants during simulations
* (gaiad) [\#4343](https://github.com/cosmos/cosmos-sdk/issues/4343) Upgrade toolchain to Go 1.12.5.
* (gaiacli) [\#4068](https://github.com/cosmos/cosmos-sdk/issues/4068) Remove redundant account check on `gaiacli`
* (gaiacli) [\#4227](https://github.com/cosmos/cosmos-sdk/issues/4227) Support for Ledger App v1.5
* (rest) [\#2007](https://github.com/cosmos/cosmos-sdk/issues/2007) Return 200 status code on empty results
* (rest) [\#4123](https://github.com/cosmos/cosmos-sdk/issues/4123) Fix typo, url error and outdated command description of doc clients.
* (rest) [\#4129](https://github.com/cosmos/cosmos-sdk/issues/4129) Translate doc clients to chinese.
* (rest) [\#4141](https://github.com/cosmos/cosmos-sdk/issues/4141) Fix /txs/encode endpoint
