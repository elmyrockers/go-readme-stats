package main

import (
	"os"
	"context"
	"fmt"
	"slices"
	// "github.com/davecgh/go-spew/spew"

	"github.com/google/go-github/v60/github"
)




//------------------------------------------------------------------------------------------------------
type stats struct {
	client *github.Client
}

func newStats() *stats {
	token := os.Getenv( "GITHUB_TOKEN" )
	client := github.NewClient(nil).WithAuthToken( token )
	return &stats{ client: client }
}

func ( s *stats ) get_repos( visibility string, sort string ) []*github.Repository {
	// Set options
		options := &github.RepositoryListByAuthenticatedUserOptions{
		    Visibility:  visibility,
		    Affiliation: "owner",
		    Sort:        sort, //pushed & full_name
		    ListOptions: github.ListOptions{PerPage: 100},
		}

	// Get list of repos
		ctx := context.Background()
		repos, _, err := s.client.Repositories.ListByAuthenticatedUser(ctx, options)
		if err != nil {
			fmt.Println( err )
		}
	return repos
}


type Language struct {
	Name  string
	Bytes int
	Percentage float64
}
func ( s *stats ) most_used_languages( count int ) []Language  {
	repos := s.get_repos( "public", "full_name" )
	ctx := context.Background()

	// Get list of languages with their size
		languages := map[string]int{}
		totalBytes := 0
		for _, repo := range repos {
			if repo.GetFork() { continue } //skip fork

			owner := repo.GetOwner().GetLogin()
			name := repo.GetName()
			langData, _, _ := s.client.Repositories.ListLanguages(ctx, owner, name)
			for lang, bytes := range langData {
				languages[ lang ] += bytes
				totalBytes += bytes
			}
		}

	// Sort languages with their size (from biggest to smallest)
		sortedLanguages := []Language{}
		for name, bytes := range languages {
			sortedLanguages = append( sortedLanguages,Language{
				Name: name,
				Bytes: bytes,
				Percentage: (float64(bytes) / float64(totalBytes)) * 100,
			})
		}
		slices.SortFunc( sortedLanguages, func(a, b Language) int {
			return b.Bytes - a.Bytes
		})

	// Get Most Used of Languages
		sortedLanguages = sortedLanguages[: count ]
	return sortedLanguages
}