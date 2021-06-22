package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	godelNode, godelNodeErr := NewGodelNode(ctx)
	if godelNodeErr != nil {
		panic(fmt.Errorf("Fatal error godel node: %s \n", godelNodeErr))
	}

	startErr := godelNode.Start(ctx)
	if startErr != nil {
		panic(fmt.Errorf("Fatal error godel node start: %s \n", startErr))
	}

	log.Fatal(godelNode.ServeApi())
}
