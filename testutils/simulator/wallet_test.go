package simulator

import (
	"encoding/hex"
	"github.com/Qitmeer/qng/config"
	"github.com/Qitmeer/qng/testutils/simulator/testprivatekey"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestImportRawKey(t *testing.T) {
	node, err := StartMockNode(func(cfg *config.Config) error {
		cfg.DebugLevel = "trace"
		cfg.DebugPrintOrigins = true
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	defer node.Stop()

	pk, err := node.pb.Build()
	if err != nil {
		t.Fatal(err)
	}
	pkHex := hex.EncodeToString(pk)
	eaddr, err := node.GetPrivateWalletManagerAPI().ImportRawKey(pkHex, testprivatekey.Password)
	if err != nil {
		t.Fatal(err)
	}
	accounts, err := node.GetPrivateWalletManagerAPI().ListAccount()
	if err != nil {
		t.Fatal(err)
	}
	accountsM := accounts.([]map[string]interface{})
	ret := accountsM[1]["address"].(common.Address)
	if ret.Cmp(eaddr) != 0 {
		t.Fatalf("%s != %s", ret.String(), eaddr.String())
	}
}
