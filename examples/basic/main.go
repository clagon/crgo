package main

import (
	"context"
	"fmt"
	"log"
	"os"

	crgo "github.com/clagon/crgo"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: basic PLAYER_TAG")
	}
	client, err := crgo.NewClient(os.Getenv("CLASH_ROYALE_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	player, err := client.GetPlayer(context.Background(), os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s (%s)\n", player.Name, player.Tag)
}
