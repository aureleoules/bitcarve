package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/thedevsaddam/gojsonq"
)

// DecryptTX decrypts data carved in a tx
func DecryptTX() {
	chain := "bitcoin"
	if Network == TestNetwork {
		chain = "bitcoin/testnet"
	}
	// Fetch data
	route := "https://api.blockchair.com/" + chain + "/dashboards/transaction/" + txid

	resp, err := http.Get(route)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Retrieve and cast outputs
	jq := gojsonq.New().FromString(string(body))
	outputs := jq.From("data." + txid + ".outputs").Get().([]interface{})

	// Store buffers in data
	var data []byte
	for _, o := range outputs {
		data = append(data, RetrieveAddressData(o.(map[string]interface{})["recipient"].(string))...)
	}

	// Write to output
	err = ioutil.WriteFile(outputFile, data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Retrieved data.")
}

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

	fmt.Println(inputs)
	output := Exec(BitcoinBin, BitcoinArgs, "createrawtransaction", inputs, outputs)
	ioutil.WriteFile("raw.tx", []byte(output), 0644)

	return output
}

// SignRawTX signes a bitcoin transaction
func SignRawTX(tx string) string {
	output := Exec(BitcoinBin, BitcoinArgs, "signrawtransactionwithkey", tx, "[\""+privateKey+"\"]")
	var response map[string]interface{}
	err := json.Unmarshal([]byte(output), &response)
	if err != nil {
		panic(err)
	}
	signed := response["hex"].(string)
	ioutil.WriteFile("signed.tx", []byte(signed), 0644)
	fmt.Println("Signed.")
	return signed
}

// BroadcastTX sends the signed tx to the bitcoin network
func BroadcastTX(tx string) string {
	output := Exec(BitcoinBin, BitcoinArgs, "sendrawtransaction", tx)
	fmt.Println(output)
	return output
}
