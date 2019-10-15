package aeternity

import (
	"errors"
	"fmt"

	"github.com/aeternity/aepp-sdk-go/v6/account"
	"github.com/aeternity/aepp-sdk-go/v6/config"
	"github.com/aeternity/aepp-sdk-go/v6/naet"
	"github.com/aeternity/aepp-sdk-go/v6/transactions"
)

func findVMABIVersion(nodeVersion, compilerBackend string) (VMVersion, ABIVersion uint16, err error) {
	if nodeVersion[0] == '5' && compilerBackend == "fate" {
		return 5, 3, nil
	} else if nodeVersion[0] == '5' && compilerBackend == "aevm" {
		return 6, 1, nil
	} else if nodeVersion[0] == '4' {
		return 4, 1, nil
	} else {
		return 0, 0, errors.New("Other node versions unsupported")
	}
}

type compileencoder interface {
	naet.CompileContracter
	naet.EncodeCalldataer
}

// CreateContract lets one deploy a contract with minimum fuss.
func CreateContract(n naet.NodeInterface, c compileencoder, acc *account.Account, source, function string, args []string, backend string) (signedTxStr, hash, signature string, blockHeight uint64, blockHash string, err error) {
	status, err := n.GetStatus()
	if err != nil {
		return
	}
	networkID := *status.NetworkID

	bytecode, err := c.CompileContract(source, backend)
	if err != nil {
		return
	}
	calldata, err := c.EncodeCalldata(source, function, args, backend)
	if err != nil {
		return
	}

	VMVersion, ABIVersion, err := findVMABIVersion(*status.NodeVersion, backend)
	if err != nil {
		return
	}
	ttlNonceGetter := GenerateGetTTLNonce(GenerateGetTTL(n), GenerateGetNextNonce(n))
	ttl, nonce, err := ttlNonceGetter(acc.Address, config.Client.TTL)
	if err != nil {
		return
	}

	createTx := transactions.NewContractCreateTx(acc.Address, nonce, bytecode, VMVersion, ABIVersion, config.Client.Contracts.Deposit, config.Client.Contracts.Amount, config.Client.Contracts.GasLimit, config.Client.GasPrice, config.Client.Fee, ttl, calldata)
	transactions.CalculateFee(createTx)
	createTxStr, _ := transactions.SerializeTx(createTx)
	fmt.Printf("%+v\n", createTx)
	fmt.Println(createTxStr)
	signedTxStr, hash, signature, blockHeight, blockHash, err = SignBroadcastWaitTransaction(createTx, acc, n.(*naet.Node), networkID, config.Client.WaitBlocks)
	if err != nil {
		return
	}
	return
}