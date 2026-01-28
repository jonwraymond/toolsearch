# Examples

## Deterministic ranking

```go
searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{})
idx := toolindex.NewInMemoryIndex(toolindex.IndexOptions{Searcher: searcher})

results, _ := idx.Search("repo", 10)
// results are deterministic across runs
```

## Custom boosts

```go
searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{
  NameBoost:      5,
  NamespaceBoost: 3,
  TagsBoost:      1,
})
```
