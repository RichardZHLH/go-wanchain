package slotleader

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/wanchain/go-wanchain/pos/posconfig"
	"github.com/wanchain/go-wanchain/pos/util/convert"

	"github.com/btcsuite/btcd/btcec"
	"github.com/wanchain/go-wanchain/crypto"
	"github.com/wanchain/go-wanchain/pos/posdb"
	"github.com/wanchain/go-wanchain/rlp"
)

func TestSlotLeaderSelectionGetInstance(t *testing.T) {
	posdb.GetDb().DbInit("test")
	slot := GetSlotLeaderSelection()
	if slot == nil {
		t.Fail()
	}
}

func TestPublicKeyCompress(t *testing.T) {
	privKey, _ := crypto.GenerateKey()

	fmt.Println("Is on curve: ", crypto.S256().IsOnCurve(privKey.X, privKey.Y))

	fmt.Println("public key:", hex.EncodeToString(crypto.FromECDSAPub(&privKey.PublicKey)))

	pk := btcec.PublicKey(privKey.PublicKey)

	fmt.Println("public key uncompress:", hex.EncodeToString(pk.SerializeUncompressed()), "len: ", len(pk.SerializeUncompressed()))

	fmt.Println("public key compress:", hex.EncodeToString(pk.SerializeCompressed()), "len: ", len(pk.SerializeCompressed()))

	keyCompress := pk.SerializeCompressed()

	key, _ := btcec.ParsePubKey(keyCompress, btcec.S256())

	pKey := ecdsa.PublicKey(*key)

	fmt.Println("public key:", hex.EncodeToString(crypto.FromECDSAPub(&pKey)))
}

func TestRlpEncodeAndDecode(t *testing.T) {

	privKey, _ := crypto.GenerateKey()
	pk := btcec.PublicKey(privKey.PublicKey)
	keyCompress := pk.SerializeCompressed()

	var test = [][]byte{
		new(big.Int).SetInt64(1).Bytes(),
		keyCompress,
		keyCompress,
	}

	fmt.Println("before encode:", test)

	buf, err := rlp.EncodeToBytes(test)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("encode buf: ", hex.EncodeToString(buf))

	fmt.Println("encode len: ", len(buf))

	var output [][]byte
	rlp.DecodeBytes(buf, &output)

	fmt.Println("after decode:", output)
}

func TestAbiPack2(t *testing.T) {

}

// TestByteToString is test for bytes compare with string() convert
func TestByteToString(t *testing.T) {
	testBytes := make([]byte, 0)
	for i := 0; i < 255; i++ {
		testBytes = append(testBytes, byte(i))
	}
	fmt.Println("bytes: ", testBytes)
	fmt.Println("string: ", string(testBytes))
	fmt.Println("string len:", len(string(testBytes)))

	testBytes2 := make([]byte, 0)
	for i := 0; i < 255; i++ {
		testBytes2 = append(testBytes2, byte(i))
	}

	if string(testBytes) != string(testBytes2) {
		t.Fail()
	}
}

func TestNumToString(t *testing.T) {
	value, err := hex.DecodeString("0")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(value)
}

func TestCompare(t *testing.T) {
	epID := []byte{84}
	epochID := uint64(84)
	idxID := []byte{1}
	index := uint64(1)

	fmt.Println(hex.EncodeToString(epID))
	fmt.Println(hex.EncodeToString(idxID))
	fmt.Println(hex.EncodeToString(convert.Uint64ToBytes(epochID)))
	fmt.Println(hex.EncodeToString(convert.Uint64ToBytes(index)))

	if hex.EncodeToString(epID) == hex.EncodeToString(convert.Uint64ToBytes(epochID)) &&
		hex.EncodeToString(idxID) == hex.EncodeToString(convert.Uint64ToBytes(index)) {
		return
	}

	t.Fail()
}

