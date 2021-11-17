// Copyright 2018 Wanchain Foundation Ltd
// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"fmt"
	"math/big"

	"github.com/wanchain/go-wanchain/common"
)

var (
	MainnetGenesisHash = common.HexToHash("0x0376899c001618fc7d5ab4f31cfd7f57ca3a896ccc1581a57d8f129ecf40b840") // Mainnet genesis hash to enforce below configs on
	TestnetGenesisHash = common.HexToHash("0xa37b811609a9d1e898fb49b3901728023e5e72e18e58643d9a7a82db483bfeb0") // Testnet genesis hash to enforce below configs on
	PlutoGenesisHash   = common.HexToHash("0x4ff7e18e5842c540f49d827e67894c39783ef1e0494ea52f569db0bcf63786e6") // Pluto genesis hash to enforce below configs on

	InternalGenesisHash = common.HexToHash("0xb1dc31a86510003c23b9ddee0e194775807262529b8dafa6dc23d9315364d2b3")
)

const MainnetPow2PosUpgradeBlockNumber = 4046000
const TestnetPow2PosUpgradeBlockNumber = 3560000
const InternalPow2PosUpgradeBlockNumber = 10

const MAINNET_CHAIN_ID = 1
const TESTNET_CHAIN_ID = 3

// INTERNAL_CHAIN_ID is Private chain pow -> pos mode chainId
const INTERNAL_CHAIN_ID = 4

// PLUTO_CHAIN_ID is Private chain pos mode chainId
const PLUTO_CHAIN_ID = 6

// PLUTODEV_CHAIN_ID is Private chain pos mode single node chainId
const PLUTODEV_CHAIN_ID = 6

// JUPITER_MAINNET_CHAIN_ID is mainnet chainId after jupiter fork
const JUPITER_MAINNET_CHAIN_ID = 888

// JUPITER_TESTNET_CHAIN_ID is testnet chainId after jupiter fork
const JUPITER_TESTNET_CHAIN_ID = 999
const JUPITER_INTERNAL_CHAIN_ID = 777
//const JUPITER_PLUTO_CHAIN_ID = 6
const JUPITER_PLUTO_CHAIN_ID = 666
//const JUPITER_PLUTODEV_CHAIN_ID = 6
const JUPITER_PLUTODEV_CHAIN_ID = 555

