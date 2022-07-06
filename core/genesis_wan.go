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

package core

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/ppw"
	"github.com/ethereum/go-ethereum/params"
)

//go:generate gencodec -type Genesis -field-override genesisSpecMarshaling -out gen_genesis.go
//go:generate gencodec -type GenesisAccount -field-override genesisAccountMarshaling -out gen_genesis_account.go

func jsonPrealloc(data string) GenesisAlloc {
	var ga GenesisAlloc
	if err := json.Unmarshal([]byte(data), &ga); err != nil {
		panic(err)
	}
	return ga
}

// DefaultGenesisBlock returns the Ethereum main net genesis block.
func DefaultGenesisBlock() *Genesis {
	return &Genesis{
		Config:     params.MainnetChainConfig,
		Nonce:      98,
		ExtraData:  hexutil.MustDecode(ppw.GetMainNetPpwSignStr()),
		GasLimit:   0x2fefd8,
		Difficulty: big.NewInt(1048576),
		//Difficulty: big.NewInt(17179869184),
		Alloc: jsonPrealloc(wanchainAllocJson),
	}
}

// DefaultTestnetGenesisBlock returns the Ropsten network genesis block.
func DefaultTestnetGenesisBlock() *Genesis {
	return &Genesis{
		Config:     params.TestnetChainConfig,
		Nonce:      28, //same with the version
		ExtraData:  hexutil.MustDecode(ppw.GetTestNetPpwSignStr()),
		GasLimit:   0x2fefd8,
		Difficulty: big.NewInt(1048576),
		Alloc:      jsonPrealloc(wanchainTestAllocJson),
	}
}

// DefaultInternalGenesisBlock returns the Rinkeby network genesis block.
func DefaultInternalGenesisBlock() *Genesis {
	return &Genesis{
		Config:     params.InternalChainConfig,
		Nonce:      20,
		ExtraData:  hexutil.MustDecode(ppw.GetInternalNetPpwSignStr()),
		GasLimit:   0x2fefd8,
		Difficulty: big.NewInt(1),
		Alloc:      jsonPrealloc(wanchainTestAllocJson),
	}
}

// DefaultPlutoGenesisBlock returns the Pluto network genesis block.

func DefaultPlutoGenesisBlock() *Genesis {
	return &Genesis{
		Config:     params.PlutoChainConfig,
		Timestamp:  0x59f83144,
		ExtraData:  hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000e8ffc3d0c02c0bfc39b139fa49e2c5475f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
		GasLimit:   0x47b760,
		Difficulty: big.NewInt(1),
		Alloc:      jsonPrealloc(PlutoAllocJson),
	}
}

func DefaultPPOWTestingGenesisBlock() *Genesis {

	key, _ := crypto.HexToECDSA("f1572f76b75b40a7da72d6f2ee7fda3d1189c2d28f0a2f096347055abe344d7f")
	coinbase := crypto.PubkeyToAddress(key.PublicKey)
	return &Genesis{
		Coinbase:   coinbase,
		Config:     params.TestChainConfig,
		Nonce:      66,
		ExtraData:  hexutil.MustDecode("0xf9b32578b4420a36f132db32b56f3831a7cc1804810524175efa012446103d1a04c9f4263a962accdb05642eabc8347ec78e21bdf0d906ba579d423ab5eb9bf02a924367ed9d4f86dfcb1c572cd9a4f80036805b6846f26ac35f2a7d7eda4a2a58f08e8ef073d4e52c506f3f288faa9db1c1e5ae0f1e70f8c38eb01bce9bcb61327532dc5a540da4cf484ae57e98bc5a465c1d2afa6b9376709a525981f53d493a46ef1eb55428b3b88a222d80d23531054ef51dbd100cf8286136659a7d63a38a154e28dbf3e0fd"),
		GasLimit:   0x47b760,
		Difficulty: big.NewInt(1),
		Alloc:      jsonPrealloc(wanchainPPOWTestAllocJson),
	}
}

