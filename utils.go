package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os/exec"
	"strings"

	"github.com/btcsuite/btcutil/base58"
)

// Exec binary and return stdout
func Exec(cmd string, args ...string) string {
	command := exec.Command(cmd, args...)
	// set var to get the output
	var out bytes.Buffer

	// set the output to our variable
	command.Stdout = &out
	err := command.Run()
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	return strings.TrimSpace(out.String())
}

// Address computes the base58 public address
func Address(data []byte, network byte) string {
	address := append([]byte{network}, data...)
	hash := sha256.Sum256(address)
	hash = sha256.Sum256(hash[:])

	checksum := hash[0:4]

	address = append(address, checksum...)
	return base58.Encode(address)
}
