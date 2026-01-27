package toolsearch_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jonwraymond/toolindex"
	"github.com/jonwraymond/toolmodel"
	"github.com/jonwraymond/toolsearch"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// TestExample_Basic verifies the basic example works correctly.
// Mirrors: example/basic/main.go
func TestExample_Basic(t *testing.T) {
	searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{})
	defer func() {
		if err := searcher.Close(); err != nil {
			t.Fatalf("close failed: %v", err)
		}
	}()

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

	// Test 1: Search for git-related tools
	t.Run("search_git", func(t *testing.T) {
		results, err := searcher.Search("git", 10, docs)
		if err != nil {
			t.Fatalf("Search error: %v", err)
		}
		if len(results) < 2 {
			t.Errorf("expected at least 2 git results, got %d", len(results))
		}
		// Git tools should rank first
		for _, r := range results[:2] {
			if r.Namespace != "git" {
				t.Errorf("expected git namespace, got %s", r.Namespace)
			}
		}
	})

	// Test 2: Search for containers
	t.Run("search_containers", func(t *testing.T) {
		results, err := searcher.Search("containers", 10, docs)
		if err != nil {
			t.Fatalf("Search error: %v", err)
		}
		if len(results) == 0 {
			t.Fatal("expected results for 'containers'")
		}
		if results[0].ID != "docker:ps" {
			t.Errorf("expected docker:ps first, got %s", results[0].ID)
		}
	})

	// Test 3: No matches
	t.Run("no_matches", func(t *testing.T) {
		results, err := searcher.Search("terraform", 10, docs)
		if err != nil {
			t.Fatalf("Search error: %v", err)
		}
		if len(results) != 0 {
			t.Errorf("expected 0 results for 'terraform', got %d", len(results))
		}
	})

	// Test 4: Empty query returns first N
	t.Run("empty_query", func(t *testing.T) {
		results, err := searcher.Search("", 2, docs)
		if err != nil {
			t.Fatalf("Search error: %v", err)
		}
		if len(results) != 2 {
			t.Errorf("expected 2 results, got %d", len(results))
		}
	})
}

// TestExample_CustomConfig verifies custom configuration works correctly.
// Mirrors: example/custom_config/main.go
func TestExample_CustomConfig(t *testing.T) {
	docs := []toolindex.SearchDoc{
		{
			ID:      "ci:deploy",
			DocText: "deploy application to production continuous integration",
			Summary: toolindex.Summary{
				ID:               "ci:deploy",
				Name:             "deploy",
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
				ShortDescription: "Gradually deploy new version",
				Tags:             []string{"deployment"},
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

	// Test 1: Default config - name matches rank higher
	t.Run("default_config_name_boost", func(t *testing.T) {
		searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{})
		defer func() {
			if err := searcher.Close(); err != nil {
				t.Fatalf("close failed: %v", err)
			}
		}()

		results, err := searcher.Search("deploy", 10, docs)
		if err != nil {
			t.Fatalf("Search error: %v", err)
		}
		if len(results) == 0 {
			t.Fatal("expected results")
		}
		// ci:deploy should rank first because "deploy" is in the name
		if results[0].ID != "ci:deploy" {
			t.Errorf("expected ci:deploy first (name match), got %s", results[0].ID)
		}
	})

	// Test 2: High name boost amplifies effect
	t.Run("high_name_boost", func(t *testing.T) {
		searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{
			NameBoost:      10,
			NamespaceBoost: 1,
			TagsBoost:      1,
		})
		defer func() {
			if err := searcher.Close(); err != nil {
				t.Fatalf("close failed: %v", err)
			}
		}()

		results, err := searcher.Search("deploy", 10, docs)
		if err != nil {
			t.Fatalf("Search error: %v", err)
		}
		if len(results) == 0 {
			t.Fatal("expected results")
		}
		if results[0].ID != "ci:deploy" {
			t.Errorf("expected ci:deploy first with high name boost, got %s", results[0].ID)
		}
	})

	// Test 3: MaxDocs limits indexed documents
	t.Run("max_docs_limit", func(t *testing.T) {
		searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{
			MaxDocs: 2,
		})
		defer func() {
			if err := searcher.Close(); err != nil {
				t.Fatalf("close failed: %v", err)
			}
		}()

		longDocs := make([]toolindex.SearchDoc, 4)
		for i := range longDocs {
			longDocs[i] = toolindex.SearchDoc{
				ID:      fmt.Sprintf("tool:%d", i),
				DocText: strings.Repeat("keyword ", 100),
				Summary: toolindex.Summary{
					ID:               fmt.Sprintf("tool:%d", i),
					Name:             fmt.Sprintf("tool%d", i),
					ShortDescription: "A tool",
				},
			}
		}

		results, err := searcher.Search("keyword", 10, longDocs)
		if err != nil {
			t.Fatalf("Search error: %v", err)
		}
		// Should be limited by MaxDocs=2
		if len(results) > 2 {
			t.Errorf("expected at most 2 results (MaxDocs), got %d", len(results))
		}
	})

	// Test 4: MaxDocTextLen truncates long descriptions
	t.Run("max_doc_text_len", func(t *testing.T) {
		searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{
			MaxDocTextLen: 50,
		})
		defer func() {
			if err := searcher.Close(); err != nil {
				t.Fatalf("close failed: %v", err)
			}
		}()

		// "uniqueword" is past the truncation point
		longDoc := []toolindex.SearchDoc{
			{
				ID:      "long-doc",
				DocText: strings.Repeat("padding ", 100) + "uniqueword",
				Summary: toolindex.Summary{ID: "long-doc", Name: "LongDoc"},
			},
		}

		results, err := searcher.Search("uniqueword", 10, longDoc)
		if err != nil {
			t.Fatalf("Search error: %v", err)
		}
		// Should not find "uniqueword" since it's truncated
		if len(results) != 0 {
			t.Errorf("expected 0 results (word truncated), got %d", len(results))
		}
	})
}

