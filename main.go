package main

import (
	"os"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.Write([]byte("A variable identifier must be given as the first and only argument!"))
		os.Exit(-1)
	}
	variableIdentifier := os.Args[1]

	RetrieveSecret(variableIdentifier, ioutil.ReadFile)
}

func printAndExitIfError(err error) {
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}
}
