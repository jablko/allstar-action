package main

import (
	"context"
	"os"

	"github.com/google/go-github/v50/github"
	"github.com/ossf/allstar/pkg/enforce"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("INPUT_TOKEN")})
	tc := oauth2.NewClient(ctx, ts)
	_, err := enforce.EnforceAll(ctx, (*GHClients)(github.NewClient(tc)), "", os.Getenv("GITHUB_REPOSITORY"))
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Unexpected error enforcing policies.")
	}
}

type GHClients github.Client

func (g *GHClients) Get(i int64) (*github.Client, error) {
	return (*github.Client)(g), nil
}

func (g *GHClients) LogCacheSize() {}
