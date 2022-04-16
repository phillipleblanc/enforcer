package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

type Inputs struct {
	Token     string
	Owner     string
	Repo      string
	Workspace string
}

func main() {
	inputs := getInputs()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: inputs.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	ghClient := github.NewClient(tc)

	issues, _, err := ghClient.Issues.ListByRepo(context.Background(), inputs.Owner, inputs.Repo, &github.IssueListByRepoOptions{State: "open"})
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, issue := range issues {
		if issue.Title != nil {
			fmt.Printf("%s\n", *issue.Title)
		}
	}
}

func getInputs() *Inputs {
	env := []string{"GITHUB_TOKEN", "GITHUB_REPOSITORY", "GITHUB_WORKSPACE"}
	for _, e := range env {
		if os.Getenv(e) == "" {
			log.Fatal(e + " is empty")
		}
	}
	inputs := &Inputs{}

	repoInputs := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")
	if len(repoInputs) != 2 {
		log.Fatalf("unexpected GITHUB_REPOSITORY: %s", os.Getenv("GITHUB_REPOSITORY"))
	}

	inputs.Owner = repoInputs[0]
	inputs.Repo = repoInputs[1]
	inputs.Workspace = os.Getenv("GITHUB_WORKSPACE")
	inputs.Token = os.Getenv("GITHUB_TOKEN")

	return inputs
}
