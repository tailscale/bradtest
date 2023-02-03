package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v50/github"
	"github.com/gregjones/httpcache"
	"golang.org/x/oauth2"
)

/*
GITHUB_WORKFLOW_REF=tailscale/bradtest2/.github/workflows/main.yml@refs/pull/1/merge
POWERSHELL_DISTRIBUTION_CHANNEL=GitHub-Actions-ubuntu22
DOTNET_MULTILEVEL_LOOKUP=0
GITHUB_REF=refs/pull/1/merge
GITHUB_WORKFLOW_SHA=d55e97cee545de9b0031b32a43da38ee41973e4f
GITHUB_REF_NAME=1/merge
GITHUB_REPOSITORY=tailscale/bradtest2

commit d55e97cee545de9b0031b32a43da38ee41973e4f
Merge: 8c21aab 87dbd86
Author: Brad Fitzpatrick <brad@danga.com>
Date:   Fri Feb 3 00:16:18 2023 +0000

    Merge 87dbd86448f35f8610b7e88224a2b1f4c0646d3b into 8c21aab3397fcc3582bd2249d581f2946b9855fd


*/

func main() {
	// env, _ := json.MarshalIndent(os.Environ(), "", "  ")
	// envs := strings.ReplaceAll(string(env), "TOKEN", "NOT_A_THING")
	// fmt.Printf("Hello from Go. Env is: %s\n", envs)
	// fmt.Printf("Argv is %q\n", os.Args)

	gitHubToken := os.Getenv("GITHUB_TOKEN")
	if gitHubToken == "" && len(os.Args) >= 2 {
		gitHubToken = os.Args[1]
	}

	fmt.Printf("length of the token: %v\n", len(gitHubToken))

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitHubToken},
	)
	tc := &http.Client{
		Transport: &oauth2.Transport{
			Base:   httpcache.NewMemoryCacheTransport(),
			Source: ts,
		},
	}
	ctx := context.Background()
	client := github.NewClient(tc)

	repoEnv := os.Getenv("GITHUB_REPOSITORY")
	org, repo, ok := strings.Cut(repoEnv, "/")
	if !ok {
		log.Fatalf("unexpected GITHUB_REPOSITORY env of %q", repoEnv)
	}
	ref := os.Getenv("GITHUB_REF")
	if ref == "" {
		log.Fatalf("no GITHUB_REF env")
	}

	repoCommit, _, err := client.Repositories.GetCommit(ctx, org, repo, ref, &github.ListOptions{})
	if err != nil {
		log.Fatalf("GetCommit: %v", err)
	}
	commitJ, _ := json.MarshalIndent(repoCommit, "", "  ")
	fmt.Printf("Got commit: %s\n", commitJ)
	//client.Repositories.CreateStatus(ctx, "tailscale", "bradtest2", ""
}
