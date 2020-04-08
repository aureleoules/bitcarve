package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// CreateRawTX builds a raw bitcoin tx
func CreateRawTX() string {
	inputs := "[{\"txid\":\"" + utxo + "\",\"vout\":" + strconv.Itoa(vout) + "}] "

	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	outputs := "["

	const offset = 20
	for i := 0; i < len(data); i += offset {
		encoded := data[i : i+20]
		address := Address(encoded, Network)
		outputs += "{\"" + address + "\":" + fmt.Sprintf("%.8f", float64(amount)/100000000) + "}"

		if i+offset < len(data) {
			outputs += ","
		}
	}

	outputs += "]"

	output := Exec(BitcoinBin, BitcoinArgs, "createrawtransaction", inputs, outputs)
	return output
}

// SignRawTX signes a bitcoin transaction
func SignRawTX(tx string) string {
	output := Exec(BitcoinBin, BitcoinArgs, "signrawtransactionwithkey", tx, "[\""+privateKey+"\"]")
	var response map[string]string
	json.Unmarshal([]byte(output), &response)
	return response["hex"]
}

// BroadcastTX sends the signed tx to the bitcoin network
func BroadcastTX(tx string) string {
	return Exec(BitcoinBin, BitcoinArgs, "sendrawtransaction", tx)
}
