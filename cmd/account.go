// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"regexp"
	"runtime"
	"sync"

	"github.com/aeternity/aepp-sdk-go/v8/account"
	"github.com/aeternity/aepp-sdk-go/v8/config"
	"github.com/aeternity/aepp-sdk-go/v8/naet"
	"github.com/aeternity/aepp-sdk-go/v8/transactions"

	"github.com/spf13/cobra"
)

var (
	waitForTx       bool
	spendTxPayload  string
	printPrivateKey bool
	accountFileName string
	password        string
	fee             string // leave it as a string because viper cannot parse it directly into a BigInt
	ttl             uint64
	nonce           uint64
	regex           bool
)

// accountCmd implements the account command
var accountCmd = &cobra.Command{
	Use:   "account PRIVATE_KEY_PATH",
	Short: "Interact with a account",
	Long:  ``,
}

// addressCmd implements the account address subcommand
var addressCmd = &cobra.Command{
	Use:   "address ACCOUNT_KEYSTORE",
	Short: "Print the aeternity account address",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE:  addressFunc,
}

func getPassword() (p string, err error) {
	if len(password) != 0 {
		return password, nil
	}
	p, err = AskPassword("Enter the password to unlock the keystore: ")
	if err != nil {
		return "", err
	}
	return p, nil
}

func addressFunc(cmd *cobra.Command, args []string) error {
	p, err := getPassword()
	if err != nil {
		return err
	}

	// load the account
	account, err := account.LoadFromKeyStoreFile(args[0], p)
	if err != nil {
		return err
	}

	Pp("Account address", account.Address)
	if printPrivateKey {
		if AskYes("Are you sure you want to print your private key? This could be insecure.", false) {
			Pp("Account private key", account.SigningKeyToHexString())
		}
	}

	return nil
}

// createCmd implements the account generate subcommand
var createCmd = &cobra.Command{
	Use:   "create ACCOUNT_KEYSTORE",
	Short: "Create a new account",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE:  createFunc,
}

func createFunc(cmd *cobra.Command, args []string) (err error) {
	acc, _ := account.New()
	p, err := getPassword()
	if err != nil {
		return err
	}
	accountFileName = args[0]

	// check if a name was given
	f, err := account.StoreToKeyStoreFile(acc, p, accountFileName)
	if err != nil {
		return err
	}

	Pp(
		"Wallet path", f,
		"Account address", acc.Address,
	)

	return nil
}

// balanceCmd implements the account balance subcommand
var balanceCmd = &cobra.Command{
	Use:   "balance ACCOUNT_KEYSTORE",
	Short: "Get the balance of an account",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		node := newAeNode()
		return balanceFunc(node, args)
	},
}

func balanceFunc(conn naet.GetAccounter, args []string) (err error) {
	p, err := getPassword()

	// load the account
	account, err := account.LoadFromKeyStoreFile(args[0], p)
	if err != nil {
		return err
	}

	a, err := conn.GetAccount(account.Address)
	if err != nil {
		return err
	}

	PrintObject("account", a)
	return nil
}

// signCmd implements the account sign subcommand
var signCmd = &cobra.Command{
	Use:   "sign ACCOUNT_KEYSTORE UNSIGNED_TRANSACTION",
	Short: "Sign the input (e.g. a transaction)",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	RunE:  signFunc,
}

func signFunc(cmd *cobra.Command, args []string) (err error) {
	p, err := getPassword()

	// load the account
	account, err := account.LoadFromKeyStoreFile(args[0], p)
	if err != nil {
		return err
	}

	txUnsignedBase64 := args[1]
	tx, err := transactions.DeserializeTxStr(txUnsignedBase64)
	if err != nil {
		return err
	}

	txSigned, txHash, signature, err := transactions.SignHashTx(account, tx, config.Node.NetworkID)
	if err != nil {
		return err
	}
	txSignedB64, err := transactions.SerializeTx(txSigned)
	if err != nil {
		return err
	}

	Pp(
		"Signing account address", account.Address,
		"Signature", signature,
		"Unsigned", txUnsignedBase64,
		"Signed", txSignedB64,
		"Hash", txHash,
	)
	return nil
}

// saveCmd implements the account save subcommand
var saveCmd = &cobra.Command{
	Use:   "save ACCOUNT_KEYSTORE ACCOUNT_HEX_STRING",
	Short: "Save an account from a hex string to a keystore file",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	RunE:  saveFunc,
}

func saveFunc(cmd *cobra.Command, args []string) (err error) {
	accountFileName := args[0]
	acc, err := account.FromHexString(args[1])
	if err != nil {
		return err
	}

	p, err := getPassword()

	f, err := account.StoreToKeyStoreFile(acc, p, accountFileName)
	if err != nil {
		return err
	}

	Pp("Keystore path ", f)

	return nil
}

var vanityCmd = &cobra.Command{
	Use:   "vanity",
	Short: "Find an account that starts with or contains the user-specified text",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run:   vanityFunc,
}

func vanityFunc(cmd *cobra.Command, args []string) {
	var searchString string
	if regex {
		searchString = args[0]
	} else {
		searchString = fmt.Sprintf("^%s", args[0])
	}
	r, err := regexp.Compile(searchString)
	if err != nil {
		fmt.Println("Ouch! The search input ", searchString, "is not a valid regexp")
		return
	}
	fmt.Println("The search for your account matching", searchString, "has begun")

	foundAccounts := make(chan *account.Account, 5)
	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				a, _ := account.New()
				if r.MatchString(a.Address[3:]) {
					foundAccounts <- a
				}
			}
		}()
	}
	for a := range foundAccounts {
		fmt.Printf("Found account! %s \n", a.Address)
		filename := fmt.Sprintf("account.%s.json", a.Address)
		yes := AskYes(fmt.Sprintf("Save it to %s?", filename), false)
		if yes {
			pw, err := AskPassword("To encrypt the keystore, please provide a password")
			if err != nil {
				fmt.Println(err)
			}

			path, err := account.StoreToKeyStoreFile(a, pw, filename)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Account saved in", path)

			continueSearch := AskYes("Would you like to continue searching?", false)
			if !continueSearch {
				return
			}
		}
	}
	wg.Wait()
}

func init() {
	RootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(addressCmd)
	accountCmd.AddCommand(createCmd)
	accountCmd.AddCommand(saveCmd)
	accountCmd.AddCommand(balanceCmd)
	accountCmd.AddCommand(signCmd)
	accountCmd.AddCommand(vanityCmd)
	accountCmd.PersistentFlags().StringVar(&password, "password", "", "Read account password from stdin [WARN: this method is not secure]")
	// account address flags
	addressCmd.Flags().BoolVar(&printPrivateKey, "private-key", false, "Print the private key as hex string")
	vanityCmd.Flags().BoolVar(&regex, "regex", false, "Search using a regular expression that can match anywhere within the address instead of a string that matches at the beginning")
}
