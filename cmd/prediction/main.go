package main

import (
	"github.com/streamingfast/sparkle/cli"
	"github.com/streamingfast/sparkle/subgraph"
	"github.com/biswapcom/sparkle-biswap/prediction"
)

func main() {
	subgraph.MainSubgraphDef = prediction.Definition
	cli.Execute()
}