// NOT_JUPITER_CHAIN_ID is used for compare
const NOT_JUPITER_CHAIN_ID = 0xffffffff

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainId: big.NewInt(MAINNET_CHAIN_ID),

		//HomesteadBlock: big.NewInt(1150000),
		//DAOForkBlock:   big.NewInt(1920000),
		//DAOForkSupport: true,
		//EIP150Block:    big.NewInt(2463000),
		//EIP150Hash:     common.HexToHash("0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0"),
		//EIP155Block:    big.NewInt(2675000),
		//EIP158Block:    big.NewInt(2675000),

		//ByzantiumBlock: big.NewInt(4370000),
		ByzantiumBlock: big.NewInt(0),
		Ethash:         new(EthashConfig),
		PosFirstBlock:  big.NewInt(MainnetPow2PosUpgradeBlockNumber), // set as n * epoch_length
		IsPosActive:    false,
		Pluto: &PlutoConfig{
			Period: 10,
			Epoch:  100,
		},
	}

	WanchainChainConfig = &ChainConfig{
		ChainId: big.NewInt(MAINNET_CHAIN_ID),
		//HomesteadBlock: big.NewInt(0),
		//DAOForkBlock:   nil,
		//DAOForkSupport: true,
		//EIP150Block:    big.NewInt(0),
		//EIP150Hash:     common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d"),
		//EIP155Block:    big.NewInt(0),
		//EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(0),
		Ethash:         new(EthashConfig),
		PosFirstBlock:  big.NewInt(MainnetPow2PosUpgradeBlockNumber), // set as n * epoch_length
		IsPosActive:    false,                                        // when pos running, the state changed to true by program
		Pluto: &PlutoConfig{
			Period: 10,
			Epoch:  100,
		},
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainId: big.NewInt(TESTNET_CHAIN_ID),
		//HomesteadBlock: big.NewInt(0),
		//DAOForkBlock:   nil,
		//DAOForkSupport: true,
		//EIP150Block:    big.NewInt(0),
		//EIP150Hash:     common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d"),
		//EIP155Block:    big.NewInt(10),
		//EIP158Block:    big.NewInt(10),
		ByzantiumBlock: big.NewInt(0),

		Ethash:        new(EthashConfig),
		PosFirstBlock: big.NewInt(TestnetPow2PosUpgradeBlockNumber), // set as n * epoch_length
		IsPosActive:   false,
		Pluto: &PlutoConfig{
			Period: 10,
			Epoch:  100,
		},
	}

	// RinkebyChainConfig contains the chain parameters to run a node on the Rinkeby test network.
	InternalChainConfig = &ChainConfig{
		ChainId: big.NewInt(INTERNAL_CHAIN_ID),
		//HomesteadBlock: big.NewInt(1),
		//DAOForkBlock:   nil,
		//DAOForkSupport: true,
		//EIP150Block:    big.NewInt(2),
		//EIP150Hash:     common.HexToHash("0x9b095b36c15eaf13044373aef8ee0bd3a382a5abb92e402afa44b8249c3a90e9"),
		//EIP155Block:    big.NewInt(3),
		//EIP158Block:    big.NewInt(3),
		ByzantiumBlock: big.NewInt(0),

		Ethash:        new(EthashConfig),
		PosFirstBlock: big.NewInt(InternalPow2PosUpgradeBlockNumber), // set as n * epoch_length
		IsPosActive:   false,
		Pluto: &PlutoConfig{
			Period: 10,
			Epoch:  100,
		},
	}
	// PlutoChainConfig contains the chain parameters to run a node on the Pluto test network.
	PlutoChainConfig = &ChainConfig{
		ChainId: big.NewInt(PLUTO_CHAIN_ID),
		//HomesteadBlock: big.NewInt(0),
		//DAOForkBlock:   nil,
		//DAOForkSupport: true,
		//EIP150Block:    big.NewInt(0),
		//EIP150Hash:     common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		//EIP155Block:    big.NewInt(3),
		//EIP158Block:    big.NewInt(3),
		ByzantiumBlock: big.NewInt(0),
		PosFirstBlock:  big.NewInt(1),
		IsPosActive:    true,

		Pluto: &PlutoConfig{
			Period: 10,
			Epoch:  100,
		},
	}

	// AllProtocolChanges contains every protocol change (EIPs)
	// introduced and accepted by the Ethereum core developers.
	//
	// This configuration is intentionally not using keyed fields.
	// This configuration must *always* have all forks enabled, which
	// means that all fields must be set at all times. This forces
	// anyone adding flags to the config to also have to set these
	// fields.
	AllProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), big.NewInt(100), false, new(EthashConfig), nil, nil}

	TestChainConfig = &ChainConfig{
		ChainId:        big.NewInt(MAINNET_CHAIN_ID),
		ByzantiumBlock: big.NewInt(0),
		Ethash:         new(EthashConfig),
		PosFirstBlock:  big.NewInt(TestnetPow2PosUpgradeBlockNumber), // set as n * epoch_length
		IsPosActive:    false,
	}

	TestRules = TestChainConfig.Rules(new(big.Int))

	noStaking = false
)

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainId *big.Int `json:"chainId"` // Chain id identifies the current chain and is used for replay protection

	//HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	//DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	//DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	//EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	//EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	//EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	//EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock *big.Int `json:"byzantiumBlock,omitempty"` // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	PosFirstBlock  *big.Int `json:"posFirstBlock,omitempty"`
	IsPosActive    bool     `json:"isPosActive,omitempty"`

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
	Pluto  *PlutoConfig  `json:"pluto,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// PlutoConfig is the consensus engine configs for proof-of-authority based sealing.
type PlutoConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *PlutoConfig) String() string {
	return "pluto"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	case c.Pluto != nil:
		engine = c.Pluto
	default:
		engine = "unknown"
	}
	//return fmt.Sprintf("{ChainID: %v Homestead: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Engine: %v}",
	return fmt.Sprintf("{ChainID: %v Byzantium: %v Engine: %v}",
		c.ChainId,
		//c.HomesteadBlock,
		//c.DAOForkBlock,
		//c.DAOForkSupport,
		//c.EIP150Block,
		//c.EIP155Block,
		//c.EIP158Block,

		c.ByzantiumBlock,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
//func (c *ChainConfig) IsHomestead(num *big.Int) bool {
//	return isForked(c.HomesteadBlock, num)
//}

// IsDAO returns whether num is either equal to the DAO fork block or greater.
//func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
//	return isForked(c.DAOForkBlock, num)
//}

//func (c *ChainConfig) IsEIP150(num *big.Int) bool {
//	return isForked(c.EIP150Block, num)
//}

//func (c *ChainConfig) IsEIP155(num *big.Int) bool {
//	return isForked(c.EIP155Block, num)
//}

//func (c *ChainConfig) IsEIP158(num *big.Int) bool {
//	return isForked(c.EIP158Block, num)
//}

//func (c *ChainConfig) IsByzantium(num *big.Int) bool {
//	return isForked(c.ByzantiumBlock, num)
//}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	//if num == nil {
	//	return GasTableHomestead
	//}
	//switch {

	//case c.IsEIP158(num):
	//	return GasTableEIP158
	/*
		case c.IsEIP150(num):
			return GasTableEIP150
	*/
	//default:
	//	return GasTableHomestead
	//}

	return GasTableEIP158
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)
	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {

	//if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
	//	return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	//}

	//if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
	//	return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	//}

	//if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
	//	return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	//}

	//if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
	//	return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	//}

	//if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
	//	return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	//}
	//
	//if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
	//	return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	//}

	//if c.IsEIP158(head) && !configNumEqual(c.ChainId, newcfg.ChainId) {
	//	return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	//}

	//if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
	//	return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	//}

	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntatic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainId *big.Int
	//IsHomestead, IsEIP150, IsEIP155, IsEIP158 bool
	//IsByzantium                               bool
}

func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainId := c.ChainId
	if chainId == nil {
		chainId = new(big.Int)
	}
	//return Rules{ChainId: new(big.Int).Set(chainId), IsHomestead: /*c.IsHomestead(num)*/false, IsEIP150: false/*c.IsEIP150(num)*/, IsEIP155: false/*c.IsEIP155(num)*/, IsEIP158:false/* c.IsEIP158(num)*/, IsByzantium: c.IsByzantium(num)}

	return Rules{ChainId: new(big.Int).Set(chainId)}
}

func (c *ChainConfig) SetPosActive() {
	c.IsPosActive = true
	SetPosActive(c.IsPosActive)
}

func (c *ChainConfig) IsPosBlockNumber(n *big.Int) bool {
	return n.Cmp(c.PosFirstBlock) >= 0
}

var (
	isPosActive    = false
	TestnetChainId = TestnetChainConfig.ChainId.Int64()
	MainnetChainId = WanchainChainConfig.ChainId.Int64()
)

func IsPosActive() bool {
	return isPosActive
}

func SetPosActive(active bool) {
	isPosActive = active
}

func IsNoStaking() bool {
	return noStaking
}
func SetNoStaking() {
	noStaking = true
}

func JupiterChainId(chainId uint64) uint64 {
	if chainId == MAINNET_CHAIN_ID {
		return JUPITER_MAINNET_CHAIN_ID
	}

	if chainId == TESTNET_CHAIN_ID {
		return JUPITER_TESTNET_CHAIN_ID
	}

	if chainId == INTERNAL_CHAIN_ID {
		return JUPITER_INTERNAL_CHAIN_ID
	}

	if chainId == PLUTO_CHAIN_ID {
		return JUPITER_PLUTO_CHAIN_ID
	}

	if chainId == PLUTODEV_CHAIN_ID {
		return JUPITER_PLUTODEV_CHAIN_ID
	}

	return NOT_JUPITER_CHAIN_ID
}

func IsOldChainId(chainId uint64) bool {
	if chainId == MAINNET_CHAIN_ID {
		return true
	}

	if chainId == TESTNET_CHAIN_ID {
		return true
	}

	if chainId == INTERNAL_CHAIN_ID {
		return true
	}

	if chainId == PLUTO_CHAIN_ID {
		return true
	}

	if chainId == PLUTODEV_CHAIN_ID {
		return true
	}

	return false
}
