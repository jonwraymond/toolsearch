# toolsearch

`toolsearch` provides higher-quality search strategies for `toolindex` while
keeping heavy dependencies out of the core registry. Today it ships a BM25
searcher backed by Bleve.

## Key APIs

- `BM25Searcher` (implements `toolindex.Searcher`)
- `BM25Config` for ranking and safety caps

## Quickstart

```go
searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{})
idx := toolindex.NewInMemoryIndex(toolindex.IndexOptions{Searcher: searcher})
```

## Next

- Ranking pipeline: `architecture.md`
- Configuration and caps: `usage.md`
- Examples: `examples.md`
