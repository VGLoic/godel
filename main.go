package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("-- Creating Godel Node --")
	godelNode, godelNodeErr := NewGodelNode(ctx)
	if godelNodeErr != nil {
		panic(fmt.Errorf("Fatal error godel node: %s \n", godelNodeErr))
	}

	fmt.Println("-- Starting Godel Node --")
	startErr := godelNode.Start(ctx)
	if startErr != nil {
		panic(fmt.Errorf("Fatal error godel node start: %s \n", startErr))
	}

	fmt.Println("-- Serving the API --")
	log.Fatal(godelNode.ServeApi())
}
