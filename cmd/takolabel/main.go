package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"takolabel"
	"takolabel/config"
)

func main() {
	ctx := context.Background()

	viper.SetConfigName("takolabel")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	githubToken := viper.GetString("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	baseUrl := viper.GetString("BASE_URL")
	var client *github.Client
	if baseUrl != "" {
		client, err = github.NewEnterpriseClient(baseUrl, baseUrl, tc)
		if err != nil {
			panic(fmt.Errorf("error setting ghe client: %s", err))
		}
	} else {
		client = github.NewClient(tc)
	}

	viper.SetConfigName("labels")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	var labels []config.Label
	err = viper.UnmarshalKey("labels", &labels)
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	var repositories []config.Repository
	err = viper.UnmarshalKey("repositories", &repositories)
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	for _, repository := range repositories {
		for _, label := range labels {
			_, err := takolabel.CreateLabel(ctx, client.Issues, label, repository)
			if err != nil {
				fmt.Printf("error creating label \"%s\" for repository \"%s\": %s\n", label.Name, repository.Org+"/"+repository.Repo, err)
			} else {
				fmt.Printf("created label \"%s\" for repository \"%s\"\n", label.Name, repository.Org+"/"+repository.Repo)
			}
		}
	}
}
