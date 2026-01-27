// Example: Full integration with toolindex
//
// This example demonstrates:
// - Injecting BM25Searcher into toolindex.IndexOptions
// - Real-world tool catalog setup
// - Using toolindex.Search() with BM25 ranking underneath
//
// Run: go run ./example/toolindex_integration/
package main

import (
	"fmt"
	"log"

	"github.com/jonwraymond/toolindex"
	"github.com/jonwraymond/toolmodel"
	"github.com/jonwraymond/toolsearch"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	// Create a BM25 searcher with custom config
	searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{
		NameBoost:      3,
		NamespaceBoost: 2,
		TagsBoost:      2,
	})

	// Inject the BM25 searcher into toolindex
	idx := toolindex.NewInMemoryIndex(toolindex.IndexOptions{
		Searcher: searcher,
	})

	// Register realistic tools
	tools := []struct {
		tool    toolmodel.Tool
		backend toolmodel.ToolBackend
	}{
		{
			tool: toolmodel.Tool{
				Tool:      newMCPTool("git_status", "Show the working tree status"),
				Namespace: "git",
				Tags:      []string{"vcs", "version-control"},
			},
			backend: mcpBackend("git-mcp"),
		},
		{
			tool: toolmodel.Tool{
				Tool:      newMCPTool("git_commit", "Record changes to the repository"),
				Namespace: "git",
				Tags:      []string{"vcs", "version-control"},
			},
			backend: mcpBackend("git-mcp"),
		},
		{
			tool: toolmodel.Tool{
				Tool:      newMCPTool("git_push", "Update remote refs along with associated objects"),
				Namespace: "git",
				Tags:      []string{"vcs", "version-control", "remote"},
			},
			backend: mcpBackend("git-mcp"),
		},
		{
			tool: toolmodel.Tool{
				Tool:      newMCPTool("docker_ps", "List containers"),
				Namespace: "docker",
				Tags:      []string{"containers", "devops"},
			},
			backend: mcpBackend("docker-mcp"),
		},
		{
			tool: toolmodel.Tool{
				Tool:      newMCPTool("docker_build", "Build an image from a Dockerfile"),
				Namespace: "docker",
				Tags:      []string{"containers", "devops", "images"},
			},
			backend: mcpBackend("docker-mcp"),
		},
		{
			tool: toolmodel.Tool{
				Tool:      newMCPTool("kubectl_get", "Display one or many resources"),
				Namespace: "kubectl",
				Tags:      []string{"kubernetes", "k8s", "devops"},
			},
			backend: mcpBackend("k8s-mcp"),
		},
		{
			tool: toolmodel.Tool{
				Tool:      newMCPTool("kubectl_apply", "Apply a configuration to a resource"),
				Namespace: "kubectl",
				Tags:      []string{"kubernetes", "k8s", "devops"},
			},
			backend: mcpBackend("k8s-mcp"),
		},
	}

	// Register all tools
	for _, t := range tools {
		if err := idx.RegisterTool(t.tool, t.backend); err != nil {
			log.Fatalf("Failed to register %s: %v", t.tool.Name, err)
		}
	}

	// Now search through toolindex (uses BM25 underneath)
	fmt.Println("=== Search via toolindex: 'git' ===")
	results, err := idx.Search("git", 10)
	if err != nil {
		log.Fatal(err)
	}
	printResults(results)

	fmt.Println("\n=== Search via toolindex: 'container' ===")
	results, err = idx.Search("container", 10)
	if err != nil {
		log.Fatal(err)
	}
	printResults(results)

	fmt.Println("\n=== Search via toolindex: 'devops' ===")
	results, err = idx.Search("devops", 10)
	if err != nil {
		log.Fatal(err)
	}
	printResults(results)

	// List all namespaces
	fmt.Println("\n=== Registered Namespaces ===")
	namespaces, err := idx.ListNamespaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, ns := range namespaces {
		fmt.Printf("  - %s\n", ns)
	}
}

// newMCPTool creates a minimal MCP tool for demonstration
func newMCPTool(name, description string) mcp.Tool {
	return mcp.Tool{
		Name:        name,
		Description: description,
		InputSchema: map[string]any{"type": "object"},
	}
}

// mcpBackend creates an MCP backend for a server
func mcpBackend(serverName string) toolmodel.ToolBackend {
	return toolmodel.ToolBackend{
		Kind: toolmodel.BackendKindMCP,
		MCP:  &toolmodel.MCPBackend{ServerName: serverName},
	}
}

func printResults(results []toolindex.Summary) {
	if len(results) == 0 {
		fmt.Println("  No results found")
		return
	}
	for i, r := range results {
		fmt.Printf("  %d. %s - %s\n", i+1, r.ID, r.ShortDescription)
	}
}
