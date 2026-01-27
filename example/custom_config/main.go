// Example: Custom BM25Config with field boosting and safety limits
//
// This example demonstrates:
// - Custom BM25Config with field boosting
// - MaxDocs and MaxDocTextLen safety limits
// - How boosting affects ranking (name vs description matches)
//
// Run: go run ./example/custom_config/
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jonwraymond/toolindex"
	"github.com/jonwraymond/toolsearch"
)

func main() {
	// Create sample docs where "deploy" appears in different fields
	docs := []toolindex.SearchDoc{
		{
			ID:      "ci:deploy",
			DocText: "deploy application to production continuous integration",
			Summary: toolindex.Summary{
				ID:               "ci:deploy",
				Name:             "deploy", // "deploy" in name
				Namespace:        "ci",
				ShortDescription: "Deploy application to production",
				Tags:             []string{"ci", "cd"},
			},
		},
		{
			ID:      "ops:rollout",
			DocText: "rollout deploy new version gradually canary deployment",
			Summary: toolindex.Summary{
				ID:               "ops:rollout",
				Name:             "rollout",
				Namespace:        "ops",
				ShortDescription: "Gradually deploy new version", // "deploy" in description
				Tags:             []string{"deployment"},         // "deploy" in tags
			},
		},
		{
			ID:      "k8s:apply",
			DocText: "apply kubernetes manifest deploy resources yaml",
			Summary: toolindex.Summary{
				ID:               "k8s:apply",
				Name:             "apply",
				Namespace:        "k8s",
				ShortDescription: "Apply a configuration to deploy resources",
				Tags:             []string{"kubernetes"},
			},
		},
	}

	// Default config: NameBoost=3, NamespaceBoost=2, TagsBoost=2
	fmt.Println("=== Default Config (NameBoost=3) ===")
	defaultSearcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{})
	defer func() {
		if err := defaultSearcher.Close(); err != nil {
			log.Printf("close failed: %v", err)
		}
	}()

	results, err := defaultSearcher.Search("deploy", 10, docs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Search: 'deploy'")
	printResults(results)
	fmt.Println("(ci:deploy ranks first because 'deploy' matches the tool name)")

	// High name boost: name matches dominate
	fmt.Println("\n=== High Name Boost (NameBoost=10) ===")
	highNameBoost := toolsearch.NewBM25Searcher(toolsearch.BM25Config{
		NameBoost:      10,
		NamespaceBoost: 1,
		TagsBoost:      1,
	})
	defer func() {
		if err := highNameBoost.Close(); err != nil {
			log.Printf("close failed: %v", err)
		}
	}()

	results, err = highNameBoost.Search("deploy", 10, docs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Search: 'deploy'")
	printResults(results)
	fmt.Println("(name matches strongly preferred)")

	// Demonstrate safety limits
	fmt.Println("\n=== Safety Limits (MaxDocs=2, MaxDocTextLen=50) ===")
	limitedSearcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{
		MaxDocs:       2,              // Only index first 2 docs
		MaxDocTextLen: 50,             // Truncate long descriptions
	})
	defer func() {
		if err := limitedSearcher.Close(); err != nil {
			log.Printf("close failed: %v", err)
		}
	}()

	// Create docs with long descriptions
	longDocs := make([]toolindex.SearchDoc, 4)
	for i := range longDocs {
		longDocs[i] = toolindex.SearchDoc{
			ID:      fmt.Sprintf("tool:%d", i),
			DocText: strings.Repeat("keyword ", 100), // Very long
			Summary: toolindex.Summary{
				ID:               fmt.Sprintf("tool:%d", i),
				Name:             fmt.Sprintf("tool%d", i),
				ShortDescription: "A tool",
			},
		}
	}

	results, err = limitedSearcher.Search("keyword", 10, longDocs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created 4 docs, but MaxDocs=2 limits indexing\n")
	fmt.Printf("Found %d results (max 2 indexed)\n", len(results))
}

func printResults(results []toolindex.Summary) {
	for i, r := range results {
		fmt.Printf("  %d. %s (%s)\n", i+1, r.ID, r.Name)
	}
}