func TestProof(t *testing.T) {

	type Test struct {
		Proof    [][]byte
		ProofMeg [][]byte
	}

	key1, _ := crypto.GenerateKey()
	key2, _ := crypto.GenerateKey()

	a := &Test{Proof: [][]byte{big.NewInt(999).Bytes(), big.NewInt(111).Bytes()}, ProofMeg: [][]byte{crypto.FromECDSAPub(&key1.PublicKey), crypto.FromECDSAPub(&key2.PublicKey)}}

	fmt.Println(a)

	buf, err := rlp.EncodeToBytes(a)
	if err != nil {
		t.Fail()
	}

	fmt.Println(hex.EncodeToString(buf))

	var b Test

	err = rlp.DecodeBytes(buf, &b)
	if err != nil {
		t.Fail()
	}

	fmt.Println(b)

}

func TestCRSave(t *testing.T) {

	fmt.Printf("hello world\n\n\n")

	info := ""

	i := 1
	info += fmt.Sprintf("hello world %d \n\n\n", i)

	fmt.Print(info)

	cr := make([]*big.Int, 100)
	for i := 0; i < 100; i++ {
		key, _ := crypto.GenerateKey()
		fmt.Println(hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)))
		cr[i] = key.D
	}

	buf, err := rlp.EncodeToBytes(cr)
	if err != nil {
		t.Fail()
	}

	fmt.Println("buf len:", len(buf))

	var crOut []*big.Int
	err = rlp.DecodeBytes(buf, &crOut)
	if err != nil {
		t.Fail()
	}

	for i := 0; i < 100; i++ {
		if cr[i].String() != crOut[i].String() {
			t.Fail()
		}
	}

}

func TestArraySave(t *testing.T) {

	fmt.Printf("TestArraySave\n\n\n")
	var sendtrans [posconfig.EpochLeaderCount]bool
	for index := range sendtrans {
		sendtrans[index] = false
	}
	fmt.Println(sendtrans)

	sendtrans[0] = true
	sendtrans[posconfig.EpochLeaderCount-1] = true

	bytes, err := rlp.EncodeToBytes(sendtrans)
	if err != nil {
		t.Error(err.Error())
	}

	db := posdb.NewDb("testArraySave")
	db.Put(uint64(0), "TestArraySave", bytes)

	var sendtransGet [posconfig.EpochLeaderCount]bool
	bytesGet, err := db.Get(uint64(0), "TestArraySave")
	if err != nil {
		t.Error(err.Error())
	}
	err = rlp.DecodeBytes(bytesGet, &sendtransGet)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(sendtransGet)
}

func TestGetEpoch0LeadersPK(t *testing.T) {
	SlsInit()
	s := GetSlotLeaderSelection()
	leadersPK := s.getEpoch0LeadersPK()
	if len(leadersPK) != posconfig.EpochLeaderCount {
		t.Fail()
	}

	for _, epLeader := range leadersPK {
		if !crypto.S256().IsOnCurve(epLeader.X, epLeader.Y) {
			t.Error("PK not on the S256 curve")
		}
	}
}

func TestGetPreEpochLeadersPK(t *testing.T) {
	SlsInit()
	s := GetSlotLeaderSelection()
	pks, err := s.getPreEpochLeadersPK(0)
	if err != nil {
		t.Error(err.Error())
	}
	if len(pks) != posconfig.EpochLeaderCount {
		t.Fail()
	}
}

func TestGetSMAPieces(t *testing.T) {
	SlsInit()
	s := GetSlotLeaderSelection()
	pks, isGenesis, err := s.getSMAPieces(0)
	if err != nil {
		t.Error(err.Error())
	}
	if !isGenesis {
		t.Fail()
	}

	if len(pks) != posconfig.EpochLeaderCount {
		t.Fail()
	}
}

