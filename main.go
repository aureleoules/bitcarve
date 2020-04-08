package main

import "fmt"

// PubKeyHashSize is the byte size of a public key's hash
const PubKeyHashSize = 20

// MainNetwork byte
const MainNetwork byte = 0x00

// TestNetwork byte
const TestNetwork byte = 0x6f

// Network byte
var Network = MainNetwork

// BitcoinBin bicoin binary
var BitcoinBin = "bitcoin-cli"

// BitcoinArgs arguments for binary
var BitcoinArgs = "-chain=main"

func init() {
	ParseArgs()
	if network == "testnet" {
		Network = TestNetwork
		BitcoinArgs = "-chain=test"
	}
}

func main() {
	if decrypt {
		// Decrypt TX
		DecryptTX()
	} else {
		// Carve on the network
		raw := CreateRawTX()
		signed := SignRawTX(raw)
		txid := BroadcastTX(signed)
		fmt.Println("Data successfully carved on the Bitcoin network.\nTxID: " + txid)
	}

}
