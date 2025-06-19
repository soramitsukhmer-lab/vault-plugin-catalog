package main

import (
	"log"
	"os"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/command"
)

func main() {
	if err := command.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
