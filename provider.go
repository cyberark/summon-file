package main

import (
	"os"
)

func RetrieveSecret(variableIdentifier string, readFile func(filename string) ([]byte, error)) {
	secretValueBytes, err := readFile(variableIdentifier)

	printAndExitIfError(err)

	os.Stdout.Write(secretValueBytes)
}
