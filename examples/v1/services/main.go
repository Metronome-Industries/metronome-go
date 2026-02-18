package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	listServices(client)
}

// listServices lists all available Metronome services.
func listServices(client metronome.Client) {
	resp, err := client.V1.Services.List(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Services: %+v\n", resp)
}