/*
1 wanfoundation address:1０％,６％,３％

2.wanminer addresses: 10%

3.wansolder addresses:１０％,10%,10%,10%,10%,1%

4.wanteam address:10%,5%,5%
*/
const wanchainAllocJson = `{
		"0x4A2a82864c2Db7091142C2078543FaB0286251a9": {"balance": "21000000000000000000000000"},
		"0x0dA4512bB81fA50891533d9A956238dCC219aBCf": {"balance": "12600000000000000000000000"},
		"0xD209fe3073Ca6B61A1a8A7b895e1F4aD997103E1": {"balance":  "6300000000000000000000000"},

		"0x01d0689001F18637993948e487a15BF3064b16e4": {"balance": "21000000000000000000000000"},

		"0xb3a5c9789A4d882BceF63abBe9B7893aC505bf60": {"balance": "21000000000000000000000000"},
		"0xc57FeeC601d5A473fE9d1D70Af26ac639e0c61a1": {"balance": "21000000000000000000000000"},
		"0xEeCABC0900998aFeE0B52438a6003F2388c78A62": {"balance": "21000000000000000000000000"},
		"0x2dC9A6A04Bc004a8f68f0e886a463AeF23D43030": {"balance": "21000000000000000000000000"},
		"0x5866dD6794B8996E5bC745D508AC6901FF3b0427": {"balance": "21000000000000000000000000"},
		"0x89442477dC39A2503E30D1f8d7FFD4Ea5f87a2aF": {"balance":  "2100000000000000000000000"},

		"0xae8d9B975eC8df8359eA79e50e89b18601816aC3": {"balance": "21000000000000000000000000"},
		"0x53D81A644a0d1081D6C6E8B25f807C2cFb6edE35": {"balance": "10500000000000000000000000"},
		"0x3B9289124f04194F0b3C4F8F862fE1Fbac59c978": {"balance": "10500000000000000000000000"}
}`

//miner reward
//public sale
//Team holding
//Foundation operation
const wanchainTestAllocJson = `{
	  "0x4cb79c7868cd88629df6d4fa8637dda83d13ef27": {"balance": "21000000000000000000000000"},
	  "0xeb71d33d5c7cf05d9177934200c51efa53057c27": {"balance": "107100000000000000000000000"},
      "0x6b4683cafa549d9f4c06815a2397cef5a540b919": {"balance": "42000000000000000000000000"},
	  "0xbb9003ca8226f411811dd16a3f1a2c1b3f71825d": {"balance": "39900000000000000000000000"}
}`

const wanchainPPOWTestAllocJson = `{
	  "0xbd100cf8286136659a7d63a38a154e28dbf3e0fd": {"balance": "3000000000000000000000000000"},
	  "0xF9b32578b4420a36F132DB32b56f3831A7CC1804": {"balance": "3000000000000000000000000000"},
	  "0x1631447d041f929595a9c7b0c9c0047de2e76186": {"balance": "1000"}
}`

const wanchainPPOWDevAllocJson = `{
	  "0x2d0e7c0813a51d3bd1d08246af2a8a7a57d8922e": {"balance": "1000000000000000000000"},
	  "0x8b179c2b542f47bb2fb2dc40a3cf648aaae1df16": {"balance": "1000000000000000000000"},
	  "0x7a22d4e2dc5c135c4e8a7554157446e206835b05": {"balance": "3000000000000000000000000000"}
}`

var PlutoAllocJson = `{
	"0x2d0e7c0813a51d3bd1d08246af2a8a7a57d8922e": {"balance": "10000000000000000000000000000000000000",
		"staking":{
			"amount":"5000000000000000000000000",
			"s256pk":"0x04d7dffe5e06d2c7024d9bb93f675b8242e71901ee66a1bfe3fe5369324c0a75bf6f033dc4af65f5d0fe7072e98788fcfa670919b5bdc046f1ca91f28dff59db70",
			"bn256pk":"0x2b18624b28a714be61ca85bbbcc22dd47717494493fdb86fcae04928119fb00d15cd5544da7c76d347d4d27b2552e5044ae0e4cff2836ad7bb0de6afe46124a9"
		}
	  },
	"0xe20bfe3c8777036ca0ab03f3bcbbfb438d97dd91": {"balance": "100000000000000000000"},
	"0x9da26fc2e1d6ad9fdd46138906b0104ae68a65d8": {"balance": "100000000000000000000"},
 	"0xEf73Eaa714dC9a58B0990c40a01F4C0573599959": {"balance": "10000000000000000000000000000000000000"},
	"0xb752027021340f2fec33abc91daf198915bbbbbb": {"balance": "10000000000000000000000000000000000000"},
	"0xe8ffc3d0c02c0bfc39b139fa49e2c5475f000000": {"balance": "10000000000000000000000000000000000000"},
	"0x1f233a8cfef8a0b8b84b318d1305e3bcd0074b99": {"balance": "10000000000000000000000000000000000000"},
	"0x2c8C7dA82306377ae85BA2EED8b377E1269E5004": {"balance": "10000000000000000000000000000000000000"}
}`
