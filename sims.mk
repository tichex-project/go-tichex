#!/usr/bin/make -f

########################################
### Simulations

SIMAPP = github.com/tichex-project/go-tichex/app

sim-tichex-nondeterminism:
	@echo "Running nondeterminism test..."
	@go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminism -SimulationEnabled=true -v -timeout 10m

sim-tichex-custom-genesis-fast:
	@echo "Running custom genesis simulation..."
	@echo "By default, ${HOME}/.tichexd/config/genesis.json will be used."
	@go test -mod=readonly github.com/cosmos/tichex/app -run TestFullTichexSimulation -SimulationGenesis=${HOME}/.tichexd/config/genesis.json \
		-SimulationEnabled=true -SimulationNumBlocks=100 -SimulationBlockSize=200 -SimulationCommit=true -SimulationSeed=99 -SimulationPeriod=5 -v -timeout 24h

sim-tichex-fast:
	@echo "Running quick Tichex simulation. This may take several minutes..."
	@go test -mod=readonly github.com/tichex-project/go-tichex/app -run TestFullTichexSimulation -SimulationEnabled=true -SimulationNumBlocks=100 -SimulationBlockSize=200 -SimulationCommit=true -SimulationSeed=99 -SimulationPeriod=5 -v -timeout 24h

sim-tichex-import-export: runsim
	@echo "Running Tichex import/export simulation. This may take several minutes..."
	$(GOPATH)/bin/runsim 25 5 TestTichexImportExport

sim-tichex-simulation-after-import: runsim
	@echo "Running Tichex simulation-after-import. This may take several minutes..."
	$(GOPATH)/bin/runsim 25 5 TestTichexSimulationAfterImport

sim-tichex-custom-genesis-multi-seed: runsim
	@echo "Running multi-seed custom genesis simulation..."
	@echo "By default, ${HOME}/.tichexd/config/genesis.json will be used."
	$(GOPATH)/bin/runsim -g ${HOME}/.tichexd/config/genesis.json 400 5 TestFullTichexSimulation

sim-tichex-multi-seed: runsim
	@echo "Running multi-seed Tichex simulation. This may take awhile!"
	$(GOPATH)/bin/runsim 400 5 TestFullTichexSimulation

sim-benchmark-invariants:
	@echo "Running simulation invariant benchmarks..."
	@go test -mod=readonly github.com/tichex-project/go-tichex/app -benchmem -bench=BenchmarkInvariants -run=^$ \
	-SimulationEnabled=true -SimulationNumBlocks=1000 -SimulationBlockSize=200 \
	-SimulationCommit=true -SimulationSeed=57 -v -timeout 24h

SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 200
SIM_COMMIT ?= true
sim-tichex-benchmark:
	@echo "Running Tichex benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ github.com/tichex-project/go-tichex/app -bench ^BenchmarkFullTichexSimulation$$  \
		-SimulationEnabled=true -SimulationNumBlocks=$(SIM_NUM_BLOCKS) -SimulationBlockSize=$(SIM_BLOCK_SIZE) -SimulationCommit=$(SIM_COMMIT) -timeout 24h

sim-tichex-profile:
	@echo "Running Tichex benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ github.com/ctichex-project/go-tichex/app -bench ^BenchmarkFullTichexSimulation$$ \
		-SimulationEnabled=true -SimulationNumBlocks=$(SIM_NUM_BLOCKS) -SimulationBlockSize=$(SIM_BLOCK_SIZE) -SimulationCommit=$(SIM_COMMIT) -timeout 24h -cpuprofile cpu.out -memprofile mem.out


.PHONY: runsim sim-tichex-nondeterminism sim-tichex-custom-genesis-fast sim-tichex-fast sim-tichex-import-export \
	sim-tichex-simulation-after-import sim-tichex-custom-genesis-multi-seed sim-tichex-multi-seed \
	sim-benchmark-invariants sim-tichex-benchmark sim-tichex-profile
