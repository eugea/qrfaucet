package main

import (
	"fmt"

	"github.com/celestiaorg/celestia-node/nodebuilder"
	"github.com/celestiaorg/celestia-node/nodebuilder/node"
)

func main() {
	cfg := nodebuilder.DefaultConfig(node.Light)
	fmt.Println("cfg", cfg)
}
