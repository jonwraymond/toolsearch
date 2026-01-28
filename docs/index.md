# toolsearch

`toolsearch` provides higher-quality search strategies for `toolindex` while
keeping heavy dependencies out of the core registry. Today it ships a BM25
searcher backed by Bleve.

## Motivation

- **Keep toolindex small** while enabling stronger ranking
- **Experiment safely** with search without changing core behavior
- **Deterministic results** for stable agent behavior

## Key APIs

- `BM25Searcher` (implements `toolindex.Searcher`)
- `BM25Config` for ranking and safety caps

## Quickstart

```go
searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{})
idx := toolindex.NewInMemoryIndex(toolindex.IndexOptions{Searcher: searcher})
```

## Usability notes

- Deterministic ordering avoids jitter in tool selection
- Index fingerprinting avoids rebuilds on no-op updates

## Next

- Ranking pipeline: `architecture.md`
- Configuration and caps: `usage.md`
- Examples: `examples.md`
- Design Notes: `design-notes.md`
- User Journey: `user-journey.md`
