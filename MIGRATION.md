# Migration Guide: toolsearch to tooldiscovery/search

This document provides instructions for migrating from the deprecated
`github.com/jonwraymond/toolsearch` package to
`github.com/jonwraymond/tooldiscovery/search`.

## Why the Migration?

The `toolsearch` package has been consolidated into the `tooldiscovery`
repository as part of the ApertureStack simplification effort. The search
functionality now lives alongside discovery, indexing, and registry features
in a single cohesive package.

## Import Path Changes

Update your Go imports as follows:

| Old Import | New Import |
|------------|------------|
| `github.com/jonwraymond/toolsearch` | `github.com/jonwraymond/tooldiscovery/search` |

## Code Changes

### Basic Usage

**Before:**

```go
import (
    "github.com/jonwraymond/toolindex"
    "github.com/jonwraymond/toolsearch"
)

searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{})

idx := toolindex.NewInMemoryIndex(toolindex.IndexOptions{
    Searcher: searcher,
})
```

**After:**

```go
import (
    "github.com/jonwraymond/tooldiscovery/search"
)

searcher := search.NewBM25Searcher(search.BM25Config{})

// Use with tooldiscovery's integrated index
```

### Configuration

The `BM25Config` struct and its fields remain the same:

```go
// Configuration options are unchanged
config := search.BM25Config{
    MaxDocs:       1000,
    MaxDocTextLen: 4096,
}
searcher := search.NewBM25Searcher(config)
```

## Dependency Updates

Update your `go.mod`:

```bash
# Remove old dependency
go mod edit -droprequire github.com/jonwraymond/toolsearch

# Add new dependency
go get github.com/jonwraymond/tooldiscovery/search@latest

# Tidy up
go mod tidy
```

## Timeline

- **Now**: Begin migrating to `tooldiscovery/search`
- **Next minor release**: `toolsearch` will be archived
- **Future**: No further updates to `toolsearch`

## Support

For questions or issues during migration, please open an issue in the
[tooldiscovery repository](https://github.com/jonwraymond/tooldiscovery/issues).
