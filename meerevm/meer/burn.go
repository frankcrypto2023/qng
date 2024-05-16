package meer

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/Qitmeer/qng/core/address"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const RELEASE_CONTRACT_ADDR = "0x1000000000000000000000000000000000000000"

type BurnDetail struct {
	Order  int64  `json:"order"`
	Height int64  `json:"height"`
	From   string `json:"from"`
	Amount int64  `json:"amount"`
	Time   int64  `json:"time"`
}
type BurnTimes int

type BurnerAdressHash160 string

type BurnerTimes map[BurnerAdressHash160]BurnTimes

func (bt *BurnTimes) ToInt() int {
	return int(*bt)
}
func (bah *BurnerAdressHash160) ToString() string {
	return string(*bah)
}

func (bah *BurnerTimes) SortKeys(callback func(keys []string)) {
	keys := make([]string, 0)
	for k := range *bah {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	callback(keys)
}

// 2022/08/17 20:35:36 MmQitmeerMainNetHonorAddressXY9JH2y burn amount 408011208230864
// 2022/08/17 20:35:36 MmQitmeerMainNetGuardAddressXd7b76q burn amount 514790066054534
// 2022/08/17 20:35:36 All burn amount 922801274285398
// 2022/08/14 17:43:57 end height 910000
// 2022/08/14 17:43:57 end order 1013260
// 2022/08/14 17:43:57 end blockhash efc89d8b4ef5733b6e566d9f06c0596075100f8406d3a9b581c74d42fb99dd79
// 2022/08/14 17:43:57 pow meer amount (1013260 /10) * 12 * 10 = 1013260 * 12 = 12159120
// all amount 1215912000000000+922801274285398 = 2138713274285398

func BuildBurnBalance(burnStr string) map[common.Hash]common.Hash {
	storage := map[common.Hash]common.Hash{}
	burnList := map[string][]BurnDetail{}
	jsonR := strings.NewReader(burnStr)
	if err := json.NewDecoder(jsonR).Decode(&burnList); err != nil {
		panic(err)
	}

	burnTimes := BurnerTimes{}

	allBurnAmount := uint64(0)
	burnM := map[string]uint64{}
	keys := make([]string, 0)
	for k := range burnList {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		burnRecords := burnList[k]
		for _, burnDetail := range burnRecords {
			burnerAddr, err := address.DecodeAddress(burnDetail.From)
			if err != nil {
				panic(burnDetail.From + "meer address err" + err.Error())
			}

			h16 := burnerAddr.Hash160()
			burnerAddrHash160 := BurnerAdressHash160(hex.EncodeToString(h16[:]))

			// storage the mapping key value on storage slot
			burnerTimes := burnTimes[burnerAddrHash160]
			key, value := BuildMappingStorageSlot(burnerAddrHash160.ToString(), burnerTimes.ToInt(), 0, big.NewInt(burnDetail.Amount))
			storage[key] = value
			key, value = BuildMappingStorageSlot(burnerAddrHash160.ToString(), burnerTimes.ToInt(), 1, big.NewInt(burnDetail.Time))
			storage[key] = value
			key, value = BuildMappingStorageSlot(burnerAddrHash160.ToString(), burnerTimes.ToInt(), 2, big.NewInt(burnDetail.Order))
			storage[key] = value
			key, value = BuildMappingStorageSlot(burnerAddrHash160.ToString(), burnerTimes.ToInt(), 3, big.NewInt(burnDetail.Height))
			storage[key] = value

			burnTimes[burnerAddrHash160]++
			allBurnAmount += uint64(burnDetail.Amount)
			burnM[k] += uint64(burnDetail.Amount)
		}
	}
	burnTimes.SortKeys(func(keys []string) {
		for _, v := range keys {
			burnerTimes := burnTimes[BurnerAdressHash160(v)]
			// storage the mapping key length on storage slot
			key, value := BuildStorageSlot(v, 0, burnerTimes.ToInt())
			storage[key] = value
		}
	})

	for k, v := range burnM {
		log.Trace(k, "burn amount", v)
	}
	log.Debug("All burn amount", allBurnAmount)
	return storage
}

/*
*

	solidity code like:
	struct BurnDetail {
	    uint amount;
	    uint time;
	    uint order;
	    uint height;
	}
	mapping(string => BurnDetail[]) burnUsers;
	@param mapKey is burnUsers key of user's address hash160 string
	@param keyPosition is the BurnDetail[] index , slot storage position 0-1-2-3...
	@param valuePosition is the BurnDetail fields storage position 0-1-2-3, just the field order
	@param mapVal is the actual value of the BurnDetail field

*
*/
func BuildMappingStorageSlot(mapKey string, keyPosition, valuePosition int, mapVal *big.Int) (key common.Hash, value common.Hash) {
	s := mapKey + fmt.Sprintf("%064x", big.NewInt(1))
	b, _ := hex.DecodeString(s)
	keyHash := crypto.Keccak256(b)
	s = fmt.Sprintf("%064x", big.NewInt(int64(keyPosition))) + hex.EncodeToString(keyHash)
	b, _ = hex.DecodeString(s)
	keyHash = crypto.Keccak256(b)
	key0Big := new(big.Int).Add(new(big.Int).SetBytes(keyHash), big.NewInt(int64(valuePosition)))
	key = common.HexToHash(fmt.Sprintf("%064x", key0Big))
	value = common.HexToHash(fmt.Sprintf("%064x", mapVal))
	return
}

/*
*

	solidity code like:
	struct BurnDetail {
	    uint amount;
	    uint time;
	    uint order;
	    uint height;
	}
	mapping(string => BurnDetail[]) burnUsers;
	@param mapKey is burnUsers key of user's address hash160 string
	@param keyPosition is the mapping first position for recording the length of BurnDetail[]
	@param valueLength is the length of the BurnDetail[]

*
*/
func BuildStorageSlot(mapKey string, keyPosition, valueLength int) (key common.Hash, value common.Hash) {
	b, _ := hex.DecodeString(mapKey + fmt.Sprintf("%064x", big.NewInt(int64(keyPosition))))
	keyHash := crypto.Keccak256(b)
	key = common.BytesToHash(keyHash)
	value = common.HexToHash(fmt.Sprintf("%064x", valueLength))
	return
}
