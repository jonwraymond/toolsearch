// Example: Basic usage of toolsearch BM25Searcher
//
// This example demonstrates:
// - Creating a searcher with default config
// - Searching with a simple query
// - Handling results and empty results
// - Cleanup with Close()
//
// Run: go run ./example/basic/
package main

import (
	"fmt"
	"log"

	"github.com/jonwraymond/toolindex"
	"github.com/jonwraymond/toolsearch"
)

func main() {
	// Create a BM25 searcher with default configuration.
	// Defaults: NameBoost=3, NamespaceBoost=2, TagsBoost=2
	searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{})
	defer searcher.Close()

	// Create sample tool documents (normally provided by toolindex)
	docs := []toolindex.SearchDoc{
		{
			ID:      "git:status",
			DocText: "git status show working tree status version control",
			Summary: toolindex.Summary{
				ID:               "git:status",
				Name:             "status",
				Namespace:        "git",
				ShortDescription: "Show the working tree status",
				Tags:             []string{"vcs", "git"},
			},
		},
		{
			ID:      "git:commit",
			DocText: "git commit save changes to repository version control",
			Summary: toolindex.Summary{
				ID:               "git:commit",
				Name:             "commit",
				Namespace:        "git",
				ShortDescription: "Record changes to the repository",
				Tags:             []string{"vcs", "git"},
			},
		},
		{
			ID:      "docker:ps",
			DocText: "docker ps list containers running processes",
			Summary: toolindex.Summary{
				ID:               "docker:ps",
				Name:             "ps",
				Namespace:        "docker",
				ShortDescription: "List containers",
				Tags:             []string{"containers", "docker"},
			},
		},
		{
			ID:      "kubectl:get",
			DocText: "kubectl get display resources kubernetes pods services",
			Summary: toolindex.Summary{
				ID:               "kubectl:get",
				Name:             "get",
				Namespace:        "kubectl",
				ShortDescription: "Display one or many resources",
				Tags:             []string{"kubernetes", "k8s"},
			},
		},
	}

	// Search for git-related tools
	fmt.Println("=== Search: 'git' ===")
	results, err := searcher.Search("git", 10, docs)
	if err != nil {
		log.Fatal(err)
	}
	printResults(results)

	// Search for container tools
	fmt.Println("\n=== Search: 'containers' ===")
	results, err = searcher.Search("containers", 10, docs)
	if err != nil {
		log.Fatal(err)
	}
	printResults(results)

	// Search with no matches
	fmt.Println("\n=== Search: 'terraform' (no matches) ===")
	results, err = searcher.Search("terraform", 10, docs)
	if err != nil {
		log.Fatal(err)
	}
	if len(results) == 0 {
		fmt.Println("No results found")
	}

	// Empty query returns first N documents
	fmt.Println("\n=== Empty query (returns first 2) ===")
	results, err = searcher.Search("", 2, docs)
	if err != nil {
		log.Fatal(err)
	}
	printResults(results)
}

func printResults(results []toolindex.Summary) {
	for i, r := range results {
		fmt.Printf("%d. %s - %s\n", i+1, r.ID, r.ShortDescription)
	}
}
