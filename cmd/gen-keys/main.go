package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/celo-org/celo-blockchain/mycelo/env"
)

func main() {
	var (
		mnemonic = flag.String("mnemonic", "", "mnemonic")
	)
	flag.Parse()

	if *mnemonic == "" {
		fmt.Fprintln(os.Stderr, fmt.Errorf("mnemonic args is required"))
		os.Exit(1)
	}

	ac := &env.AccountsConfig{
		Mnemonic:             *mnemonic,
		NumDeveloperAccounts: 10,
	}
	accounts := ac.DeveloperAccounts()
	for i, account := range accounts {
		fmt.Printf("[%d] address: %s, private key: %s\n", i, account.Address.Hex(), account.PrivateKeyHex())
	}
}