func TestDump(t *testing.T) {
	SlsInit()
	s := GetSlotLeaderSelection()
	epochID := s.getWorkingEpochID()
	s.setWorkingEpochID(2)

	epochLeaderAllBytes := make([]byte, 65*posconfig.EpochLeaderCount)
	for i := 0; i < posconfig.EpochLeaderCount; i++ {
		prvKey, _ := crypto.GenerateKey()
		pubKeyByes := crypto.FromECDSAPub(&prvKey.PublicKey)
		copy(epochLeaderAllBytes[i*65:], pubKeyByes[:])
	}
	posdb.GetDb().Put(2, EpochLeaders, epochLeaderAllBytes[:])
	posdb.GetDb().Put(1, EpochLeaders, epochLeaderAllBytes[:])

	posconfig.SelfTestMode = true
	s.dumpData()

	s.setWorkingEpochID(epochID)
	posconfig.SelfTestMode = false
}

func TestClearData(t *testing.T) {
	SlsInit()
	s := GetSlotLeaderSelection()
	epochID := s.getWorkingEpochID()
	s.setWorkingEpochID(2)

	epochLeaderAllBytes := make([]byte, 65*posconfig.EpochLeaderCount)
	for i := 0; i < posconfig.EpochLeaderCount; i++ {
		prvKey, _ := crypto.GenerateKey()
		pubKeyByes := crypto.FromECDSAPub(&prvKey.PublicKey)
		copy(epochLeaderAllBytes[i*65:], pubKeyByes[:])
	}
	posdb.GetDb().Put(2, EpochLeaders, epochLeaderAllBytes[:])
	posdb.GetDb().Put(1, EpochLeaders, epochLeaderAllBytes[:])

	posconfig.SelfTestMode = true
	s.buildEpochLeaderGroup(1)

	s.dumpData()
	s.clearData()
	s.dumpData()

	s.setWorkingEpochID(epochID)
	posconfig.SelfTestMode = false
}

func TestGetSlotLeader(t *testing.T) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	t.Log("Current dir path ", dir)
	os.RemoveAll(path.Join(dir, "sl_leader_test"))
	posdb.GetDb().DbInit(path.Join(dir, "sl_leader_test"))
	//posdb.NewDb(path.Join(dir, "sl_leader_test"))

	SlsInit()
	s := GetSlotLeaderSelection()
	pk, err := s.GetSlotLeader(0, 1)
	if err != nil {
		t.Fail()
	}

	if pkBytes, _ := hex.DecodeString(posconfig.GenesisPK); !bytes.Equal(pkBytes, crypto.FromECDSAPub(pk)) {
		t.Fail()
	}

	pkGenesisBytes, _ := hex.DecodeString(posconfig.GenesisPK)
	// ErrSlotIDOutOfRange
	_, err = s.GetSlotLeader(1, posconfig.SlotCount)
	if err == nil {
		t.Logf("should err:%v,but nil", err.Error())
		t.Fail()
	}

	//ErrSlotLeaderGroupNotReady

	_, err = s.GetSlotLeader(1, posconfig.SlotCount-1)
	if err == nil {
		t.Fail()
	}

	for i := 0; i < posconfig.SlotCount; i++ {
		_, err = posdb.GetDb().PutWithIndex(1, uint64(i), SlotLeader, pkGenesisBytes)
		if err != nil {
			t.Error(err.Error())
			t.Fail()
		}
	}

	s.slotCreateStatus[1] = false

	pkSelected, err := s.GetSlotLeader(1, uint64(posconfig.SlotCount-1))
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if !bytes.Equal(pkGenesisBytes, crypto.FromECDSAPub(pkSelected)) {
		t.Logf("should pk:%v, now pk:%v", pkGenesisBytes, crypto.FromECDSAPub(pkSelected))
		t.Fail()
	} else {
		t.Logf("\nshould pk:\t%v\nnow pk:\t%v", hex.EncodeToString(pkGenesisBytes),
			hex.EncodeToString(crypto.FromECDSAPub(pkSelected)))
	}

	os.RemoveAll(path.Join(dir, "sl_leader_test"))

}
