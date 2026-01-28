# toolsearch

`toolsearch` provides higher-quality search strategies for `toolindex` while
keeping heavy dependencies out of the core registry. Today it ships a BM25
searcher backed by Bleve.

## What this library provides

- `BM25Searcher` implementing `toolindex.Searcher`
- Deterministic ranking and tie-breaking
- Index caching + fingerprinting

## Quickstart

```go
searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{})
idx := toolindex.NewInMemoryIndex(toolindex.IndexOptions{Searcher: searcher})
```

## Next

- Ranking pipeline: `architecture.md`
- Configuration and safety caps: `usage.md`
- Examples: `examples.md`
