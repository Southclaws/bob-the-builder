package main

import (
	"context"
	"os"

	"github.com/Southclaws/bob-the-builder/bot"
	"github.com/Southclaws/bob-the-builder/builder"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		panic("no token found in env var 'DISCORD_TOKEN'")
	}

	b, err := bot.New(token, &builder.RandomTeamBuilder{})
	if err != nil {
		panic(err)
	}

	err = b.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
