package github

import (
	"context"

	hub "github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

type GitHub interface {
	Init() error
	Client() *hub.Client
}

type github struct {
	token  string
	client *hub.Client
}

func NewGitHub(token string) GitHub {
	return &github{
		token: token,
	}
}

func (g *github) Init() error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.token},
	)
	tc := oauth2.NewClient(ctx, ts)

	g.client = hub.NewClient(tc)
	_, _, err := g.client.Zen(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (g *github) Client() *hub.Client {
	return g.client
}
