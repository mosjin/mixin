package bsv

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/MixinNetwork/mixin/crypto"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/btcsuite/btcd/chaincfg"
)

var (
	BitcoinSVChainBase string
	BitcoinSVChainId   crypto.Hash
)

func init() {
	BitcoinSVChainBase = "574388fd-b93f-4034-a682-01c2bc095d17"
	BitcoinSVChainId = crypto.NewHash([]byte(BitcoinSVChainBase))
}

func VerifyAssetKey(assetKey string) error {
	if assetKey == BitcoinSVChainBase {
		return nil
	}
	return fmt.Errorf("invalid bitcoin sv asset key %s", assetKey)
}

func VerifyAddress(address string) error {
	if strings.TrimSpace(address) != address {
		return fmt.Errorf("invalid bitcoin sv address %s", address)
	}

	address = strings.TrimSpace(address)
	params := &chaincfg.MainNetParams
	_, netID, err := base58.CheckDecode(address)
	if err != nil {
		return fmt.Errorf("invalid bitcoin sv address %s %s", address, err)
	}
	if netID != params.PubKeyHashAddrID {
		return fmt.Errorf("invalid bitcoin sv address %s", address)
	}

	bsvAddress, err := btcutil.DecodeAddress(address, params)
	if err != nil {
		return fmt.Errorf("invalid bitcoin sv address %s %s", address, err)
	}
	if bsvAddress.String() != address {
		return fmt.Errorf("invalid bitcoin sv address %s", address)
	}
	return nil
}

func VerifyTransactionHash(hash string) error {
	if len(hash) != 64 {
		return fmt.Errorf("invalid bitcoin sv transaction hash %s", hash)
	}
	if strings.ToLower(hash) != hash {
		return fmt.Errorf("invalid bitcoin sv transaction hash %s", hash)
	}
	h, err := hex.DecodeString(hash)
	if err != nil {
		return fmt.Errorf("invalid bitcoin sv transaction hash %s %s", hash, err.Error())
	}
	if len(h) != 32 {
		return fmt.Errorf("invalid bitcoin sv transaction hash %s", hash)
	}
	return nil
}

func GenerateAssetId(assetKey string) crypto.Hash {
	switch assetKey {
	case BitcoinSVChainBase:
		return BitcoinSVChainId
	default:
		panic(assetKey)
	}
}
