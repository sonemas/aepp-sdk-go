package integrationtest

import (
	"reflect"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v5/config"
	"github.com/aeternity/aepp-sdk-go/v5/naet"
	"gotest.tools/golden"
)

func TestCompiler(t *testing.T) {
	simplestorageSource := "simplestorage.aes"
	simplestorageBytecode := "simplestorage_bytecode.txt"
	simplestorageCalldata := "simplestorage_init42.txt"
	identitySource := "identity.aes"

	c := naet.NewCompiler("http://localhost:3080", false)
	t.Run("APIVersion", func(t *testing.T) {
		_, err := c.APIVersion()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("CompileContract", func(t *testing.T) {
		compiled, err := c.CompileContract(string(golden.Get(t, simplestorageSource)), config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
		golden.Assert(t, compiled, simplestorageBytecode)
	})
	t.Run("CompileErrorDeserialization", func(t *testing.T) {
		// Compiler v4.0.0-rc4 and onwards will return a
		// CompileContractForbidden type for compile errors. If the error type
		// changes in the future and cannot be deserialized, reflect.TypeOf(err)
		// will be of type json.UnmarshalError instead, and this test will fail.
		wontcompileSource := `contract Test =
		entrypoint some_test(ae_address: address) =
			// let test = String.concat("\x19Ethereum Signed Message:\n52", Address.to_str(ae_address))
			ae_addres
	`

		_, err := c.CompileContract(wontcompileSource, config.Compiler.Backend)
		errtype := reflect.TypeOf(err).String()
		if errtype != "*operations.CompileContractForbidden" {
			t.Error(err)
		}
	})
	t.Run("DecodeCallResult", func(t *testing.T) {
		// taken from contract_test.go
		_, err := c.DecodeCallResult("ok", "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACr8s/aY", "main", string(golden.Get(t, identitySource)), config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("DecodeCalldataBytecode", func(t *testing.T) {
		_, err := c.DecodeCalldataBytecode(string(golden.Get(t, simplestorageBytecode)), string(golden.Get(t, simplestorageCalldata)), config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("DecodeCalldataSource", func(t *testing.T) {
		_, err := c.DecodeCalldataSource(string(golden.Get(t, simplestorageSource)), "init", string(golden.Get(t, simplestorageCalldata)), config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("DecodeData", func(t *testing.T) {
		// taken from testnet Contract Call Tx th_toPLrggySMKVecSkEdy7QYF7VEQ4nANAdSiwNXomtwhdp6ZNw
		_, err := c.DecodeData("cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAArMtts", "int")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("EncodeCalldata SimpleStorage set(123)", func(t *testing.T) {
		encodedCalldata, err := c.EncodeCalldata(string(golden.Get(t, simplestorageSource)), "set", []string{"123"}, config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
		golden.Assert(t, encodedCalldata, "simplestorage_set123.txt")
	})
	t.Run("EncodeCalldata SimpleStorage init(42)", func(t *testing.T) {
		encodedCalldata, err := c.EncodeCalldata(string(golden.Get(t, simplestorageSource)), "init", []string{"42"}, config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
		golden.Assert(t, encodedCalldata, "simplestorage_init42.txt")
	})
	t.Run("GenerateACI", func(t *testing.T) {
		_, err := c.GenerateACI(string(golden.Get(t, simplestorageSource)), config.Compiler.Backend)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("SophiaVersion", func(t *testing.T) {
		_, err := c.SophiaVersion()
		if err != nil {
			t.Error(err)
		}
	})

}
