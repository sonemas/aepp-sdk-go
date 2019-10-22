package integrationtest

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v6/account"
	"github.com/aeternity/aepp-sdk-go/v6/aeternity"
	"github.com/aeternity/aepp-sdk-go/v6/config"
	"github.com/aeternity/aepp-sdk-go/v6/models"
	"github.com/aeternity/aepp-sdk-go/v6/naet"
	"github.com/aeternity/aepp-sdk-go/v6/transactions"
	"github.com/aeternity/aepp-sdk-go/v6/utils"
	rlp "github.com/randomshinichi/rlpae"
	"gotest.tools/golden"
)

func EncodeRLPToBytes(tx rlp.Encoder) (b []byte, err error) {
	w := new(bytes.Buffer)
	err = rlp.Encode(w, tx)
	if err != nil {
		return
	}
	return w.Bytes(), nil
}

func TestGeneralizedAccounts(t *testing.T) {
	alice, bob := setupAccounts(t)
	node := setupNetwork(t, privatenetURL, false)
	compiler := naet.NewCompiler(config.Client.Contracts.CompilerURL, false)
	ttler, _, ttlnoncer := transactions.GenerateTTLNoncer(node)

	// Take note of Bob's balance, and after this test, we expect it to have this much more AE
	amount := utils.NewIntFromUint64(5000)
	expected := new(big.Int)
	bobState, err := node.GetAccount(bob.Address)
	if err != nil {
		expected.Set(amount)
	} else {
		bS := big.Int(bobState.Balance)
		expected.Add(&bS, amount)
	}

	authorizeSource := string(golden.Get(t, "authorize.aes"))
	// Read the auth contract from a file, compile and prepare its init() calldata
	authBytecode, err := compiler.CompileContract(authorizeSource, config.Compiler.Backend)
	if err != nil {
		t.Fatal(err)
	}
	authInitCalldata, err := compiler.EncodeCalldata(authorizeSource, "init", []string{alice.Address}, config.Compiler.Backend)
	if err != nil {
		t.Fatal(err)
	}

	// Create throwaway test account, fund it and ensure it is a POA
	testAccount, err := account.New()
	if err != nil {
		t.Fatal(err)
	}
	fundAccount(t, node, alice, testAccount, big.NewInt(1000000000000000000))
	testAccountState, err := node.GetAccount(testAccount.Address)
	if err != nil {
		t.Fatal(err)
	}
	if testAccountState.Kind != "basic" {
		t.Fatalf("%s is supposed to be a basic account but wasn't", testAccount.Address)
	}

	// GAAttachTx
	// Create a Contract{} struct from the compiled bytecode to get its authfunc hash
	auth, err := models.NewContractFromString(authBytecode)
	if err != nil {
		t.Fatal(err)
	}
	gaTx, err := transactions.NewGAAttachTx(testAccount.Address, authBytecode, auth.TypeInfo[0].FuncHash, config.Client.Contracts.VMVersion, config.Client.Contracts.ABIVersion, config.Client.Contracts.GasLimit, config.Client.GasPrice, authInitCalldata, ttlnoncer)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _, _, _, err = aeternity.SignBroadcastWaitTransaction(gaTx, testAccount, node, networkID, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}

	// The test account should now be a generalized account
	checkGeneralizedAccount := func() {
		testAccountState, err = node.GetAccount(testAccount.Address)
		if err != nil {
			t.Fatal(err)
		}
	}
	delay(checkGeneralizedAccount)
	if testAccountState.Kind != "generalized" {
		t.Fatalf("%s was supposed to be a generalized account but isn't", testAccount.Address)
	}

	// GAMetaTx
	// spendTx will be wrapped in a SignedTx with 0 signatures before being
	// included in GAMetaTx. The constructor NewGAMetaTx() does this for you.
	// authData is authorize(3)
	authData := "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACBGtXufEG2HuMYcRcNwsGAeqymslunKf692bHnvwI5K6wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAU3aKBNm"
	zeroNoncer := func(accountID string) (nonce uint64, err error) {
		return 0, nil
	}
	metaTxTTLNoncer := transactions.CreateTTLNoncer(ttler, zeroNoncer)
	spendTx, err := transactions.NewSpendTx(testAccount.Address, bob.Address, big.NewInt(5000), []byte{}, metaTxTTLNoncer)
	gaMetaTx, err := transactions.NewGAMetaTx(testAccount.Address, authData, config.Client.Contracts.ABIVersion, config.Client.GasPrice, config.Client.GasPrice, spendTx, ttler)

	gaMetaTxFinal, hash, _, err := transactions.SignHashTx(testAccount, gaMetaTx, config.Node.NetworkID)
	if err != nil {
		t.Fatal(err)
	}

	gaMetaTxStr, err := transactions.SerializeTx(gaMetaTxFinal)
	if err != nil {
		t.Fatal(err)
	}
	err = node.PostTransaction(gaMetaTxStr, hash)
	if err != nil {
		t.Fatal(err)
	}

	// check bob.Address, make sure it got the SpendTx
	getBobsAccount := func() {
		bobState, err = node.GetAccount(bob.Address)
		if err != nil {
			t.Fatalf("Couldn't get Bob's account data: %v", err)
		}
	}
	delay(getBobsAccount)
	b := big.Int(bobState.Balance)

	if expected.Cmp(&b) != 0 {
		t.Fatalf("Bob should have %v, but has %v instead", expected.String(), bobState.Balance.String())
	}
}
