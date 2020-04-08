package main

import "flag"

var network string
var amount int
var fee int
var utxo string
var vout int
var file string
var privateKey string
var carve bool

var decrypt bool
var txid string
var outputFile string

// ParseArgs from cli
func ParseArgs() {
	flag.StringVar(&network, "network", "main", "network")
	flag.IntVar(&amount, "amount", 1, "sats per address")

	flag.IntVar(&fee, "fee", 1, "fee sats/byte")

	flag.StringVar(&utxo, "utxo", "", "utxo")
	flag.IntVar(&vout, "vout", 0, "output id in tx")
	flag.StringVar(&file, "file", "", "file to carve")
	flag.StringVar(&privateKey, "key", "", "private key in base 58")
	flag.BoolVar(&carve, "carve", true, "carve")

	flag.BoolVar(&decrypt, "decrypt", false, "decrypt")
	flag.StringVar(&txid, "txid", "", "TxID to decrypt")
	flag.StringVar(&outputFile, "output", "output", "File output")

	flag.Parse()
}