// TestExample_ToolindexIntegration verifies full toolindex integration works.
// Mirrors: example/toolindex_integration/main.go
func TestExample_ToolindexIntegration(t *testing.T) {
	// Create a BM25 searcher
	searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{
		NameBoost:      3,
		NamespaceBoost: 2,
		TagsBoost:      2,
	})

	// Inject into toolindex
	idx := toolindex.NewInMemoryIndex(toolindex.IndexOptions{
		Searcher: searcher,
	})

	// Helper to create MCP tools
	newMCPTool := func(name, description string) mcp.Tool {
		return mcp.Tool{
			Name:        name,
			Description: description,
			InputSchema: map[string]any{"type": "object"},
		}
	}
	mcpBackend := func(serverName string) toolmodel.ToolBackend {
		return toolmodel.ToolBackend{
			Kind: toolmodel.BackendKindMCP,
			MCP:  &toolmodel.MCPBackend{ServerName: serverName},
		}
	}

	// Register tools
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
				Tool:      newMCPTool("docker_ps", "List containers"),
				Namespace: "docker",
				Tags:      []string{"containers", "devops"},
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
	}

	for _, tt := range tools {
		if err := idx.RegisterTool(tt.tool, tt.backend); err != nil {
			t.Fatalf("Failed to register %s: %v", tt.tool.Name, err)
		}
	}

	// Test 1: Search for git tools
	t.Run("search_git", func(t *testing.T) {
		results, err := idx.Search("git", 10)
		if err != nil {
			t.Fatalf("Search error: %v", err)
		}
		if len(results) < 2 {
			t.Errorf("expected at least 2 git results, got %d", len(results))
		}
		// All results should be git-related
		for _, r := range results {
			if r.Namespace != "git" {
				t.Errorf("expected git namespace, got %s for %s", r.Namespace, r.ID)
			}
		}
	})

	// Test 2: Search for container tools
	t.Run("search_containers", func(t *testing.T) {
		results, err := idx.Search("containers", 10)
		if err != nil {
			t.Fatalf("Search error: %v", err)
		}
		if len(results) == 0 {
			t.Fatal("expected results for 'containers'")
		}
		// docker_ps should be found
		found := false
		for _, r := range results {
			if r.Name == "docker_ps" {
				found = true
				break
			}
		}
		if !found {
			t.Error("expected docker_ps in containers results")
		}
	})

	// Test 3: Search for devops tools (tag-based)
	t.Run("search_devops_tag", func(t *testing.T) {
		results, err := idx.Search("devops", 10)
		if err != nil {
			t.Fatalf("Search error: %v", err)
		}
		// Should find docker and kubectl tools (both have devops tag)
		if len(results) < 2 {
			t.Errorf("expected at least 2 devops-tagged results, got %d", len(results))
		}
	})

	// Test 4: List namespaces
	t.Run("list_namespaces", func(t *testing.T) {
		namespaces, err := idx.ListNamespaces()
		if err != nil {
			t.Fatalf("ListNamespaces error: %v", err)
		}
		expected := map[string]bool{"git": true, "docker": true, "kubectl": true}
		for _, ns := range namespaces {
			delete(expected, ns)
		}
		if len(expected) > 0 {
			t.Errorf("missing namespaces: %v", expected)
		}
	})
}
