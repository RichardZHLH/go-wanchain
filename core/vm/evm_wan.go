// Copyright 2014 The go-ethereum Authors
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

package vm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// add by jacob
func (evm *EVM) Time() *big.Int {
	return evm.Context.Time
}

func (evm *EVM) BlockNumber() *big.Int {
	return evm.Context.BlockNumber
}

func IsWanchainPrecompiled(addr common.Address, contract *Contract, evm *EVM) (PrecompiledContract, bool) { // TODO delete it????
	switch addr {
	case wanCoinPrecompileAddr:
		return &wanCoinSC{contract, evm}, true
	case wanStampPrecompileAddr:
		return &wanchainStampSC{contract, evm}, true
	case WanCscPrecompileAddr:
		return &PosStaking{contract, evm}, true
	case PosControlPrecompileAddr:
		return &PosControl{contract, evm}, true
	case slotLeaderPrecompileAddr:
		return &slotLeaderSC{contract, evm}, true
	case randomBeaconPrecompileAddr:
		return &RandomBeaconContract{contract, evm}, true
	case SolEnhancePrecompileAddr:
		return &SolEnhance{contract, evm}, true
	case s256AddPrecompileAddr:
		return &s256Add{contract, evm}, true
	case s256ScalarMulPrecompileAddr:
		return &s256ScalarMul{contract, evm}, true
	case sha3fipsPrecompileAddr:
		if evm.chainRules.IsLondon {
			return &sha3fips{contract, evm}, true
		} else {
			return nil, false
		}
	case ecrecoverPublicKeyPrecompileAddr:
		if evm.chainRules.IsLondon {
			return &ecrecoverPublicKey{contract, evm}, true
		} else {
			return nil, false
		}
	default:
		return nil, false
	}
}
func (evm *EVM) precompile(addr common.Address, caller ContractRef, value *big.Int, gas uint64) (PrecompiledContract, bool) {
	p, ok := evm.precompileEth(addr)
	if !ok {
		contract := NewContract(caller, AccountRef(addr), value, gas)
		return IsWanchainPrecompiled(addr, contract, evm)
	}
	return p, ok
}
